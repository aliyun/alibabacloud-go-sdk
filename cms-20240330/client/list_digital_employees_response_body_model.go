// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDigitalEmployees(v []*ListDigitalEmployeesResponseBodyDigitalEmployees) *ListDigitalEmployeesResponseBody
	GetDigitalEmployees() []*ListDigitalEmployeesResponseBodyDigitalEmployees
	SetMaxResults(v int32) *ListDigitalEmployeesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDigitalEmployeesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDigitalEmployeesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListDigitalEmployeesResponseBody
	GetTotal() *int32
}

type ListDigitalEmployeesResponseBody struct {
	DigitalEmployees []*ListDigitalEmployeesResponseBodyDigitalEmployees `json:"digitalEmployees,omitempty" xml:"digitalEmployees,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 56
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDigitalEmployeesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesResponseBody) GetDigitalEmployees() []*ListDigitalEmployeesResponseBodyDigitalEmployees {
	return s.DigitalEmployees
}

func (s *ListDigitalEmployeesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDigitalEmployeesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDigitalEmployeesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDigitalEmployeesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDigitalEmployeesResponseBody) SetDigitalEmployees(v []*ListDigitalEmployeesResponseBodyDigitalEmployees) *ListDigitalEmployeesResponseBody {
	s.DigitalEmployees = v
	return s
}

func (s *ListDigitalEmployeesResponseBody) SetMaxResults(v int32) *ListDigitalEmployeesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDigitalEmployeesResponseBody) SetNextToken(v string) *ListDigitalEmployeesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDigitalEmployeesResponseBody) SetRequestId(v string) *ListDigitalEmployeesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDigitalEmployeesResponseBody) SetTotal(v int32) *ListDigitalEmployeesResponseBody {
	s.Total = &v
	return s
}

func (s *ListDigitalEmployeesResponseBody) Validate() error {
	if s.DigitalEmployees != nil {
		for _, item := range s.DigitalEmployees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDigitalEmployeesResponseBodyDigitalEmployees struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-11-04T08:08:57Z
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
	DisplayName  *string                                                     `json:"displayName,omitempty" xml:"displayName,omitempty"`
	EmployeeType *string                                                     `json:"employeeType,omitempty" xml:"employeeType,omitempty"`
	Knowledges   *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges `json:"knowledges,omitempty" xml:"knowledges,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// acs:ram::12345678912:role/testrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	Tags    []*Tag  `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-05-07T02:26:01Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListDigitalEmployeesResponseBodyDigitalEmployees) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesResponseBodyDigitalEmployees) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetDefaultRule() *string {
	return s.DefaultRule
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetDescription() *string {
	return s.Description
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetEmployeeType() *string {
	return s.EmployeeType
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetKnowledges() *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges {
	return s.Knowledges
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetName() *string {
	return s.Name
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetTags() []*Tag {
	return s.Tags
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetCreateTime(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.CreateTime = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetDefaultRule(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.DefaultRule = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetDescription(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.Description = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetDisplayName(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.DisplayName = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetEmployeeType(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.EmployeeType = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetKnowledges(v *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.Knowledges = v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetName(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.Name = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetResourceGroupId(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetRoleArn(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.RoleArn = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetTags(v []*Tag) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.Tags = v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) SetUpdateTime(v string) *ListDigitalEmployeesResponseBodyDigitalEmployees {
	s.UpdateTime = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployees) Validate() error {
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
	return nil
}

type ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges struct {
	Bailian []*ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
	Sop     []map[string]interface{}                                             `json:"sop,omitempty" xml:"sop,omitempty" type:"Repeated"`
}

func (s ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) GetBailian() []*ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian {
	return s.Bailian
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) GetSop() []map[string]interface{} {
	return s.Sop
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) SetBailian(v []*ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges {
	s.Bailian = v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) SetSop(v []map[string]interface{}) *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges {
	s.Sop = v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledges) Validate() error {
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

type ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian struct {
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

func (s ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) GetAttributes() *string {
	return s.Attributes
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) GetIndexId() *string {
	return s.IndexId
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) GetRegion() *string {
	return s.Region
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) SetAttributes(v string) *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian {
	s.Attributes = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) SetIndexId(v string) *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian {
	s.IndexId = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) SetRegion(v string) *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian {
	s.Region = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) SetWorkspaceId(v string) *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian {
	s.WorkspaceId = &v
	return s
}

func (s *ListDigitalEmployeesResponseBodyDigitalEmployeesKnowledgesBailian) Validate() error {
	return dara.Validate(s)
}
