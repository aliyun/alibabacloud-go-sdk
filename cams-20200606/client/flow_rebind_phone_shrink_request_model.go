// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowRebindPhoneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCode(v string) *FlowRebindPhoneShrinkRequest
	GetChannelCode() *string
	SetChannelType(v string) *FlowRebindPhoneShrinkRequest
	GetChannelType() *string
	SetFlowCode(v string) *FlowRebindPhoneShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *FlowRebindPhoneShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *FlowRebindPhoneShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumbersShrink(v string) *FlowRebindPhoneShrinkRequest
	GetPhoneNumbersShrink() *string
	SetResourceOwnerAccount(v string) *FlowRebindPhoneShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlowRebindPhoneShrinkRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *FlowRebindPhoneShrinkRequest
	GetWabaId() *string
}

type FlowRebindPhoneShrinkRequest struct {
	// The message channel code, which is the channel ID. You can view the channel ID on the [Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-8c8*********
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The flow code. You can view it on the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. On the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) page, click the flow name to open the flow builder canvas and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of phone numbers, PageIds, or AccountIds<props="intl">, or ServiceIds under the channel instance.
	PhoneNumbersShrink   *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The WABA account ID, PageId, or AccountId<props="intl">, or ServiceId.
	//
	// - If ChannelType is set to WHATSAPP, specify the WABA account ID. You can view the WABA account ID on the Channel Management > Manage > WABA Management page.
	//
	// - If ChannelType is not set to WHATSAPP, specify the PageId for MESSENGER, the AccountId for INSTAGRAM<props="intl">, or the ServiceId for VIBER.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1952************
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s FlowRebindPhoneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlowRebindPhoneShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlowRebindPhoneShrinkRequest) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *FlowRebindPhoneShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *FlowRebindPhoneShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *FlowRebindPhoneShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowRebindPhoneShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlowRebindPhoneShrinkRequest) GetPhoneNumbersShrink() *string {
	return s.PhoneNumbersShrink
}

func (s *FlowRebindPhoneShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlowRebindPhoneShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlowRebindPhoneShrinkRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *FlowRebindPhoneShrinkRequest) SetChannelCode(v string) *FlowRebindPhoneShrinkRequest {
	s.ChannelCode = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetChannelType(v string) *FlowRebindPhoneShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetFlowCode(v string) *FlowRebindPhoneShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetFlowVersion(v string) *FlowRebindPhoneShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetOwnerId(v int64) *FlowRebindPhoneShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetPhoneNumbersShrink(v string) *FlowRebindPhoneShrinkRequest {
	s.PhoneNumbersShrink = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetResourceOwnerAccount(v string) *FlowRebindPhoneShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetResourceOwnerId(v int64) *FlowRebindPhoneShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) SetWabaId(v string) *FlowRebindPhoneShrinkRequest {
	s.WabaId = &v
	return s
}

func (s *FlowRebindPhoneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
