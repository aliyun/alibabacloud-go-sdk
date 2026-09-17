// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResourceQuotaUsageRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGatewayResourceQuotaUsageRequest struct {
}

func (s GetGatewayResourceQuotaUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageRequest) Validate() error {
	return dara.Validate(s)
}
