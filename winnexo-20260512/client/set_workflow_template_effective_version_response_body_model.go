// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowTemplateEffectiveVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetDescription() *string
	SetEffectiveVersion(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetGmtModified() *string
	SetMessage(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetMessage() *string
	SetName(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetName() *string
	SetRequestId(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetRequestId() *string
	SetSteps(v []*SetWorkflowTemplateEffectiveVersionResponseBodySteps) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetSteps() []*SetWorkflowTemplateEffectiveVersionResponseBodySteps
	SetTemplateId(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetTemplateId() *string
	SetVersion(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetVersion() *string
	SetVersionDescription(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetVersionDescription() *string
	SetVersionStatus(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetVersionStatus() *string
}

type SetWorkflowTemplateEffectiveVersionResponseBody struct {
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
	Steps []*SetWorkflowTemplateEffectiveVersionResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
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
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The change description of the version. An empty value indicates that no description is provided.
	//
	// example:
	//
	// Added a review step
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	// The version status. Valid values:
	//
	// - DRAFT: Draft.
	//
	// - SUBMITTED: Submitted.
	//
	// example:
	//
	// SUBMITTED
	VersionStatus *string `json:"versionStatus,omitempty" xml:"versionStatus,omitempty"`
}

func (s SetWorkflowTemplateEffectiveVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowTemplateEffectiveVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetName() *string {
	return s.Name
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetSteps() []*SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	return s.Steps
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetVersion() *string {
	return s.Version
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetVersionStatus() *string {
	return s.VersionStatus
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetCode(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Code = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetCreatedByUserId(v int64) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetDescription(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Description = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetEffectiveVersion(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetGmtCreate(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetGmtModified(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.GmtModified = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetMessage(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Message = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetName(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Name = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetRequestId(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetSteps(v []*SetWorkflowTemplateEffectiveVersionResponseBodySteps) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Steps = v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetTemplateId(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.TemplateId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetVersion(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Version = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetVersionDescription(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetVersionStatus(v string) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.VersionStatus = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) Validate() error {
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

type SetWorkflowTemplateEffectiveVersionResponseBodySteps struct {
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
	// The name of the digital worker executor. This parameter is required when ownerType is set to AGENT.
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// The type of the executor. Valid values:
	//
	// - MEMBER: Workspace member.
	//
	// - AGENT: Digital worker.
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
	// The list of user IDs of the owners (formerly reviewers). This parameter is valid only for steps with a digital worker executor and is always empty for steps with a member executor.
	//
	// example:
	//
	// example
	ReviewerUserIds []*int64 `json:"reviewerUserIds,omitempty" xml:"reviewerUserIds,omitempty" type:"Repeated"`
	// The order of the step.
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The array of ordered step description segments. Supported types include text, mention, skill, app, files, and web_search. When plain text is entered, a single text segment is returned. The fields and nested attachments of each segment are retained as-is by the backend.
	//
	// example:
	//
	// [{"type":"text","content":"Organize input materials"}]
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of structured step description segments. Supported types include text, mention, skill, app, files, and web_search. The fields and nested attachments of each segment are retained as-is by the backend.
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

func (s SetWorkflowTemplateEffectiveVersionResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowTemplateEffectiveVersionResponseBodySteps) GoString() string {
	return s.String()
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetModelMode(v string) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetName(v string) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.Name = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetOwnerAgentName(v string) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetOwnerType(v string) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetOwnerUserId(v int64) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetReviewerUserIds(v []*int64) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetSortOrder(v int64) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetSpecification(v string) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetSpecificationJson(v []map[string]interface{}) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.SpecificationJson = v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetStepId(v string) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
