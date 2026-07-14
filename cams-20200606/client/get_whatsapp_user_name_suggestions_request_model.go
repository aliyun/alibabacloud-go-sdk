// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappUserNameSuggestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetWhatsappUserNameSuggestionsRequest
	GetCustSpaceId() *string
	SetPhoneNumber(v string) *GetWhatsappUserNameSuggestionsRequest
	GetPhoneNumber() *string
}

type GetWhatsappUserNameSuggestionsRequest struct {
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
	// cams-dk**
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

func (s GetWhatsappUserNameSuggestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappUserNameSuggestionsRequest) GoString() string {
	return s.String()
}

func (s *GetWhatsappUserNameSuggestionsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetWhatsappUserNameSuggestionsRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetWhatsappUserNameSuggestionsRequest) SetCustSpaceId(v string) *GetWhatsappUserNameSuggestionsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetWhatsappUserNameSuggestionsRequest) SetPhoneNumber(v string) *GetWhatsappUserNameSuggestionsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetWhatsappUserNameSuggestionsRequest) Validate() error {
	return dara.Validate(s)
}
