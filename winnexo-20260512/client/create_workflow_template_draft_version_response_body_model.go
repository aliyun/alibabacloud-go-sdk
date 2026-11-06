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
	SetEdges(v []map[string]interface{}) *CreateWorkflowTemplateDraftVersionResponseBody
	GetEdges() []map[string]interface{}
	SetEffectiveVersion(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetGmtModified() *string
	SetLayout(v map[string]interface{}) *CreateWorkflowTemplateDraftVersionResponseBody
	GetLayout() map[string]interface{}
	SetMessage(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetMessage() *string
	SetName(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetName() *string
	SetRequestId(v string) *CreateWorkflowTemplateDraftVersionResponseBody
	GetRequestId() *string
	SetSchemaVersion(v int64) *CreateWorkflowTemplateDraftVersionResponseBody
	GetSchemaVersion() *int64
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
	// Standard process for handling customer complaints
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 节点依赖边；每项包含from/to，引用稳定stepId
	//
	// example:
	//
	// []
	Edges []map[string]interface{} `json:"edges,omitempty" xml:"edges,omitempty" type:"Repeated"`
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
	// The name of the step. This parameter can be left empty for a draft, but is required when the template is submitted.
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
	// [{\\"stepId\\":\\"collect\\",\\"sortOrder\\":1,\\"name\\":\\"收集客户信息\\"}]
	Steps []*CreateWorkflowTemplateDraftVersionResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The ID of the template.
	//
	// example:
	//
	// tpl_9401
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

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetEdges() []map[string]interface{} {
	return s.Edges
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

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetLayout() map[string]interface{} {
	return s.Layout
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

func (s *CreateWorkflowTemplateDraftVersionResponseBody) GetSchemaVersion() *int64 {
	return s.SchemaVersion
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

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetEdges(v []map[string]interface{}) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Edges = v
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

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetLayout(v map[string]interface{}) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.Layout = v
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

func (s *CreateWorkflowTemplateDraftVersionResponseBody) SetSchemaVersion(v int64) *CreateWorkflowTemplateDraftVersionResponseBody {
	s.SchemaVersion = &v
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
	// The name of the step. This parameter can be left empty for a draft, but is required when the template is submitted.
	//
	// example:
	//
	// Customer complaint handling process
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the agent owner.
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// The type of the owner. Valid values: MEMBER (manual) and AGENT (agent).
	//
	// example:
	//
	// MEMBER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The user ID of the manual owner.
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
	// The order of the step.
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The display text derived from ordered segments. Fixed titles are not added, and the text is not stored independently.
	//
	// example:
	//
	// Collect customer information and forward it to the owner for review
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of ordered segments. This parameter supports text, mention, skill, app, files, and web_search, and retains reference identities and order.
	//
	// example:
	//
	// [{\\"type\\":\\"text\\",\\"text\\":\\"收集客户信息\\"}]
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

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) GetReviewOutput() *bool {
	return s.ReviewOutput
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

func (s *CreateWorkflowTemplateDraftVersionResponseBodySteps) SetReviewOutput(v bool) *CreateWorkflowTemplateDraftVersionResponseBodySteps {
	s.ReviewOutput = &v
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
