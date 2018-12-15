// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package graceful

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Graceful tries to properly shut down a web server upon interrupts.
func Graceful(srv *http.Server, listen func(), sig ...os.Signal) {
	idle := make(chan struct{})

	go func() {
		i := make(chan os.Signal, 1)
		signal.Notify(i, sig...)

		<-i

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("server listening on port '%s' gracefully shutdown\n",
				srv.Addr)
		}

		close(idle)
	}()

	listen()

	<-idle
}
