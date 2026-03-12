// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *UpdateSkillResponseBody
	GetContent() map[string]interface{}
	SetDbtypes(v []*string) *UpdateSkillResponseBody
	GetDbtypes() []*string
	SetDescription(v string) *UpdateSkillResponseBody
	GetDescription() *string
	SetId(v string) *UpdateSkillResponseBody
	GetId() *string
	SetName(v string) *UpdateSkillResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateSkillResponseBody
	GetRequestId() *string
	SetSkillType(v string) *UpdateSkillResponseBody
	GetSkillType() *string
	SetUpdatedAt(v string) *UpdateSkillResponseBody
	GetUpdatedAt() *string
}

type UpdateSkillResponseBody struct {
	// The content of the skill.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database engines.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The description of the skill. It can be up to 1000 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the skill.
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the skill, which can contain only lowercase letters, numbers, and hyphens.
	//
	// example:
	//
	// sql-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill type.
	//
	// example:
	//
	// user
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// The update time of the skill.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s UpdateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateSkillResponseBody) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *UpdateSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillResponseBody) GetSkillType() *string {
	return s.SkillType
}

func (s *UpdateSkillResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateSkillResponseBody) SetContent(v map[string]interface{}) *UpdateSkillResponseBody {
	s.Content = v
	return s
}

func (s *UpdateSkillResponseBody) SetDbtypes(v []*string) *UpdateSkillResponseBody {
	s.Dbtypes = v
	return s
}

func (s *UpdateSkillResponseBody) SetDescription(v string) *UpdateSkillResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateSkillResponseBody) SetId(v string) *UpdateSkillResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateSkillResponseBody) SetName(v string) *UpdateSkillResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateSkillResponseBody) SetRequestId(v string) *UpdateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillResponseBody) SetSkillType(v string) *UpdateSkillResponseBody {
	s.SkillType = &v
	return s
}

func (s *UpdateSkillResponseBody) SetUpdatedAt(v string) *UpdateSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
