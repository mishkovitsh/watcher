const express = require('express')
const app = express()
const r = require('./routers/r')
const dotenv = require('dotenv').config()
app.use('/api/', r)



try {
    app.listen(process.env.port, () => {
        console.log('1')
    })
} 
catch (err) {
    console.log(err)
}