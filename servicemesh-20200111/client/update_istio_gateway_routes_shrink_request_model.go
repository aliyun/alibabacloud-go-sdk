// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioGatewayRoutesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateIstioGatewayRoutesShrinkRequest
	GetDescription() *string
	SetGatewayRouteShrink(v string) *UpdateIstioGatewayRoutesShrinkRequest
	GetGatewayRouteShrink() *string
	SetIstioGatewayName(v string) *UpdateIstioGatewayRoutesShrinkRequest
	GetIstioGatewayName() *string
	SetPriority(v int32) *UpdateIstioGatewayRoutesShrinkRequest
	GetPriority() *int32
	SetServiceMeshId(v string) *UpdateIstioGatewayRoutesShrinkRequest
	GetServiceMeshId() *string
	SetStatus(v int32) *UpdateIstioGatewayRoutesShrinkRequest
	GetStatus() *int32
}

type UpdateIstioGatewayRoutesShrinkRequest struct {
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be updated for the ASM gateway.
	GatewayRouteShrink *string `json:"GatewayRoute,omitempty" xml:"GatewayRoute,omitempty"`
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The priority of the routing rule. The value of this parameter is an integer. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values:
	//
	// 	- `0`: The routing rule is valid.
	//
	// 	- `1`: The routing rule is invalid.
	//
	// 	- `2`: An error occurs during the creation or update of the routing rule.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIstioGatewayRoutesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) GetGatewayRouteShrink() *string {
	return s.GatewayRouteShrink
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetDescription(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetGatewayRouteShrink(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.GatewayRouteShrink = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetIstioGatewayName(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetPriority(v int32) *UpdateIstioGatewayRoutesShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetServiceMeshId(v string) *UpdateIstioGatewayRoutesShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) SetStatus(v int32) *UpdateIstioGatewayRoutesShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateIstioGatewayRoutesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
