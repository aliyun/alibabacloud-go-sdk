// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkflowTemplateResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *CreateWorkflowTemplateResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *CreateWorkflowTemplateResponseBody
	GetDescription() *string
	SetEffectiveVersion(v string) *CreateWorkflowTemplateResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *CreateWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *CreateWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetMessage(v string) *CreateWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *CreateWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *CreateWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSteps(v []*CreateWorkflowTemplateResponseBodySteps) *CreateWorkflowTemplateResponseBody
	GetSteps() []*CreateWorkflowTemplateResponseBodySteps
	SetTemplateId(v string) *CreateWorkflowTemplateResponseBody
	GetTemplateId() *string
	SetVersion(v string) *CreateWorkflowTemplateResponseBody
	GetVersion() *string
	SetVersionDescription(v string) *CreateWorkflowTemplateResponseBody
	GetVersionDescription() *string
	SetVersionStatus(v string) *CreateWorkflowTemplateResponseBody
	GetVersionStatus() *string
}

type CreateWorkflowTemplateResponseBody struct {
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
	Steps []*CreateWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The template ID.
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The version number of the returned content.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The change description for this version. If this field is empty for v1, the template description is displayed.
	//
	// example:
	//
	// Initial template creation
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	// The version status. Valid values: DRAFT: Draft. SUBMITTED: Submitted.
	//
	// example:
	//
	// SUBMITTED
	VersionStatus *string `json:"versionStatus,omitempty" xml:"versionStatus,omitempty"`
}

func (s CreateWorkflowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkflowTemplateResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *CreateWorkflowTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkflowTemplateResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *CreateWorkflowTemplateResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateWorkflowTemplateResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateWorkflowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkflowTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowTemplateResponseBody) GetSteps() []*CreateWorkflowTemplateResponseBodySteps {
	return s.Steps
}

func (s *CreateWorkflowTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowTemplateResponseBody) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkflowTemplateResponseBody) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateWorkflowTemplateResponseBody) GetVersionStatus() *string {
	return s.VersionStatus
}

func (s *CreateWorkflowTemplateResponseBody) SetCode(v string) *CreateWorkflowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetCreatedByUserId(v int64) *CreateWorkflowTemplateResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetDescription(v string) *CreateWorkflowTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetEffectiveVersion(v string) *CreateWorkflowTemplateResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetGmtCreate(v string) *CreateWorkflowTemplateResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetGmtModified(v string) *CreateWorkflowTemplateResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetMessage(v string) *CreateWorkflowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetName(v string) *CreateWorkflowTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetRequestId(v string) *CreateWorkflowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetSteps(v []*CreateWorkflowTemplateResponseBodySteps) *CreateWorkflowTemplateResponseBody {
	s.Steps = v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetTemplateId(v string) *CreateWorkflowTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetVersion(v string) *CreateWorkflowTemplateResponseBody {
	s.Version = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetVersionDescription(v string) *CreateWorkflowTemplateResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) SetVersionStatus(v string) *CreateWorkflowTemplateResponseBody {
	s.VersionStatus = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBody) Validate() error {
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

type CreateWorkflowTemplateResponseBodySteps struct {
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
	// The owner type. Valid values: MEMBER: Human. AGENT: Agent.
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
	// The list of reviewer user IDs.
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
	// The array of structured segments for step instructions. Supports text, mention, skill, app, files, and web_search. Segment fields and nested attachments are preserved as returned by the backend.
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

func (s CreateWorkflowTemplateResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowTemplateResponseBodySteps) GoString() string {
	return s.String()
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *CreateWorkflowTemplateResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetModelMode(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetName(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.Name = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetOwnerAgentName(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetOwnerType(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetOwnerUserId(v int64) *CreateWorkflowTemplateResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetReviewerUserIds(v []*int64) *CreateWorkflowTemplateResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetSortOrder(v int64) *CreateWorkflowTemplateResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetSpecification(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetSpecificationJson(v []map[string]interface{}) *CreateWorkflowTemplateResponseBodySteps {
	s.SpecificationJson = v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) SetStepId(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
