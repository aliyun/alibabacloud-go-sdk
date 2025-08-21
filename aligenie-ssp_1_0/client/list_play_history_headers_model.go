// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlayHistoryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListPlayHistoryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListPlayHistoryHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListPlayHistoryHeaders
	GetAuthorization() *string
}

type ListPlayHistoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListPlayHistoryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListPlayHistoryHeaders) GoString() string {
	return s.String()
}

func (s *ListPlayHistoryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListPlayHistoryHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListPlayHistoryHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListPlayHistoryHeaders) SetCommonHeaders(v map[string]*string) *ListPlayHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPlayHistoryHeaders) SetXAcsAligenieAccessToken(v string) *ListPlayHistoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListPlayHistoryHeaders) SetAuthorization(v string) *ListPlayHistoryHeaders {
	s.Authorization = &v
	return s
}

func (s *ListPlayHistoryHeaders) Validate() error {
	return dara.Validate(s)
}
