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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	Tag []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateServiceInstanceRequestTag struct {
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
