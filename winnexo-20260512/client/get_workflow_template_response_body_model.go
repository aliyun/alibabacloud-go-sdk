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
}

type GetWorkflowTemplateResponseBody struct {
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
	Steps []*GetWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// 模板ID
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
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

func (s *GetWorkflowTemplateResponseBodySteps) SetStepId(v string) *GetWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *GetWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
