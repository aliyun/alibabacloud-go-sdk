// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayWebsocketUsageRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGatewayWebsocketUsageRequest struct {
}

func (s GetGatewayWebsocketUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayWebsocketUsageRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayWebsocketUsageRequest) Validate() error {
	return dara.Validate(s)
}
