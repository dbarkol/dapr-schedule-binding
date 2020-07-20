const express = require('express')
const bodyParser = require('body-parser')
const request = require('request')
const app = express()
app.use(bodyParser.json({ type: 'application/*+json' }));

const port = 3000

app.post('/schedule-trigger', (req, res) => {
    request('https://api.chucknorris.io/jokes/random?category=sport', 
        { json: true }, 
        function(error, response, body){
            console.log('body:', body.value); 
    });

    console.log("schedule trigger invoked: " + new Date().toISOString());
});

app.listen(port, () => console.log(`app listening on port ${port}!`))	
