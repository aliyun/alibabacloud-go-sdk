// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateTicketHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateTicketHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateTicketHeaders
	GetAuthorization() *string
}

type UpdateTicketHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateTicketHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTicketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateTicketHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateTicketHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateTicketHeaders) SetCommonHeaders(v map[string]*string) *UpdateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTicketHeaders) SetXAcsAligenieAccessToken(v string) *UpdateTicketHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateTicketHeaders) SetAuthorization(v string) *UpdateTicketHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateTicketHeaders) Validate() error {
	return dara.Validate(s)
}
