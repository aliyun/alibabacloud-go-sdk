// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowBindPhoneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCode(v string) *FlowBindPhoneShrinkRequest
	GetChannelCode() *string
	SetChannelType(v string) *FlowBindPhoneShrinkRequest
	GetChannelType() *string
	SetFlowCode(v string) *FlowBindPhoneShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *FlowBindPhoneShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *FlowBindPhoneShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumbersShrink(v string) *FlowBindPhoneShrinkRequest
	GetPhoneNumbersShrink() *string
	SetResourceOwnerAccount(v string) *FlowBindPhoneShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlowBindPhoneShrinkRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *FlowBindPhoneShrinkRequest
	GetWabaId() *string
}

type FlowBindPhoneShrinkRequest struct {
	// Message channel Code
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// Message channel Type
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
	// 1
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
	// 示例值
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s FlowBindPhoneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlowBindPhoneShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlowBindPhoneShrinkRequest) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *FlowBindPhoneShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *FlowBindPhoneShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *FlowBindPhoneShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowBindPhoneShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlowBindPhoneShrinkRequest) GetPhoneNumbersShrink() *string {
	return s.PhoneNumbersShrink
}

func (s *FlowBindPhoneShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlowBindPhoneShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlowBindPhoneShrinkRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *FlowBindPhoneShrinkRequest) SetChannelCode(v string) *FlowBindPhoneShrinkRequest {
	s.ChannelCode = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetChannelType(v string) *FlowBindPhoneShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetFlowCode(v string) *FlowBindPhoneShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetFlowVersion(v string) *FlowBindPhoneShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetOwnerId(v int64) *FlowBindPhoneShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetPhoneNumbersShrink(v string) *FlowBindPhoneShrinkRequest {
	s.PhoneNumbersShrink = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetResourceOwnerAccount(v string) *FlowBindPhoneShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetResourceOwnerId(v int64) *FlowBindPhoneShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) SetWabaId(v string) *FlowBindPhoneShrinkRequest {
	s.WabaId = &v
	return s
}

func (s *FlowBindPhoneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
