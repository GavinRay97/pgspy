package pgspy

import "os"

// Start will start a new server
func Start() {
	// Read DATABASE_PORT and PROXY_PORT from environment variables
	var databasePort string = getEnv("DATABASE_PORT", "5432")
	var proxyPort string = getEnv("PROXY_PORT", "5433")

	pgHost := "0.0.0.0" + ":" + databasePort
	proxyHost := "0.0.0.0" + ":" + proxyPort

	proxy := NewProxy(pgHost, proxyHost)
	queryWatcher := QueryWatcher{}
	proxy.OnMessage = queryWatcher.OnMessage
	// proxy.OnMessage = func(msg PostgresMessage) {
	// 	if msg.Outgoing {
	// 		log.Infof("-> %s\n", msg.TypeIdentifier)
	// 	} else {
	// 		log.Infof("<- %s \n %s \n", msg.TypeIdentifier, msg.Payload)
	// 	}
	// }
	proxy.Start()
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
