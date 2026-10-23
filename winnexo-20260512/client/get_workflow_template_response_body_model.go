// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkflowTemplateResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *GetWorkflowTemplateResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *GetWorkflowTemplateResponseBody
	GetDescription() *string
	SetEffectiveVersion(v string) *GetWorkflowTemplateResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *GetWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetMessage(v string) *GetWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *GetWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *GetWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSteps(v []*GetWorkflowTemplateResponseBodySteps) *GetWorkflowTemplateResponseBody
	GetSteps() []*GetWorkflowTemplateResponseBodySteps
	SetTemplateId(v string) *GetWorkflowTemplateResponseBody
	GetTemplateId() *string
	SetVersion(v string) *GetWorkflowTemplateResponseBody
	GetVersion() *string
	SetVersionDescription(v string) *GetWorkflowTemplateResponseBody
	GetVersionDescription() *string
	SetVersionStatus(v string) *GetWorkflowTemplateResponseBody
	GetVersionStatus() *string
}

type GetWorkflowTemplateResponseBody struct {
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
	// The name of the template.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
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
	Steps []*GetWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
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
	// The change description for this version. If this field is empty for v1, the template description can be displayed.
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

func (s GetWorkflowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkflowTemplateResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *GetWorkflowTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetWorkflowTemplateResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *GetWorkflowTemplateResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetWorkflowTemplateResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetWorkflowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *GetWorkflowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowTemplateResponseBody) GetSteps() []*GetWorkflowTemplateResponseBodySteps {
	return s.Steps
}

func (s *GetWorkflowTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetWorkflowTemplateResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetWorkflowTemplateResponseBody) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *GetWorkflowTemplateResponseBody) GetVersionStatus() *string {
	return s.VersionStatus
}

func (s *GetWorkflowTemplateResponseBody) SetCode(v string) *GetWorkflowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetCreatedByUserId(v int64) *GetWorkflowTemplateResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetDescription(v string) *GetWorkflowTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetEffectiveVersion(v string) *GetWorkflowTemplateResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetGmtCreate(v string) *GetWorkflowTemplateResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetGmtModified(v string) *GetWorkflowTemplateResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetMessage(v string) *GetWorkflowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetName(v string) *GetWorkflowTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetRequestId(v string) *GetWorkflowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetSteps(v []*GetWorkflowTemplateResponseBodySteps) *GetWorkflowTemplateResponseBody {
	s.Steps = v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetTemplateId(v string) *GetWorkflowTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetVersion(v string) *GetWorkflowTemplateResponseBody {
	s.Version = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetVersionDescription(v string) *GetWorkflowTemplateResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) SetVersionStatus(v string) *GetWorkflowTemplateResponseBody {
	s.VersionStatus = &v
	return s
}

func (s *GetWorkflowTemplateResponseBody) Validate() error {
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

type GetWorkflowTemplateResponseBodySteps struct {
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
	// The name of the agent owner.
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
	// The order of the step.
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The processing instructions for the step.
	//
	// example:
	//
	// example
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of structured segments for the step instructions. Supported segment types include text, mention, skill, app, files, and web_search. The fields and nested attachments of each segment are preserved exactly as returned by the backend.
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

func (s GetWorkflowTemplateResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTemplateResponseBodySteps) GoString() string {
	return s.String()
}

func (s *GetWorkflowTemplateResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *GetWorkflowTemplateResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *GetWorkflowTemplateResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *GetWorkflowTemplateResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetWorkflowTemplateResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *GetWorkflowTemplateResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *GetWorkflowTemplateResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *GetWorkflowTemplateResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *GetWorkflowTemplateResponseBodySteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *GetWorkflowTemplateResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *GetWorkflowTemplateResponseBodySteps) SetModelMode(v string) *GetWorkflowTemplateResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetName(v string) *GetWorkflowTemplateResponseBodySteps {
	s.Name = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetOwnerAgentName(v string) *GetWorkflowTemplateResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetOwnerType(v string) *GetWorkflowTemplateResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetOwnerUserId(v int64) *GetWorkflowTemplateResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetReviewerUserIds(v []*int64) *GetWorkflowTemplateResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetSortOrder(v int64) *GetWorkflowTemplateResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetSpecification(v string) *GetWorkflowTemplateResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetSpecificationJson(v []map[string]interface{}) *GetWorkflowTemplateResponseBodySteps {
	s.SpecificationJson = v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) SetStepId(v string) *GetWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
