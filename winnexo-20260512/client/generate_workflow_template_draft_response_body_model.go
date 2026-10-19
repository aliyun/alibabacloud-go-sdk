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
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 生成的非持久化模板草稿
	//
	// example:
	//
	// example
	Draft *GenerateWorkflowTemplateDraftResponseBodyDraft `json:"draft,omitempty" xml:"draft,omitempty" type:"Struct"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
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
	// 模板描述
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 模板名称
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 有序步骤
	//
	// example:
	//
	// example
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

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetName() *string {
	return s.Name
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) GetSteps() []*GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	return s.Steps
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetDescription(v string) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Description = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraft) SetName(v string) *GenerateWorkflowTemplateDraftResponseBodyDraft {
	s.Name = &v
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
	// 模型档位
	//
	// example:
	//
	// standard
	ModelMode *string `json:"modelMode,omitempty" xml:"modelMode,omitempty"`
	// 步骤名称
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 智能体负责人名称
	//
	// example:
	//
	// example
	OwnerAgentName *string `json:"ownerAgentName,omitempty" xml:"ownerAgentName,omitempty"`
	// 负责人类型：MEMBER 人工 / AGENT 智能体
	//
	// example:
	//
	// MEMBER
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// 人工负责人用户ID
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// 审核人用户ID列表
	//
	// example:
	//
	// example
	ReviewerUserIds []*int64 `json:"reviewerUserIds,omitempty" xml:"reviewerUserIds,omitempty" type:"Repeated"`
	// 步骤顺序
	//
	// example:
	//
	// 1
	SortOrder *int64 `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// 步骤处理说明
	//
	// example:
	//
	// example
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 步骤ID
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

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) GetSpecification() *string {
	return s.Specification
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

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) SetStepId(v string) *GenerateWorkflowTemplateDraftResponseBodyDraftSteps {
	s.StepId = &v
	return s
}

func (s *GenerateWorkflowTemplateDraftResponseBodyDraftSteps) Validate() error {
	return dara.Validate(s)
}
