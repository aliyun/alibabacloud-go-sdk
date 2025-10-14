// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPricingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PricingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAirticketAccessToken(v string) *PricingHeaders
	GetXAcsAirticketAccessToken() *string
	SetXAcsAirticketLanguage(v string) *PricingHeaders
	GetXAcsAirticketLanguage() *string
}

type PricingHeaders struct {
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

func (s PricingHeaders) String() string {
	return dara.Prettify(s)
}

func (s PricingHeaders) GoString() string {
	return s.String()
}

func (s *PricingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PricingHeaders) GetXAcsAirticketAccessToken() *string {
	return s.XAcsAirticketAccessToken
}

func (s *PricingHeaders) GetXAcsAirticketLanguage() *string {
	return s.XAcsAirticketLanguage
}

func (s *PricingHeaders) SetCommonHeaders(v map[string]*string) *PricingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PricingHeaders) SetXAcsAirticketAccessToken(v string) *PricingHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *PricingHeaders) SetXAcsAirticketLanguage(v string) *PricingHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

func (s *PricingHeaders) Validate() error {
	return dara.Validate(s)
}
