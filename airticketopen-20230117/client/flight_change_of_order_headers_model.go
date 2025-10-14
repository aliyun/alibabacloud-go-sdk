// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightChangeOfOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightChangeOfOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAirticketAccessToken(v string) *FlightChangeOfOrderHeaders
	GetXAcsAirticketAccessToken() *string
	SetXAcsAirticketLanguage(v string) *FlightChangeOfOrderHeaders
	GetXAcsAirticketLanguage() *string
}

type FlightChangeOfOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s FlightChangeOfOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightChangeOfOrderHeaders) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightChangeOfOrderHeaders) GetXAcsAirticketAccessToken() *string {
	return s.XAcsAirticketAccessToken
}

func (s *FlightChangeOfOrderHeaders) GetXAcsAirticketLanguage() *string {
	return s.XAcsAirticketLanguage
}

func (s *FlightChangeOfOrderHeaders) SetCommonHeaders(v map[string]*string) *FlightChangeOfOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightChangeOfOrderHeaders) SetXAcsAirticketAccessToken(v string) *FlightChangeOfOrderHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *FlightChangeOfOrderHeaders) SetXAcsAirticketLanguage(v string) *FlightChangeOfOrderHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

func (s *FlightChangeOfOrderHeaders) Validate() error {
	return dara.Validate(s)
}
