// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentBySkillGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListAgentBySkillGroupIdRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListAgentBySkillGroupIdRequest
	GetInstanceId() *string
	SetSkillGroupId(v int64) *ListAgentBySkillGroupIdRequest
	GetSkillGroupId() *int64
}

type ListAgentBySkillGroupIdRequest struct {
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
	// 666666
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListAgentBySkillGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentBySkillGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListAgentBySkillGroupIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentBySkillGroupIdRequest) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *ListAgentBySkillGroupIdRequest) SetClientToken(v string) *ListAgentBySkillGroupIdRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetInstanceId(v string) *ListAgentBySkillGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetSkillGroupId(v int64) *ListAgentBySkillGroupIdRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
