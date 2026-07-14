// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhatsappUserNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UpdateWhatsappUserNameRequest
	GetCustSpaceId() *string
	SetPhoneNumber(v string) *UpdateWhatsappUserNameRequest
	GetPhoneNumber() *string
	SetTransferAction(v string) *UpdateWhatsappUserNameRequest
	GetTransferAction() *string
	SetUsername(v string) *UpdateWhatsappUserNameRequest
	GetUsername() *string
}

type UpdateWhatsappUserNameRequest struct {
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the space ID on the
	//
	// <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)
	//
	// <props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList)
	//
	// page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-kskd****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The business phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800***
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The transfer action that controls what happens when the requested username is currently used by another business phone number in the same business asset portfolio. For example, use this parameter when you want to move an existing username to another phone number. Valid values:
	//
	//  - none (default): does not transfer the account. If another business phone number in the same business asset portfolio already uses this username, the request fails with error code 147005.
	//
	//  - force_transfer: transfers the account from the other business phone number to this business phone number. The account is removed from the other phone number and assigned to this phone number.
	//
	// example:
	//
	// none
	TransferAction *string `json:"TransferAction,omitempty" xml:"TransferAction,omitempty"`
	// Whatsapp user name
	//
	// This parameter is required.
	//
	// example:
	//
	// alibaba
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateWhatsappUserNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhatsappUserNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhatsappUserNameRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateWhatsappUserNameRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateWhatsappUserNameRequest) GetTransferAction() *string {
	return s.TransferAction
}

func (s *UpdateWhatsappUserNameRequest) GetUsername() *string {
	return s.Username
}

func (s *UpdateWhatsappUserNameRequest) SetCustSpaceId(v string) *UpdateWhatsappUserNameRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateWhatsappUserNameRequest) SetPhoneNumber(v string) *UpdateWhatsappUserNameRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateWhatsappUserNameRequest) SetTransferAction(v string) *UpdateWhatsappUserNameRequest {
	s.TransferAction = &v
	return s
}

func (s *UpdateWhatsappUserNameRequest) SetUsername(v string) *UpdateWhatsappUserNameRequest {
	s.Username = &v
	return s
}

func (s *UpdateWhatsappUserNameRequest) Validate() error {
	return dara.Validate(s)
}
