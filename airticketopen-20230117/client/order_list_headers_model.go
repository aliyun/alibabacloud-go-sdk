// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *OrderListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAirticketAccessToken(v string) *OrderListHeaders
	GetXAcsAirticketAccessToken() *string
	SetXAcsAirticketLanguage(v string) *OrderListHeaders
	GetXAcsAirticketLanguage() *string
}

type OrderListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s OrderListHeaders) String() string {
	return dara.Prettify(s)
}

func (s OrderListHeaders) GoString() string {
	return s.String()
}

func (s *OrderListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *OrderListHeaders) GetXAcsAirticketAccessToken() *string {
	return s.XAcsAirticketAccessToken
}

func (s *OrderListHeaders) GetXAcsAirticketLanguage() *string {
	return s.XAcsAirticketLanguage
}

func (s *OrderListHeaders) SetCommonHeaders(v map[string]*string) *OrderListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrderListHeaders) SetXAcsAirticketAccessToken(v string) *OrderListHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *OrderListHeaders) SetXAcsAirticketLanguage(v string) *OrderListHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

func (s *OrderListHeaders) Validate() error {
	return dara.Validate(s)
}
