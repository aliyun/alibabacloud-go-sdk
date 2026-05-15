// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateAgentRequest
	GetAccountName() *string
	SetClientToken(v string) *CreateAgentRequest
	GetClientToken() *string
	SetDisplayName(v string) *CreateAgentRequest
	GetDisplayName() *string
	SetInstanceId(v string) *CreateAgentRequest
	GetInstanceId() *string
	SetSkillGroupId(v []*int64) *CreateAgentRequest
	GetSkillGroupId() []*int64
	SetSkillGroupIdList(v []*int64) *CreateAgentRequest
	GetSkillGroupIdList() []*int64
}

type CreateAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAgentRequest) GetSkillGroupId() []*int64 {
	return s.SkillGroupId
}

func (s *CreateAgentRequest) GetSkillGroupIdList() []*int64 {
	return s.SkillGroupIdList
}

func (s *CreateAgentRequest) SetAccountName(v string) *CreateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAgentRequest) SetClientToken(v string) *CreateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentRequest) SetDisplayName(v string) *CreateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentRequest) SetInstanceId(v string) *CreateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupId(v []*int64) *CreateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupIdList(v []*int64) *CreateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	return dara.Validate(s)
}
