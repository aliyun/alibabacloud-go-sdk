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
	SetMultiWabaPhoneNumbers(v []*FlowBindPhoneRequestMultiWabaPhoneNumbers) *FlowBindPhoneRequest
	GetMultiWabaPhoneNumbers() []*FlowBindPhoneRequestMultiWabaPhoneNumbers
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
	// The message channel code, which is the channel ID. View the channel ID in the [Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement) page.
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
	// The flow code. View the flow code in the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. Click the flow name in the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) page to enter the flow builder canvas and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The multi-WABA binding configuration.
	MultiWabaPhoneNumbers []*FlowBindPhoneRequestMultiWabaPhoneNumbers `json:"MultiWabaPhoneNumbers,omitempty" xml:"MultiWabaPhoneNumbers,omitempty" type:"Repeated"`
	OwnerId               *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of phone numbers, PageIds, or AccountIds<props="intl">, or ServiceIds under the channel instance.
	PhoneNumbers         []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The WABA account ID, PageId, or AccountId<props="intl">, or ServiceId.
	//
	// - If the ChannelType parameter is set to WHATSAPP, specify the WABA account ID. View the WABA account ID in [**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement) > **Manage*	- > **WABA Management**.
	//
	// - If the ChannelType parameter is not set to WHATSAPP, specify the PageId for MESSENGER, the AccountId for INSTAGRAM<props="intl">, or the ServiceId for VIBER.
	//
	// example:
	//
	// 1952************
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

func (s *FlowBindPhoneRequest) GetMultiWabaPhoneNumbers() []*FlowBindPhoneRequestMultiWabaPhoneNumbers {
	return s.MultiWabaPhoneNumbers
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

func (s *FlowBindPhoneRequest) SetMultiWabaPhoneNumbers(v []*FlowBindPhoneRequestMultiWabaPhoneNumbers) *FlowBindPhoneRequest {
	s.MultiWabaPhoneNumbers = v
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
	if s.MultiWabaPhoneNumbers != nil {
		for _, item := range s.MultiWabaPhoneNumbers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FlowBindPhoneRequestMultiWabaPhoneNumbers struct {
	// The channel code.
	//
	// example:
	//
	// 示例值示例值
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// The list of phone numbers.
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// wabaId
	//
	// example:
	//
	// 示例值示例值
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s FlowBindPhoneRequestMultiWabaPhoneNumbers) String() string {
	return dara.Prettify(s)
}

func (s FlowBindPhoneRequestMultiWabaPhoneNumbers) GoString() string {
	return s.String()
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) GetWabaId() *string {
	return s.WabaId
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) SetChannelCode(v string) *FlowBindPhoneRequestMultiWabaPhoneNumbers {
	s.ChannelCode = &v
	return s
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) SetPhoneNumbers(v []*string) *FlowBindPhoneRequestMultiWabaPhoneNumbers {
	s.PhoneNumbers = v
	return s
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) SetWabaId(v string) *FlowBindPhoneRequestMultiWabaPhoneNumbers {
	s.WabaId = &v
	return s
}

func (s *FlowBindPhoneRequestMultiWabaPhoneNumbers) Validate() error {
	return dara.Validate(s)
}
