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
	SetEdges(v []map[string]interface{}) *GetWorkflowTemplateResponseBody
	GetEdges() []map[string]interface{}
	SetEffectiveVersion(v string) *GetWorkflowTemplateResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *GetWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetLayout(v map[string]interface{}) *GetWorkflowTemplateResponseBody
	GetLayout() map[string]interface{}
	SetMessage(v string) *GetWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *GetWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *GetWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSchemaVersion(v int64) *GetWorkflowTemplateResponseBody
	GetSchemaVersion() *int64
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
	// Standard process for handling customer complaints
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 节点依赖边；每项包含from/to，引用稳定stepId；显式空数组表示没有依赖
	//
	// example:
	//
	// [{"from":"collect","to":"deliver"}]
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
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Customer complaint handling process
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
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
	Steps []*GetWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The ID of the template.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The optimistic version number.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version description.
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

func (s *GetWorkflowTemplateResponseBody) GetEdges() []map[string]interface{} {
	return s.Edges
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

func (s *GetWorkflowTemplateResponseBody) GetLayout() map[string]interface{} {
	return s.Layout
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

func (s *GetWorkflowTemplateResponseBody) GetSchemaVersion() *int64 {
	return s.SchemaVersion
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

func (s *GetWorkflowTemplateResponseBody) SetEdges(v []map[string]interface{}) *GetWorkflowTemplateResponseBody {
	s.Edges = v
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

func (s *GetWorkflowTemplateResponseBody) SetLayout(v map[string]interface{}) *GetWorkflowTemplateResponseBody {
	s.Layout = v
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

func (s *GetWorkflowTemplateResponseBody) SetSchemaVersion(v int64) *GetWorkflowTemplateResponseBody {
	s.SchemaVersion = &v
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
	// Customer complaint handling process
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
	// The sort order of the step.
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The processing instructions for the step.
	//
	// example:
	//
	// Collect customer information and submit it to the owner for review
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The array of structured segments for the step instructions. Supported types include text, mention, skill, app, files, and web_search. The fields and nested attachments of each segment are preserved exactly as returned by the backend.
	//
	// example:
	//
	// [{"type":"text","content":"整理输入资料"}]
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

func (s *GetWorkflowTemplateResponseBodySteps) GetReviewOutput() *bool {
	return s.ReviewOutput
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

func (s *GetWorkflowTemplateResponseBodySteps) SetReviewOutput(v bool) *GetWorkflowTemplateResponseBodySteps {
	s.ReviewOutput = &v
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
