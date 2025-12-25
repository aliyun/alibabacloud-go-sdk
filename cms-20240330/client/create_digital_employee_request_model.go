// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetRoleArn(v string) *CreateDigitalEmployeeRequest
	GetRoleArn() *string
}

type CreateDigitalEmployeeRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::12345678912:role/testrole
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
}

func (s CreateDigitalEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeRequest) GoString() string {
	return s.String()
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

func (s *CreateDigitalEmployeeRequest) GetRoleArn() *string {
	return s.RoleArn
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

func (s *CreateDigitalEmployeeRequest) SetRoleArn(v string) *CreateDigitalEmployeeRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateDigitalEmployeeRequest) Validate() error {
	if s.Knowledges != nil {
		if err := s.Knowledges.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDigitalEmployeeRequestKnowledges struct {
	Bailian []*CreateDigitalEmployeeRequestKnowledgesBailian `json:"bailian,omitempty" xml:"bailian,omitempty" type:"Repeated"`
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

func (s *CreateDigitalEmployeeRequestKnowledges) SetBailian(v []*CreateDigitalEmployeeRequestKnowledgesBailian) *CreateDigitalEmployeeRequestKnowledges {
	s.Bailian = v
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
