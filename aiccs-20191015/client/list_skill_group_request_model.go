// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v int32) *ListSkillGroupRequest
	GetChannelType() *int32
	SetClientToken(v string) *ListSkillGroupRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListSkillGroupRequest
	GetInstanceId() *string
}

type ListSkillGroupRequest struct {
	// The skill group channel type. Valid values:
	//
	// - **0**: Returns all skill groups.
	//
	// - **1**: Voice skill group.
	//
	// - **2**: Chat skill group.
	//
	// - **3**: Chat and voice skill group.
	//
	// - **4**: Ticket skill group.
	//
	// - **5**: Voice and ticket skill group.
	//
	// - **6**: Chat and ticket skill group.
	//
	// - **7**: Chat, voice, and ticket skill group.
	//
	// example:
	//
	// 2
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Unique ID of the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-****-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupRequest) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *ListSkillGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupRequest) SetChannelType(v int32) *ListSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *ListSkillGroupRequest) SetClientToken(v string) *ListSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSkillGroupRequest) SetInstanceId(v string) *ListSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
