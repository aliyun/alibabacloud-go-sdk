// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPilotEipResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipId(v string) *ModifyPilotEipResourceRequest
	GetEipId() *string
	SetIsCanary(v bool) *ModifyPilotEipResourceRequest
	GetIsCanary() *bool
	SetOperation(v string) *ModifyPilotEipResourceRequest
	GetOperation() *string
	SetServiceMeshId(v string) *ModifyPilotEipResourceRequest
	GetServiceMeshId() *string
}

type ModifyPilotEipResourceRequest struct {
	// example:
	//
	// eip-hp36jpqq5eged********
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// The type of the Istio Pilot with which you want to associate the EIPs. Valid values:
	//
	// 	- `true`: Bind an EIP to the Istio Pilot during canary release (only valid during the canary release).
	//
	// 	- `false`: Bind an EIP to the Istio Pilot in stable state.
	//
	// example:
	//
	// false
	IsCanary *bool `json:"IsCanary,omitempty" xml:"IsCanary,omitempty"`
	// The type of the operation that you want to perform. Valid values:
	//
	// 	- `UnBindEip`: unbinds an elastic IP address (EIP) from the Istio Pilot.
	//
	// 	- `BindEip`: binds an EIP to the Istio Pilot.
	//
	// example:
	//
	// BindEip
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// c1f5a67154bec40629c2698ec********
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ModifyPilotEipResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPilotEipResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyPilotEipResourceRequest) GetEipId() *string {
	return s.EipId
}

func (s *ModifyPilotEipResourceRequest) GetIsCanary() *bool {
	return s.IsCanary
}

func (s *ModifyPilotEipResourceRequest) GetOperation() *string {
	return s.Operation
}

func (s *ModifyPilotEipResourceRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *ModifyPilotEipResourceRequest) SetEipId(v string) *ModifyPilotEipResourceRequest {
	s.EipId = &v
	return s
}

func (s *ModifyPilotEipResourceRequest) SetIsCanary(v bool) *ModifyPilotEipResourceRequest {
	s.IsCanary = &v
	return s
}

func (s *ModifyPilotEipResourceRequest) SetOperation(v string) *ModifyPilotEipResourceRequest {
	s.Operation = &v
	return s
}

func (s *ModifyPilotEipResourceRequest) SetServiceMeshId(v string) *ModifyPilotEipResourceRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *ModifyPilotEipResourceRequest) Validate() error {
	return dara.Validate(s)
}
