// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthorizedSecurityGroupRulesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListGatewayAuthorizedSecurityGroupRulesRequest struct {
}

func (s ListGatewayAuthorizedSecurityGroupRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizedSecurityGroupRulesRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizedSecurityGroupRulesRequest) Validate() error {
	return dara.Validate(s)
}
