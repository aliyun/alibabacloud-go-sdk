// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingRelationsForFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *ListBindingRelationsForFlowVersionRequest
	GetChannelType() *string
	SetFlowCode(v string) *ListBindingRelationsForFlowVersionRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *ListBindingRelationsForFlowVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListBindingRelationsForFlowVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListBindingRelationsForFlowVersionRequest
	GetResourceOwnerId() *int64
}

type ListBindingRelationsForFlowVersionRequest struct {
	// Channel type. Values:
	//
	// - INSTAGRAM
	//
	// - WHATSAPP
	//
	// - MESSENGER
	//
	//
	// <props="intl">- VIBER
	//
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Process code. View the process code in the [Flow Editor](https://chatapp.console.aliyun.com/ChatFlowBuilder) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode             *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListBindingRelationsForFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindingRelationsForFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *ListBindingRelationsForFlowVersionRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListBindingRelationsForFlowVersionRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ListBindingRelationsForFlowVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListBindingRelationsForFlowVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListBindingRelationsForFlowVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListBindingRelationsForFlowVersionRequest) SetChannelType(v string) *ListBindingRelationsForFlowVersionRequest {
	s.ChannelType = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionRequest) SetFlowCode(v string) *ListBindingRelationsForFlowVersionRequest {
	s.FlowCode = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionRequest) SetOwnerId(v int64) *ListBindingRelationsForFlowVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionRequest) SetResourceOwnerAccount(v string) *ListBindingRelationsForFlowVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionRequest) SetResourceOwnerId(v int64) *ListBindingRelationsForFlowVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListBindingRelationsForFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
