package main

type configModel struct {
	mongoUri    string
	mongoDb     string
	tokenSecret string
	tokenExp    string
	serveUri    string
}

var config = configModel{
	mongoUri:    "mongodb://admin:Damba11a@ds253567.mlab.com:53567/react-recipes", // Mongo Uri example: mongodb://admin:123456@localhost:37812/react-recipes
	mongoDb:     "react-recipes",                                        // DB name
	tokenSecret: "secret",                                               // Secret to use in Tokens
	tokenExp:    "1h",                                                   // Expiration of Token
	serveUri:    ":53567",                                                // Serve
}
