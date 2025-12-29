// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveChildAccountAuthHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveChildAccountAuthHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *RemoveChildAccountAuthHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *RemoveChildAccountAuthHeaders
	GetAuthorization() *string
}

type RemoveChildAccountAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveChildAccountAuthHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveChildAccountAuthHeaders) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveChildAccountAuthHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *RemoveChildAccountAuthHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RemoveChildAccountAuthHeaders) SetCommonHeaders(v map[string]*string) *RemoveChildAccountAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveChildAccountAuthHeaders) SetXAcsAligenieAccessToken(v string) *RemoveChildAccountAuthHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RemoveChildAccountAuthHeaders) SetAuthorization(v string) *RemoveChildAccountAuthHeaders {
	s.Authorization = &v
	return s
}

func (s *RemoveChildAccountAuthHeaders) Validate() error {
	return dara.Validate(s)
}
