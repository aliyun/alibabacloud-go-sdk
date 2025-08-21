// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeEnhanceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCodeEnhanceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetCodeEnhanceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetCodeEnhanceHeaders
	GetAuthorization() *string
}

type GetCodeEnhanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCodeEnhanceHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceHeaders) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCodeEnhanceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetCodeEnhanceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetCodeEnhanceHeaders) SetCommonHeaders(v map[string]*string) *GetCodeEnhanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCodeEnhanceHeaders) SetXAcsAligenieAccessToken(v string) *GetCodeEnhanceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCodeEnhanceHeaders) SetAuthorization(v string) *GetCodeEnhanceHeaders {
	s.Authorization = &v
	return s
}

func (s *GetCodeEnhanceHeaders) Validate() error {
	return dara.Validate(s)
}
