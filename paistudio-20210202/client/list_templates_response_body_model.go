// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetTemplateData(v []*ListTemplatesResponseBodyTemplateData) *ListTemplatesResponseBody
	GetTemplateData() []*ListTemplatesResponseBodyTemplateData
	SetTotalCount(v int32) *ListTemplatesResponseBody
	GetTotalCount() *int32
}

type ListTemplatesResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateData []*ListTemplatesResponseBodyTemplateData `json:"TemplateData,omitempty" xml:"TemplateData,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetTemplateData() []*ListTemplatesResponseBodyTemplateData {
	return s.TemplateData
}

func (s *ListTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplateData(v []*ListTemplatesResponseBodyTemplateData) *ListTemplatesResponseBody {
	s.TemplateData = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) Validate() error {
	if s.TemplateData != nil {
		for _, item := range s.TemplateData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplatesResponseBodyTemplateData struct {
	Template     *ListTemplatesResponseBodyTemplateDataTemplate     `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	TemplateTag  *ListTemplatesResponseBodyTemplateDataTemplateTag  `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Struct"`
	TemplateType *ListTemplatesResponseBodyTemplateDataTemplateType `json:"TemplateType,omitempty" xml:"TemplateType,omitempty" type:"Struct"`
}

func (s ListTemplatesResponseBodyTemplateData) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateData) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateData) GetTemplate() *ListTemplatesResponseBodyTemplateDataTemplate {
	return s.Template
}

func (s *ListTemplatesResponseBodyTemplateData) GetTemplateTag() *ListTemplatesResponseBodyTemplateDataTemplateTag {
	return s.TemplateTag
}

func (s *ListTemplatesResponseBodyTemplateData) GetTemplateType() *ListTemplatesResponseBodyTemplateDataTemplateType {
	return s.TemplateType
}

func (s *ListTemplatesResponseBodyTemplateData) SetTemplate(v *ListTemplatesResponseBodyTemplateDataTemplate) *ListTemplatesResponseBodyTemplateData {
	s.Template = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateData) SetTemplateTag(v *ListTemplatesResponseBodyTemplateDataTemplateTag) *ListTemplatesResponseBodyTemplateData {
	s.TemplateTag = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateData) SetTemplateType(v *ListTemplatesResponseBodyTemplateDataTemplateType) *ListTemplatesResponseBodyTemplateData {
	s.TemplateType = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateData) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateTag != nil {
		if err := s.TemplateTag.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateType != nil {
		if err := s.TemplateType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTemplatesResponseBodyTemplateDataTemplate struct {
	// example:
	//
	// {}
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail      *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// url://xxx
	DocLink *string `json:"DocLink,omitempty" xml:"DocLink,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// url://xxx
	ImageLink *string                  `json:"ImageLink,omitempty" xml:"ImageLink,omitempty"`
	Labels    []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name      *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// template-rbvg5wzljzjhc9ks92
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplateDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateDataTemplate) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetContent() *string {
	return s.Content
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetCreator() *string {
	return s.Creator
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetDescription() *string {
	return s.Description
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetDetail() *string {
	return s.Detail
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetDocLink() *string {
	return s.DocLink
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetImageLink() *string {
	return s.ImageLink
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetLabels() []map[string]interface{} {
	return s.Labels
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetName() *string {
	return s.Name
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetContent(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Content = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetCreator(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Creator = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetDescription(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetDetail(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Detail = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetDocLink(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.DocLink = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetGmtCreateTime(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetGmtModifiedTime(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetImageLink(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.ImageLink = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetLabels(v []map[string]interface{}) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Labels = v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetName(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) SetTemplateId(v string) *ListTemplatesResponseBodyTemplateDataTemplate {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplate) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesResponseBodyTemplateDataTemplateTag struct {
	// example:
	//
	// PyTorch
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// template-tag-rbvg5wzljzjhc9ks92
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// template-tag-type-rbvg5wzljzjhc9ks92
	TypeId *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplateDataTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateDataTemplateTag) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) GetName() *string {
	return s.Name
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) GetTagId() *string {
	return s.TagId
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) GetTypeId() *string {
	return s.TypeId
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) SetName(v string) *ListTemplatesResponseBodyTemplateDataTemplateTag {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) SetTagId(v string) *ListTemplatesResponseBodyTemplateDataTemplateTag {
	s.TagId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) SetTypeId(v string) *ListTemplatesResponseBodyTemplateDataTemplateTag {
	s.TypeId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateTag) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesResponseBodyTemplateDataTemplateType struct {
	// example:
	//
	// 行业分类
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// templatetagtype-rbvg5wzljzjhc9ks92
	TypeId *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
}

func (s ListTemplatesResponseBodyTemplateDataTemplateType) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplateDataTemplateType) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) GetName() *string {
	return s.Name
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) GetTypeId() *string {
	return s.TypeId
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) SetName(v string) *ListTemplatesResponseBodyTemplateDataTemplateType {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) SetTypeId(v string) *ListTemplatesResponseBodyTemplateDataTemplateType {
	s.TypeId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplateDataTemplateType) Validate() error {
	return dara.Validate(s)
}
