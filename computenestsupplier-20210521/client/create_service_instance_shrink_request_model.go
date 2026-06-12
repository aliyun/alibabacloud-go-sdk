// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceInstanceShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateServiceInstanceShrinkRequest
	GetDryRun() *bool
	SetEndTime(v string) *CreateServiceInstanceShrinkRequest
	GetEndTime() *string
	SetName(v string) *CreateServiceInstanceShrinkRequest
	GetName() *string
	SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *CreateServiceInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateServiceInstanceShrinkRequest
	GetServiceId() *string
	SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest
	GetServiceVersion() *string
	SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest
	GetSpecificationName() *string
	SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest
	GetTag() []*CreateServiceInstanceShrinkRequestTag
	SetTemplateName(v string) *CreateServiceInstanceShrinkRequest
	GetTemplateName() *string
	SetUserId(v string) *CreateServiceInstanceShrinkRequest
	GetUserId() *string
}

type CreateServiceInstanceShrinkRequest struct {
	// A client token to ensure the idempotence of the request. Generate a unique value for this parameter from your client. The token can be up to 64 characters in length and can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. A dry run checks for permissions and instance status. Valid values:
	//
	// - true: The system checks the request but does not create the service instance.
	//
	// - false: The system sends the request. If the request passes the check, the service instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The time when the service instance is released.
	//
	// > Only service providers can set this parameter for their own service instances in managed scenarios.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-12T03:13:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the service instance. The name must meet the following requirements:
	//
	// - It can be up to 64 characters in length.
	//
	// - It must start with a letter or a digit and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters used to deploy the service instance.
	//
	// > If the service instance includes information about the deployment region, specify that information in the deployment parameters.
	//
	// example:
	//
	// {
	//
	//       "RegionId": "cn-hangzhou"
	//
	//       "NodeCount": 3,
	//
	//       "SystemDiskSize": 40,
	//
	//       "InstancePassword": "******"
	//
	// }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the region. Valid values:
	//
	// - cn-hangzhou: China (Hangzhou)
	//
	// - ap-southeast-1: Singapore
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The custom tags.
	Tag []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template. Specify this parameter if the service supports multiple templates.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceInstanceShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceInstanceShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateServiceInstanceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceInstanceShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateServiceInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceInstanceShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceInstanceShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *CreateServiceInstanceShrinkRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *CreateServiceInstanceShrinkRequest) GetTag() []*CreateServiceInstanceShrinkRequestTag {
	return s.Tag
}

func (s *CreateServiceInstanceShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateServiceInstanceShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetDryRun(v bool) *CreateServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEndTime(v string) *CreateServiceInstanceShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetName(v string) *CreateServiceInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetUserId(v string) *CreateServiceInstanceShrinkRequest {
	s.UserId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServiceInstanceShrinkRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceInstanceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
