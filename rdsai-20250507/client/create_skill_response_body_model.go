// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *CreateSkillResponseBody
	GetContent() map[string]interface{}
	SetCreatedAt(v string) *CreateSkillResponseBody
	GetCreatedAt() *string
	SetDbtypes(v []*string) *CreateSkillResponseBody
	GetDbtypes() []*string
	SetDescription(v string) *CreateSkillResponseBody
	GetDescription() *string
	SetId(v string) *CreateSkillResponseBody
	GetId() *string
	SetName(v string) *CreateSkillResponseBody
	GetName() *string
	SetRequestId(v string) *CreateSkillResponseBody
	GetRequestId() *string
	SetSkillType(v string) *CreateSkillResponseBody
	GetSkillType() *string
}

type CreateSkillResponseBody struct {
	// The database engine-specific content.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The creation time of the skill.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of database engines.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The description of the skill.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the skill.
	//
	// example:
	//
	// 82cf3d62-0add-47bd-869f-877131f7****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// query-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the skill.
	//
	// example:
	//
	// user
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
}

func (s CreateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *CreateSkillResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateSkillResponseBody) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *CreateSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillResponseBody) GetSkillType() *string {
	return s.SkillType
}

func (s *CreateSkillResponseBody) SetContent(v map[string]interface{}) *CreateSkillResponseBody {
	s.Content = v
	return s
}

func (s *CreateSkillResponseBody) SetCreatedAt(v string) *CreateSkillResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateSkillResponseBody) SetDbtypes(v []*string) *CreateSkillResponseBody {
	s.Dbtypes = v
	return s
}

func (s *CreateSkillResponseBody) SetDescription(v string) *CreateSkillResponseBody {
	s.Description = &v
	return s
}

func (s *CreateSkillResponseBody) SetId(v string) *CreateSkillResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSkillResponseBody) SetName(v string) *CreateSkillResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSkillResponseBody) SetRequestId(v string) *CreateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillResponseBody) SetSkillType(v string) *CreateSkillResponseBody {
	s.SkillType = &v
	return s
}

func (s *CreateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
