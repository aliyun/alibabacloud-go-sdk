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
	SetSandboxNetworkPolicy(v *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) *UpdateDigitalEmployeeRequest
	GetSandboxNetworkPolicy() *UpdateDigitalEmployeeRequestSandboxNetworkPolicy
	SetToolPolicy(v *UpdateDigitalEmployeeRequestToolPolicy) *UpdateDigitalEmployeeRequest
	GetToolPolicy() *UpdateDigitalEmployeeRequestToolPolicy
}

type UpdateDigitalEmployeeRequest struct {
	// The attributes.
	Attributes map[string]*string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The default rule of the digital employee.
	//
	// example:
	//
	// test
	DefaultRule *string `json:"defaultRule,omitempty" xml:"defaultRule,omitempty"`
	// The description of the digital employee.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the digital employee.
	//
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The list of knowledge bases.
	Knowledges *UpdateDigitalEmployeeRequestKnowledges `json:"knowledges,omitempty" xml:"knowledges,omitempty" type:"Struct"`
	// The ARN of the RAM role.
	//
	// example:
	//
	// acs:ram::12345678912:role/testrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// The list of CIDRs or IP addresses that are allowed to be accessed.
	//
	// example:
	//
	// {"allowFqdns":["api.example.com"],"allowCidrs":["1.2.3.0/24","8.8.8.8"],"enableAcl":false}
	SandboxNetworkPolicy *UpdateDigitalEmployeeRequestSandboxNetworkPolicy `json:"sandboxNetworkPolicy,omitempty" xml:"sandboxNetworkPolicy,omitempty" type:"Struct"`
	// The security policy configuration for tool calling of the digital employee.
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

func (s *UpdateDigitalEmployeeRequest) GetSandboxNetworkPolicy() *UpdateDigitalEmployeeRequestSandboxNetworkPolicy {
	return s.SandboxNetworkPolicy
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

func (s *UpdateDigitalEmployeeRequest) SetSandboxNetworkPolicy(v *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) *UpdateDigitalEmployeeRequest {
	s.SandboxNetworkPolicy = v
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
	if s.SandboxNetworkPolicy != nil {
		if err := s.SandboxNetworkPolicy.Validate(); err != nil {
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
	// The list of Bailian knowledge bases.
	Bailian []*UpdateDigitalEmployeeRequestKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
	// The list of SOP knowledge bases.
	Sop []map[string]interface{} `json:"sop,omitempty" xml:"sop,omitempty" type:"Repeated"`
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
	// The attributes of the knowledge base.
	//
	// example:
	//
	// test
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The Bailian index ID.
	//
	// example:
	//
	// index-xxxx
	IndexId *string `json:"indexId,omitempty" xml:"indexId,omitempty"`
	// The region of the knowledge base.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The Bailian workspace ID.
	//
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

type UpdateDigitalEmployeeRequestSandboxNetworkPolicy struct {
	// The list of CIDRs or IP addresses that are allowed to be accessed.
	AllowCidrs []*string `json:"allowCidrs,omitempty" xml:"allowCidrs,omitempty" type:"Repeated"`
	// The list of FQDNs that are allowed to be accessed.
	AllowFqdns []*string `json:"allowFqdns,omitempty" xml:"allowFqdns,omitempty" type:"Repeated"`
	// Specifies whether to enable the sandbox network ACL.
	//
	// example:
	//
	// false
	EnableAcl *bool `json:"enableAcl,omitempty" xml:"enableAcl,omitempty"`
}

func (s UpdateDigitalEmployeeRequestSandboxNetworkPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequestSandboxNetworkPolicy) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) GetAllowCidrs() []*string {
	return s.AllowCidrs
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) GetAllowFqdns() []*string {
	return s.AllowFqdns
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) GetEnableAcl() *bool {
	return s.EnableAcl
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) SetAllowCidrs(v []*string) *UpdateDigitalEmployeeRequestSandboxNetworkPolicy {
	s.AllowCidrs = v
	return s
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) SetAllowFqdns(v []*string) *UpdateDigitalEmployeeRequestSandboxNetworkPolicy {
	s.AllowFqdns = v
	return s
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) SetEnableAcl(v bool) *UpdateDigitalEmployeeRequestSandboxNetworkPolicy {
	s.EnableAcl = &v
	return s
}

func (s *UpdateDigitalEmployeeRequestSandboxNetworkPolicy) Validate() error {
	return dara.Validate(s)
}

type UpdateDigitalEmployeeRequestToolPolicy struct {
	// The security policy configuration for Aliyun CLI tool calling.
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
	// Specifies whether to enable the policy.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of Aliyun CLI tool policy statements.
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
	// RAM action
	//
	// example:
	//
	// ["log:GetProject","log:CreateDashboard"]
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The API version. This parameter is deprecated.
	//
	// example:
	//
	// 2024-03-30
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The execution policy.
	//
	// example:
	//
	// user_ack
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The cloud service code.
	//
	// example:
	//
	// Cms
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
