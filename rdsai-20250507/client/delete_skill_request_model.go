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
	SetWorkspaceId(v string) *DeleteSkillRequest
	GetWorkspaceId() *string
}

type DeleteSkillRequest struct {
	// The unique identifier of the Skill.
	//
	// example:
	//
	// 1fbb6d8b-8845-4e65-871e-48bc6830****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The ContextDB workspace ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *DeleteSkillRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteSkillRequest) SetSkillId(v string) *DeleteSkillRequest {
	s.SkillId = &v
	return s
}

func (s *DeleteSkillRequest) SetWorkspaceId(v string) *DeleteSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}
