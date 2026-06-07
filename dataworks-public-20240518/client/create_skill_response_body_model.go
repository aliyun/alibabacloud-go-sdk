// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSkillResponseBody
	GetRequestId() *string
	SetSkill(v *CreateSkillResponseBodySkill) *CreateSkillResponseBody
	GetSkill() *CreateSkillResponseBodySkill
}

type CreateSkillResponseBody struct {
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Skill     *CreateSkillResponseBodySkill `json:"Skill,omitempty" xml:"Skill,omitempty" type:"Struct"`
}

func (s CreateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillResponseBody) GetSkill() *CreateSkillResponseBodySkill {
	return s.Skill
}

func (s *CreateSkillResponseBody) SetRequestId(v string) *CreateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillResponseBody) SetSkill(v *CreateSkillResponseBodySkill) *CreateSkillResponseBody {
	s.Skill = v
	return s
}

func (s *CreateSkillResponseBody) Validate() error {
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSkillResponseBodySkill struct {
	// example:
	//
	// -
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// 123456
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TENANT
	Visibility      *string                                      `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VisibilityScope *CreateSkillResponseBodySkillVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s CreateSkillResponseBodySkill) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBodySkill) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBodySkill) GetBody() *string {
	return s.Body
}

func (s *CreateSkillResponseBodySkill) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateSkillResponseBodySkill) GetDescription() *string {
	return s.Description
}

func (s *CreateSkillResponseBodySkill) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CreateSkillResponseBodySkill) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *CreateSkillResponseBodySkill) GetModifierId() *string {
	return s.ModifierId
}

func (s *CreateSkillResponseBodySkill) GetName() *string {
	return s.Name
}

func (s *CreateSkillResponseBodySkill) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateSkillResponseBodySkill) GetVisibilityScope() *CreateSkillResponseBodySkillVisibilityScope {
	return s.VisibilityScope
}

func (s *CreateSkillResponseBodySkill) SetBody(v string) *CreateSkillResponseBodySkill {
	s.Body = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetCreatorId(v string) *CreateSkillResponseBodySkill {
	s.CreatorId = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetDescription(v string) *CreateSkillResponseBodySkill {
	s.Description = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetGmtCreateTime(v string) *CreateSkillResponseBodySkill {
	s.GmtCreateTime = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetGmtModifiedTime(v string) *CreateSkillResponseBodySkill {
	s.GmtModifiedTime = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetModifierId(v string) *CreateSkillResponseBodySkill {
	s.ModifierId = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetName(v string) *CreateSkillResponseBodySkill {
	s.Name = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetVisibility(v string) *CreateSkillResponseBodySkill {
	s.Visibility = &v
	return s
}

func (s *CreateSkillResponseBodySkill) SetVisibilityScope(v *CreateSkillResponseBodySkillVisibilityScope) *CreateSkillResponseBodySkill {
	s.VisibilityScope = v
	return s
}

func (s *CreateSkillResponseBodySkill) Validate() error {
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSkillResponseBodySkillVisibilityScope struct {
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateSkillResponseBodySkillVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponseBodySkillVisibilityScope) GoString() string {
	return s.String()
}

func (s *CreateSkillResponseBodySkillVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *CreateSkillResponseBodySkillVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateSkillResponseBodySkillVisibilityScope) SetProjectIds(v []*string) *CreateSkillResponseBodySkillVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *CreateSkillResponseBodySkillVisibilityScope) SetUserIds(v []*string) *CreateSkillResponseBodySkillVisibilityScope {
	s.UserIds = v
	return s
}

func (s *CreateSkillResponseBodySkillVisibilityScope) Validate() error {
	return dara.Validate(s)
}
