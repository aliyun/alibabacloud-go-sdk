// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplateVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkflowTemplateVersionsResponseBody
	GetCode() *string
	SetEffectiveVersion(v string) *ListWorkflowTemplateVersionsResponseBody
	GetEffectiveVersion() *string
	SetItems(v []*ListWorkflowTemplateVersionsResponseBodyItems) *ListWorkflowTemplateVersionsResponseBody
	GetItems() []*ListWorkflowTemplateVersionsResponseBodyItems
	SetMessage(v string) *ListWorkflowTemplateVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkflowTemplateVersionsResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *ListWorkflowTemplateVersionsResponseBody
	GetTemplateId() *string
}

type ListWorkflowTemplateVersionsResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The current effective version number. An empty value indicates that only a draft exists.
	//
	// example:
	//
	// v1
	EffectiveVersion *string `json:"effectiveVersion,omitempty" xml:"effectiveVersion,omitempty"`
	// All versions, sorted in descending order by version number.
	//
	// example:
	//
	// example
	Items []*ListWorkflowTemplateVersionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s ListWorkflowTemplateVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplateVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkflowTemplateVersionsResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *ListWorkflowTemplateVersionsResponseBody) GetItems() []*ListWorkflowTemplateVersionsResponseBodyItems {
	return s.Items
}

func (s *ListWorkflowTemplateVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowTemplateVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowTemplateVersionsResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWorkflowTemplateVersionsResponseBody) SetCode(v string) *ListWorkflowTemplateVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBody) SetEffectiveVersion(v string) *ListWorkflowTemplateVersionsResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBody) SetItems(v []*ListWorkflowTemplateVersionsResponseBodyItems) *ListWorkflowTemplateVersionsResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBody) SetMessage(v string) *ListWorkflowTemplateVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBody) SetRequestId(v string) *ListWorkflowTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBody) SetTemplateId(v string) *ListWorkflowTemplateVersionsResponseBody {
	s.TemplateId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowTemplateVersionsResponseBodyItems struct {
	// The ID of the user who created the version.
	//
	// example:
	//
	// 1
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// The template description for this version. This can be displayed when version v1 has no version description.
	//
	// example:
	//
	// Automatically generate weekly reports
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The creation time of the version in ISO 8601 format.
	//
	// example:
	//
	// 2026-10-19T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time of the version in ISO 8601 format.
	//
	// example:
	//
	// 2026-10-19T10:00:00+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The submission time of the version in ISO 8601 format. This is empty for drafts.
	//
	// example:
	//
	// 2026-10-19T10:00:00+08:00
	GmtSubmitted *string `json:"gmtSubmitted,omitempty" xml:"gmtSubmitted,omitempty"`
	// Indicates whether this is the current unique effective version.
	//
	// example:
	//
	// false
	IsEffective *bool `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	// The template name for this version.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The submitted version number on which this version is based. This is empty when v1 is created. Saving or submitting a draft retains the original source version.
	//
	// example:
	//
	// v1
	SourceVersion *string `json:"sourceVersion,omitempty" xml:"sourceVersion,omitempty"`
	// The version status. Valid values:
	//
	// - DRAFT: Draft.
	//
	// - SUBMITTED: Submitted.
	//
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of steps in this version.
	//
	// example:
	//
	// 3
	StepCount *int64 `json:"stepCount,omitempty" xml:"stepCount,omitempty"`
	// The ID of the user who submitted the version. This is empty for drafts.
	//
	// example:
	//
	// 1
	SubmittedByUserId *int64 `json:"submittedByUserId,omitempty" xml:"submittedByUserId,omitempty"`
	// The ID of the user who last modified the version.
	//
	// example:
	//
	// 1
	UpdatedByUserId *int64 `json:"updatedByUserId,omitempty" xml:"updatedByUserId,omitempty"`
	// The version number.
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The persisted version change description. An empty value indicates that it is left blank.
	//
	// example:
	//
	// Added an approval step
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s ListWorkflowTemplateVersionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplateVersionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetGmtSubmitted() *string {
	return s.GmtSubmitted
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetIsEffective() *bool {
	return s.IsEffective
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetSourceVersion() *string {
	return s.SourceVersion
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetStepCount() *int64 {
	return s.StepCount
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetSubmittedByUserId() *int64 {
	return s.SubmittedByUserId
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetUpdatedByUserId() *int64 {
	return s.UpdatedByUserId
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetVersion() *string {
	return s.Version
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetCreatedByUserId(v int64) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.CreatedByUserId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetDescription(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetGmtCreate(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetGmtModified(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetGmtSubmitted(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.GmtSubmitted = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetIsEffective(v bool) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.IsEffective = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetName(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetSourceVersion(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.SourceVersion = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetStatus(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetStepCount(v int64) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.StepCount = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetSubmittedByUserId(v int64) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.SubmittedByUserId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetUpdatedByUserId(v int64) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.UpdatedByUserId = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetVersion(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.Version = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) SetVersionDescription(v string) *ListWorkflowTemplateVersionsResponseBodyItems {
	s.VersionDescription = &v
	return s
}

func (s *ListWorkflowTemplateVersionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
