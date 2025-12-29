// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicInfoQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateBasicInfoQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateBasicInfoQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateBasicInfoQAHeaders
	GetAuthorization() *string
}

type UpdateBasicInfoQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateBasicInfoQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicInfoQAHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateBasicInfoQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateBasicInfoQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateBasicInfoQAHeaders) SetCommonHeaders(v map[string]*string) *UpdateBasicInfoQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBasicInfoQAHeaders) SetXAcsAligenieAccessToken(v string) *UpdateBasicInfoQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateBasicInfoQAHeaders) SetAuthorization(v string) *UpdateBasicInfoQAHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateBasicInfoQAHeaders) Validate() error {
	return dara.Validate(s)
}
