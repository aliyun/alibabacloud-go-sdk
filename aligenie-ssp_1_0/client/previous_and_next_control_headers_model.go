// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviousAndNextControlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PreviousAndNextControlHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PreviousAndNextControlHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PreviousAndNextControlHeaders
	GetAuthorization() *string
}

type PreviousAndNextControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PreviousAndNextControlHeaders) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlHeaders) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PreviousAndNextControlHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PreviousAndNextControlHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PreviousAndNextControlHeaders) SetCommonHeaders(v map[string]*string) *PreviousAndNextControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreviousAndNextControlHeaders) SetXAcsAligenieAccessToken(v string) *PreviousAndNextControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PreviousAndNextControlHeaders) SetAuthorization(v string) *PreviousAndNextControlHeaders {
	s.Authorization = &v
	return s
}

func (s *PreviousAndNextControlHeaders) Validate() error {
	return dara.Validate(s)
}
