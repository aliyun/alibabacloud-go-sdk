// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *CreateDigitalEmployeeRequest
	GetAttributes() map[string]*string
	SetDefaultRule(v string) *CreateDigitalEmployeeRequest
	GetDefaultRule() *string
	SetDescription(v string) *CreateDigitalEmployeeRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateDigitalEmployeeRequest
	GetDisplayName() *string
	SetKnowledges(v *CreateDigitalEmployeeRequestKnowledges) *CreateDigitalEmployeeRequest
	GetKnowledges() *CreateDigitalEmployeeRequestKnowledges
	SetName(v string) *CreateDigitalEmployeeRequest
	GetName() *string
	SetResourceGroupId(v string) *CreateDigitalEmployeeRequest
	GetResourceGroupId() *string
	SetRoleArn(v string) *CreateDigitalEmployeeRequest
	GetRoleArn() *string
	SetTags(v []*Tag) *CreateDigitalEmployeeRequest
	GetTags() []*Tag
	SetToolPolicy(v *CreateDigitalEmployeeRequestToolPolicy) *CreateDigitalEmployeeRequest
	GetToolPolicy() *CreateDigitalEmployeeRequestToolPolicy
}

type CreateDigitalEmployeeRequest struct {
	Attributes map[string]*string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// test
	DefaultRule *string `json:"defaultRule,omitempty" xml:"defaultRule,omitempty"`
	// example:
	//
	// aaa
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// digial-employee-test
	DisplayName *string                                 `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Knowledges  *CreateDigitalEmployeeRequestKnowledges `json:"knowledges,omitempty" xml:"knowledges,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::12345678912:role/testrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Tags    []*Tag  `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 数字员工工具调用安全策略配置。
	//
	// example:
	//
	// {"aliyun":{"enable":true,"statements":[{"decision":"user_ack","product":"Sls","apiVersion":"2020-12-30","actions":["log:GetProject","log:CreateDashboard"]}]}}
	ToolPolicy *CreateDigitalEmployeeRequestToolPolicy `json:"toolPolicy,omitempty" xml:"toolPolicy,omitempty" type:"Struct"`
}

func (s CreateDigitalEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeRequest) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *CreateDigitalEmployeeRequest) GetDefaultRule() *string {
	return s.DefaultRule
}

func (s *CreateDigitalEmployeeRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDigitalEmployeeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateDigitalEmployeeRequest) GetKnowledges() *CreateDigitalEmployeeRequestKnowledges {
	return s.Knowledges
}

func (s *CreateDigitalEmployeeRequest) GetName() *string {
	return s.Name
}

func (s *CreateDigitalEmployeeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDigitalEmployeeRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateDigitalEmployeeRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *CreateDigitalEmployeeRequest) GetToolPolicy() *CreateDigitalEmployeeRequestToolPolicy {
	return s.ToolPolicy
}

