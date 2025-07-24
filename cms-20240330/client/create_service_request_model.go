// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v string) *CreateServiceRequest
	GetAttributes() *string
	SetDescription(v string) *CreateServiceRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateServiceRequest
	GetDisplayName() *string
	SetPid(v string) *CreateServiceRequest
	GetPid() *string
	SetServiceName(v string) *CreateServiceRequest
	GetServiceName() *string
	SetServiceStatus(v string) *CreateServiceRequest
	GetServiceStatus() *string
	SetServiceType(v string) *CreateServiceRequest
	GetServiceType() *string
}

type CreateServiceRequest struct {
	// example:
	//
	// {"language":"java"}
	Attributes  *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// bx3udsi5ie@ed2ba6beebdb6de
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mag_test
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// example:
	//
	// Created
	ServiceStatus *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TRACE
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) GetAttributes() *string {
	return s.Attributes
}

func (s *CreateServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateServiceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateServiceRequest) GetPid() *string {
	return s.Pid
}

func (s *CreateServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceRequest) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *CreateServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateServiceRequest) SetAttributes(v string) *CreateServiceRequest {
	s.Attributes = &v
	return s
}

func (s *CreateServiceRequest) SetDescription(v string) *CreateServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateServiceRequest) SetDisplayName(v string) *CreateServiceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateServiceRequest) SetPid(v string) *CreateServiceRequest {
	s.Pid = &v
	return s
}

func (s *CreateServiceRequest) SetServiceName(v string) *CreateServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceRequest) SetServiceStatus(v string) *CreateServiceRequest {
	s.ServiceStatus = &v
	return s
}

func (s *CreateServiceRequest) SetServiceType(v string) *CreateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateServiceRequest) Validate() error {
	return dara.Validate(s)
}
