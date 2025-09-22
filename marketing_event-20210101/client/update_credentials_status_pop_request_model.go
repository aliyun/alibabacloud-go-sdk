// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialsStatusPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCredentialsStatusPopRequest
	GetCode() *string
	SetProxyRecipientName(v string) *UpdateCredentialsStatusPopRequest
	GetProxyRecipientName() *string
	SetProxyRecipientPhoneNumber(v string) *UpdateCredentialsStatusPopRequest
	GetProxyRecipientPhoneNumber() *string
	SetReceiptLocation(v string) *UpdateCredentialsStatusPopRequest
	GetReceiptLocation() *string
	SetTime(v string) *UpdateCredentialsStatusPopRequest
	GetTime() *string
}

type UpdateCredentialsStatusPopRequest struct {
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 张三
	ProxyRecipientName *string `json:"ProxyRecipientName,omitempty" xml:"ProxyRecipientName,omitempty"`
	// example:
	//
	// 14787627188
	ProxyRecipientPhoneNumber *string `json:"ProxyRecipientPhoneNumber,omitempty" xml:"ProxyRecipientPhoneNumber,omitempty"`
	// example:
	//
	// Z30
	ReceiptLocation *string `json:"ReceiptLocation,omitempty" xml:"ReceiptLocation,omitempty"`
	// example:
	//
	// 429005111100000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s UpdateCredentialsStatusPopRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialsStatusPopRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialsStatusPopRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateCredentialsStatusPopRequest) GetProxyRecipientName() *string {
	return s.ProxyRecipientName
}

func (s *UpdateCredentialsStatusPopRequest) GetProxyRecipientPhoneNumber() *string {
	return s.ProxyRecipientPhoneNumber
}

func (s *UpdateCredentialsStatusPopRequest) GetReceiptLocation() *string {
	return s.ReceiptLocation
}

func (s *UpdateCredentialsStatusPopRequest) GetTime() *string {
	return s.Time
}

func (s *UpdateCredentialsStatusPopRequest) SetCode(v string) *UpdateCredentialsStatusPopRequest {
	s.Code = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetProxyRecipientName(v string) *UpdateCredentialsStatusPopRequest {
	s.ProxyRecipientName = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetProxyRecipientPhoneNumber(v string) *UpdateCredentialsStatusPopRequest {
	s.ProxyRecipientPhoneNumber = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetReceiptLocation(v string) *UpdateCredentialsStatusPopRequest {
	s.ReceiptLocation = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) SetTime(v string) *UpdateCredentialsStatusPopRequest {
	s.Time = &v
	return s
}

func (s *UpdateCredentialsStatusPopRequest) Validate() error {
	return dara.Validate(s)
}
