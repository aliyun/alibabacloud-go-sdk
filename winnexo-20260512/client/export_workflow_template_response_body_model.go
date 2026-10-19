// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowTemplateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportWorkflowTemplateResponseBody
  GetCode() *string 
  SetMessage(v string) *ExportWorkflowTemplateResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportWorkflowTemplateResponseBody
  GetRequestId() *string 
  SetTemplate(v *ExportWorkflowTemplateResponseBodyTemplate) *ExportWorkflowTemplateResponseBody
  GetTemplate() *ExportWorkflowTemplateResponseBodyTemplate 
}

type ExportWorkflowTemplateResponseBody struct {
  // 业务状态码；成功为200
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
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
  // 可回流导入的精简模板定义
  // 
  // example:
  // 
  // example
  Template *ExportWorkflowTemplateResponseBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
}

func (s ExportWorkflowTemplateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowTemplateResponseBody) GoString() string {
  return s.String()
}

func (s *ExportWorkflowTemplateResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportWorkflowTemplateResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportWorkflowTemplateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportWorkflowTemplateResponseBody) GetTemplate() *ExportWorkflowTemplateResponseBodyTemplate  {
  return s.Template
}

func (s *ExportWorkflowTemplateResponseBody) SetCode(v string) *ExportWorkflowTemplateResponseBody {
  s.Code = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBody) SetMessage(v string) *ExportWorkflowTemplateResponseBody {
  s.Message = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBody) SetRequestId(v string) *ExportWorkflowTemplateResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBody) SetTemplate(v *ExportWorkflowTemplateResponseBodyTemplate) *ExportWorkflowTemplateResponseBody {
  s.Template = v
  return s
}

func (s *ExportWorkflowTemplateResponseBody) Validate() error {
  if s.Template != nil {
    if err := s.Template.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportWorkflowTemplateResponseBodyTemplate struct {
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
  // 规范化步骤
  // 
  // example:
  // 
  // example
  Steps []*ExportWorkflowTemplateResponseBodyTemplateSteps `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s ExportWorkflowTemplateResponseBodyTemplate) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowTemplateResponseBodyTemplate) GoString() string {
  return s.String()
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetDescription() *string  {
  return s.Description
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetName() *string  {
  return s.Name
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetSteps() []*ExportWorkflowTemplateResponseBodyTemplateSteps  {
  return s.Steps
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetDescription(v string) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Description = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetName(v string) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Name = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetSteps(v []*ExportWorkflowTemplateResponseBodyTemplateSteps) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Steps = v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) Validate() error {
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

type ExportWorkflowTemplateResponseBodyTemplateSteps struct {
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

func (s ExportWorkflowTemplateResponseBodyTemplateSteps) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowTemplateResponseBodyTemplateSteps) GoString() string {
  return s.String()
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetModelMode() *string  {
  return s.ModelMode
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetName() *string  {
  return s.Name
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetOwnerAgentName() *string  {
  return s.OwnerAgentName
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetOwnerType() *string  {
  return s.OwnerType
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetOwnerUserId() *int64  {
  return s.OwnerUserId
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetReviewerUserIds() []*int64  {
  return s.ReviewerUserIds
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetSortOrder() *int64  {
  return s.SortOrder
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetSpecification() *string  {
  return s.Specification
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetStepId() *string  {
  return s.StepId
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetModelMode(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.ModelMode = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetName(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.Name = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetOwnerAgentName(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.OwnerAgentName = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetOwnerType(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.OwnerType = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetOwnerUserId(v int64) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.OwnerUserId = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetReviewerUserIds(v []*int64) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.ReviewerUserIds = v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetSortOrder(v int64) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.SortOrder = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetSpecification(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.Specification = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetStepId(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.StepId = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) Validate() error {
  return dara.Validate(s)
}

