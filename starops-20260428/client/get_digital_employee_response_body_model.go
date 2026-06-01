// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *GetDigitalEmployeeResponseBody
	GetAttributes() map[string]*string
	SetCreateTime(v string) *GetDigitalEmployeeResponseBody
	GetCreateTime() *string
	SetDefaultRule(v string) *GetDigitalEmployeeResponseBody
	GetDefaultRule() *string
	SetDescription(v string) *GetDigitalEmployeeResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetDigitalEmployeeResponseBody
	GetDisplayName() *string
	SetEmployeeType(v string) *GetDigitalEmployeeResponseBody
	GetEmployeeType() *string
	SetKnowledges(v *GetDigitalEmployeeResponseBodyKnowledges) *GetDigitalEmployeeResponseBody
	GetKnowledges() *GetDigitalEmployeeResponseBodyKnowledges
	SetName(v string) *GetDigitalEmployeeResponseBody
	GetName() *string
	SetRegionId(v string) *GetDigitalEmployeeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetDigitalEmployeeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetDigitalEmployeeResponseBody
	GetResourceGroupId() *string
	SetRoleArn(v string) *GetDigitalEmployeeResponseBody
	GetRoleArn() *string
	SetTags(v []*Tag) *GetDigitalEmployeeResponseBody
	GetTags() []*Tag
	SetToolPolicy(v *GetDigitalEmployeeResponseBodyToolPolicy) *GetDigitalEmployeeResponseBody
	GetToolPolicy() *GetDigitalEmployeeResponseBodyToolPolicy
	SetUpdateTime(v string) *GetDigitalEmployeeResponseBody
	GetUpdateTime() *string
}

type GetDigitalEmployeeResponseBody struct {
	Attributes map[string]*string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
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
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// custom
	EmployeeType *string                                   `json:"employeeType,omitempty" xml:"employeeType,omitempty"`
	Knowledges   *GetDigitalEmployeeResponseBodyKnowledges `json:"knowledges,omitempty" xml:"knowledges,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
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
	ToolPolicy *GetDigitalEmployeeResponseBodyToolPolicy `json:"toolPolicy,omitempty" xml:"toolPolicy,omitempty" type:"Struct"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-02-18T02:25:06Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDigitalEmployeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponseBody) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *GetDigitalEmployeeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDigitalEmployeeResponseBody) GetDefaultRule() *string {
	return s.DefaultRule
}

func (s *GetDigitalEmployeeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDigitalEmployeeResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDigitalEmployeeResponseBody) GetEmployeeType() *string {
	return s.EmployeeType
}

func (s *GetDigitalEmployeeResponseBody) GetKnowledges() *GetDigitalEmployeeResponseBodyKnowledges {
	return s.Knowledges
}

func (s *GetDigitalEmployeeResponseBody) GetName() *string {
	return s.Name
}

func (s *GetDigitalEmployeeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDigitalEmployeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDigitalEmployeeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetDigitalEmployeeResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetDigitalEmployeeResponseBody) GetTags() []*Tag {
	return s.Tags
}

func (s *GetDigitalEmployeeResponseBody) GetToolPolicy() *GetDigitalEmployeeResponseBodyToolPolicy {
	return s.ToolPolicy
}

