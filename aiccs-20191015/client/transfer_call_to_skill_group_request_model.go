// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferCallToSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *TransferCallToSkillGroupRequest
	GetAccountName() *string
	SetCallId(v string) *TransferCallToSkillGroupRequest
	GetCallId() *string
	SetClientToken(v string) *TransferCallToSkillGroupRequest
	GetClientToken() *string
	SetConnectionId(v string) *TransferCallToSkillGroupRequest
	GetConnectionId() *string
	SetHoldConnectionId(v string) *TransferCallToSkillGroupRequest
	GetHoldConnectionId() *string
	SetInstanceId(v string) *TransferCallToSkillGroupRequest
	GetInstanceId() *string
	SetIsSingleTransfer(v bool) *TransferCallToSkillGroupRequest
	GetIsSingleTransfer() *bool
	SetJobId(v string) *TransferCallToSkillGroupRequest
	GetJobId() *string
	SetSkillGroupId(v int64) *TransferCallToSkillGroupRequest
	GetSkillGroupId() *int64
	SetType(v int32) *TransferCallToSkillGroupRequest
	GetType() *int32
}

type TransferCallToSkillGroupRequest struct {
	// The agent account name, which is the phone number or mailbox entered during account registration. It is unique within the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Hotline session ID.
	//
	// example:
	//
	// 7719786
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Unique ID for the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connId in the WebSocket after an inbound call.
	//
	// example:
	//
	// 7719788
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The holdConnId in the WebSocket after an inbound call (required only for two-step transfer).
	//
	// example:
	//
	// 0
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
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
	// Default value: **true**. Valid values:
	//
	// - **true**: Single-step transfer.
	//
	// - **false**: Two-step transfer.
	//
	// example:
	//
	// true
	IsSingleTransfer *bool `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
	// The jobId in the WebSocket after an inbound call.
	//
	// example:
	//
	// 7719787
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 356543
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Default value: **1**. Valid values:
	//
	// - **1**: Single-step transfer.
	//
	// - **2**: Two-step transfer.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TransferCallToSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferCallToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *TransferCallToSkillGroupRequest) GetCallId() *string {
	return s.CallId
}

func (s *TransferCallToSkillGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TransferCallToSkillGroupRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *TransferCallToSkillGroupRequest) GetHoldConnectionId() *string {
	return s.HoldConnectionId
}

func (s *TransferCallToSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransferCallToSkillGroupRequest) GetIsSingleTransfer() *bool {
	return s.IsSingleTransfer
}

func (s *TransferCallToSkillGroupRequest) GetJobId() *string {
	return s.JobId
}

func (s *TransferCallToSkillGroupRequest) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *TransferCallToSkillGroupRequest) GetType() *int32 {
	return s.Type
}

func (s *TransferCallToSkillGroupRequest) SetAccountName(v string) *TransferCallToSkillGroupRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetCallId(v string) *TransferCallToSkillGroupRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetClientToken(v string) *TransferCallToSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetHoldConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetInstanceId(v string) *TransferCallToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetIsSingleTransfer(v bool) *TransferCallToSkillGroupRequest {
	s.IsSingleTransfer = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetJobId(v string) *TransferCallToSkillGroupRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetSkillGroupId(v int64) *TransferCallToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetType(v int32) *TransferCallToSkillGroupRequest {
	s.Type = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
