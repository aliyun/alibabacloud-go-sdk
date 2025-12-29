// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChildAccountAuthHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ChildAccountAuthHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ChildAccountAuthHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ChildAccountAuthHeaders
	GetAuthorization() *string
}

type ChildAccountAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ChildAccountAuthHeaders) String() string {
	return dara.Prettify(s)
}

func (s ChildAccountAuthHeaders) GoString() string {
	return s.String()
}

func (s *ChildAccountAuthHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ChildAccountAuthHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ChildAccountAuthHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ChildAccountAuthHeaders) SetCommonHeaders(v map[string]*string) *ChildAccountAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChildAccountAuthHeaders) SetXAcsAligenieAccessToken(v string) *ChildAccountAuthHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ChildAccountAuthHeaders) SetAuthorization(v string) *ChildAccountAuthHeaders {
	s.Authorization = &v
	return s
}

func (s *ChildAccountAuthHeaders) Validate() error {
	return dara.Validate(s)
}
