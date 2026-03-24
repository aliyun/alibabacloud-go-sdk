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
	SetResourceGroupId(v string) *CreateServiceRequest
	GetResourceGroupId() *string
	SetServiceName(v string) *CreateServiceRequest
	GetServiceName() *string
	SetServiceStatus(v string) *CreateServiceRequest
	GetServiceStatus() *string
	SetServiceType(v string) *CreateServiceRequest
	GetServiceType() *string
	SetTags(v []*CreateServiceRequestTags) *CreateServiceRequest
	GetTags() []*CreateServiceRequestTags
}

type CreateServiceRequest struct {
	// The extended properties.
	//
	// example:
	//
	// {"language":"java"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The service description. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// mag测试应用
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// mag测试应用
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The application ID. You do not typically need to specify this parameter.
	//
	// example:
	//
	// bx3udsi5ie@ed2ba6beebdb6de
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxxxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// mag_test
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The service status. Do not specify this parameter when you create a service.
	//
	// example:
	//
	// Created
	ServiceStatus *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	// The service type.
	//
	// This parameter is required.
	//
	// example:
	//
	// TRACE
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// An array of tags.
	Tags []*CreateServiceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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

func (s *CreateServiceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *CreateServiceRequest) GetTags() []*CreateServiceRequestTags {
	return s.Tags
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

func (s *CreateServiceRequest) SetResourceGroupId(v string) *CreateServiceRequest {
	s.ResourceGroupId = &v
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

func (s *CreateServiceRequest) SetTags(v []*CreateServiceRequestTags) *CreateServiceRequest {
	s.Tags = v
	return s
}

func (s *CreateServiceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServiceRequestTags struct {
	// The `key` of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The `value` of the tag.
	//
	// example:
	//
	// prod
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateServiceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateServiceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateServiceRequestTags) SetKey(v string) *CreateServiceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateServiceRequestTags) SetValue(v string) *CreateServiceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateServiceRequestTags) Validate() error {
	return dara.Validate(s)
}
