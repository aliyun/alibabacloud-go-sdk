// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportWorkflowTemplateResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *ImportWorkflowTemplateResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *ImportWorkflowTemplateResponseBody
	GetDescription() *string
	SetEdges(v []map[string]interface{}) *ImportWorkflowTemplateResponseBody
	GetEdges() []map[string]interface{}
	SetEffectiveVersion(v string) *ImportWorkflowTemplateResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *ImportWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *ImportWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetLayout(v map[string]interface{}) *ImportWorkflowTemplateResponseBody
	GetLayout() map[string]interface{}
	SetMessage(v string) *ImportWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *ImportWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *ImportWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSchemaVersion(v int64) *ImportWorkflowTemplateResponseBody
	GetSchemaVersion() *int64
	SetSteps(v []*ImportWorkflowTemplateResponseBodySteps) *ImportWorkflowTemplateResponseBody
	GetSteps() []*ImportWorkflowTemplateResponseBodySteps
	SetTemplateId(v string) *ImportWorkflowTemplateResponseBody
	GetTemplateId() *string
	SetVersion(v string) *ImportWorkflowTemplateResponseBody
	GetVersion() *string
	SetVersionStatus(v string) *ImportWorkflowTemplateResponseBody
	GetVersionStatus() *string
}

type ImportWorkflowTemplateResponseBody struct {
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
	// Standard process for handling customer complaints
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 节点依赖边；每项包含from/to，引用稳定stepId；显式空数组表示没有依赖
	//
	// example:
	//
	// [{"from":"collect","to":"deliver"}]
	Edges []map[string]interface{} `json:"edges,omitempty" xml:"edges,omitempty" type:"Repeated"`
	// The current effective version number of the template. An empty value indicates that only a draft exists and it cannot be used for work items. For imported templates, a v1 draft is created and this field is empty.
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
	// The template name.
	//
	// example:
	//
	// Customer Complaint Handling Process
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
	Steps []*ImportWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// The template ID.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The version number generated for this import.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version status. Valid values:
	//
	// - DRAFT: draft.
	//
	// - SUBMITTED: submitted.
	//
	// The value is fixed to DRAFT for imported templates.
	//
	// example:
	//
	// DRAFT
	VersionStatus *string `json:"versionStatus,omitempty" xml:"versionStatus,omitempty"`
}

func (s ImportWorkflowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportWorkflowTemplateResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *ImportWorkflowTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ImportWorkflowTemplateResponseBody) GetEdges() []map[string]interface{} {
	return s.Edges
}

func (s *ImportWorkflowTemplateResponseBody) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *ImportWorkflowTemplateResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ImportWorkflowTemplateResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ImportWorkflowTemplateResponseBody) GetLayout() map[string]interface{} {
	return s.Layout
}

func (s *ImportWorkflowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportWorkflowTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *ImportWorkflowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportWorkflowTemplateResponseBody) GetSchemaVersion() *int64 {
	return s.SchemaVersion
}

func (s *ImportWorkflowTemplateResponseBody) GetSteps() []*ImportWorkflowTemplateResponseBodySteps {
	return s.Steps
}

func (s *ImportWorkflowTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ImportWorkflowTemplateResponseBody) GetVersion() *string {
	return s.Version
}

func (s *ImportWorkflowTemplateResponseBody) GetVersionStatus() *string {
	return s.VersionStatus
}

func (s *ImportWorkflowTemplateResponseBody) SetCode(v string) *ImportWorkflowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetCreatedByUserId(v int64) *ImportWorkflowTemplateResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetDescription(v string) *ImportWorkflowTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetEdges(v []map[string]interface{}) *ImportWorkflowTemplateResponseBody {
	s.Edges = v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetEffectiveVersion(v string) *ImportWorkflowTemplateResponseBody {
	s.EffectiveVersion = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetGmtCreate(v string) *ImportWorkflowTemplateResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetGmtModified(v string) *ImportWorkflowTemplateResponseBody {
	s.GmtModified = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetLayout(v map[string]interface{}) *ImportWorkflowTemplateResponseBody {
	s.Layout = v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetMessage(v string) *ImportWorkflowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetName(v string) *ImportWorkflowTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetRequestId(v string) *ImportWorkflowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetSchemaVersion(v int64) *ImportWorkflowTemplateResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetSteps(v []*ImportWorkflowTemplateResponseBodySteps) *ImportWorkflowTemplateResponseBody {
	s.Steps = v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetTemplateId(v string) *ImportWorkflowTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetVersion(v string) *ImportWorkflowTemplateResponseBody {
	s.Version = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetVersionStatus(v string) *ImportWorkflowTemplateResponseBody {
	s.VersionStatus = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) Validate() error {
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

type ImportWorkflowTemplateResponseBodySteps struct {
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
	// Customer Complaint Handling Process
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee executor.
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// The executor type. Valid values:
	//
	// - MEMBER: workspace member.
	//
	// - AGENT: digital employee.
	//
	// example:
	//
	// MEMBER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The user ID of the employee executor.
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
	// The list of user IDs of the persons in charge. This field is valid only for steps with a digital employee executor and is always empty for steps with a member executor.
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
	// Collect customer information and submit it to the person in charge for review
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The ordered array of fragments, which supports text, mention, skill, app, files, and web_search. The referenced identities and order are retained.
	//
	// example:
	//
	// [{\\"type\\":\\"text\\",\\"text\\":\\"收集客户信息\\"}]
	SpecificationJson []map[string]interface{} `json:"specificationJson,omitempty" xml:"specificationJson,omitempty" type:"Repeated"`
	// The step ID.
	//
	// example:
	//
	// collect
	StepId *string `json:"stepId,omitempty" xml:"stepId,omitempty"`
}

func (s ImportWorkflowTemplateResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateResponseBodySteps) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetReviewOutput() *bool {
	return s.ReviewOutput
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetModelMode(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetName(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.Name = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetOwnerAgentName(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetOwnerType(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetOwnerUserId(v int64) *ImportWorkflowTemplateResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetReviewOutput(v bool) *ImportWorkflowTemplateResponseBodySteps {
	s.ReviewOutput = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetReviewerUserIds(v []*int64) *ImportWorkflowTemplateResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetSortOrder(v int64) *ImportWorkflowTemplateResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetSpecification(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetSpecificationJson(v []map[string]interface{}) *ImportWorkflowTemplateResponseBodySteps {
	s.SpecificationJson = v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) SetStepId(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
