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
}

type CreateWorkflowTemplateResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建人用户ID
	//
	// example:
	//
	// 1
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// 模板描述
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间(ISO8601)
	//
	// example:
	//
	// example
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间(ISO8601)
	//
	// example:
	//
	// example
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 模板名称
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 有序步骤
	//
	// example:
	//
	// example
	Steps []*CreateWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// 模板ID
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
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

func (s *CreateWorkflowTemplateResponseBodySteps) SetStepId(v string) *CreateWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *CreateWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
