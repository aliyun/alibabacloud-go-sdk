// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowBindPhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelCode(v string) *FlowBindPhoneRequest
	GetChannelCode() *string
	SetChannelType(v string) *FlowBindPhoneRequest
	GetChannelType() *string
	SetFlowCode(v string) *FlowBindPhoneRequest
	GetFlowCode() *string
	SetFlowVersion(v string) *FlowBindPhoneRequest
	GetFlowVersion() *string
	SetOwnerId(v int64) *FlowBindPhoneRequest
	GetOwnerId() *int64
	SetPhoneNumbers(v []*string) *FlowBindPhoneRequest
	GetPhoneNumbers() []*string
	SetResourceOwnerAccount(v string) *FlowBindPhoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FlowBindPhoneRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *FlowBindPhoneRequest
	GetWabaId() *string
}

type FlowBindPhoneRequest struct {
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
	PhoneNumbers         []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// WABA account ID, or PageId for other channel types, etc.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s FlowBindPhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s FlowBindPhoneRequest) GoString() string {
	return s.String()
}

func (s *FlowBindPhoneRequest) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *FlowBindPhoneRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *FlowBindPhoneRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *FlowBindPhoneRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowBindPhoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FlowBindPhoneRequest) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *FlowBindPhoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FlowBindPhoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FlowBindPhoneRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *FlowBindPhoneRequest) SetChannelCode(v string) *FlowBindPhoneRequest {
	s.ChannelCode = &v
	return s
}

func (s *FlowBindPhoneRequest) SetChannelType(v string) *FlowBindPhoneRequest {
	s.ChannelType = &v
	return s
}

func (s *FlowBindPhoneRequest) SetFlowCode(v string) *FlowBindPhoneRequest {
	s.FlowCode = &v
	return s
}

func (s *FlowBindPhoneRequest) SetFlowVersion(v string) *FlowBindPhoneRequest {
	s.FlowVersion = &v
	return s
}

func (s *FlowBindPhoneRequest) SetOwnerId(v int64) *FlowBindPhoneRequest {
	s.OwnerId = &v
	return s
}

func (s *FlowBindPhoneRequest) SetPhoneNumbers(v []*string) *FlowBindPhoneRequest {
	s.PhoneNumbers = v
	return s
}

func (s *FlowBindPhoneRequest) SetResourceOwnerAccount(v string) *FlowBindPhoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FlowBindPhoneRequest) SetResourceOwnerId(v int64) *FlowBindPhoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FlowBindPhoneRequest) SetWabaId(v string) *FlowBindPhoneRequest {
	s.WabaId = &v
	return s
}

func (s *FlowBindPhoneRequest) Validate() error {
	return dara.Validate(s)
}
