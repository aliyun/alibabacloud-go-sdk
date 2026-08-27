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
	SetMultiWabaPhoneNumbers(v []*FlowRebindPhoneRequestMultiWabaPhoneNumbers) *FlowRebindPhoneRequest
	GetMultiWabaPhoneNumbers() []*FlowRebindPhoneRequestMultiWabaPhoneNumbers
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
	// The message channel code, which is the channel ID. You can view the channel ID in the [Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement) console.
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
	// The flow code. You can view this in the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow version. In the [Flow Builder](https://chatapp.console.aliyun.com/ChatFlowBuilder) console, click the flow name to open the flow editor canvas and view the flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The multi-WABA binding configurations.
	MultiWabaPhoneNumbers []*FlowRebindPhoneRequestMultiWabaPhoneNumbers `json:"MultiWabaPhoneNumbers,omitempty" xml:"MultiWabaPhoneNumbers,omitempty" type:"Repeated"`
	OwnerId               *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of phone numbers, PageIds, or AccountIds<props="intl">, or ServiceIds under the channel instance.
	PhoneNumbers         []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The WABA account ID, PageId, or AccountId<props="intl">, or ServiceId.
	//
	// - If ChannelType is set to WHATSAPP, specify the WABA account ID. You can view the WABA account ID in Channel Management > Manage > WABA Management.
	//
	// - If ChannelType is not set to WHATSAPP, specify the PageId for MESSENGER, the AccountId for INSTAGRAM<props="intl">, or the ServiceId for VIBER.
	//
	// example:
	//
	// 1952************
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

func (s *FlowRebindPhoneRequest) GetMultiWabaPhoneNumbers() []*FlowRebindPhoneRequestMultiWabaPhoneNumbers {
	return s.MultiWabaPhoneNumbers
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

func (s *FlowRebindPhoneRequest) SetMultiWabaPhoneNumbers(v []*FlowRebindPhoneRequestMultiWabaPhoneNumbers) *FlowRebindPhoneRequest {
	s.MultiWabaPhoneNumbers = v
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

type FlowRebindPhoneRequestMultiWabaPhoneNumbers struct {
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

func (s FlowRebindPhoneRequestMultiWabaPhoneNumbers) String() string {
	return dara.Prettify(s)
}

func (s FlowRebindPhoneRequestMultiWabaPhoneNumbers) GoString() string {
	return s.String()
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) GetWabaId() *string {
	return s.WabaId
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) SetChannelCode(v string) *FlowRebindPhoneRequestMultiWabaPhoneNumbers {
	s.ChannelCode = &v
	return s
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) SetPhoneNumbers(v []*string) *FlowRebindPhoneRequestMultiWabaPhoneNumbers {
	s.PhoneNumbers = v
	return s
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) SetWabaId(v string) *FlowRebindPhoneRequestMultiWabaPhoneNumbers {
	s.WabaId = &v
	return s
}

func (s *FlowRebindPhoneRequestMultiWabaPhoneNumbers) Validate() error {
	return dara.Validate(s)
}
