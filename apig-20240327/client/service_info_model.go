// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetExpressType(v string) *ServiceInfo
	GetExpressType() *string
	SetGroupName(v string) *ServiceInfo
	GetGroupName() *string
	SetName(v string) *ServiceInfo
	GetName() *string
	SetNamespace(v string) *ServiceInfo
	GetNamespace() *string
	SetPaiWorkspaceId(v string) *ServiceInfo
	GetPaiWorkspaceId() *string
	SetPaiWorkspaceName(v string) *ServiceInfo
	GetPaiWorkspaceName() *string
	SetPorts(v []*ServiceInfoPorts) *ServiceInfo
	GetPorts() []*ServiceInfoPorts
	SetQualifier(v string) *ServiceInfo
	GetQualifier() *string
	SetServiceId(v string) *ServiceInfo
	GetServiceId() *string
	SetSourceType(v string) *ServiceInfo
	GetSourceType() *string
	SetStatus(v string) *ServiceInfo
	GetStatus() *string
	SetVersions(v []*ServiceInfoVersions) *ServiceInfo
	GetVersions() []*ServiceInfoVersions
}

type ServiceInfo struct {
	ExpressType      *string                `json:"expressType,omitempty" xml:"expressType,omitempty"`
	GroupName        *string                `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Name             *string                `json:"name,omitempty" xml:"name,omitempty"`
	Namespace        *string                `json:"namespace,omitempty" xml:"namespace,omitempty"`
	PaiWorkspaceId   *string                `json:"paiWorkspaceId,omitempty" xml:"paiWorkspaceId,omitempty"`
	PaiWorkspaceName *string                `json:"paiWorkspaceName,omitempty" xml:"paiWorkspaceName,omitempty"`
	Ports            []*ServiceInfoPorts    `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
	Qualifier        *string                `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	ServiceId        *string                `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	SourceType       *string                `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Status           *string                `json:"status,omitempty" xml:"status,omitempty"`
	Versions         []*ServiceInfoVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ServiceInfo) String() string {
	return dara.Prettify(s)
}

func (s ServiceInfo) GoString() string {
	return s.String()
}

func (s *ServiceInfo) GetExpressType() *string {
	return s.ExpressType
}

func (s *ServiceInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *ServiceInfo) GetName() *string {
	return s.Name
}

func (s *ServiceInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *ServiceInfo) GetPaiWorkspaceId() *string {
	return s.PaiWorkspaceId
}

func (s *ServiceInfo) GetPaiWorkspaceName() *string {
	return s.PaiWorkspaceName
}

func (s *ServiceInfo) GetPorts() []*ServiceInfoPorts {
	return s.Ports
}

func (s *ServiceInfo) GetQualifier() *string {
	return s.Qualifier
}

func (s *ServiceInfo) GetServiceId() *string {
	return s.ServiceId
}

func (s *ServiceInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *ServiceInfo) GetStatus() *string {
	return s.Status
}

func (s *ServiceInfo) GetVersions() []*ServiceInfoVersions {
	return s.Versions
}

func (s *ServiceInfo) SetExpressType(v string) *ServiceInfo {
	s.ExpressType = &v
	return s
}

func (s *ServiceInfo) SetGroupName(v string) *ServiceInfo {
	s.GroupName = &v
	return s
}

func (s *ServiceInfo) SetName(v string) *ServiceInfo {
	s.Name = &v
	return s
}

func (s *ServiceInfo) SetNamespace(v string) *ServiceInfo {
	s.Namespace = &v
	return s
}

func (s *ServiceInfo) SetPaiWorkspaceId(v string) *ServiceInfo {
	s.PaiWorkspaceId = &v
	return s
}

func (s *ServiceInfo) SetPaiWorkspaceName(v string) *ServiceInfo {
	s.PaiWorkspaceName = &v
	return s
}

func (s *ServiceInfo) SetPorts(v []*ServiceInfoPorts) *ServiceInfo {
	s.Ports = v
	return s
}

func (s *ServiceInfo) SetQualifier(v string) *ServiceInfo {
	s.Qualifier = &v
	return s
}

func (s *ServiceInfo) SetServiceId(v string) *ServiceInfo {
	s.ServiceId = &v
	return s
}

func (s *ServiceInfo) SetSourceType(v string) *ServiceInfo {
	s.SourceType = &v
	return s
}

func (s *ServiceInfo) SetStatus(v string) *ServiceInfo {
	s.Status = &v
	return s
}

func (s *ServiceInfo) SetVersions(v []*ServiceInfoVersions) *ServiceInfo {
	s.Versions = v
	return s
}

func (s *ServiceInfo) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ServiceInfoPorts struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Port     *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ServiceInfoPorts) String() string {
	return dara.Prettify(s)
}

func (s ServiceInfoPorts) GoString() string {
	return s.String()
}

func (s *ServiceInfoPorts) GetName() *string {
	return s.Name
}

func (s *ServiceInfoPorts) GetPort() *int32 {
	return s.Port
}

func (s *ServiceInfoPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ServiceInfoPorts) SetName(v string) *ServiceInfoPorts {
	s.Name = &v
	return s
}

func (s *ServiceInfoPorts) SetPort(v int32) *ServiceInfoPorts {
	s.Port = &v
	return s
}

func (s *ServiceInfoPorts) SetProtocol(v string) *ServiceInfoPorts {
	s.Protocol = &v
	return s
}

func (s *ServiceInfoPorts) Validate() error {
	return dara.Validate(s)
}

type ServiceInfoVersions struct {
	Labels []*ServiceInfoVersionsLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	Name   *string                      `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ServiceInfoVersions) String() string {
	return dara.Prettify(s)
}

func (s ServiceInfoVersions) GoString() string {
	return s.String()
}

func (s *ServiceInfoVersions) GetLabels() []*ServiceInfoVersionsLabels {
	return s.Labels
}

func (s *ServiceInfoVersions) GetName() *string {
	return s.Name
}

func (s *ServiceInfoVersions) SetLabels(v []*ServiceInfoVersionsLabels) *ServiceInfoVersions {
	s.Labels = v
	return s
}

func (s *ServiceInfoVersions) SetName(v string) *ServiceInfoVersions {
	s.Name = &v
	return s
}

func (s *ServiceInfoVersions) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ServiceInfoVersionsLabels struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ServiceInfoVersionsLabels) String() string {
	return dara.Prettify(s)
}

func (s ServiceInfoVersionsLabels) GoString() string {
	return s.String()
}

func (s *ServiceInfoVersionsLabels) GetKey() *string {
	return s.Key
}

func (s *ServiceInfoVersionsLabels) GetValue() *string {
	return s.Value
}

func (s *ServiceInfoVersionsLabels) SetKey(v string) *ServiceInfoVersionsLabels {
	s.Key = &v
	return s
}

func (s *ServiceInfoVersionsLabels) SetValue(v string) *ServiceInfoVersionsLabels {
	s.Value = &v
	return s
}

func (s *ServiceInfoVersionsLabels) Validate() error {
	return dara.Validate(s)
}
