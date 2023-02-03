### run command :
```
go run main.go 
```
### path [home](show all data on db) : localhost:3000/
example : 
```
[
    {
        "id": "1",
        "active_power": "100",
        "power_input": "200"
    },
    {
        "id": "10",
        "active_power": "81",
        "power_input": "318"
    },
    {
        "id": "100",
        "active_power": "720",
        "power_input": "783"
    },
]
```

### path [sum](show sum all field on db) : localhost:3000/sum
example : 
```
{
    "active_power": "474932",
    "power_input": "446476"
}
```

### path [sum](show sum active_power field) : localhost:3000/sum/active_power
example : 
```
{
    "active_power": "474932"
}
```

### path [sum](show sum power_input field) : localhost:3000/sum/power_input
example : 
```
{
    "power_input": "446476"
}
```