// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *GetSkillGroupRequest
	GetSkillGroupId() *string
}

type GetSkillGroupRequest struct {
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

func (s GetSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *GetSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *GetSkillGroupRequest) SetInstanceId(v string) *GetSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupRequest) SetSkillGroupId(v string) *GetSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *GetSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
