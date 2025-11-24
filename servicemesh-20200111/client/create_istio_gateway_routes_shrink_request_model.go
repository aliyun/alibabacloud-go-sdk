// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayRoutesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateIstioGatewayRoutesShrinkRequest
	GetDescription() *string
	SetGatewayRouteShrink(v string) *CreateIstioGatewayRoutesShrinkRequest
	GetGatewayRouteShrink() *string
	SetIstioGatewayName(v string) *CreateIstioGatewayRoutesShrinkRequest
	GetIstioGatewayName() *string
	SetPriority(v int32) *CreateIstioGatewayRoutesShrinkRequest
	GetPriority() *int32
	SetServiceMeshId(v string) *CreateIstioGatewayRoutesShrinkRequest
	GetServiceMeshId() *string
	SetStatus(v int32) *CreateIstioGatewayRoutesShrinkRequest
	GetStatus() *int32
}

type CreateIstioGatewayRoutesShrinkRequest struct {
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the routing rule to be created for the ASM gateway.
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

func (s CreateIstioGatewayRoutesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIstioGatewayRoutesShrinkRequest) GetGatewayRouteShrink() *string {
	return s.GatewayRouteShrink
}

func (s *CreateIstioGatewayRoutesShrinkRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *CreateIstioGatewayRoutesShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateIstioGatewayRoutesShrinkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateIstioGatewayRoutesShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetDescription(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetGatewayRouteShrink(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.GatewayRouteShrink = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetIstioGatewayName(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetPriority(v int32) *CreateIstioGatewayRoutesShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetServiceMeshId(v string) *CreateIstioGatewayRoutesShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) SetStatus(v int32) *CreateIstioGatewayRoutesShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateIstioGatewayRoutesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
