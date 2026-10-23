// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWorkflowTemplateVersionResponseBody
	GetCode() *string
	SetEffectiveVersion(v string) *DeleteWorkflowTemplateVersionResponseBody
	GetEffectiveVersion() *string
	SetItems(v []*DeleteWorkflowTemplateVersionResponseBodyItems) *DeleteWorkflowTemplateVersionResponseBody
	GetItems() []*DeleteWorkflowTemplateVersionResponseBodyItems
	SetMessage(v string) *DeleteWorkflowTemplateVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkflowTemplateVersionResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *DeleteWorkflowTemplateVersionResponseBody
	GetTemplateId() *string
}

type DeleteWorkflowTemplateVersionResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The current active version number. An empty value indicates that only a draft exists.
	//
	// example:
	//
	// v1
	EffectiveVersion *string `json:"effectiveVersion,omitempty" xml:"effectiveVersion,omitempty"`
	// All remaining versions after deletion, sorted in descending order by version number.
	//
	// example:
	//
	// example
	Items []*DeleteWorkflowTemplateVersionResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
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

func (s DeleteWorkflowTemplateVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWorkflowTemplateVersionResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *DeleteWorkflowTemplateVersionResponseBody) GetItems() []*DeleteWorkflowTemplateVersionResponseBodyItems {
	return s.Items
}

func (s *DeleteWorkflowTemplateVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkflowTemplateVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowTemplateVersionResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWorkflowTemplateVersionResponseBody) SetCode(v string) *DeleteWorkflowTemplateVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBody) SetEffectiveVersion(v string) *DeleteWorkflowTemplateVersionResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBody) SetItems(v []*DeleteWorkflowTemplateVersionResponseBodyItems) *DeleteWorkflowTemplateVersionResponseBody {
	s.Items = v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBody) SetMessage(v string) *DeleteWorkflowTemplateVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBody) SetRequestId(v string) *DeleteWorkflowTemplateVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBody) SetTemplateId(v string) *DeleteWorkflowTemplateVersionResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBody) Validate() error {
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

type DeleteWorkflowTemplateVersionResponseBodyItems struct {
	// The ID of the user who created the version.
	//
	// example:
	//
	// 1
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// The description of the template for this version. This can be used for display when version v1 has no version description.
	//
	// example:
	//
	// Automatically generate weekly reports
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the version was created. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-10-19T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the version was modified. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-10-19T10:00:00+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The time when the version was submitted. The time is in the ISO 8601 format. This is empty for a draft.
	//
	// example:
	//
	// 2026-10-19T10:00:00+08:00
	GmtSubmitted *string `json:"gmtSubmitted,omitempty" xml:"gmtSubmitted,omitempty"`
	// Indicates whether the version is the only current active version.
	//
	// example:
	//
	// false
	IsEffective *bool `json:"isEffective,omitempty" xml:"isEffective,omitempty"`
	// The name of the template for this version.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The submitted version number on which this version is based. This is empty for a newly created v1 version. Saving or submitting a draft retains the original source version.
	//
	// example:
	//
	// v1
	SourceVersion *string `json:"sourceVersion,omitempty" xml:"sourceVersion,omitempty"`
	// The version status. Valid values: DRAFT (draft) and SUBMITTED (submitted).
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
	// The ID of the user who submitted the version. This is empty for a draft.
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
	// The persisted version change description. An empty value indicates that no description is provided.
	//
	// example:
	//
	// Added an approval step
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s DeleteWorkflowTemplateVersionResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateVersionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetGmtSubmitted() *string {
	return s.GmtSubmitted
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetIsEffective() *bool {
	return s.IsEffective
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetSourceVersion() *string {
	return s.SourceVersion
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetStepCount() *int64 {
	return s.StepCount
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetSubmittedByUserId() *int64 {
	return s.SubmittedByUserId
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetUpdatedByUserId() *int64 {
	return s.UpdatedByUserId
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetVersion() *string {
	return s.Version
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetCreatedByUserId(v int64) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.CreatedByUserId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetDescription(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetGmtCreate(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetGmtModified(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetGmtSubmitted(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.GmtSubmitted = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetIsEffective(v bool) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.IsEffective = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetName(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetSourceVersion(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.SourceVersion = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetStatus(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetStepCount(v int64) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.StepCount = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetSubmittedByUserId(v int64) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.SubmittedByUserId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetUpdatedByUserId(v int64) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.UpdatedByUserId = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetVersion(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.Version = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) SetVersionDescription(v string) *DeleteWorkflowTemplateVersionResponseBodyItems {
	s.VersionDescription = &v
	return s
}

func (s *DeleteWorkflowTemplateVersionResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
