// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveSkillGroupRequest
	GetClientToken() *string
	SetInstanceId(v string) *RemoveSkillGroupRequest
	GetInstanceId() *string
	SetSkillGroupId(v string) *RemoveSkillGroupRequest
	GetSkillGroupId() *string
}

type RemoveSkillGroupRequest struct {
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s RemoveSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *RemoveSkillGroupRequest) SetClientToken(v string) *RemoveSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetInstanceId(v string) *RemoveSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetSkillGroupId(v string) *RemoveSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
