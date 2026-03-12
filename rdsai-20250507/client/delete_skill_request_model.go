// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *DeleteSkillRequest
	GetSkillId() *string
}

type DeleteSkillRequest struct {
	// The unique identifier of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1fbb6d8b-8845-4e65-871e-48bc6830****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s DeleteSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *DeleteSkillRequest) SetSkillId(v string) *DeleteSkillRequest {
	s.SkillId = &v
	return s
}

func (s *DeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}
