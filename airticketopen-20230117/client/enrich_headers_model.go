// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrichHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EnrichHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsAirticketAccessToken(v string) *EnrichHeaders
  GetXAcsAirticketAccessToken() *string 
  SetXAcsAirticketLanguage(v string) *EnrichHeaders
  GetXAcsAirticketLanguage() *string 
}

type EnrichHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // access_token
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
  XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
  // Multi-language, default is based on the buyer\\"s account configuration
  // 
  // example:
  // 
  // en_US
  XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s EnrichHeaders) String() string {
  return dara.Prettify(s)
}

func (s EnrichHeaders) GoString() string {
  return s.String()
}

func (s *EnrichHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EnrichHeaders) GetXAcsAirticketAccessToken() *string  {
  return s.XAcsAirticketAccessToken
}

func (s *EnrichHeaders) GetXAcsAirticketLanguage() *string  {
  return s.XAcsAirticketLanguage
}

func (s *EnrichHeaders) SetCommonHeaders(v map[string]*string) *EnrichHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EnrichHeaders) SetXAcsAirticketAccessToken(v string) *EnrichHeaders {
  s.XAcsAirticketAccessToken = &v
  return s
}

func (s *EnrichHeaders) SetXAcsAirticketLanguage(v string) *EnrichHeaders {
  s.XAcsAirticketLanguage = &v
  return s
}

func (s *EnrichHeaders) Validate() error {
  return dara.Validate(s)
}

