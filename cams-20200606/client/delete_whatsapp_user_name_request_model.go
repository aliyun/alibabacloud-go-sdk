// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhatsappUserNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *DeleteWhatsappUserNameRequest
	GetCustSpaceId() *string
	SetPhoneNumber(v string) *DeleteWhatsappUserNameRequest
	GetPhoneNumber() *string
}

type DeleteWhatsappUserNameRequest struct {
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the
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
	// cams-xkd***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The business phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800***
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s DeleteWhatsappUserNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhatsappUserNameRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhatsappUserNameRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteWhatsappUserNameRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DeleteWhatsappUserNameRequest) SetCustSpaceId(v string) *DeleteWhatsappUserNameRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteWhatsappUserNameRequest) SetPhoneNumber(v string) *DeleteWhatsappUserNameRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DeleteWhatsappUserNameRequest) Validate() error {
	return dara.Validate(s)
}
