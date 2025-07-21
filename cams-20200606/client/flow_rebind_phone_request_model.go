// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowRebindPhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCode(v string) *FlowRebindPhoneRequest
	GetChannelCode() *string
	SetChannelType(v string) *FlowRebindPhoneRequest
	GetChannelType() *string
	SetFlowCode(v string) *FlowRebindPhoneRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *FlowRebindPhoneRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *FlowRebindPhoneRequest
	GetOwnerId() *int64
	SetPhoneNumbers(v []*string) *FlowRebindPhoneRequest
	GetPhoneNumbers() []*string
	SetResourceOwnerAccount(v string) *FlowRebindPhoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlowRebindPhoneRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *FlowRebindPhoneRequest
	GetWabaId() *string
}

type FlowRebindPhoneRequest struct {
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
	PhoneNumbers         []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// WABA account ID, or PageId for other channel types, etc.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s FlowRebindPhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s FlowRebindPhoneRequest) GoString() string {
	return s.String()
}

func (s *FlowRebindPhoneRequest) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *FlowRebindPhoneRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *FlowRebindPhoneRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *FlowRebindPhoneRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowRebindPhoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlowRebindPhoneRequest) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *FlowRebindPhoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlowRebindPhoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlowRebindPhoneRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *FlowRebindPhoneRequest) SetChannelCode(v string) *FlowRebindPhoneRequest {
	s.ChannelCode = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetChannelType(v string) *FlowRebindPhoneRequest {
	s.ChannelType = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetFlowCode(v string) *FlowRebindPhoneRequest {
	s.FlowCode = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetFlowVersion(v string) *FlowRebindPhoneRequest {
	s.FlowVersion = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetOwnerId(v int64) *FlowRebindPhoneRequest {
	s.OwnerId = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetPhoneNumbers(v []*string) *FlowRebindPhoneRequest {
	s.PhoneNumbers = v
	return s
}

func (s *FlowRebindPhoneRequest) SetResourceOwnerAccount(v string) *FlowRebindPhoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetResourceOwnerId(v int64) *FlowRebindPhoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlowRebindPhoneRequest) SetWabaId(v string) *FlowRebindPhoneRequest {
	s.WabaId = &v
	return s
}

func (s *FlowRebindPhoneRequest) Validate() error {
	return dara.Validate(s)
}
