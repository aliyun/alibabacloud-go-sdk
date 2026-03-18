// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstateAuthenticationTokenByConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReinstateAuthenticationTokenByConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReinstateAuthenticationTokenByConsumerResponse
	GetStatusCode() *int32
}

type ReinstateAuthenticationTokenByConsumerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ReinstateAuthenticationTokenByConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s ReinstateAuthenticationTokenByConsumerResponse) GoString() string {
	return s.String()
}

func (s *ReinstateAuthenticationTokenByConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReinstateAuthenticationTokenByConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReinstateAuthenticationTokenByConsumerResponse) SetHeaders(v map[string]*string) *ReinstateAuthenticationTokenByConsumerResponse {
	s.Headers = v
	return s
}

func (s *ReinstateAuthenticationTokenByConsumerResponse) SetStatusCode(v int32) *ReinstateAuthenticationTokenByConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinstateAuthenticationTokenByConsumerResponse) Validate() error {
	return dara.Validate(s)
}
