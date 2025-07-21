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
	// Message channel code
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// Message channel type
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Flow code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// Flow version
	//
	// example:
	//
	// 示例值
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone numbers or PageIds under the channel instance, etc.
	PhoneNumbersShrink   *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// WABA account ID, or PageId for other channel types, etc.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
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
