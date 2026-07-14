// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappUserNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetWhatsappUserNameRequest
	GetCustSpaceId() *string
	SetPhoneNumber(v string) *GetWhatsappUserNameRequest
	GetPhoneNumber() *string
}

type GetWhatsappUserNameRequest struct {
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the [Channel Management](https://chatapp.console.aliyun.com/CustomerList) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-kd***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The business phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800**
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetWhatsappUserNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappUserNameRequest) GoString() string {
	return s.String()
}

func (s *GetWhatsappUserNameRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetWhatsappUserNameRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetWhatsappUserNameRequest) SetCustSpaceId(v string) *GetWhatsappUserNameRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetWhatsappUserNameRequest) SetPhoneNumber(v string) *GetWhatsappUserNameRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetWhatsappUserNameRequest) Validate() error {
	return dara.Validate(s)
}
