// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullCashierHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PullCashierHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PullCashierHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PullCashierHeaders
	GetAuthorization() *string
}

type PullCashierHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PullCashierHeaders) String() string {
	return dara.Prettify(s)
}

func (s PullCashierHeaders) GoString() string {
	return s.String()
}

func (s *PullCashierHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PullCashierHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PullCashierHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PullCashierHeaders) SetCommonHeaders(v map[string]*string) *PullCashierHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullCashierHeaders) SetXAcsAligenieAccessToken(v string) *PullCashierHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PullCashierHeaders) SetAuthorization(v string) *PullCashierHeaders {
	s.Authorization = &v
	return s
}

func (s *PullCashierHeaders) Validate() error {
	return dara.Validate(s)
}
