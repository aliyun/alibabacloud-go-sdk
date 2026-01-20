// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type UpdateDigitalEmployeeRequest struct {
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
}

func (s UpdateDigitalEmployeeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeRequest) GoString() string {
	return s.String()
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

func (s *UpdateDigitalEmployeeRequest) Validate() error {
	if s.Knowledges != nil {
		if err := s.Knowledges.Validate(); err != nil {
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
