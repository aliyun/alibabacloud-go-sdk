// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplates(v []*ListJobTemplatesResponseBodyJobTemplates) *ListJobTemplatesResponseBody
	GetJobTemplates() []*ListJobTemplatesResponseBodyJobTemplates
	SetPageNumber(v int32) *ListJobTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobTemplatesResponseBody
	GetTotalCount() *int32
}

type ListJobTemplatesResponseBody struct {
	// The job templates.
	JobTemplates []*ListJobTemplatesResponseBodyJobTemplates `json:"JobTemplates,omitempty" xml:"JobTemplates,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of templates.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBody) GetJobTemplates() []*ListJobTemplatesResponseBodyJobTemplates {
	return s.JobTemplates
}

func (s *ListJobTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobTemplatesResponseBody) SetJobTemplates(v []*ListJobTemplatesResponseBodyJobTemplates) *ListJobTemplatesResponseBody {
	s.JobTemplates = v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageNumber(v int32) *ListJobTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageSize(v int32) *ListJobTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetRequestId(v string) *ListJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTotalCount(v int32) *ListJobTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobTemplatesResponseBody) Validate() error {
	if s.JobTemplates != nil {
		for _, item := range s.JobTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobTemplatesResponseBodyJobTemplates struct {
	// The creation time. The example uses the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-04-20 11:09:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The template description.
	//
	// example:
	//
	// Demo
	JobTemplateDescription *string `json:"JobTemplateDescription,omitempty" xml:"JobTemplateDescription,omitempty"`
	// The job template ID.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// my-template
	JobTemplateName *string `json:"JobTemplateName,omitempty" xml:"JobTemplateName,omitempty"`
	// The job template status.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job tags.
	Tags []*ListJobTemplatesResponseBodyJobTemplatesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The update time. The example uses the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-04-20 11:09:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListJobTemplatesResponseBodyJobTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesResponseBodyJobTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetJobTemplateDescription() *string {
	return s.JobTemplateDescription
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetJobTemplateName() *string {
	return s.JobTemplateName
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetStatus() *string {
	return s.Status
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetTags() []*ListJobTemplatesResponseBodyJobTemplatesTags {
	return s.Tags
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetCreateTime(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetJobTemplateDescription(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.JobTemplateDescription = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetJobTemplateId(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.JobTemplateId = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetJobTemplateName(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.JobTemplateName = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetStatus(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.Status = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetTags(v []*ListJobTemplatesResponseBodyJobTemplatesTags) *ListJobTemplatesResponseBodyJobTemplates {
	s.Tags = v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetUpdateTime(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) Validate() error {
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

type ListJobTemplatesResponseBodyJobTemplatesTags struct {
	// The job tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The job tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListJobTemplatesResponseBodyJobTemplatesTags) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesResponseBodyJobTemplatesTags) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyJobTemplatesTags) GetKey() *string {
	return s.Key
}

func (s *ListJobTemplatesResponseBodyJobTemplatesTags) GetValue() *string {
	return s.Value
}

func (s *ListJobTemplatesResponseBodyJobTemplatesTags) SetKey(v string) *ListJobTemplatesResponseBodyJobTemplatesTags {
	s.Key = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplatesTags) SetValue(v string) *ListJobTemplatesResponseBodyJobTemplatesTags {
	s.Value = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplatesTags) Validate() error {
	return dara.Validate(s)
}