func (s *GetDigitalEmployeeResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDigitalEmployeeResponseBody) SetAttributes(v map[string]*string) *GetDigitalEmployeeResponseBody {
	s.Attributes = v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetCreateTime(v string) *GetDigitalEmployeeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetDefaultRule(v string) *GetDigitalEmployeeResponseBody {
	s.DefaultRule = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetDescription(v string) *GetDigitalEmployeeResponseBody {
	s.Description = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetDisplayName(v string) *GetDigitalEmployeeResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetEmployeeType(v string) *GetDigitalEmployeeResponseBody {
	s.EmployeeType = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetKnowledges(v *GetDigitalEmployeeResponseBodyKnowledges) *GetDigitalEmployeeResponseBody {
	s.Knowledges = v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetName(v string) *GetDigitalEmployeeResponseBody {
	s.Name = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetRegionId(v string) *GetDigitalEmployeeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetRequestId(v string) *GetDigitalEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetResourceGroupId(v string) *GetDigitalEmployeeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetRoleArn(v string) *GetDigitalEmployeeResponseBody {
	s.RoleArn = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetTags(v []*Tag) *GetDigitalEmployeeResponseBody {
	s.Tags = v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetToolPolicy(v *GetDigitalEmployeeResponseBodyToolPolicy) *GetDigitalEmployeeResponseBody {
	s.ToolPolicy = v
	return s
}

func (s *GetDigitalEmployeeResponseBody) SetUpdateTime(v string) *GetDigitalEmployeeResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetDigitalEmployeeResponseBody) Validate() error {
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

type GetDigitalEmployeeResponseBodyKnowledges struct {
	Bailian []*GetDigitalEmployeeResponseBodyKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
	Sop     []map[string]interface{}                           `json:"sop,omitempty" xml:"sop,omitempty" type:"Repeated"`
}

func (s GetDigitalEmployeeResponseBodyKnowledges) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponseBodyKnowledges) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponseBodyKnowledges) GetBailian() []*GetDigitalEmployeeResponseBodyKnowledgesBailian {
	return s.Bailian
}

func (s *GetDigitalEmployeeResponseBodyKnowledges) GetSop() []map[string]interface{} {
	return s.Sop
}

func (s *GetDigitalEmployeeResponseBodyKnowledges) SetBailian(v []*GetDigitalEmployeeResponseBodyKnowledgesBailian) *GetDigitalEmployeeResponseBodyKnowledges {
	s.Bailian = v
	return s
}

func (s *GetDigitalEmployeeResponseBodyKnowledges) SetSop(v []map[string]interface{}) *GetDigitalEmployeeResponseBodyKnowledges {
	s.Sop = v
	return s
}

func (s *GetDigitalEmployeeResponseBodyKnowledges) Validate() error {
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

type GetDigitalEmployeeResponseBodyKnowledgesBailian struct {
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

func (s GetDigitalEmployeeResponseBodyKnowledgesBailian) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponseBodyKnowledgesBailian) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) GetAttributes() *string {
	return s.Attributes
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) GetIndexId() *string {
	return s.IndexId
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) GetRegion() *string {
	return s.Region
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) SetAttributes(v string) *GetDigitalEmployeeResponseBodyKnowledgesBailian {
	s.Attributes = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) SetIndexId(v string) *GetDigitalEmployeeResponseBodyKnowledgesBailian {
	s.IndexId = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) SetRegion(v string) *GetDigitalEmployeeResponseBodyKnowledgesBailian {
	s.Region = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) SetWorkspaceId(v string) *GetDigitalEmployeeResponseBodyKnowledgesBailian {
	s.WorkspaceId = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyKnowledgesBailian) Validate() error {
	return dara.Validate(s)
}

type GetDigitalEmployeeResponseBodyToolPolicy struct {
	// Aliyun MCP 工具调用安全策略配置。
	//
	// example:
	//
	// {"enable":true,"statements":[{"decision":"user_ack","product":"Sls","apiVersion":"2020-12-30","actions":["log:GetProject","log:CreateDashboard"]}]}
	Aliyun *GetDigitalEmployeeResponseBodyToolPolicyAliyun `json:"aliyun,omitempty" xml:"aliyun,omitempty" type:"Struct"`
}

func (s GetDigitalEmployeeResponseBodyToolPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponseBodyToolPolicy) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponseBodyToolPolicy) GetAliyun() *GetDigitalEmployeeResponseBodyToolPolicyAliyun {
	return s.Aliyun
}

func (s *GetDigitalEmployeeResponseBodyToolPolicy) SetAliyun(v *GetDigitalEmployeeResponseBodyToolPolicyAliyun) *GetDigitalEmployeeResponseBodyToolPolicy {
	s.Aliyun = v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicy) Validate() error {
	if s.Aliyun != nil {
		if err := s.Aliyun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDigitalEmployeeResponseBodyToolPolicyAliyun struct {
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
	Statements []*GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements `json:"statements,omitempty" xml:"statements,omitempty" type:"Repeated"`
}

func (s GetDigitalEmployeeResponseBodyToolPolicyAliyun) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponseBodyToolPolicyAliyun) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyun) GetEnable() *bool {
	return s.Enable
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyun) GetStatements() []*GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements {
	return s.Statements
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyun) SetEnable(v bool) *GetDigitalEmployeeResponseBodyToolPolicyAliyun {
	s.Enable = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyun) SetStatements(v []*GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) *GetDigitalEmployeeResponseBodyToolPolicyAliyun {
	s.Statements = v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyun) Validate() error {
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

type GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements struct {
	// Aliyun OpenAPI Action 列表，格式为 product:ApiName、product:Prefix	- 或 product:*。
	//
	// example:
	//
	// ["log:GetProject","log:CreateDashboard"]
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 本条语句对应的 Aliyun OpenAPI API 版本。
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// Sls
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) GetActions() []*string {
	return s.Actions
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) GetDecision() *string {
	return s.Decision
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) GetProduct() *string {
	return s.Product
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) SetActions(v []*string) *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements {
	s.Actions = v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) SetApiVersion(v string) *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements {
	s.ApiVersion = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) SetDecision(v string) *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements {
	s.Decision = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) SetProduct(v string) *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements {
	s.Product = &v
	return s
}

func (s *GetDigitalEmployeeResponseBodyToolPolicyAliyunStatements) Validate() error {
	return dara.Validate(s)
}
