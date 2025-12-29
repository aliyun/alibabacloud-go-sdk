// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDisPlayModesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddOrUpdateDisPlayModesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddOrUpdateDisPlayModesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddOrUpdateDisPlayModesHeaders
	GetAuthorization() *string
}

type AddOrUpdateDisPlayModesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateDisPlayModesHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDisPlayModesHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddOrUpdateDisPlayModesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddOrUpdateDisPlayModesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddOrUpdateDisPlayModesHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateDisPlayModesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateDisPlayModesHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateDisPlayModesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateDisPlayModesHeaders) SetAuthorization(v string) *AddOrUpdateDisPlayModesHeaders {
	s.Authorization = &v
	return s
}

func (s *AddOrUpdateDisPlayModesHeaders) Validate() error {
	return dara.Validate(s)
}
