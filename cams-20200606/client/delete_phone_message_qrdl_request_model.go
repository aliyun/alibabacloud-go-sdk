// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhoneMessageQrdlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *DeletePhoneMessageQrdlRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *DeletePhoneMessageQrdlRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *DeletePhoneMessageQrdlRequest
	GetPhoneNumber() *string
	SetQrdlCode(v string) *DeletePhoneMessageQrdlRequest
	GetQrdlCode() *string
	SetResourceOwnerAccount(v string) *DeletePhoneMessageQrdlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePhoneMessageQrdlRequest
	GetResourceOwnerId() *int64
}

type DeletePhoneMessageQrdlRequest struct {
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// example:
	//
	// cams-kei****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number. You can view the phone number on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **WABA Management*	- > **Number Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86158********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The QR code encoding. This is the QrdlCode returned by the [CreatePhoneMessageQrdl](https://help.aliyun.com/document_detail/2638749.html) operation when you created the message QR code, or the QrdlCode returned by the [UpdatePhoneMessageQrdl](https://help.aliyun.com/document_detail/2638746.html) operation when you updated the QR code.
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

func (s DeletePhoneMessageQrdlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePhoneMessageQrdlRequest) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeletePhoneMessageQrdlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePhoneMessageQrdlRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DeletePhoneMessageQrdlRequest) GetQrdlCode() *string {
	return s.QrdlCode
}

func (s *DeletePhoneMessageQrdlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePhoneMessageQrdlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePhoneMessageQrdlRequest) SetCustSpaceId(v string) *DeletePhoneMessageQrdlRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetOwnerId(v int64) *DeletePhoneMessageQrdlRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetPhoneNumber(v string) *DeletePhoneMessageQrdlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetQrdlCode(v string) *DeletePhoneMessageQrdlRequest {
	s.QrdlCode = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetResourceOwnerAccount(v string) *DeletePhoneMessageQrdlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) SetResourceOwnerId(v int64) *DeletePhoneMessageQrdlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePhoneMessageQrdlRequest) Validate() error {
	return dara.Validate(s)
}
