# Temp data, will pull live data later
$Client = "XWare"
$Date = "21/02/23"
$Server = "x-01"
$Space = "22"

# Pull data together
$loggeddata = @{ 'client' = $Client; 'date' = $Date; 'server' = $Server; 'space' = $Space}
$Uri = 'http://sendit-01.local:8080/loggeddatas'
$Header = @{ 'Content-Type' = 'application/json'}

# Send the data in json format
Invoke-RestMethod -Uri $Uri -Headers $Header -Method Post -Body ($loggeddata|ConvertTo-Json)