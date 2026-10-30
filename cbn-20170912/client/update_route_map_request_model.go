// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRouteMapRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateRouteMapRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateRouteMapRequest
	GetDryRun() *bool
	SetName(v string) *UpdateRouteMapRequest
	GetName() *string
	SetPriority(v int32) *UpdateRouteMapRequest
	GetPriority() *int32
	SetRouteMapId(v string) *UpdateRouteMapRequest
	GetRouteMapId() *string
}

type UpdateRouteMapRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the route map.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without updating the route map. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the route map is updated.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the route map.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Policy priority of the route map. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
	//
	// > Policy priority of each route map within an Alibaba Cloud account must be unique. When a route map is applied, the system starts to match routes against the rules of the route map that has the highest priority (the smallest priority value) and then proceeds to route maps with lower priorities based on the rule priorities. Therefore, when you specify policy priority of a route map, make sure that policy priority meets your expected matching order.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the route map.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-abcdedfghij****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s UpdateRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRouteMapRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRouteMapRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateRouteMapRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRouteMapRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRouteMapRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *UpdateRouteMapRequest) SetClientToken(v string) *UpdateRouteMapRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRouteMapRequest) SetDescription(v string) *UpdateRouteMapRequest {
	s.Description = &v
	return s
}

func (s *UpdateRouteMapRequest) SetDryRun(v bool) *UpdateRouteMapRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRouteMapRequest) SetName(v string) *UpdateRouteMapRequest {
	s.Name = &v
	return s
}

func (s *UpdateRouteMapRequest) SetPriority(v int32) *UpdateRouteMapRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRouteMapRequest) SetRouteMapId(v string) *UpdateRouteMapRequest {
	s.RouteMapId = &v
	return s
}

func (s *UpdateRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
