const restify = require('restify');
const routes = require('./routes');

let app = restify.createServer();


routes.init(app);

app.use(restify.plugins.queryParser());

export default app;
