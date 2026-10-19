// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWorkflowTemplateResponseBody
	GetCode() *string
	SetCreatedByUserId(v int64) *UpdateWorkflowTemplateResponseBody
	GetCreatedByUserId() *int64
	SetDescription(v string) *UpdateWorkflowTemplateResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *UpdateWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *UpdateWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetMessage(v string) *UpdateWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *UpdateWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSteps(v []*UpdateWorkflowTemplateResponseBodySteps) *UpdateWorkflowTemplateResponseBody
	GetSteps() []*UpdateWorkflowTemplateResponseBodySteps
	SetTemplateId(v string) *UpdateWorkflowTemplateResponseBody
	GetTemplateId() *string
}

type UpdateWorkflowTemplateResponseBody struct {
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
	Steps []*UpdateWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// 模板ID
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UpdateWorkflowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWorkflowTemplateResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *UpdateWorkflowTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkflowTemplateResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateWorkflowTemplateResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateWorkflowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowTemplateResponseBody) GetSteps() []*UpdateWorkflowTemplateResponseBodySteps {
	return s.Steps
}

func (s *UpdateWorkflowTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateResponseBody) SetCode(v string) *UpdateWorkflowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetCreatedByUserId(v int64) *UpdateWorkflowTemplateResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetDescription(v string) *UpdateWorkflowTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetGmtCreate(v string) *UpdateWorkflowTemplateResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetGmtModified(v string) *UpdateWorkflowTemplateResponseBody {
	s.GmtModified = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetMessage(v string) *UpdateWorkflowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetName(v string) *UpdateWorkflowTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetRequestId(v string) *UpdateWorkflowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetSteps(v []*UpdateWorkflowTemplateResponseBodySteps) *UpdateWorkflowTemplateResponseBody {
	s.Steps = v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) SetTemplateId(v string) *UpdateWorkflowTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBody) Validate() error {
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

type UpdateWorkflowTemplateResponseBodySteps struct {
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

func (s UpdateWorkflowTemplateResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateResponseBodySteps) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetModelMode() *string {
	return s.ModelMode
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetName() *string {
	return s.Name
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetOwnerAgentName() *string {
	return s.OwnerAgentName
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetOwnerType() *string {
	return s.OwnerType
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetSpecification() *string {
	return s.Specification
}

func (s *UpdateWorkflowTemplateResponseBodySteps) GetStepId() *string {
	return s.StepId
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetModelMode(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.ModelMode = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetName(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetOwnerAgentName(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.OwnerAgentName = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetOwnerType(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.OwnerType = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetOwnerUserId(v int64) *UpdateWorkflowTemplateResponseBodySteps {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetReviewerUserIds(v []*int64) *UpdateWorkflowTemplateResponseBodySteps {
	s.ReviewerUserIds = v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetSortOrder(v int64) *UpdateWorkflowTemplateResponseBodySteps {
	s.SortOrder = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetSpecification(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.Specification = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) SetStepId(v string) *UpdateWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *UpdateWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