func (s *CreateDigitalEmployeeRequest) SetAttributes(v map[string]*string) *CreateDigitalEmployeeRequest {
	s.Attributes = v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetDefaultRule(v string) *CreateDigitalEmployeeRequest {
	s.DefaultRule = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetDescription(v string) *CreateDigitalEmployeeRequest {
	s.Description = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetDisplayName(v string) *CreateDigitalEmployeeRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetKnowledges(v *CreateDigitalEmployeeRequestKnowledges) *CreateDigitalEmployeeRequest {
	s.Knowledges = v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetName(v string) *CreateDigitalEmployeeRequest {
	s.Name = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetResourceGroupId(v string) *CreateDigitalEmployeeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetRoleArn(v string) *CreateDigitalEmployeeRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetTags(v []*Tag) *CreateDigitalEmployeeRequest {
	s.Tags = v
	return s
}

func (s *CreateDigitalEmployeeRequest) SetToolPolicy(v *CreateDigitalEmployeeRequestToolPolicy) *CreateDigitalEmployeeRequest {
	s.ToolPolicy = v
	return s
}

func (s *CreateDigitalEmployeeRequest) Validate() error {
	if s.Knowledges != nil {
		if err := s.Knowledges.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ToolPolicy != nil {
		if err := s.ToolPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDigitalEmployeeRequestKnowledges struct {
	Bailian []*CreateDigitalEmployeeRequestKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
	Sop     []map[string]interface{}                         `json:"sop,omitempty" xml:"sop,omitempty" type:"Repeated"`
}

func (s CreateDigitalEmployeeRequestKnowledges) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequestKnowledges) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeRequestKnowledges) GetBailian() []*CreateDigitalEmployeeRequestKnowledgesBailian {
	return s.Bailian
}

func (s *CreateDigitalEmployeeRequestKnowledges) GetSop() []map[string]interface{} {
	return s.Sop
}

func (s *CreateDigitalEmployeeRequestKnowledges) SetBailian(v []*CreateDigitalEmployeeRequestKnowledgesBailian) *CreateDigitalEmployeeRequestKnowledges {
	s.Bailian = v
	return s
}

func (s *CreateDigitalEmployeeRequestKnowledges) SetSop(v []map[string]interface{}) *CreateDigitalEmployeeRequestKnowledges {
	s.Sop = v
	return s
}

func (s *CreateDigitalEmployeeRequestKnowledges) Validate() error {
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

type CreateDigitalEmployeeRequestKnowledgesBailian struct {
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
	// llm-xxxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateDigitalEmployeeRequestKnowledgesBailian) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequestKnowledgesBailian) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) GetAttributes() *string {
	return s.Attributes
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) GetIndexId() *string {
	return s.IndexId
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) GetRegion() *string {
	return s.Region
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) SetAttributes(v string) *CreateDigitalEmployeeRequestKnowledgesBailian {
	s.Attributes = &v
	return s
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) SetIndexId(v string) *CreateDigitalEmployeeRequestKnowledgesBailian {
	s.IndexId = &v
	return s
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) SetRegion(v string) *CreateDigitalEmployeeRequestKnowledgesBailian {
	s.Region = &v
	return s
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) SetWorkspaceId(v string) *CreateDigitalEmployeeRequestKnowledgesBailian {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDigitalEmployeeRequestKnowledgesBailian) Validate() error {
	return dara.Validate(s)
}

type CreateDigitalEmployeeRequestToolPolicy struct {
	// Aliyun MCP 工具调用安全策略配置。
	//
	// example:
	//
	// {"enable":true,"statements":[{"decision":"user_ack","product":"Sls","apiVersion":"2020-12-30","actions":["log:GetProject","log:CreateDashboard"]}]}
	Aliyun *CreateDigitalEmployeeRequestToolPolicyAliyun `json:"aliyun,omitempty" xml:"aliyun,omitempty" type:"Struct"`
}

func (s CreateDigitalEmployeeRequestToolPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequestToolPolicy) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeRequestToolPolicy) GetAliyun() *CreateDigitalEmployeeRequestToolPolicyAliyun {
	return s.Aliyun
}

func (s *CreateDigitalEmployeeRequestToolPolicy) SetAliyun(v *CreateDigitalEmployeeRequestToolPolicyAliyun) *CreateDigitalEmployeeRequestToolPolicy {
	s.Aliyun = v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicy) Validate() error {
	if s.Aliyun != nil {
		if err := s.Aliyun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDigitalEmployeeRequestToolPolicyAliyun struct {
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
	Statements []*CreateDigitalEmployeeRequestToolPolicyAliyunStatements `json:"statements,omitempty" xml:"statements,omitempty" type:"Repeated"`
}

func (s CreateDigitalEmployeeRequestToolPolicyAliyun) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequestToolPolicyAliyun) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyun) GetEnable() *bool {
	return s.Enable
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyun) GetStatements() []*CreateDigitalEmployeeRequestToolPolicyAliyunStatements {
	return s.Statements
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyun) SetEnable(v bool) *CreateDigitalEmployeeRequestToolPolicyAliyun {
	s.Enable = &v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyun) SetStatements(v []*CreateDigitalEmployeeRequestToolPolicyAliyunStatements) *CreateDigitalEmployeeRequestToolPolicyAliyun {
	s.Statements = v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyun) Validate() error {
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

type CreateDigitalEmployeeRequestToolPolicyAliyunStatements struct {
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

func (s CreateDigitalEmployeeRequestToolPolicyAliyunStatements) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequestToolPolicyAliyunStatements) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) GetActions() []*string {
	return s.Actions
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) GetDecision() *string {
	return s.Decision
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) GetProduct() *string {
	return s.Product
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) SetActions(v []*string) *CreateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.Actions = v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) SetApiVersion(v string) *CreateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.ApiVersion = &v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) SetDecision(v string) *CreateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.Decision = &v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) SetProduct(v string) *CreateDigitalEmployeeRequestToolPolicyAliyunStatements {
	s.Product = &v
	return s
}

func (s *CreateDigitalEmployeeRequestToolPolicyAliyunStatements) Validate() error {
	return dara.Validate(s)
}
