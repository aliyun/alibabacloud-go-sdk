// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthorizableSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsClusterId(v string) *ListGatewayAuthorizableSecurityGroupsRequest
	GetCsClusterId() *string
}

type ListGatewayAuthorizableSecurityGroupsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// c4a21b3560fad4ec299f3e63f4***
	CsClusterId *string `json:"csClusterId,omitempty" xml:"csClusterId,omitempty"`
}

func (s ListGatewayAuthorizableSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthorizableSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthorizableSecurityGroupsRequest) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *ListGatewayAuthorizableSecurityGroupsRequest) SetCsClusterId(v string) *ListGatewayAuthorizableSecurityGroupsRequest {
	s.CsClusterId = &v
	return s
}

func (s *ListGatewayAuthorizableSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
