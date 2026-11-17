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
  // The business status code. A value of 200 indicates success.
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // The error description.
  // 
  // example:
  // 
  // The requested resource does not exist
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // The request tracking ID.
  // 
  // example:
  // 
  // 019FF406-1B10-0065-A97D-2D1920C2A03D
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // The simplified template definition that can be re-imported.
  // 
  // example:
  // 
  // {"name":"Customer Complaint Handling Process","schemaVersion":2}
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
  // The template description.
  // 
  // example:
  // 
  // Standard process for handling customer complaints
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // The node dependency edges. Each item contains from and to fields that reference stable step IDs.
  // 
  // example:
  // 
  // []
  Edges []map[string]interface{} `json:"edges,omitempty" xml:"edges,omitempty" type:"Repeated"`
  // The optional canvas layout, indexed by step ID. This only affects the display and does not determine dependencies.
  // 
  // example:
  // 
  // {"collect":{"x":0,"y":0},"review":{"x":240,"y":0}}
  Layout map[string]interface{} `json:"layout,omitempty" xml:"layout,omitempty"`
  // The template name.
  // 
  // example:
  // 
  // Customer Complaint Handling Process
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // The schema version of the template. The current value is 2, which is independent of the business version vN.
  // 
  // example:
  // 
  // 2
  SchemaVersion *int64 `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
  // The normalized steps.
  // 
  // example:
  // 
  // [{"stepId":"collect","sortOrder":1,"name":"Collect customer information"}]
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

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetEdges() []map[string]interface{}  {
  return s.Edges
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetLayout() map[string]interface{}  {
  return s.Layout
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetName() *string  {
  return s.Name
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetSchemaVersion() *int64  {
  return s.SchemaVersion
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) GetSteps() []*ExportWorkflowTemplateResponseBodyTemplateSteps  {
  return s.Steps
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetDescription(v string) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Description = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetEdges(v []map[string]interface{}) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Edges = v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetLayout(v map[string]interface{}) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Layout = v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetName(v string) *ExportWorkflowTemplateResponseBodyTemplate {
  s.Name = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplate) SetSchemaVersion(v int64) *ExportWorkflowTemplateResponseBodyTemplate {
  s.SchemaVersion = &v
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
  // Specifies whether to manually review the output. Default value: false. This can only be enabled for digital workers, and review tasks are automatically sent to the owner.
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
  // Collect customer information and submit it to the owner for review
  Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
  // An ordered array of segments. Supports text, mention, skill, app, files, and web_search. Preserves reference identities and order.
  // 
  // example:
  // 
  // [{"type":"text","content":"Organize input materials"}]
  SpecificationJson []map[string]interface{} `json:"specificationJson,omitempty" xml:"specificationJson,omitempty" type:"Repeated"`
  // The step ID.
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

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetReviewOutput() *bool  {
  return s.ReviewOutput
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

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) GetSpecificationJson() []map[string]interface{}  {
  return s.SpecificationJson
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

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetReviewOutput(v bool) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.ReviewOutput = &v
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

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetSpecificationJson(v []map[string]interface{}) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.SpecificationJson = v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) SetStepId(v string) *ExportWorkflowTemplateResponseBodyTemplateSteps {
  s.StepId = &v
  return s
}

func (s *ExportWorkflowTemplateResponseBodyTemplateSteps) Validate() error {
  return dara.Validate(s)
}

