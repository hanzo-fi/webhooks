<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/formancehq/webhooks/pkg/client"
	"github.com/formancehq/webhooks/pkg/client/models/components"
	"github.com/formancehq/webhooks/pkg/client/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithSecurity(components.Security{
			ClientID:     "<YOUR_CLIENT_ID_HERE>",
			ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
			TokenURL:     "/oauth/token",
		}),
	)

	res, err := s.Webhooks.V1.GetManyConfigs(ctx, operations.GetManyConfigsRequest{
		ID:       client.Pointer("4997257d-dfb6-445b-929c-cbe2ab182818"),
		Endpoint: client.Pointer("https://example.com"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ConfigsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->