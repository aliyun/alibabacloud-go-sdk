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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The time when the service instance was released.
	//
	// >  This parameter is available only for the service instances that are managed by service providers.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-12T03:13:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the service instance. The value must meet the following requirements:
	//
	// 	- The name cannot exceed 64 characters in length.
	//
	// 	- It can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters that are specified for service instance deployment.
	//
	// >  If you want to specify the region in which the service instance is deployed, you must specify the information in Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// 	- ap-southeast-1: Singapore
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
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The custom tags.
	Tag []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The template name. You must specify a template name if the service supports multiple templates.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The user ID.
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
	return dara.Validate(s)
}

type CreateServiceInstanceShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
