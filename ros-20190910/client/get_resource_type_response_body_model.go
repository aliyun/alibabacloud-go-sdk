// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]interface{}) *GetResourceTypeResponseBody
	GetAttributes() map[string]interface{}
	SetCreateTime(v string) *GetResourceTypeResponseBody
	GetCreateTime() *string
	SetDefaultVersionId(v string) *GetResourceTypeResponseBody
	GetDefaultVersionId() *string
	SetDescription(v string) *GetResourceTypeResponseBody
	GetDescription() *string
	SetEntityType(v string) *GetResourceTypeResponseBody
	GetEntityType() *string
	SetIsDefaultVersion(v bool) *GetResourceTypeResponseBody
	GetIsDefaultVersion() *bool
	SetLatestVersionId(v string) *GetResourceTypeResponseBody
	GetLatestVersionId() *string
	SetProperties(v map[string]interface{}) *GetResourceTypeResponseBody
	GetProperties() map[string]interface{}
	SetProvider(v string) *GetResourceTypeResponseBody
	GetProvider() *string
	SetRequestId(v string) *GetResourceTypeResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetResourceTypeResponseBody
	GetResourceType() *string
	SetSupportDriftDetection(v bool) *GetResourceTypeResponseBody
	GetSupportDriftDetection() *bool
	SetSupportScratchDetection(v bool) *GetResourceTypeResponseBody
	GetSupportScratchDetection() *bool
	SetTemplateBody(v string) *GetResourceTypeResponseBody
	GetTemplateBody() *string
	SetTotalVersionCount(v int32) *GetResourceTypeResponseBody
	GetTotalVersionCount() *int32
	SetUpdateTime(v string) *GetResourceTypeResponseBody
	GetUpdateTime() *string
}

type GetResourceTypeResponseBody struct {
	// The type of the resource.
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The creation time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-24T08:25:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default version ID.
	//
	// > This parameter is returned only if the resource type is queried.
	//
	// example:
	//
	// v1
	DefaultVersionId *string `json:"DefaultVersionId,omitempty" xml:"DefaultVersionId,omitempty"`
	// The description of the resource type.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Valid values:
	//
	// 	- Resource: regular resource. For more information, see [Resources](https://help.aliyun.com/document_detail/28863.html).
	//
	// 	- DataSource: DataSource resource. For more information, see [DataSource resources](https://help.aliyun.com/document_detail/404753.html).
	//
	// 	- module: module.
	//
	// example:
	//
	// Resource
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// > This parameter is returned only if a specific version of the resource type is queried.
	//
	// example:
	//
	// true
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The latest version ID.
	//
	// > This parameter is returned only if the resource type is queried.
	//
	// example:
	//
	// v10
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// Indicates whether the resource supports drift detection. Default value: false. Valid values:
	//
	// 	- true: Drift detection is supported.
	//
	// 	- false: Drift detection is not supported.
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// 	- ROS: The resource type is provided by Resource Orchestration Service (ROS).
	//
	// 	- Self: The resource type is provided by you.
	//
	// example:
	//
	// ROS
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The attributes of the resource.
	//
	// example:
	//
	// A28FBA2E-B6B3-5822-AA45-AB875EF23641
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The properties of the resource.
	//
	// example:
	//
	// ALIYUN::ROS::WaitConditionHandle
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the resource supports scratch detection. Default value: false. Valid values:
	//
	// 	- true: Scratch detection is supported.
	//
	// 	- false: Scratch detection is not supported.
	//
	// example:
	//
	// false
	SupportDriftDetection *bool `json:"SupportDriftDetection,omitempty" xml:"SupportDriftDetection,omitempty"`
	// The entity type. Valid values:
	//
	// 	- Resource: resources other than DataSource resources. For more information, see [Resources](https://help.aliyun.com/document_detail/28863.html).
	//
	// 	- DataSource: DataSource resources.
	//
	// example:
	//
	// false
	SupportScratchDetection *bool `json:"SupportScratchDetection,omitempty" xml:"SupportScratchDetection,omitempty"`
	// The template content in the module.
	//
	// > This parameter is returned only if a specific version of the resource type is queried.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion":"2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The total number of versions.
	//
	// > This parameter is returned only if the resource type is queried.
	//
	// example:
	//
	// 10
	TotalVersionCount *int32 `json:"TotalVersionCount,omitempty" xml:"TotalVersionCount,omitempty"`
	// The update time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-24T08:25:21
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBody) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *GetResourceTypeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetResourceTypeResponseBody) GetDefaultVersionId() *string {
	return s.DefaultVersionId
}

func (s *GetResourceTypeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetResourceTypeResponseBody) GetEntityType() *string {
	return s.EntityType
}

func (s *GetResourceTypeResponseBody) GetIsDefaultVersion() *bool {
	return s.IsDefaultVersion
}

func (s *GetResourceTypeResponseBody) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *GetResourceTypeResponseBody) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetResourceTypeResponseBody) GetProvider() *string {
	return s.Provider
}

func (s *GetResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceTypeResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceTypeResponseBody) GetSupportDriftDetection() *bool {
	return s.SupportDriftDetection
}

func (s *GetResourceTypeResponseBody) GetSupportScratchDetection() *bool {
	return s.SupportScratchDetection
}

func (s *GetResourceTypeResponseBody) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetResourceTypeResponseBody) GetTotalVersionCount() *int32 {
	return s.TotalVersionCount
}

func (s *GetResourceTypeResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetResourceTypeResponseBody) SetAttributes(v map[string]interface{}) *GetResourceTypeResponseBody {
	s.Attributes = v
	return s
}

func (s *GetResourceTypeResponseBody) SetCreateTime(v string) *GetResourceTypeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetDefaultVersionId(v string) *GetResourceTypeResponseBody {
	s.DefaultVersionId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetDescription(v string) *GetResourceTypeResponseBody {
	s.Description = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetEntityType(v string) *GetResourceTypeResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetIsDefaultVersion(v bool) *GetResourceTypeResponseBody {
	s.IsDefaultVersion = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetLatestVersionId(v string) *GetResourceTypeResponseBody {
	s.LatestVersionId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetProperties(v map[string]interface{}) *GetResourceTypeResponseBody {
	s.Properties = v
	return s
}

func (s *GetResourceTypeResponseBody) SetProvider(v string) *GetResourceTypeResponseBody {
	s.Provider = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetRequestId(v string) *GetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetResourceType(v string) *GetResourceTypeResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetSupportDriftDetection(v bool) *GetResourceTypeResponseBody {
	s.SupportDriftDetection = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetSupportScratchDetection(v bool) *GetResourceTypeResponseBody {
	s.SupportScratchDetection = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetTemplateBody(v string) *GetResourceTypeResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetTotalVersionCount(v int32) *GetResourceTypeResponseBody {
	s.TotalVersionCount = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetUpdateTime(v string) *GetResourceTypeResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetResourceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
