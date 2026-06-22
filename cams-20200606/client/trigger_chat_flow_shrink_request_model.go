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
	// The time when the event occurs. This is when the flow is triggered and is typically the time when the request is created. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1731502129000
	ClaimTimeMillis *int64 `json:"ClaimTimeMillis,omitempty" xml:"ClaimTimeMillis,omitempty"`
	// The input parameters in a key-value format. The keys must match the input parameter policy configured in the start node of the flow. To view the variable names in the start node, go to the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder), click the name of the flow, and open the orchestration canvas.
	//
	// example:
	//
	// {
	//
	//   "my_biz_data_0": "hi",
	//
	//   "my_biz_data_1": "1024"
	//
	// }
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The time when the event expires. If you specify this parameter, the trigger is canceled if the request is not processed before this time. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1731502729000
	DiscardTimeMillis *int64 `json:"DiscardTimeMillis,omitempty" xml:"DiscardTimeMillis,omitempty"`
	// The code of the flow. View the flow code on the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// A custom serial number from an external system. Use this parameter to associate the trigger with an external business process. After the flow is triggered, you can retrieve this parameter from within the flow.
	//
	// example:
	//
	// 8d4acf7e-e360-eb83-6d74-fcf9c4538fda
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A custom unique ID for the event, used to ensure idempotence. Do not include business semantics in the ID. After the flow is triggered, you can retrieve this parameter from within the flow.
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
