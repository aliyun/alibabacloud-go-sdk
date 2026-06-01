// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *UpdateDigitalEmployeeRequest
	GetAttributes() map[string]*string
	SetDefaultRule(v string) *UpdateDigitalEmployeeRequest
	GetDefaultRule() *string
	SetDescription(v string) *UpdateDigitalEmployeeRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateDigitalEmployeeRequest
	GetDisplayName() *string
	SetKnowledges(v *UpdateDigitalEmployeeRequestKnowledges) *UpdateDigitalEmployeeRequest
	GetKnowledges() *UpdateDigitalEmployeeRequestKnowledges
	SetRoleArn(v string) *UpdateDigitalEmployeeRequest
	GetRoleArn() *string
	SetToolPolicy(v *UpdateDigitalEmployeeRequestToolPolicy) *UpdateDigitalEmployeeRequest
	GetToolPolicy() *UpdateDigitalEmployeeRequestToolPolicy
}

type UpdateDigitalEmployeeRequest struct {
	Attributes map[string]*string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// test
	DefaultRule *string `json:"defaultRule,omitempty" xml:"defaultRule,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	DisplayName *string                                 `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Knowledges  *UpdateDigitalEmployeeRequestKnowledges `json:"knowledges,omitempty" xml:"knowledges,omitempty" type:"Struct"`
	// example:
	//
	// acs:ram::12345678912:role/testrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// 数字员工工具调用安全策略配置。
	//
	// example:
	//
	// {"aliyun":{"enable":true,"statements":[{"decision":"user_ack","product":"Sls","apiVersion":"2020-12-30","actions":["log:GetProject","log:CreateDashboard"]}]}}
	ToolPolicy *UpdateDigitalEmployeeRequestToolPolicy `json:"toolPolicy,omitempty" xml:"toolPolicy,omitempty" type:"Struct"`
}

func (s UpdateDigitalEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequest) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *UpdateDigitalEmployeeRequest) GetDefaultRule() *string {
	return s.DefaultRule
}

func (s *UpdateDigitalEmployeeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDigitalEmployeeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateDigitalEmployeeRequest) GetKnowledges() *UpdateDigitalEmployeeRequestKnowledges {
	return s.Knowledges
}

func (s *UpdateDigitalEmployeeRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateDigitalEmployeeRequest) GetToolPolicy() *UpdateDigitalEmployeeRequestToolPolicy {
	return s.ToolPolicy
}

func (s *UpdateDigitalEmployeeRequest) SetAttributes(v map[string]*string) *UpdateDigitalEmployeeRequest {
	s.Attributes = v
	return s
}

func (s *UpdateDigitalEmployeeRequest) SetDefaultRule(v string) *UpdateDigitalEmployeeRequest {
	s.DefaultRule = &v
	return s
}

func (s *UpdateDigitalEmployeeRequest) SetDescription(v string) *UpdateDigitalEmployeeRequest {
	s.Description = &v
	return s
}

func (s *UpdateDigitalEmployeeRequest) SetDisplayName(v string) *UpdateDigitalEmployeeRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateDigitalEmployeeRequest) SetKnowledges(v *UpdateDigitalEmployeeRequestKnowledges) *UpdateDigitalEmployeeRequest {
	s.Knowledges = v
	return s
}

func (s *UpdateDigitalEmployeeRequest) SetRoleArn(v string) *UpdateDigitalEmployeeRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateDigitalEmployeeRequest) SetToolPolicy(v *UpdateDigitalEmployeeRequestToolPolicy) *UpdateDigitalEmployeeRequest {
	s.ToolPolicy = v
	return s
}

func (s *UpdateDigitalEmployeeRequest) Validate() error {
	if s.Knowledges != nil {
		if err := s.Knowledges.Validate(); err != nil {
			return err
		}
	}
	if s.ToolPolicy != nil {
		if err := s.ToolPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDigitalEmployeeRequestKnowledges struct {
	Bailian []*UpdateDigitalEmployeeRequestKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
	Sop     []map[string]interface{}                         `json:"sop,omitempty" xml:"sop,omitempty" type:"Repeated"`
}

func (s UpdateDigitalEmployeeRequestKnowledges) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequestKnowledges) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequestKnowledges) GetBailian() []*UpdateDigitalEmployeeRequestKnowledgesBailian {
	return s.Bailian
}

func (s *UpdateDigitalEmployeeRequestKnowledges) GetSop() []map[string]interface{} {
	return s.Sop
}

func (s *UpdateDigitalEmployeeRequestKnowledges) SetBailian(v []*UpdateDigitalEmployeeRequestKnowledgesBailian) *UpdateDigitalEmployeeRequestKnowledges {
	s.Bailian = v
	return s
}

