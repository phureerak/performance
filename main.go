package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Performance struct {
	ID          *string `json:"id,omitempty"`
	ActivePower *string `json:"active_power,omitempty"`
	PowerInput  *string `json:"power_input,omitempty"`
}

func main() {
	app := fiber.New()

	db, err := sql.Open("mysql", "root:ฟกทรื@tcp(127.0.0.1:3306)/machine")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("db connected.")

	app.Get("/sum/:name", func(c *fiber.Ctx) error {
		var perform Performance
		if name := c.Params("name"); name == "active_power" {
			perform.ActivePower = String(GetSumPerformance(db, "performance.active_power"))
		} else if name == "power_input" {
			perform.PowerInput = String(GetSumPerformance(db, "performance.power_input"))
		} else if name == "all" {
			perform.ActivePower = String(GetSumPerformance(db, "performance.active_power"))
			perform.PowerInput = String(GetSumPerformance(db, "performance.power_input"))
		}

		return c.JSON(perform)
	})

	app.Get("/sum", func(c *fiber.Ctx) error {
		perform := Performance{
			ActivePower: String(GetSumPerformance(db, "performance.active_power")),
			PowerInput:  String(GetSumPerformance(db, "performance.power_input")),
		}

		return c.JSON(perform)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(GetAllPerformance(db))
	})

	app.Listen(":3000")
}

func GetSumPerformance(db *sql.DB, column string) string {
	rows, err := db.Query(fmt.Sprintf("SELECT sum(%s) FROM machine.performance", column))
	if err != nil {
		panic(err.Error())
	}

	var sum string
	for rows.Next() {
		rows.Scan(&sum)
	}
	return sum
}

func GetAllPerformance(db *sql.DB) []Performance {
	rows, err := db.Query("SELECT * FROM machine.performance")
	if err != nil {
		panic(err.Error())
	}

	var performs []Performance
	for rows.Next() {
		var id, activePower, powerInput int
		rows.Scan(&id, &activePower, &powerInput)
		performs = append(performs, Performance{
			ID:          String(fmt.Sprintf("%d", id)),
			ActivePower: String(fmt.Sprintf("%d", activePower)),
			PowerInput:  String(fmt.Sprintf("%d", powerInput)),
		})
	}
	return performs
}

func Int(data int) *int {
	return &data
}

func String(data string) *string {
	return &data
}
