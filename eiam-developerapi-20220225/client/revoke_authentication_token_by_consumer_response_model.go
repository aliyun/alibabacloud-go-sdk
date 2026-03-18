// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAuthenticationTokenByConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAuthenticationTokenByConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAuthenticationTokenByConsumerResponse
	GetStatusCode() *int32
}

type RevokeAuthenticationTokenByConsumerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RevokeAuthenticationTokenByConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAuthenticationTokenByConsumerResponse) GoString() string {
	return s.String()
}

func (s *RevokeAuthenticationTokenByConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAuthenticationTokenByConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAuthenticationTokenByConsumerResponse) SetHeaders(v map[string]*string) *RevokeAuthenticationTokenByConsumerResponse {
	s.Headers = v
	return s
}

func (s *RevokeAuthenticationTokenByConsumerResponse) SetStatusCode(v int32) *RevokeAuthenticationTokenByConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAuthenticationTokenByConsumerResponse) Validate() error {
	return dara.Validate(s)
}
