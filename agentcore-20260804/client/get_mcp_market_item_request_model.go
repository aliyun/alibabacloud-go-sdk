// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpMarketItemRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMcpMarketItemRequest struct {
}

func (s GetMcpMarketItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpMarketItemRequest) GoString() string {
	return s.String()
}

func (s *GetMcpMarketItemRequest) Validate() error {
	return dara.Validate(s)
}
