// NOTE: `package main` designates that this file is the entrypoint.
package main

// NOTE: Import our own module from the `router` sub-directory.
import "server/router"

const PORT = ":3000";

func main() {
    // NOTE: How the module name prefixes the exported function!
    router.Init(PORT);
}
