// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetSupportedZoneRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListGatewayIntranetSupportedZoneRequest struct {
}

func (s ListGatewayIntranetSupportedZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetSupportedZoneRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetSupportedZoneRequest) Validate() error {
	return dara.Validate(s)
}
