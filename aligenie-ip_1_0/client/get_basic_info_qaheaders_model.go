// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicInfoQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetBasicInfoQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetBasicInfoQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetBasicInfoQAHeaders
	GetAuthorization() *string
}

type GetBasicInfoQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetBasicInfoQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetBasicInfoQAHeaders) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetBasicInfoQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetBasicInfoQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetBasicInfoQAHeaders) SetCommonHeaders(v map[string]*string) *GetBasicInfoQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBasicInfoQAHeaders) SetXAcsAligenieAccessToken(v string) *GetBasicInfoQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetBasicInfoQAHeaders) SetAuthorization(v string) *GetBasicInfoQAHeaders {
	s.Authorization = &v
	return s
}

func (s *GetBasicInfoQAHeaders) Validate() error {
	return dara.Validate(s)
}
