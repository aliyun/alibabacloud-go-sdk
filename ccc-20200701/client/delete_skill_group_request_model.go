// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteSkillGroupRequest
	GetForce() *bool
	SetInstanceId(v string) *DeleteSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *DeleteSkillGroupRequest
	GetSkillGroupId() *string
}

type DeleteSkillGroupRequest struct {
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s DeleteSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *DeleteSkillGroupRequest) SetForce(v bool) *DeleteSkillGroupRequest {
	s.Force = &v
	return s
}

func (s *DeleteSkillGroupRequest) SetInstanceId(v string) *DeleteSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSkillGroupRequest) SetSkillGroupId(v string) *DeleteSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *DeleteSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
