// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWorkflowTemplateResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *UpdateWorkflowTemplateResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *UpdateWorkflowTemplateResponseBody
	GetDescription() *string
	SetEffectiveVersion(v string) *UpdateWorkflowTemplateResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *UpdateWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *UpdateWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetMessage(v string) *UpdateWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *UpdateWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSteps(v []*UpdateWorkflowTemplateResponseBodySteps) *UpdateWorkflowTemplateResponseBody
	GetSteps() []*UpdateWorkflowTemplateResponseBodySteps
	SetTemplateId(v string) *UpdateWorkflowTemplateResponseBody
	GetTemplateId() *string
	SetVersion(v string) *UpdateWorkflowTemplateResponseBody
	GetVersion() *string
	SetVersionDescription(v string) *UpdateWorkflowTemplateResponseBody
	GetVersionDescription() *string
	SetVersionStatus(v string) *UpdateWorkflowTemplateResponseBody
	GetVersionStatus() *string
}

type UpdateWorkflowTemplateResponseBody struct {
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
	// The template description.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The current effective version number of the template. An empty value indicates that only a draft exists and it cannot be used for work items.
	//
	// example:
	//
	// v2
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
	// The template name.
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
	Steps []*UpdateWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The template ID.
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
	// The change description of the version. An empty value indicates that no description is provided.
	//
	// example:
	//
	// Add a review step
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

func (s UpdateWorkflowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWorkflowTemplateResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *UpdateWorkflowTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowTemplateResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *UpdateWorkflowTemplateResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateWorkflowTemplateResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateWorkflowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowTemplateResponseBody) GetSteps() []*UpdateWorkflowTemplateResponseBodySteps {
	return s.Steps
}

func (s *UpdateWorkflowTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateResponseBody) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkflowTemplateResponseBody) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateWorkflowTemplateResponseBody) GetVersionStatus() *string {
	return s.VersionStatus
}

func (s *UpdateWorkflowTemplateResponseBody) SetCode(v string) *UpdateWorkflowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetCreatedByUserId(v int64) *UpdateWorkflowTemplateResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetDescription(v string) *UpdateWorkflowTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetEffectiveVersion(v string) *UpdateWorkflowTemplateResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetGmtCreate(v string) *UpdateWorkflowTemplateResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetGmtModified(v string) *UpdateWorkflowTemplateResponseBody {
	s.GmtModified = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetMessage(v string) *UpdateWorkflowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetName(v string) *UpdateWorkflowTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetRequestId(v string) *UpdateWorkflowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetSteps(v []*UpdateWorkflowTemplateResponseBodySteps) *UpdateWorkflowTemplateResponseBody {
	s.Steps = v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetTemplateId(v string) *UpdateWorkflowTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetVersion(v string) *UpdateWorkflowTemplateResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetVersionDescription(v string) *UpdateWorkflowTemplateResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetVersionStatus(v string) *UpdateWorkflowTemplateResponseBody {
	s.VersionStatus = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) Validate() error {
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

type UpdateWorkflowTemplateResponseBodySteps struct {
	// The model tier.
	//
	// example:
	//
	// standard
	ModelMode *string `json:"modelMode,omitempty" xml:"modelMode,omitempty"`
	// The step name.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the responsible agent.
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// The owner type. Valid values:
	//
	// - MEMBER: Human.
	//
	// - AGENT: Agent.
	//
	// example:
	//
	// MEMBER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The user ID of the human owner.
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// The list of user IDs of the reviewers.
	//
	// example:
	//
	// example
	ReviewerUserIds []*int64 `json:"reviewerUserIds,omitempty" xml:"reviewerUserIds,omitempty" type:"Repeated"`
	// The step order.
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The step processing instructions.
	//
	// example:
	//
	// example
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of structured segments for the step instructions. Supports text, mention, skill, app, files, and web_search. The fields and nested attachments of each segment are retained as returned by the backend.
	//
	// example:
	//
	// [{"type":"text","content":"Organize input materials"}]
	SpecificationJson []map[string]interface{} `json:"specificationJson,omitempty" xml:"specificationJson,omitempty" type:"Repeated"`
	// The step ID.
	//
	// example:
	//
	// collect
	StepId *string `json:"stepId,omitempty" xml:"stepId,omitempty"`
}

func (s UpdateWorkflowTemplateResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateResponseBodySteps) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetModelMode(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetName(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetOwnerAgentName(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetOwnerType(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetOwnerUserId(v int64) *UpdateWorkflowTemplateResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetReviewerUserIds(v []*int64) *UpdateWorkflowTemplateResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetSortOrder(v int64) *UpdateWorkflowTemplateResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetSpecification(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetSpecificationJson(v []map[string]interface{}) *UpdateWorkflowTemplateResponseBodySteps {
	s.SpecificationJson = v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetStepId(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
