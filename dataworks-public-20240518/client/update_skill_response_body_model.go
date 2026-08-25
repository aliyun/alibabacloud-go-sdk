// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSkillResponseBody
	GetRequestId() *string
	SetSkill(v *UpdateSkillResponseBodySkill) *UpdateSkillResponseBody
	GetSkill() *UpdateSkillResponseBodySkill
}

type UpdateSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Skill details.
	Skill *UpdateSkillResponseBodySkill `json:"Skill,omitempty" xml:"Skill,omitempty" type:"Struct"`
}

func (s UpdateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillResponseBody) GetSkill() *UpdateSkillResponseBodySkill {
	return s.Skill
}

func (s *UpdateSkillResponseBody) SetRequestId(v string) *UpdateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillResponseBody) SetSkill(v *UpdateSkillResponseBodySkill) *UpdateSkillResponseBody {
	s.Skill = v
	return s
}

func (s *UpdateSkillResponseBody) Validate() error {
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSkillResponseBodySkill struct {
	// The SKILL.md body content.
	//
	// example:
	//
	// Putting an elephant in a refrigerator takes three steps: open the refrigerator door, put the elephant in, and close the refrigerator door
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 123456
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The Skill description.
	//
	// example:
	//
	// Data analytics skill
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time (millisecond timestamp).
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time (millisecond timestamp).
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The last modifier ID.
	//
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The Skill name.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The visibility level.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope.
	VisibilityScope *UpdateSkillResponseBodySkillVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s UpdateSkillResponseBodySkill) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBodySkill) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBodySkill) GetBody() *string {
	return s.Body
}

func (s *UpdateSkillResponseBodySkill) GetCreatorId() *string {
	return s.CreatorId
}

func (s *UpdateSkillResponseBodySkill) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillResponseBodySkill) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *UpdateSkillResponseBodySkill) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *UpdateSkillResponseBodySkill) GetModifierId() *string {
	return s.ModifierId
}

func (s *UpdateSkillResponseBodySkill) GetName() *string {
	return s.Name
}

func (s *UpdateSkillResponseBodySkill) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateSkillResponseBodySkill) GetVisibilityScope() *UpdateSkillResponseBodySkillVisibilityScope {
	return s.VisibilityScope
}

func (s *UpdateSkillResponseBodySkill) SetBody(v string) *UpdateSkillResponseBodySkill {
	s.Body = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetCreatorId(v string) *UpdateSkillResponseBodySkill {
	s.CreatorId = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetDescription(v string) *UpdateSkillResponseBodySkill {
	s.Description = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetGmtCreateTime(v string) *UpdateSkillResponseBodySkill {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetGmtModifiedTime(v string) *UpdateSkillResponseBodySkill {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetModifierId(v string) *UpdateSkillResponseBodySkill {
	s.ModifierId = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetName(v string) *UpdateSkillResponseBodySkill {
	s.Name = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetVisibility(v string) *UpdateSkillResponseBodySkill {
	s.Visibility = &v
	return s
}

func (s *UpdateSkillResponseBodySkill) SetVisibilityScope(v *UpdateSkillResponseBodySkillVisibilityScope) *UpdateSkillResponseBodySkill {
	s.VisibilityScope = v
	return s
}

func (s *UpdateSkillResponseBodySkill) Validate() error {
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSkillResponseBodySkillVisibilityScope struct {
	// The list of visible project IDs.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of visible user IDs.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateSkillResponseBodySkillVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBodySkillVisibilityScope) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBodySkillVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *UpdateSkillResponseBodySkillVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateSkillResponseBodySkillVisibilityScope) SetProjectIds(v []*string) *UpdateSkillResponseBodySkillVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *UpdateSkillResponseBodySkillVisibilityScope) SetUserIds(v []*string) *UpdateSkillResponseBodySkillVisibilityScope {
	s.UserIds = v
	return s
}

func (s *UpdateSkillResponseBodySkillVisibilityScope) Validate() error {
	return dara.Validate(s)
}
