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
	JobTemplates []*ListJobTemplatesResponseBodyJobTemplates `json:"JobTemplates,omitempty" xml:"JobTemplates,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 本次请求的 ID，用于诊断和答疑。
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0D*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
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
	// example:
	//
	// 2
	DefaultVersion *int32 `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// example:
	//
	// job description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-01-23T07:29:06Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-03-03T05:48:02Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 20**************26
	ModifiedBy *string `json:"ModifiedBy,omitempty" xml:"ModifiedBy,omitempty"`
	// example:
	//
	// tpl1r5g9ait7****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// job-template-1772516653885
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 142388383837****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 20**************26
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 88****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobTemplatesResponseBodyJobTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesResponseBodyJobTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetDefaultVersion() *int32 {
	return s.DefaultVersion
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetModifiedBy() *string {
	return s.ModifiedBy
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetTenantId() *string {
	return s.TenantId
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetUserId() *string {
	return s.UserId
}

func (s *ListJobTemplatesResponseBodyJobTemplates) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetDefaultVersion(v int32) *ListJobTemplatesResponseBodyJobTemplates {
	s.DefaultVersion = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetDescription(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.Description = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetGmtCreateTime(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.GmtCreateTime = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetGmtModifyTime(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.GmtModifyTime = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetMetadata(v map[string]interface{}) *ListJobTemplatesResponseBodyJobTemplates {
	s.Metadata = v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetModifiedBy(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.ModifiedBy = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetTemplateId(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetTemplateName(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetTenantId(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.TenantId = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetUserId(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.UserId = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) SetWorkspaceId(v string) *ListJobTemplatesResponseBodyJobTemplates {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobTemplatesResponseBodyJobTemplates) Validate() error {
	return dara.Validate(s)
}
