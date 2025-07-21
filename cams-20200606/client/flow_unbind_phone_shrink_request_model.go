// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowUnbindPhoneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *FlowUnbindPhoneShrinkRequest
	GetChannelType() *string
	SetFlowCode(v string) *FlowUnbindPhoneShrinkRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *FlowUnbindPhoneShrinkRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *FlowUnbindPhoneShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumbersShrink(v string) *FlowUnbindPhoneShrinkRequest
	GetPhoneNumbersShrink() *string
	SetResourceOwnerAccount(v string) *FlowUnbindPhoneShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlowUnbindPhoneShrinkRequest
	GetResourceOwnerId() *int64
}

type FlowUnbindPhoneShrinkRequest struct {
	// Message channel type
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
	// 示例值
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// Flow version
	//
	// example:
	//
	// 示例值示例值示例值
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone numbers or PageIds under the channel instance, etc.
	PhoneNumbersShrink   *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s FlowUnbindPhoneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlowUnbindPhoneShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlowUnbindPhoneShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *FlowUnbindPhoneShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *FlowUnbindPhoneShrinkRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowUnbindPhoneShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlowUnbindPhoneShrinkRequest) GetPhoneNumbersShrink() *string {
	return s.PhoneNumbersShrink
}

func (s *FlowUnbindPhoneShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlowUnbindPhoneShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlowUnbindPhoneShrinkRequest) SetChannelType(v string) *FlowUnbindPhoneShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) SetFlowCode(v string) *FlowUnbindPhoneShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) SetFlowVersion(v string) *FlowUnbindPhoneShrinkRequest {
	s.FlowVersion = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) SetOwnerId(v int64) *FlowUnbindPhoneShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) SetPhoneNumbersShrink(v string) *FlowUnbindPhoneShrinkRequest {
	s.PhoneNumbersShrink = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) SetResourceOwnerAccount(v string) *FlowUnbindPhoneShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) SetResourceOwnerId(v int64) *FlowUnbindPhoneShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlowUnbindPhoneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
