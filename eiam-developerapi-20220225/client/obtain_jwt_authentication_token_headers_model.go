// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ObtainJwtAuthenticationTokenHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ObtainJwtAuthenticationTokenHeaders
	GetAuthorization() *string
}

type ObtainJwtAuthenticationTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ObtainJwtAuthenticationTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenHeaders) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ObtainJwtAuthenticationTokenHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ObtainJwtAuthenticationTokenHeaders) SetCommonHeaders(v map[string]*string) *ObtainJwtAuthenticationTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ObtainJwtAuthenticationTokenHeaders) SetAuthorization(v string) *ObtainJwtAuthenticationTokenHeaders {
	s.Authorization = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenHeaders) Validate() error {
	return dara.Validate(s)
}
