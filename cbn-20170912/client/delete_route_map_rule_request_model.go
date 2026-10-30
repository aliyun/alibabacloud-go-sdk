// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRouteMapRuleRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteRouteMapRuleRequest
	GetDryRun() *bool
	SetRouteMapRuleId(v string) *DeleteRouteMapRuleRequest
	GetRouteMapRuleId() *string
}

type DeleteRouteMapRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a check request and does not delete the route map rule. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the check, the corresponding error is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a regular request and directly deletes the route map rule after the request passes the check.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the route map rule to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmr-0a1b2c3d4e5f6a7b8
	RouteMapRuleId *string `json:"RouteMapRuleId,omitempty" xml:"RouteMapRuleId,omitempty"`
}

func (s DeleteRouteMapRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRouteMapRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRouteMapRuleRequest) GetRouteMapRuleId() *string {
	return s.RouteMapRuleId
}

func (s *DeleteRouteMapRuleRequest) SetClientToken(v string) *DeleteRouteMapRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRouteMapRuleRequest) SetDryRun(v bool) *DeleteRouteMapRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRouteMapRuleRequest) SetRouteMapRuleId(v string) *DeleteRouteMapRuleRequest {
	s.RouteMapRuleId = &v
	return s
}

func (s *DeleteRouteMapRuleRequest) Validate() error {
	return dara.Validate(s)
}
