// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateDraftVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *CreateWorkflowTemplateDraftVersionResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetDescription() *string
	SetEffectiveVersion(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetGmtModified() *string
	SetMessage(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetMessage() *string
	SetName(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetName() *string
	SetRequestId(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetRequestId() *string
	SetSteps(v []*CreateWorkflowTemplateDraftVersionResponseBodySteps) *CreateWorkflowTemplateDraftVersionResponseBody
	GetSteps() []*CreateWorkflowTemplateDraftVersionResponseBodySteps
	SetTemplateId(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetTemplateId() *string
	SetVersion(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetVersion() *string
	SetVersionDescription(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetVersionDescription() *string
	SetVersionStatus(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetVersionStatus() *string
}

type CreateWorkflowTemplateDraftVersionResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 1
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The current effective version number of the template. An empty value indicates that only a draft exists and the template cannot be used for work items.
	//
	// example:
	//
	// v1
	EffectiveVersion *string `json:"effectiveVersion,omitempty" xml:"effectiveVersion,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// example
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time in ISO 8601 format.
	//
	// example:
	//
	// example
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ordered steps.
	//
	// example:
	//
	// example
	Steps []*CreateWorkflowTemplateDraftVersionResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The ID of the template.
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The version number to which the returned content belongs.
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The change description for this version. An empty value indicates that no description is provided.
	//
	// example:
	//
	// Added a review step
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	// The version status. Valid values: DRAFT (draft) and SUBMITTED (submitted).
	//
	// example:
	//
	// DRAFT
	VersionStatus *string `json:"versionStatus,omitempty" xml:"versionStatus,omitempty"`
}

func (s CreateWorkflowTemplateDraftVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateDraftVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetSteps() []*CreateWorkflowTemplateDraftVersionResponseBodySteps {
	return s.Steps
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetVersionStatus() *string {
	return s.VersionStatus
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetCode(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetCreatedByUserId(v int64) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetDescription(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetEffectiveVersion(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetGmtCreate(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetGmtModified(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetMessage(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetName(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Name = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetRequestId(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetSteps(v []*CreateWorkflowTemplateDraftVersionResponseBodySteps) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Steps = v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetTemplateId(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetVersion(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Version = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetVersionDescription(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetVersionStatus(v string) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.VersionStatus = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBody) Validate() error {
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkflowTemplateDraftVersionResponseBodySteps struct {
	// The model tier.
	//
	// example:
	//
	// standard
	ModelMode *string `json:"modelMode,omitempty" xml:"modelMode,omitempty"`
	// The name of the step.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee executor. This parameter is required when ownerType is set to AGENT.
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// The type of the executor. Valid values: MEMBER (workspace member) and AGENT (digital employee).
	//
	// example:
	//
	// MEMBER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The user ID of the member executor. This parameter is required when ownerType is set to MEMBER.
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// The list of user IDs for the owners (formerly reviewers). This parameter is valid only for steps with a digital employee executor and is always empty for steps with a member executor.
	//
	// example:
	//
	// example
	ReviewerUserIds []*int64 `json:"reviewerUserIds,omitempty" xml:"reviewerUserIds,omitempty" type:"Repeated"`
	// The sort order of the step.
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The array of ordered step description segments. Supported types include text, mention, skill, app, files, and web_search. When plain text is entered, a single text segment is returned. The fields and nested attachments of each segment are retained as-is from the backend.
	//
	// example:
	//
	// [{"type":"text","content":"Organize input materials"}]
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of structured step description segments. Supported types include text, mention, skill, app, files, and web_search. The fields and nested attachments of each segment are retained as-is from the backend.
	//
	// example:
	//
	// [{"type":"text","content":"Organize input materials"}]
	SpecificationJson []map[string]interface{} `json:"specificationJson,omitempty" xml:"specificationJson,omitempty" type:"Repeated"`
	// The ID of the step.
	//
	// example:
	//
	// collect
	StepId *string `json:"stepId,omitempty" xml:"stepId,omitempty"`
}

func (s CreateWorkflowTemplateDraftVersionResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateDraftVersionResponseBodySteps) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetModelMode(v string) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetName(v string) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.Name = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetOwnerAgentName(v string) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetOwnerType(v string) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetOwnerUserId(v int64) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetReviewerUserIds(v []*int64) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetSortOrder(v int64) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetSpecification(v string) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetSpecificationJson(v []map[string]interface{}) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.SpecificationJson = v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetStepId(v string) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
