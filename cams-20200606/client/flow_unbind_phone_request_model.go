// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowUnbindPhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *FlowUnbindPhoneRequest
	GetChannelType() *string
	SetFlowCode(v string) *FlowUnbindPhoneRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *FlowUnbindPhoneRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *FlowUnbindPhoneRequest
	GetOwnerId() *int64
	SetPhoneNumbers(v []*string) *FlowUnbindPhoneRequest
	GetPhoneNumbers() []*string
	SetResourceOwnerAccount(v string) *FlowUnbindPhoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlowUnbindPhoneRequest
	GetResourceOwnerId() *int64
}

type FlowUnbindPhoneRequest struct {
	// The message channel type. Valid values:
	//
	// - INSTAGRAM
	//
	// - WHATSAPP
	//
	// - MESSENGER
	//
	// <props="intl">- VIBER
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The flow code. You can view the flow code on the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. You can view the flow version by going to the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) page, clicking the flow name, and entering the flow editor canvas page.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of phone numbers, PageIds, AccountIds,<props="intl"> or ServiceIds under the channel instance.
	PhoneNumbers         []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s FlowUnbindPhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s FlowUnbindPhoneRequest) GoString() string {
	return s.String()
}

func (s *FlowUnbindPhoneRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *FlowUnbindPhoneRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *FlowUnbindPhoneRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowUnbindPhoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlowUnbindPhoneRequest) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *FlowUnbindPhoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlowUnbindPhoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlowUnbindPhoneRequest) SetChannelType(v string) *FlowUnbindPhoneRequest {
	s.ChannelType = &v
	return s
}

func (s *FlowUnbindPhoneRequest) SetFlowCode(v string) *FlowUnbindPhoneRequest {
	s.FlowCode = &v
	return s
}

func (s *FlowUnbindPhoneRequest) SetFlowVersion(v string) *FlowUnbindPhoneRequest {
	s.FlowVersion = &v
	return s
}

func (s *FlowUnbindPhoneRequest) SetOwnerId(v int64) *FlowUnbindPhoneRequest {
	s.OwnerId = &v
	return s
}

func (s *FlowUnbindPhoneRequest) SetPhoneNumbers(v []*string) *FlowUnbindPhoneRequest {
	s.PhoneNumbers = v
	return s
}

func (s *FlowUnbindPhoneRequest) SetResourceOwnerAccount(v string) *FlowUnbindPhoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlowUnbindPhoneRequest) SetResourceOwnerId(v int64) *FlowUnbindPhoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlowUnbindPhoneRequest) Validate() error {
	return dara.Validate(s)
}
