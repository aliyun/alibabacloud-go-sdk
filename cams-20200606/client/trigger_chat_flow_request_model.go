// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerChatFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClaimTimeMillis(v int64) *TriggerChatFlowRequest
	GetClaimTimeMillis() *int64
	SetData(v map[string]interface{}) *TriggerChatFlowRequest
	GetData() map[string]interface{}
	SetDiscardTimeMillis(v int64) *TriggerChatFlowRequest
	GetDiscardTimeMillis() *int64
	SetFlowCode(v string) *TriggerChatFlowRequest
	GetFlowCode() *string
	SetOutId(v string) *TriggerChatFlowRequest
	GetOutId() *string
	SetOwnerId(v int64) *TriggerChatFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TriggerChatFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TriggerChatFlowRequest
	GetResourceOwnerId() *int64
	SetUuid(v string) *TriggerChatFlowRequest
	GetUuid() *string
}

type TriggerChatFlowRequest struct {
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
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s TriggerChatFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerChatFlowRequest) GoString() string {
	return s.String()
}

func (s *TriggerChatFlowRequest) GetClaimTimeMillis() *int64 {
	return s.ClaimTimeMillis
}

func (s *TriggerChatFlowRequest) GetData() map[string]interface{} {
	return s.Data
}

func (s *TriggerChatFlowRequest) GetDiscardTimeMillis() *int64 {
	return s.DiscardTimeMillis
}

func (s *TriggerChatFlowRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *TriggerChatFlowRequest) GetOutId() *string {
	return s.OutId
}

func (s *TriggerChatFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TriggerChatFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TriggerChatFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TriggerChatFlowRequest) GetUuid() *string {
	return s.Uuid
}

func (s *TriggerChatFlowRequest) SetClaimTimeMillis(v int64) *TriggerChatFlowRequest {
	s.ClaimTimeMillis = &v
	return s
}

func (s *TriggerChatFlowRequest) SetData(v map[string]interface{}) *TriggerChatFlowRequest {
	s.Data = v
	return s
}

func (s *TriggerChatFlowRequest) SetDiscardTimeMillis(v int64) *TriggerChatFlowRequest {
	s.DiscardTimeMillis = &v
	return s
}

func (s *TriggerChatFlowRequest) SetFlowCode(v string) *TriggerChatFlowRequest {
	s.FlowCode = &v
	return s
}

func (s *TriggerChatFlowRequest) SetOutId(v string) *TriggerChatFlowRequest {
	s.OutId = &v
	return s
}

func (s *TriggerChatFlowRequest) SetOwnerId(v int64) *TriggerChatFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *TriggerChatFlowRequest) SetResourceOwnerAccount(v string) *TriggerChatFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TriggerChatFlowRequest) SetResourceOwnerId(v int64) *TriggerChatFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TriggerChatFlowRequest) SetUuid(v string) *TriggerChatFlowRequest {
	s.Uuid = &v
	return s
}

func (s *TriggerChatFlowRequest) Validate() error {
	return dara.Validate(s)
}
