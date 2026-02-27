// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *UpdateSkillRequest
	GetContent() map[string]interface{}
	SetDbtypes(v []*string) *UpdateSkillRequest
	GetDbtypes() []*string
	SetDescription(v string) *UpdateSkillRequest
	GetDescription() *string
	SetName(v string) *UpdateSkillRequest
	GetName() *string
	SetSkillId(v string) *UpdateSkillRequest
	GetSkillId() *string
}

type UpdateSkillRequest struct {
	Content     map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Dbtypes     []*string              `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// sql-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8f6a2111-3828-4a9f-a3ce-51ce73c6****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s UpdateSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillRequest) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateSkillRequest) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *UpdateSkillRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateSkillRequest) SetContent(v map[string]interface{}) *UpdateSkillRequest {
	s.Content = v
	return s
}

func (s *UpdateSkillRequest) SetDbtypes(v []*string) *UpdateSkillRequest {
	s.Dbtypes = v
	return s
}

func (s *UpdateSkillRequest) SetDescription(v string) *UpdateSkillRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillRequest) SetName(v string) *UpdateSkillRequest {
	s.Name = &v
	return s
}

func (s *UpdateSkillRequest) SetSkillId(v string) *UpdateSkillRequest {
	s.SkillId = &v
	return s
}

func (s *UpdateSkillRequest) Validate() error {
	return dara.Validate(s)
}
