// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyBasedRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterPolicyBasedRouteRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterPolicyBasedRouteRequest
	GetDryRun() *bool
	SetPolicyBasedRouteId(v string) *DeleteTransitRouterPolicyBasedRouteRequest
	GetPolicyBasedRouteId() *string
}

type DeleteTransitRouterPolicyBasedRouteRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, and you must ensure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the deletion request. The system checks items such as permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the dry run, the policy-based routing rule is deleted.
	//
	// - **true**: performs a dry run only and does not delete the policy-based routing rule. The system checks whether required parameters are specified and whether the request format is valid. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the policy-based routing rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// pbr-123abczcx
	PolicyBasedRouteId *string `json:"PolicyBasedRouteId,omitempty" xml:"PolicyBasedRouteId,omitempty"`
}

func (s DeleteTransitRouterPolicyBasedRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyBasedRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) GetPolicyBasedRouteId() *string {
	return s.PolicyBasedRouteId
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) SetClientToken(v string) *DeleteTransitRouterPolicyBasedRouteRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) SetDryRun(v bool) *DeleteTransitRouterPolicyBasedRouteRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) SetPolicyBasedRouteId(v string) *DeleteTransitRouterPolicyBasedRouteRequest {
	s.PolicyBasedRouteId = &v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteRequest) Validate() error {
	return dara.Validate(s)
}
