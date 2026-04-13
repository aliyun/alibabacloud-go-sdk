// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteGatewayRequest struct {
}

func (s DeleteGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) Validate() error {
	return dara.Validate(s)
}
