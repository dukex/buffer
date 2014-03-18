# Go Buffer, Go!

```
client := NewClient("MyAccessToken")
```


## Create Update

```
client.CreateUpdate("Test My Awesome Post", []string{"profile1"}, map[string]interface{}{"now": false})
```

## Profiles

```
profiles := client.Profiles()
```
