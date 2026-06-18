// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *UpdateAgentRequest
	GetAccountName() *string
	SetClientToken(v string) *UpdateAgentRequest
	GetClientToken() *string
	SetDisplayName(v string) *UpdateAgentRequest
	GetDisplayName() *string
	SetInstanceId(v string) *UpdateAgentRequest
	GetInstanceId() *string
	SetSkillGroupId(v []*int64) *UpdateAgentRequest
	GetSkillGroupId() []*int64
	SetSkillGroupIdList(v []*int64) *UpdateAgentRequest
	GetSkillGroupIdList() []*int64
}

type UpdateAgentRequest struct {
	// The agent account name, which is the phone number or mailbox entered during account registration. It is unique within the instance.
	//
	// > Update is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Unique ID for the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Agent\\"s display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// XX测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Skill groups to which the agent belongs.
	SkillGroupId []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	// List of skill groups to which the agent belongs.
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *UpdateAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAgentRequest) GetSkillGroupId() []*int64 {
	return s.SkillGroupId
}

func (s *UpdateAgentRequest) GetSkillGroupIdList() []*int64 {
	return s.SkillGroupIdList
}

func (s *UpdateAgentRequest) SetAccountName(v string) *UpdateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateAgentRequest) SetClientToken(v string) *UpdateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentRequest) SetDisplayName(v string) *UpdateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentRequest) SetInstanceId(v string) *UpdateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupId(v []*int64) *UpdateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupIdList(v []*int64) *UpdateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
	return dara.Validate(s)
}
