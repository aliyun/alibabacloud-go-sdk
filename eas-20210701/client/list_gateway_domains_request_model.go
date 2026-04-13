// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListGatewayDomainsRequest struct {
}

func (s ListGatewayDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainsRequest) Validate() error {
	return dara.Validate(s)
}
