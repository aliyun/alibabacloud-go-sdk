// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	SetRoleArn(v string) *GetDigitalEmployeeResponseBody
	GetRoleArn() *string
	SetUpdateTime(v string) *GetDigitalEmployeeResponseBody
	GetUpdateTime() *string
}

type GetDigitalEmployeeResponseBody struct {
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
	DisplayName  *string                                   `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	// acs:ram::12345678912:role/testrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
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

func (s *GetDigitalEmployeeResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetDigitalEmployeeResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *GetDigitalEmployeeResponseBody) SetRoleArn(v string) *GetDigitalEmployeeResponseBody {
	s.RoleArn = &v
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
	return nil
}

type GetDigitalEmployeeResponseBodyKnowledges struct {
	Bailian []*GetDigitalEmployeeResponseBodyKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
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

func (s *GetDigitalEmployeeResponseBodyKnowledges) SetBailian(v []*GetDigitalEmployeeResponseBodyKnowledgesBailian) *GetDigitalEmployeeResponseBodyKnowledges {
	s.Bailian = v
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
