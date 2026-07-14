// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneMessageQrdlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UpdatePhoneMessageQrdlRequest
	GetCustSpaceId() *string
	SetGenerateQrImage(v string) *UpdatePhoneMessageQrdlRequest
	GetGenerateQrImage() *string
	SetOwnerId(v int64) *UpdatePhoneMessageQrdlRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *UpdatePhoneMessageQrdlRequest
	GetPhoneNumber() *string
	SetPrefilledMessage(v string) *UpdatePhoneMessageQrdlRequest
	GetPrefilledMessage() *string
	SetQrdlCode(v string) *UpdatePhoneMessageQrdlRequest
	GetQrdlCode() *string
	SetResourceOwnerAccount(v string) *UpdatePhoneMessageQrdlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdatePhoneMessageQrdlRequest
	GetResourceOwnerId() *int64
}

type UpdatePhoneMessageQrdlRequest struct {
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the
	//
	// <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-kei****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The type of the generated image. Valid values:
	//
	// - PNG
	//
	//  - SVG
	//
	// This parameter is required.
	//
	// example:
	//
	// SVG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number. You can view the phone number on the
	//
	// <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)
	//
	// <props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList)> **Management*	- > **WABA Management*	- > **Number Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86158********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The message content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// The QR code encoding. This is the QrdlCode returned when you called the [CreatePhoneMessageQrdl](https://help.aliyun.com/document_detail/2638749.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// D9II3***
	QrdlCode             *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdatePhoneMessageQrdlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdatePhoneMessageQrdlRequest) GetGenerateQrImage() *string {
	return s.GenerateQrImage
}

func (s *UpdatePhoneMessageQrdlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdatePhoneMessageQrdlRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdatePhoneMessageQrdlRequest) GetPrefilledMessage() *string {
	return s.PrefilledMessage
}

func (s *UpdatePhoneMessageQrdlRequest) GetQrdlCode() *string {
	return s.QrdlCode
}

func (s *UpdatePhoneMessageQrdlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdatePhoneMessageQrdlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdatePhoneMessageQrdlRequest) SetCustSpaceId(v string) *UpdatePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetGenerateQrImage(v string) *UpdatePhoneMessageQrdlRequest {
	s.GenerateQrImage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetOwnerId(v int64) *UpdatePhoneMessageQrdlRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetPhoneNumber(v string) *UpdatePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetPrefilledMessage(v string) *UpdatePhoneMessageQrdlRequest {
	s.PrefilledMessage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetQrdlCode(v string) *UpdatePhoneMessageQrdlRequest {
	s.QrdlCode = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetResourceOwnerAccount(v string) *UpdatePhoneMessageQrdlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) SetResourceOwnerId(v int64) *UpdatePhoneMessageQrdlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePhoneMessageQrdlRequest) Validate() error {
	return dara.Validate(s)
}
