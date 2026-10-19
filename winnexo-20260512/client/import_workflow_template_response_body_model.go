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
	SetGmtCreate(v string) *ImportWorkflowTemplateResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *ImportWorkflowTemplateResponseBody
	GetGmtModified() *string
	SetMessage(v string) *ImportWorkflowTemplateResponseBody
	GetMessage() *string
	SetName(v string) *ImportWorkflowTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *ImportWorkflowTemplateResponseBody
	GetRequestId() *string
	SetSteps(v []*ImportWorkflowTemplateResponseBodySteps) *ImportWorkflowTemplateResponseBody
	GetSteps() []*ImportWorkflowTemplateResponseBodySteps
	SetTemplateId(v string) *ImportWorkflowTemplateResponseBody
	GetTemplateId() *string
}

type ImportWorkflowTemplateResponseBody struct {
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
	Steps []*ImportWorkflowTemplateResponseBodySteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
	// 模板ID
	//
	// example:
	//
	// tpl_example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
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

func (s *ImportWorkflowTemplateResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ImportWorkflowTemplateResponseBody) GetGmtModified() *string {
	return s.GmtModified
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

func (s *ImportWorkflowTemplateResponseBody) GetSteps() []*ImportWorkflowTemplateResponseBodySteps {
	return s.Steps
}

func (s *ImportWorkflowTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
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

func (s *ImportWorkflowTemplateResponseBody) SetGmtCreate(v string) *ImportWorkflowTemplateResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetGmtModified(v string) *ImportWorkflowTemplateResponseBody {
	s.GmtModified = &v
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

func (s *ImportWorkflowTemplateResponseBody) SetSteps(v []*ImportWorkflowTemplateResponseBodySteps) *ImportWorkflowTemplateResponseBody {
	s.Steps = v
	return s
}

func (s *ImportWorkflowTemplateResponseBody) SetTemplateId(v string) *ImportWorkflowTemplateResponseBody {
	s.TemplateId = &v
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

func (s *ImportWorkflowTemplateResponseBodySteps) GetReviewerUserIds() []*int64 {
	return s.ReviewerUserIds
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *ImportWorkflowTemplateResponseBodySteps) GetSpecification() *string {
	return s.Specification
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

func (s *ImportWorkflowTemplateResponseBodySteps) SetStepId(v string) *ImportWorkflowTemplateResponseBodySteps {
	s.StepId = &v
	return s
}

func (s *ImportWorkflowTemplateResponseBodySteps) Validate() error {
	return dara.Validate(s)
}
