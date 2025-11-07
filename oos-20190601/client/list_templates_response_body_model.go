// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody
	GetTemplates() []*ListTemplatesResponseBodyTemplates
}

type ListTemplatesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BEF54BA-17B6-449F-A219-49ACB157E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template metadata.
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetTemplates() []*ListTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *ListTemplatesResponseBody) SetMaxResults(v int32) *ListTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesResponseBody) SetNextToken(v string) *ListTemplatesResponseBody {
	s.NextToken = &v
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

func (s *ListTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplatesResponseBodyTemplates struct {
	// The template type.
	//
	// example:
	//
	// TimerTrigger
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The template constraints.
	//
	// example:
	//
	// {
	//
	//   "InstanceTypeFamilies": ["ecs.g8y", "ecs.c8y"],
	//
	//   "ImageTypes": ["system"],
	//
	//   "OSPlatforms": ["CentOS", "Ubuntu"],
	//
	//   "OSVersions": ["CentOS7.9 64bit"]
	//
	// }
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the template.
	//
	// example:
	//
	// root(1309200)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The creation time of the template.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The template description.
	//
	// example:
	//
	// Describe instances of given status
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template was configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// Indicates whether the template is added to favorites.
	//
	// example:
	//
	// true
	IsFavorite *bool `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// The popularity of the public template. Valid values: **1-10**. A greater value indicates higher popularity. If **ShareType*	- is set to **Private**, the value of this parameter is `-1`.
	//
	// >  This parameter is valid only if **ShareType*	- is set to **Public**.
	//
	// example:
	//
	// 8
	Popularity *int32 `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	// The user who published the template.
	//
	// example:
	//
	// aliyun
	Publisher *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of a template created by a user is **Private**. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The template format. The system automatically determines whether the format of the template is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-94753deed38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The template type.
	//
	// example:
	//
	// private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The template version. The version contains the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The number of times for which the private template is executed. If **ShareType*	- is set to **Public**, the value of this parameter is `-1`.
	//
	// >  This parameter is valid only if **ShareType*	- is set to **Private**.
	//
	// example:
	//
	// 5
	TotalExecutionCount *int32 `json:"TotalExecutionCount,omitempty" xml:"TotalExecutionCount,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// root(13092000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The version name.
	//
	// example:
	//
	// v2.1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) GetCategory() *string {
	return s.Category
}

func (s *ListTemplatesResponseBodyTemplates) GetConstraints() *string {
	return s.Constraints
}

func (s *ListTemplatesResponseBodyTemplates) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListTemplatesResponseBodyTemplates) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListTemplatesResponseBodyTemplates) GetHasTrigger() *bool {
	return s.HasTrigger
}

func (s *ListTemplatesResponseBodyTemplates) GetHash() *string {
	return s.Hash
}

func (s *ListTemplatesResponseBodyTemplates) GetIsFavorite() *bool {
	return s.IsFavorite
}

func (s *ListTemplatesResponseBodyTemplates) GetPopularity() *int32 {
	return s.Popularity
}

func (s *ListTemplatesResponseBodyTemplates) GetPublisher() *string {
	return s.Publisher
}

func (s *ListTemplatesResponseBodyTemplates) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplatesResponseBodyTemplates) GetShareType() *string {
	return s.ShareType
}

func (s *ListTemplatesResponseBodyTemplates) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListTemplatesResponseBodyTemplates) GetTotalExecutionCount() *int32 {
	return s.TotalExecutionCount
}

func (s *ListTemplatesResponseBodyTemplates) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListTemplatesResponseBodyTemplates) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListTemplatesResponseBodyTemplates) GetVersionName() *string {
	return s.VersionName
}

func (s *ListTemplatesResponseBodyTemplates) SetCategory(v string) *ListTemplatesResponseBodyTemplates {
	s.Category = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetConstraints(v string) *ListTemplatesResponseBodyTemplates {
	s.Constraints = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreatedBy(v string) *ListTemplatesResponseBodyTemplates {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreatedDate(v string) *ListTemplatesResponseBodyTemplates {
	s.CreatedDate = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetHasTrigger(v bool) *ListTemplatesResponseBodyTemplates {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetHash(v string) *ListTemplatesResponseBodyTemplates {
	s.Hash = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetIsFavorite(v bool) *ListTemplatesResponseBodyTemplates {
	s.IsFavorite = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPopularity(v int32) *ListTemplatesResponseBodyTemplates {
	s.Popularity = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPublisher(v string) *ListTemplatesResponseBodyTemplates {
	s.Publisher = &v
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

func (s *ListTemplatesResponseBodyTemplates) SetTags(v map[string]interface{}) *ListTemplatesResponseBodyTemplates {
	s.Tags = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateFormat(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateFormat = &v
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

func (s *ListTemplatesResponseBodyTemplates) SetTemplateType(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTotalExecutionCount(v int32) *ListTemplatesResponseBodyTemplates {
	s.TotalExecutionCount = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdatedBy(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdatedBy = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdatedDate(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdatedDate = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetVersionName(v string) *ListTemplatesResponseBodyTemplates {
	s.VersionName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
