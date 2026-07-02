// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *GetSkillRequest
	GetSkillId() *string
}

type GetSkillRequest struct {
	// The unique ID of the Skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-06e9dca2-0ac9-4d2e-a965-e9db9c057e00
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s GetSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillRequest) SetSkillId(v string) *GetSkillRequest {
	s.SkillId = &v
	return s
}

func (s *GetSkillRequest) Validate() error {
	return dara.Validate(s)
}
