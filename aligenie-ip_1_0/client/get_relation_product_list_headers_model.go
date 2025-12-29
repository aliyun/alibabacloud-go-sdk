// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelationProductListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRelationProductListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetRelationProductListHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetRelationProductListHeaders
	GetAuthorization() *string
}

type GetRelationProductListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetRelationProductListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRelationProductListHeaders) GoString() string {
	return s.String()
}

func (s *GetRelationProductListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRelationProductListHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetRelationProductListHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetRelationProductListHeaders) SetCommonHeaders(v map[string]*string) *GetRelationProductListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelationProductListHeaders) SetXAcsAligenieAccessToken(v string) *GetRelationProductListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetRelationProductListHeaders) SetAuthorization(v string) *GetRelationProductListHeaders {
	s.Authorization = &v
	return s
}

func (s *GetRelationProductListHeaders) Validate() error {
	return dara.Validate(s)
}
