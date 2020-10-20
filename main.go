package main

import (
	"github.com/202lp2/go2/routers"
)

func main() {

	r := routers.SetupRouter()
	r.Run("localhost:8085")
}