func (s *UpdateDigitalEmployeeRequestKnowledges) SetSop(v []map[string]interface{}) *UpdateDigitalEmployeeRequestKnowledges {
	s.Sop = v
	return s
}

func (s *UpdateDigitalEmployeeRequestKnowledges) Validate() error {
	if s.Bailian != nil {
		for _, item := range s.Bailian {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDigitalEmployeeRequestKnowledgesBailian struct {
	// example:
	//
	// test
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// index-xxxx
	IndexId *string `json:"indexId,omitempty" xml:"indexId,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateDigitalEmployeeRequestKnowledgesBailian) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequestKnowledgesBailian) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) GetAttributes() *string {
	return s.Attributes
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) GetIndexId() *string {
	return s.IndexId
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) GetRegion() *string {
	return s.Region
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) SetAttributes(v string) *UpdateDigitalEmployeeRequestKnowledgesBailian {
	s.Attributes = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) SetIndexId(v string) *UpdateDigitalEmployeeRequestKnowledgesBailian {
	s.IndexId = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) SetRegion(v string) *UpdateDigitalEmployeeRequestKnowledgesBailian {
	s.Region = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) SetWorkspaceId(v string) *UpdateDigitalEmployeeRequestKnowledgesBailian {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestKnowledgesBailian) Validate() error {
	return dara.Validate(s)
}

type UpdateDigitalEmployeeRequestToolPolicy struct {
	// Aliyun MCP 工具调用安全策略配置。
	//
	// example:
	//
	// {"enable":true,"statements":[{"decision":"user_ack","product":"Sls","apiVersion":"2020-12-30","actions":["log:GetProject","log:CreateDashboard"]}]}
	Aliyun *UpdateDigitalEmployeeRequestToolPolicyAliyun `json:"aliyun,omitempty" xml:"aliyun,omitempty" type:"Struct"`
}

func (s UpdateDigitalEmployeeRequestToolPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequestToolPolicy) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequestToolPolicy) GetAliyun() *UpdateDigitalEmployeeRequestToolPolicyAliyun {
	return s.Aliyun
}

func (s *UpdateDigitalEmployeeRequestToolPolicy) SetAliyun(v *UpdateDigitalEmployeeRequestToolPolicyAliyun) *UpdateDigitalEmployeeRequestToolPolicy {
	s.Aliyun = v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicy) Validate() error {
	if s.Aliyun != nil {
		if err := s.Aliyun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDigitalEmployeeRequestToolPolicyAliyun struct {
	// 是否启用 Aliyun MCP 工具策略。
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Aliyun OpenAPI 工具策略语句列表。
	//
	// example:
	//
	// [{"decision":"user_ack","product":"Sls","apiVersion":"2020-12-30","actions":["log:GetProject","log:CreateDashboard"]}]
	Statements []*UpdateDigitalEmployeeRequestToolPolicyAliyunStatements `json:"statements,omitempty" xml:"statements,omitempty" type:"Repeated"`
}

func (s UpdateDigitalEmployeeRequestToolPolicyAliyun) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequestToolPolicyAliyun) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyun) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyun) GetStatements() []*UpdateDigitalEmployeeRequestToolPolicyAliyunStatements {
	return s.Statements
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyun) SetEnable(v bool) *UpdateDigitalEmployeeRequestToolPolicyAliyun {
	s.Enable = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyun) SetStatements(v []*UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) *UpdateDigitalEmployeeRequestToolPolicyAliyun {
	s.Statements = v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyun) Validate() error {
	if s.Statements != nil {
		for _, item := range s.Statements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDigitalEmployeeRequestToolPolicyAliyunStatements struct {
	// Aliyun OpenAPI Action 列表，格式为 product:ApiName、product:Prefix	- 或 product:*。
	//
	// example:
	//
	// ["log:GetProject","log:CreateDashboard"]
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 本条语句对应的 Aliyun OpenAPI API 版本。
	//
	// example:
	//
	// 2020-12-30
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// 命中该 API 后的执行策略。
	//
	// example:
	//
	// user_ack
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// 本条语句对应的 Aliyun OpenAPI 产品名。
	//
	// example:
	//
	// Sls
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) GetActions() []*string {
	return s.Actions
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) GetDecision() *string {
	return s.Decision
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) GetProduct() *string {
	return s.Product
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) SetActions(v []*string) *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.Actions = v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) SetApiVersion(v string) *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.ApiVersion = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) SetDecision(v string) *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.Decision = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) SetProduct(v string) *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.Product = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestToolPolicyAliyunStatements) Validate() error {
	return dara.Validate(s)
}
