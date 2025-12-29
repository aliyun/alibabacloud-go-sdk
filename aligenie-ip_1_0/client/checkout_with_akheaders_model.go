// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckoutWithAKHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckoutWithAKHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CheckoutWithAKHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CheckoutWithAKHeaders
	GetAuthorization() *string
}

type CheckoutWithAKHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckoutWithAKHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckoutWithAKHeaders) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckoutWithAKHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CheckoutWithAKHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CheckoutWithAKHeaders) SetCommonHeaders(v map[string]*string) *CheckoutWithAKHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckoutWithAKHeaders) SetXAcsAligenieAccessToken(v string) *CheckoutWithAKHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckoutWithAKHeaders) SetAuthorization(v string) *CheckoutWithAKHeaders {
	s.Authorization = &v
	return s
}

func (s *CheckoutWithAKHeaders) Validate() error {
	return dara.Validate(s)
}
