# telegram-notification-goLang
This project was created to quickly raise the service of notifications in telegrams. Every time I started making a project, I needed to raise the server of logs, notifications and other things, if the server of logs - I did it a long time ago, then notifications - you can study in this repo. Everything is simple, there is no complex code, it is important to raise the API and start using

## Instruction:
    1. Change .env.example to .env 
        1.1 telegram_id= # U telegram bot id
        1.2 port=8080 # port for start API
    2. Insert you token and port to .env file
    3. Insert you Ids chat to send to data.json

The API has two parameters, one is required, the other is not. Mandatory - message, optional - this is the IDs of the chats, if there are none, then it will take from the data.son, if there is, it will take from the body from request.

Curl Example:
```
curl -XPOST -H "Content-type: application/json" -d '{"message" : "test Message" }' 'http://localhost:8080/notify/send'

response: {"Success": true, "Message": "Message sended"}

curl -XPOST -H "Content-type: application/json" -d '{"message" : "test Message", "Ids" : ["696300339"] }' 'http://localhost:8080/notify/send'

response: {"Success": true, "Message": "Message sended"}
```

# Donate:

    BTC:  192TC7d7ZRYJQbQnAWvMpkccnBNQN1ae6R
    ETH:  0x7d1082d952f4d584ae2910e14018f4dce7495c74
    LTC:  MLx6wmFjXfBTKj6JfB5NXaiKjNLeEntRoZ
    DOGE: DHCjW71EWBzvv43XPXVJc491brcBJXXq88

# author: 

    Name:          Nikita
    Company:       SmartWorld
    Position:      TeamLead
    Mail:          n.vtorushin@inbox.ru
    TG:            @nikitavoryet
    Year of birth: 1999
    FullStack:     JS/GO

# In the plans :
```
- [ ] Docker
- [ ] Use DB for IDs chat
- [ ] CRUD Ids chat
- [ ] Role-based access
```

# License

MIT

**Free Software, Hell Yeah!**