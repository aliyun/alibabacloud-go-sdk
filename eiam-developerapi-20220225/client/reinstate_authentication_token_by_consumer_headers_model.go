// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstateAuthenticationTokenByConsumerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ReinstateAuthenticationTokenByConsumerHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ReinstateAuthenticationTokenByConsumerHeaders
	GetAuthorization() *string
}

type ReinstateAuthenticationTokenByConsumerHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ReinstateAuthenticationTokenByConsumerHeaders) String() string {
	return dara.Prettify(s)
}

func (s ReinstateAuthenticationTokenByConsumerHeaders) GoString() string {
	return s.String()
}

func (s *ReinstateAuthenticationTokenByConsumerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ReinstateAuthenticationTokenByConsumerHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ReinstateAuthenticationTokenByConsumerHeaders) SetCommonHeaders(v map[string]*string) *ReinstateAuthenticationTokenByConsumerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReinstateAuthenticationTokenByConsumerHeaders) SetAuthorization(v string) *ReinstateAuthenticationTokenByConsumerHeaders {
	s.Authorization = &v
	return s
}

func (s *ReinstateAuthenticationTokenByConsumerHeaders) Validate() error {
	return dara.Validate(s)
}
