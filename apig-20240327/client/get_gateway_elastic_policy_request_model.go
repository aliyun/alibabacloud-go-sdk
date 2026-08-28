// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayElasticPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGatewayElasticPolicyRequest struct {
}

func (s GetGatewayElasticPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayElasticPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayElasticPolicyRequest) Validate() error {
	return dara.Validate(s)
}
