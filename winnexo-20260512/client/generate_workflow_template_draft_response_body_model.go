// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWorkflowTemplateDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateWorkflowTemplateDraftResponseBody
	GetCode() *string
	SetDraft(v *GenerateWorkflowTemplateDraftResponseBodyDraft) *GenerateWorkflowTemplateDraftResponseBody
	GetDraft() *GenerateWorkflowTemplateDraftResponseBodyDraft
	SetMessage(v string) *GenerateWorkflowTemplateDraftResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateWorkflowTemplateDraftResponseBody
	GetRequestId() *string
}

type GenerateWorkflowTemplateDraftResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The generated non-persistent template draft.
	//
	// example:
	//
	// {\\"schemaVersion\\":2,\\"steps\\":[]}
	Draft *GenerateWorkflowTemplateDraftResponseBodyDraft `json:"draft,omitempty" xml:"draft,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateWorkflowTemplateDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateWorkflowTemplateDraftResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWorkflowTemplateDraftResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateWorkflowTemplateDraftResponseBody) GetDraft() *GenerateWorkflowTemplateDraftResponseBodyDraft {
	return s.Draft
}

func (s *GenerateWorkflowTemplateDraftResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateWorkflowTemplateDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateWorkflowTemplateDraftResponseBody) SetCode(v string) *GenerateWorkflowTemplateDraftResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBody) SetDraft(v *GenerateWorkflowTemplateDraftResponseBodyDraft) *GenerateWorkflowTemplateDraftResponseBody {
	s.Draft = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBody) SetMessage(v string) *GenerateWorkflowTemplateDraftResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBody) SetRequestId(v string) *GenerateWorkflowTemplateDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBody) Validate() error {
	if s.Draft != nil {
		if err := s.Draft.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateWorkflowTemplateDraftResponseBodyDraft struct {
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
	// 可选画布布局，按stepId索引；只影响展示，不决定依赖
	//
	// example:
	//
	// {"collect":{"x":0,"y":0},"review":{"x":240,"y":0}}
	Layout map[string]interface{} `json:"layout,omitempty" xml:"layout,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Customer complaint handling process
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Steps []*GenerateWorkflowTemplateDraftResponseBodyDraftSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s GenerateWorkflowTemplateDraftResponseBodyDraft) String() string {
	return dara.Prettify(s)
}

func (s GenerateWorkflowTemplateDraftResponseBodyDraft) GoString() string {
	return s.String()
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetDescription() *string {
	return s.Description
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetEdges() []map[string]interface{} {
	return s.Edges
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetLayout() map[string]interface{} {
	return s.Layout
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetName() *string {
	return s.Name
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetSchemaVersion() *int64 {
	return s.SchemaVersion
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetSteps() []*GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	return s.Steps
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetDescription(v string) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Description = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetEdges(v []map[string]interface{}) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Edges = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetLayout(v map[string]interface{}) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Layout = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetName(v string) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Name = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetSchemaVersion(v int64) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.SchemaVersion = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetSteps(v []*GenerateWorkflowTemplateDraftResponseBodyDraftSteps) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Steps = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) Validate() error {
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

type GenerateWorkflowTemplateDraftResponseBodyDraftSteps struct {
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
	// The type of the owner. Valid values:
	//
	// - MEMBER: human
	//
	// - AGENT: agent
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
	// The processing instructions of the step.
	//
	// example:
	//
	// Collect customer information and submit it to the owner for review
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// The ordered array of fragments. Supported types include text, mention, skill, app, files, and web_search. The referenced identities and order are preserved.
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

func (s GenerateWorkflowTemplateDraftResponseBodyDraftSteps) String() string {
	return dara.Prettify(s)
}

func (s GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GoString() string {
	return s.String()
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetName() *string {
	return s.Name
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetReviewOutput() *bool {
	return s.ReviewOutput
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetSpecification() *string {
	return s.Specification
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetSpecificationJson() []map[string]interface{} {
	return s.SpecificationJson
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetStepId() *string {
	return s.StepId
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetModelMode(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.ModelMode = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetName(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.Name = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetOwnerAgentName(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.OwnerAgentName = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetOwnerType(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.OwnerType = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetOwnerUserId(v int64) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.OwnerUserId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetReviewOutput(v bool) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.ReviewOutput = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetReviewerUserIds(v []*int64) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.ReviewerUserIds = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetSortOrder(v int64) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.SortOrder = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetSpecification(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.Specification = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetSpecificationJson(v []map[string]interface{}) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.SpecificationJson = v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetStepId(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.StepId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) Validate() error {
	return dara.Validate(s)
}
