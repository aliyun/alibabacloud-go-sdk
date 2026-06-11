// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayQuotaRuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteGatewayQuotaRuleRequest struct {
}

func (s DeleteGatewayQuotaRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayQuotaRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayQuotaRuleRequest) Validate() error {
	return dara.Validate(s)
}
