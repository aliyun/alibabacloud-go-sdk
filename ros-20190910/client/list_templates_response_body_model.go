// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody
	GetTemplates() []*ListTemplatesResponseBodyTemplates
	SetTotalCount(v int32) *ListTemplatesResponseBody
	GetTotalCount() *int32
}

type ListTemplatesResponseBody struct {
	// The page number of the template list.
	//
	// Start value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paginated query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C3A8413B-1F16-4DED-AC3E-61A00718DE8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of templates.
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of templates.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetTemplates() []*ListTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *ListTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplatesResponseBody) SetPageNumber(v int32) *ListTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesResponseBodyTemplates struct {
	// Supplementary information for public templates.
	//
	// example:
	//
	// {"DeploymentDuration":null,"Title":"Self-Built_ElasticSearch_Snapshot_Saved_To_OSS","Labels":{"ResourceTypes":["ALIYUN::ECS::Instance","ALIYUN::ECS::SecurityGroup","ALIYUN::ECS::VPC","ALIYUN::ECS::VSwitch","ALIYUN::OSS::Bucket","ALIYUN::ROS::WaitCondition","ALIYUN::ROS::WaitConditionHandle"],"DeployTypes":["ROS"],"ApplicationScenes":["其他"]},"Provider":"ROS","Categories":["Solution"]}
	AdditionalInfo map[string]interface{} `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2019-10-15T08:17:14.000000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Template description.
	//
	// example:
	//
	// test-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID of the Alibaba Cloud account to which the template belongs.
	//
	// example:
	//
	// 151266687691****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing type of the template.
	//
	// Values:
	//
	// - Private: The template is owned by the user themselves.
	//
	// - Shared: The template is shared by other users.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// Tags of the template.
	Tags []*ListTemplatesResponseBodyTemplatesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ARN of the template.
	//
	// example:
	//
	// acs:ros:*:151266687691****:template/a52f81be-496f-4e1c-a286-8852ab54****
	TemplateARN *string `json:"TemplateARN,omitempty" xml:"TemplateARN,omitempty"`
	// Template ID.
	//
	// example:
	//
	// 4d4f5aa2-3260-4e47-863b-763fbb12****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template name.
	//
	// example:
	//
	// demo
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Link to the template
	//
	// example:
	//
	// https://ros-public-templates.oss-cn-hangzhou.aliyuncs.com/demo.yml
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
	// Latest template version name.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The last update time of the template.
	//
	// example:
	//
	// 2019-10-15T08:17:14.000000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) GetAdditionalInfo() map[string]interface{} {
	return s.AdditionalInfo
}

func (s *ListTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListTemplatesResponseBodyTemplates) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListTemplatesResponseBodyTemplates) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplatesResponseBodyTemplates) GetShareType() *string {
	return s.ShareType
}

func (s *ListTemplatesResponseBodyTemplates) GetTags() []*ListTemplatesResponseBodyTemplatesTags {
	return s.Tags
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateARN() *string {
	return s.TemplateARN
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListTemplatesResponseBodyTemplates) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTemplatesResponseBodyTemplates) SetAdditionalInfo(v map[string]interface{}) *ListTemplatesResponseBodyTemplates {
	s.AdditionalInfo = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetOwnerId(v string) *ListTemplatesResponseBodyTemplates {
	s.OwnerId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetResourceGroupId(v string) *ListTemplatesResponseBodyTemplates {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetShareType(v string) *ListTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTags(v []*ListTemplatesResponseBodyTemplatesTags) *ListTemplatesResponseBodyTemplates {
	s.Tags = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateARN(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateARN = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateUrl(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateUrl = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesResponseBodyTemplatesTags struct {
	// Tag key of the template.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value of the template.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplatesResponseBodyTemplatesTags) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplatesTags) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplatesTags) GetKey() *string {
	return s.Key
}

func (s *ListTemplatesResponseBodyTemplatesTags) GetValue() *string {
	return s.Value
}

func (s *ListTemplatesResponseBodyTemplatesTags) SetKey(v string) *ListTemplatesResponseBodyTemplatesTags {
	s.Key = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesTags) SetValue(v string) *ListTemplatesResponseBodyTemplatesTags {
	s.Value = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesTags) Validate() error {
	return dara.Validate(s)
}
