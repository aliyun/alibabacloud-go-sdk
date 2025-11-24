// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioRouteAdditionalStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateIstioRouteAdditionalStatusRequest
	GetDescription() *string
	SetIstioGatewayName(v string) *UpdateIstioRouteAdditionalStatusRequest
	GetIstioGatewayName() *string
	SetPriority(v int32) *UpdateIstioRouteAdditionalStatusRequest
	GetPriority() *int32
	SetRouteName(v string) *UpdateIstioRouteAdditionalStatusRequest
	GetRouteName() *string
	SetServiceMeshId(v string) *UpdateIstioRouteAdditionalStatusRequest
	GetServiceMeshId() *string
	SetStatus(v int32) *UpdateIstioRouteAdditionalStatusRequest
	GetStatus() *int32
}

type UpdateIstioRouteAdditionalStatusRequest struct {
	// The description of the routing rule.
	//
	// example:
	//
	// demo route
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// http-test
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The status of the routing rule. Valid values: 0: The routing rule is valid. 1: The routing rule is invalid. 2: An error occurs during the creation or update of the routing rule.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateIstioRouteAdditionalStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioRouteAdditionalStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateIstioRouteAdditionalStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIstioRouteAdditionalStatusRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *UpdateIstioRouteAdditionalStatusRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateIstioRouteAdditionalStatusRequest) GetRouteName() *string {
	return s.RouteName
}

func (s *UpdateIstioRouteAdditionalStatusRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateIstioRouteAdditionalStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetDescription(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.Description = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetIstioGatewayName(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetPriority(v int32) *UpdateIstioRouteAdditionalStatusRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetRouteName(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.RouteName = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetServiceMeshId(v string) *UpdateIstioRouteAdditionalStatusRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) SetStatus(v int32) *UpdateIstioRouteAdditionalStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusRequest) Validate() error {
	return dara.Validate(s)
}
