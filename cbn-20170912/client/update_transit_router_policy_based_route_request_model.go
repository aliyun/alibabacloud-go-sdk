// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPolicyBasedRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTransitRouterPolicyBasedRouteRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTransitRouterPolicyBasedRouteRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateTransitRouterPolicyBasedRouteRequest
	GetDryRun() *bool
	SetName(v string) *UpdateTransitRouterPolicyBasedRouteRequest
	GetName() *string
	SetPolicyBasedRouteId(v string) *UpdateTransitRouterPolicyBasedRouteRequest
	GetPolicyBasedRouteId() *string
}

type UpdateTransitRouterPolicyBasedRouteRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client. Make sure that the value is unique among different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the policy-based route rule.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: sends a check request and does not update the policy-based route rule. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request and directly updates the policy-based route rule after the check succeeds.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the policy-based route rule.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance ID of the policy-based routing rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// pbr-123abczcx
	PolicyBasedRouteId *string `json:"PolicyBasedRouteId,omitempty" xml:"PolicyBasedRouteId,omitempty"`
}

func (s UpdateTransitRouterPolicyBasedRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPolicyBasedRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) GetPolicyBasedRouteId() *string {
	return s.PolicyBasedRouteId
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) SetClientToken(v string) *UpdateTransitRouterPolicyBasedRouteRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) SetDescription(v string) *UpdateTransitRouterPolicyBasedRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) SetDryRun(v bool) *UpdateTransitRouterPolicyBasedRouteRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) SetName(v string) *UpdateTransitRouterPolicyBasedRouteRequest {
	s.Name = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) SetPolicyBasedRouteId(v string) *UpdateTransitRouterPolicyBasedRouteRequest {
	s.PolicyBasedRouteId = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteRequest) Validate() error {
	return dara.Validate(s)
}
