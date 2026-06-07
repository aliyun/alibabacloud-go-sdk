// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSkillResponseBody
	GetRequestId() *string
	SetSkill(v *GetSkillResponseBodySkill) *GetSkillResponseBody
	GetSkill() *GetSkillResponseBodySkill
}

type GetSkillResponseBody struct {
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Skill     *GetSkillResponseBodySkill `json:"Skill,omitempty" xml:"Skill,omitempty" type:"Struct"`
}

func (s GetSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillResponseBody) GetSkill() *GetSkillResponseBodySkill {
	return s.Skill
}

func (s *GetSkillResponseBody) SetRequestId(v string) *GetSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillResponseBody) SetSkill(v *GetSkillResponseBodySkill) *GetSkillResponseBody {
	s.Skill = v
	return s
}

func (s *GetSkillResponseBody) Validate() error {
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillResponseBodySkill struct {
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// https://your-bucket.oss-cn-hangzhou.aliyuncs.com/xxx.zip?Expires=...&Signature=...
	BundleUrl *string `json:"BundleUrl,omitempty" xml:"BundleUrl,omitempty"`
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
	//
	// example:
	//
	// 1780555634000
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
	Visibility      *string                                   `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VisibilityScope *GetSkillResponseBodySkillVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s GetSkillResponseBodySkill) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBodySkill) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBodySkill) GetBody() *string {
	return s.Body
}

func (s *GetSkillResponseBodySkill) GetBundleUrl() *string {
	return s.BundleUrl
}

func (s *GetSkillResponseBodySkill) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetSkillResponseBodySkill) GetDescription() *string {
	return s.Description
}

func (s *GetSkillResponseBodySkill) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetSkillResponseBodySkill) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetSkillResponseBodySkill) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetSkillResponseBodySkill) GetName() *string {
	return s.Name
}

func (s *GetSkillResponseBodySkill) GetVisibility() *string {
	return s.Visibility
}

func (s *GetSkillResponseBodySkill) GetVisibilityScope() *GetSkillResponseBodySkillVisibilityScope {
	return s.VisibilityScope
}

func (s *GetSkillResponseBodySkill) SetBody(v string) *GetSkillResponseBodySkill {
	s.Body = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetBundleUrl(v string) *GetSkillResponseBodySkill {
	s.BundleUrl = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetCreatorId(v string) *GetSkillResponseBodySkill {
	s.CreatorId = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetDescription(v string) *GetSkillResponseBodySkill {
	s.Description = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetGmtCreateTime(v string) *GetSkillResponseBodySkill {
	s.GmtCreateTime = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetGmtModifiedTime(v string) *GetSkillResponseBodySkill {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetModifierId(v string) *GetSkillResponseBodySkill {
	s.ModifierId = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetName(v string) *GetSkillResponseBodySkill {
	s.Name = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetVisibility(v string) *GetSkillResponseBodySkill {
	s.Visibility = &v
	return s
}

func (s *GetSkillResponseBodySkill) SetVisibilityScope(v *GetSkillResponseBodySkillVisibilityScope) *GetSkillResponseBodySkill {
	s.VisibilityScope = v
	return s
}

func (s *GetSkillResponseBodySkill) Validate() error {
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillResponseBodySkillVisibilityScope struct {
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetSkillResponseBodySkillVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBodySkillVisibilityScope) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBodySkillVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *GetSkillResponseBodySkillVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetSkillResponseBodySkillVisibilityScope) SetProjectIds(v []*string) *GetSkillResponseBodySkillVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *GetSkillResponseBodySkillVisibilityScope) SetUserIds(v []*string) *GetSkillResponseBodySkillVisibilityScope {
	s.UserIds = v
	return s
}

func (s *GetSkillResponseBodySkillVisibilityScope) Validate() error {
	return dara.Validate(s)
}
