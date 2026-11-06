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
	SetEdges(v []map[string]interface{}) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetEdges() []map[string]interface{}
	SetEffectiveVersion(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetGmtModified() *string
	SetLayout(v map[string]interface{}) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetLayout() map[string]interface{}
	SetMessage(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetMessage() *string
	SetName(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetName() *string
	SetRequestId(v string) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetRequestId() *string
	SetSchemaVersion(v int64) *SetWorkflowTemplateEffectiveVersionResponseBody
	GetSchemaVersion() *int64
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
	// 10001
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// The template description.
	//
	// example:
	//
	// Standard process for handling customer complaints
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 节点依赖边；每项包含from/to，引用稳定stepId
	//
	// example:
	//
	// []
	Edges []map[string]interface{} `json:"edges,omitempty" xml:"edges,omitempty" type:"Repeated"`
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
	// 2026-11-05T08:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 可选画布布局，按stepId索引；只影响展示，不决定依赖
	//
	// example:
	//
	// {"collect":{"x":0,"y":0},"review":{"x":240,"y":0}}
	Layout map[string]interface{} `json:"layout,omitempty" xml:"layout,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The step name. This field can be empty for drafts but is required upon submission.
	//
	// example:
	//
	// Customer complaint handling process
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 模板结构版本，当前为2，与业务版本vN独立
	//
	// example:
	//
	// 2
	SchemaVersion *int64 `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The ordered steps.
	//
	// example:
	//
	// [{"stepId":"collect","sortOrder":1,"name":"收集客户信息"}]
	Steps []*SetWorkflowTemplateEffectiveVersionResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The template ID.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The version number of the returned content.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The change description for this version. An empty value indicates that no description is provided.
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetEdges() []map[string]interface{} {
	return s.Edges
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetLayout() map[string]interface{} {
	return s.Layout
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) GetSchemaVersion() *int64 {
	return s.SchemaVersion
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetEdges(v []map[string]interface{}) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Edges = v
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetLayout(v map[string]interface{}) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.Layout = v
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBody) SetSchemaVersion(v int64) *SetWorkflowTemplateEffectiveVersionResponseBody {
	s.SchemaVersion = &v
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
	// The step name. This field can be empty for drafts but is required upon submission.
	//
	// example:
	//
	// Collect customer information
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
	// 是否人工审核产出物，默认关闭；仅数字员工可开启，审核待办自动发负责人
	//
	// example:
	//
	// false
	ReviewOutput *bool `json:"reviewOutput,omitempty" xml:"reviewOutput,omitempty"`
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
	// The display text derived from ordered segments. Fixed titles are not added, and the text is not stored independently.
	//
	// example:
	//
	// Collect customer information and submit it to the person in charge for review
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of ordered segments. Supports text, mention, skill, app, files, and web_search. Retains reference identities and order.
	//
	// example:
	//
	// [{"type":"text","text":"收集客户信息"}]
	SpecificationJson []map[string]interface{} `json:"specificationJson,omitempty" xml:"specificationJson,omitempty" type:"Repeated"`
	// The step ID.
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) GetReviewOutput() *bool {
	return s.ReviewOutput
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

func (s *SetWorkflowTemplateEffectiveVersionResponseBodySteps) SetReviewOutput(v bool) *SetWorkflowTemplateEffectiveVersionResponseBodySteps {
	s.ReviewOutput = &v
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
