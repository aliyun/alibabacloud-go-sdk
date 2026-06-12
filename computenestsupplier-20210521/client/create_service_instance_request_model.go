// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceInstanceRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateServiceInstanceRequest
	GetDryRun() *bool
	SetEndTime(v string) *CreateServiceInstanceRequest
	GetEndTime() *string
	SetName(v string) *CreateServiceInstanceRequest
	GetName() *string
	SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *CreateServiceInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateServiceInstanceRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateServiceInstanceRequest
	GetServiceId() *string
	SetServiceVersion(v string) *CreateServiceInstanceRequest
	GetServiceVersion() *string
	SetSpecificationName(v string) *CreateServiceInstanceRequest
	GetSpecificationName() *string
	SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest
	GetTag() []*CreateServiceInstanceRequestTag
	SetTemplateName(v string) *CreateServiceInstanceRequest
	GetTemplateName() *string
	SetUserId(v string) *CreateServiceInstanceRequest
	GetUserId() *string
}

type CreateServiceInstanceRequest struct {
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	Tag []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreateServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceInstanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateServiceInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceInstanceRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceInstanceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceInstanceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *CreateServiceInstanceRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *CreateServiceInstanceRequest) GetTag() []*CreateServiceInstanceRequestTag {
	return s.Tag
}

func (s *CreateServiceInstanceRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateServiceInstanceRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetDryRun(v bool) *CreateServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEndTime(v string) *CreateServiceInstanceRequest {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetName(v string) *CreateServiceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceVersion(v string) *CreateServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationName(v string) *CreateServiceInstanceRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetUserId(v string) *CreateServiceInstanceRequest {
	s.UserId = &v
	return s
}

func (s *CreateServiceInstanceRequest) Validate() error {
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

type CreateServiceInstanceRequestTag struct {
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

func (s CreateServiceInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
