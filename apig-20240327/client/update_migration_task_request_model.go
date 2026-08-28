// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterNamespace(v string) *UpdateMigrationTaskRequest
	GetClusterNamespace() *string
	SetDescription(v string) *UpdateMigrationTaskRequest
	GetDescription() *string
	SetServiceName(v string) *UpdateMigrationTaskRequest
	GetServiceName() *string
	SetSlbId(v string) *UpdateMigrationTaskRequest
	GetSlbId() *string
	SetSwitchType(v string) *UpdateMigrationTaskRequest
	GetSwitchType() *string
	SetTarget(v string) *UpdateMigrationTaskRequest
	GetTarget() *string
	SetVirtualServices(v []*UpdateMigrationTaskRequestVirtualServices) *UpdateMigrationTaskRequest
	GetVirtualServices() []*UpdateMigrationTaskRequestVirtualServices
	SetWeight(v int32) *UpdateMigrationTaskRequest
	GetWeight() *int32
}

type UpdateMigrationTaskRequest struct {
	// example:
	//
	// default
	ClusterNamespace *string `json:"clusterNamespace,omitempty" xml:"clusterNamespace,omitempty"`
	// example:
	//
	// 迁移测试
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// nginx-ingress-lb
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// example:
	//
	// lb-bp1xxxx
	SlbId *string `json:"slbId,omitempty" xml:"slbId,omitempty"`
	// example:
	//
	// SLB
	SwitchType *string `json:"switchType,omitempty" xml:"switchType,omitempty"`
	// example:
	//
	// Task
	Target          *string                                      `json:"target,omitempty" xml:"target,omitempty"`
	VirtualServices []*UpdateMigrationTaskRequestVirtualServices `json:"virtualServices,omitempty" xml:"virtualServices,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s UpdateMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskRequest) GetClusterNamespace() *string {
	return s.ClusterNamespace
}

func (s *UpdateMigrationTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMigrationTaskRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateMigrationTaskRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *UpdateMigrationTaskRequest) GetSwitchType() *string {
	return s.SwitchType
}

func (s *UpdateMigrationTaskRequest) GetTarget() *string {
	return s.Target
}

func (s *UpdateMigrationTaskRequest) GetVirtualServices() []*UpdateMigrationTaskRequestVirtualServices {
	return s.VirtualServices
}

func (s *UpdateMigrationTaskRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateMigrationTaskRequest) SetClusterNamespace(v string) *UpdateMigrationTaskRequest {
	s.ClusterNamespace = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetDescription(v string) *UpdateMigrationTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetServiceName(v string) *UpdateMigrationTaskRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetSlbId(v string) *UpdateMigrationTaskRequest {
	s.SlbId = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetSwitchType(v string) *UpdateMigrationTaskRequest {
	s.SwitchType = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetTarget(v string) *UpdateMigrationTaskRequest {
	s.Target = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetVirtualServices(v []*UpdateMigrationTaskRequestVirtualServices) *UpdateMigrationTaskRequest {
	s.VirtualServices = v
	return s
}

func (s *UpdateMigrationTaskRequest) SetWeight(v int32) *UpdateMigrationTaskRequest {
	s.Weight = &v
	return s
}

func (s *UpdateMigrationTaskRequest) Validate() error {
	if s.VirtualServices != nil {
		for _, item := range s.VirtualServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMigrationTaskRequestVirtualServices struct {
	// example:
	//
	// 80
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// rsp-xxxx
	VirtualServiceGroupId *string `json:"virtualServiceGroupId,omitempty" xml:"virtualServiceGroupId,omitempty"`
	// example:
	//
	// 80-tcp
	VirtualServiceGroupName *string `json:"virtualServiceGroupName,omitempty" xml:"virtualServiceGroupName,omitempty"`
}

func (s UpdateMigrationTaskRequestVirtualServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskRequestVirtualServices) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskRequestVirtualServices) GetPort() *int32 {
	return s.Port
}

func (s *UpdateMigrationTaskRequestVirtualServices) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateMigrationTaskRequestVirtualServices) GetVirtualServiceGroupId() *string {
	return s.VirtualServiceGroupId
}

func (s *UpdateMigrationTaskRequestVirtualServices) GetVirtualServiceGroupName() *string {
	return s.VirtualServiceGroupName
}

func (s *UpdateMigrationTaskRequestVirtualServices) SetPort(v int32) *UpdateMigrationTaskRequestVirtualServices {
	s.Port = &v
	return s
}

func (s *UpdateMigrationTaskRequestVirtualServices) SetProtocol(v string) *UpdateMigrationTaskRequestVirtualServices {
	s.Protocol = &v
	return s
}

func (s *UpdateMigrationTaskRequestVirtualServices) SetVirtualServiceGroupId(v string) *UpdateMigrationTaskRequestVirtualServices {
	s.VirtualServiceGroupId = &v
	return s
}

func (s *UpdateMigrationTaskRequestVirtualServices) SetVirtualServiceGroupName(v string) *UpdateMigrationTaskRequestVirtualServices {
	s.VirtualServiceGroupName = &v
	return s
}

func (s *UpdateMigrationTaskRequestVirtualServices) Validate() error {
	return dara.Validate(s)
}
