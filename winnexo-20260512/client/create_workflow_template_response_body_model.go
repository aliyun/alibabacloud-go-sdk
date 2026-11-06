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
	SetEdges(v []map[string]interface{}) *CreateWorkflowTemplateResponseBody
	GetEdges() []map[string]interface{}
	SetEffectiveVersion(v string) *CreateWorkflowTemplateResponseBody
	GetEffectiveVersion() *string
	SetGmtCreate(v string) *CreateWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *CreateWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetLayout(v map[string]interface{}) *CreateWorkflowTemplateResponseBody
	GetLayout() map[string]interface{}
	SetMessage(v string) *CreateWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *CreateWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *CreateWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSchemaVersion(v int64) *CreateWorkflowTemplateResponseBody
	GetSchemaVersion() *int64
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
	// Standard process for handling customer complaints
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 节点依赖边；每项包含from/to，引用稳定stepId；显式空数组表示没有依赖
	//
	// example:
	//
	// [{"from":"collect","to":"deliver"}]
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
	// The template name.
	//
	// example:
	//
	// Customer Complaint Handling Process
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
	Steps []*CreateWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
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
	// The change description of the version. If this field is empty for v1, the template description can be displayed.
	//
	// example:
	//
	// Initial template creation
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

func (s *CreateWorkflowTemplateResponseBody) GetEdges() []map[string]interface{} {
	return s.Edges
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

func (s *CreateWorkflowTemplateResponseBody) GetLayout() map[string]interface{} {
	return s.Layout
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

func (s *CreateWorkflowTemplateResponseBody) GetSchemaVersion() *int64 {
	return s.SchemaVersion
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

func (s *CreateWorkflowTemplateResponseBody) SetEdges(v []map[string]interface{}) *CreateWorkflowTemplateResponseBody {
	s.Edges = v
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

func (s *CreateWorkflowTemplateResponseBody) SetLayout(v map[string]interface{}) *CreateWorkflowTemplateResponseBody {
	s.Layout = v
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

func (s *CreateWorkflowTemplateResponseBody) SetSchemaVersion(v int64) *CreateWorkflowTemplateResponseBody {
	s.SchemaVersion = &v
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
	// Customer Complaint Handling Process
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the agent in charge.
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// The type of the person in charge. Valid values:
	//
	// - MEMBER: Human.
	//
	// - AGENT: Agent.
	//
	// example:
	//
	// MEMBER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The user ID of the human person in charge.
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
	// [10001]
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
	// The ordered array of fragments. Supported types include text, mention, skill, app, files, and web_search. The referenced identities and order are preserved.
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

func (s *CreateWorkflowTemplateResponseBodySteps) GetReviewOutput() *bool {
	return s.ReviewOutput
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

func (s *CreateWorkflowTemplateResponseBodySteps) SetReviewOutput(v bool) *CreateWorkflowTemplateResponseBodySteps {
	s.ReviewOutput = &v
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
