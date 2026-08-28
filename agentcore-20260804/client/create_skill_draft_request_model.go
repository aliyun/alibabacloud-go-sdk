// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateSkillDraftRequestBody) *CreateSkillDraftRequest
	GetBody() *CreateSkillDraftRequestBody
}

type CreateSkillDraftRequest struct {
	// The request body.
	Body *CreateSkillDraftRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateSkillDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillDraftRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillDraftRequest) GetBody() *CreateSkillDraftRequestBody {
	return s.Body
}

func (s *CreateSkillDraftRequest) SetBody(v *CreateSkillDraftRequestBody) *CreateSkillDraftRequest {
	s.Body = v
	return s
}

func (s *CreateSkillDraftRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSkillDraftRequestBody struct {
	// The version from which to fork the draft. If not specified, a new Skill is created.
	//
	// example:
	//
	// 0.0.1
	BasedOnVersion *string `json:"basedOnVersion,omitempty" xml:"basedOnVersion,omitempty"`
	// The commit message.
	//
	// example:
	//
	// Initial version
	CommitMsg *string `json:"commitMsg,omitempty" xml:"commitMsg,omitempty"`
	// The Skill card JSON string that contains the complete Skill information.
	//
	// example:
	//
	// {"name":"customer-service-skill","description":"..."}
	SkillCard *string `json:"skillCard,omitempty" xml:"skillCard,omitempty"`
	// The Skill name.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer-service-skill
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The draft version number to assign. If not specified, the version number is automatically incremented.
	//
	// example:
	//
	// 0.0.2
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s CreateSkillDraftRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillDraftRequestBody) GoString() string {
	return s.String()
}

func (s *CreateSkillDraftRequestBody) GetBasedOnVersion() *string {
	return s.BasedOnVersion
}

func (s *CreateSkillDraftRequestBody) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *CreateSkillDraftRequestBody) GetSkillCard() *string {
	return s.SkillCard
}

func (s *CreateSkillDraftRequestBody) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateSkillDraftRequestBody) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateSkillDraftRequestBody) SetBasedOnVersion(v string) *CreateSkillDraftRequestBody {
	s.BasedOnVersion = &v
	return s
}

func (s *CreateSkillDraftRequestBody) SetCommitMsg(v string) *CreateSkillDraftRequestBody {
	s.CommitMsg = &v
	return s
}

func (s *CreateSkillDraftRequestBody) SetSkillCard(v string) *CreateSkillDraftRequestBody {
	s.SkillCard = &v
	return s
}

func (s *CreateSkillDraftRequestBody) SetSkillName(v string) *CreateSkillDraftRequestBody {
	s.SkillName = &v
	return s
}

func (s *CreateSkillDraftRequestBody) SetTargetVersion(v string) *CreateSkillDraftRequestBody {
	s.TargetVersion = &v
	return s
}

func (s *CreateSkillDraftRequestBody) Validate() error {
	return dara.Validate(s)
}
