// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstateAuthenticationTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ReinstateAuthenticationTokenHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ReinstateAuthenticationTokenHeaders
	GetAuthorization() *string
}

type ReinstateAuthenticationTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ReinstateAuthenticationTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s ReinstateAuthenticationTokenHeaders) GoString() string {
	return s.String()
}

func (s *ReinstateAuthenticationTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ReinstateAuthenticationTokenHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ReinstateAuthenticationTokenHeaders) SetCommonHeaders(v map[string]*string) *ReinstateAuthenticationTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReinstateAuthenticationTokenHeaders) SetAuthorization(v string) *ReinstateAuthenticationTokenHeaders {
	s.Authorization = &v
	return s
}

func (s *ReinstateAuthenticationTokenHeaders) Validate() error {
	return dara.Validate(s)
}
