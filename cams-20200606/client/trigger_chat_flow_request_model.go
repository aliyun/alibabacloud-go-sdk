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
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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
