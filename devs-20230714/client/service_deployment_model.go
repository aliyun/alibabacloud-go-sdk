// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceDeployment interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *ServiceDeployment
	GetCreatedTime() *string
	SetDescription(v string) *ServiceDeployment
	GetDescription() *string
	SetEnvironmentDeploymentName(v string) *ServiceDeployment
	GetEnvironmentDeploymentName() *string
	SetKind(v string) *ServiceDeployment
	GetKind() *string
	SetLabels(v map[string]*string) *ServiceDeployment
	GetLabels() map[string]*string
	SetName(v string) *ServiceDeployment
	GetName() *string
	SetStatus(v *ServiceDeploymentStatus) *ServiceDeployment
	GetStatus() *ServiceDeploymentStatus
	SetUid(v string) *ServiceDeployment
	GetUid() *string
}

type ServiceDeployment struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// commit by xxx.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// my-environment-deployment
	EnvironmentDeploymentName *string `json:"environmentDeploymentName,omitempty" xml:"environmentDeploymentName,omitempty"`
	// example:
	//
	// Deployment
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-deployment
	Name   *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Status *ServiceDeploymentStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ServiceDeployment) String() string {
	return dara.Prettify(s)
}

func (s ServiceDeployment) GoString() string {
	return s.String()
}

func (s *ServiceDeployment) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ServiceDeployment) GetDescription() *string {
	return s.Description
}

func (s *ServiceDeployment) GetEnvironmentDeploymentName() *string {
	return s.EnvironmentDeploymentName
}

func (s *ServiceDeployment) GetKind() *string {
	return s.Kind
}

func (s *ServiceDeployment) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ServiceDeployment) GetName() *string {
	return s.Name
}

func (s *ServiceDeployment) GetStatus() *ServiceDeploymentStatus {
	return s.Status
}

func (s *ServiceDeployment) GetUid() *string {
	return s.Uid
}

func (s *ServiceDeployment) SetCreatedTime(v string) *ServiceDeployment {
	s.CreatedTime = &v
	return s
}

func (s *ServiceDeployment) SetDescription(v string) *ServiceDeployment {
	s.Description = &v
	return s
}

func (s *ServiceDeployment) SetEnvironmentDeploymentName(v string) *ServiceDeployment {
	s.EnvironmentDeploymentName = &v
	return s
}

func (s *ServiceDeployment) SetKind(v string) *ServiceDeployment {
	s.Kind = &v
	return s
}

func (s *ServiceDeployment) SetLabels(v map[string]*string) *ServiceDeployment {
	s.Labels = v
	return s
}

func (s *ServiceDeployment) SetName(v string) *ServiceDeployment {
	s.Name = &v
	return s
}

func (s *ServiceDeployment) SetStatus(v *ServiceDeploymentStatus) *ServiceDeployment {
	s.Status = v
	return s
}

func (s *ServiceDeployment) SetUid(v string) *ServiceDeployment {
	s.Uid = &v
	return s
}

func (s *ServiceDeployment) Validate() error {
	return dara.Validate(s)
}
