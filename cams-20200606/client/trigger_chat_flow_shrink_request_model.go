// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerChatFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClaimTimeMillis(v int64) *TriggerChatFlowShrinkRequest
	GetClaimTimeMillis() *int64
	SetDataShrink(v string) *TriggerChatFlowShrinkRequest
	GetDataShrink() *string
	SetDiscardTimeMillis(v int64) *TriggerChatFlowShrinkRequest
	GetDiscardTimeMillis() *int64
	SetFlowCode(v string) *TriggerChatFlowShrinkRequest
	GetFlowCode() *string
	SetOutId(v string) *TriggerChatFlowShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *TriggerChatFlowShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TriggerChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TriggerChatFlowShrinkRequest
	GetResourceOwnerId() *int64
	SetUuid(v string) *TriggerChatFlowShrinkRequest
	GetUuid() *string
}

type TriggerChatFlowShrinkRequest struct {
	// The declared occurrence time of the event, usually the time when the request was constructed, in milliseconds timestamp.
	//
	// example:
	//
	// 1731502129000
	ClaimTimeMillis *int64 `json:"ClaimTimeMillis,omitempty" xml:"ClaimTimeMillis,omitempty"`
	// Input parameters in Key-Value format. The Key must match the input strategy configured at the start node of your flow.
	//
	// example:
	//
	// {"my_biz_data_0": "hi", "my_biz_data_1": "1024"}
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The time when the event should be discarded, i.e., the expiration time. If this field is specified, the message will be discarded if it exceeds this time, in milliseconds timestamp.
	//
	// example:
	//
	// 1731502729000
	DiscardTimeMillis *int64 `json:"DiscardTimeMillis,omitempty" xml:"DiscardTimeMillis,omitempty"`
	// Flow code.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4912c16943b4dfba44bd6fedacf8c70
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// External system transaction number, used to associate with external business system transactions. You can retrieve this parameter within the flow after triggering.
	//
	// example:
	//
	// 8d4acf7e-e360-eb83-6d74-fcf9c4538fda
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Unique event ID used for idempotent triggers. Do not include any business semantics; you can retrieve this parameter within the flow after triggering.
	//
	// example:
	//
	// c68622e6-5f0d-c8a4-af41-e949c2a7580e
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s TriggerChatFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerChatFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *TriggerChatFlowShrinkRequest) GetClaimTimeMillis() *int64 {
	return s.ClaimTimeMillis
}

func (s *TriggerChatFlowShrinkRequest) GetDataShrink() *string {
	return s.DataShrink
}

func (s *TriggerChatFlowShrinkRequest) GetDiscardTimeMillis() *int64 {
	return s.DiscardTimeMillis
}

func (s *TriggerChatFlowShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *TriggerChatFlowShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *TriggerChatFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TriggerChatFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TriggerChatFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TriggerChatFlowShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *TriggerChatFlowShrinkRequest) SetClaimTimeMillis(v int64) *TriggerChatFlowShrinkRequest {
	s.ClaimTimeMillis = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetDataShrink(v string) *TriggerChatFlowShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetDiscardTimeMillis(v int64) *TriggerChatFlowShrinkRequest {
	s.DiscardTimeMillis = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetFlowCode(v string) *TriggerChatFlowShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetOutId(v string) *TriggerChatFlowShrinkRequest {
	s.OutId = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetOwnerId(v int64) *TriggerChatFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetResourceOwnerAccount(v string) *TriggerChatFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetResourceOwnerId(v int64) *TriggerChatFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) SetUuid(v string) *TriggerChatFlowShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *TriggerChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
