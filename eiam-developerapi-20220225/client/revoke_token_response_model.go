// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeTokenResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *RevokeTokenResponse
	GetBody() map[string]interface{}
}

type RevokeTokenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenResponse) GoString() string {
	return s.String()
}

func (s *RevokeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeTokenResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *RevokeTokenResponse) SetHeaders(v map[string]*string) *RevokeTokenResponse {
	s.Headers = v
	return s
}

func (s *RevokeTokenResponse) SetStatusCode(v int32) *RevokeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTokenResponse) SetBody(v map[string]interface{}) *RevokeTokenResponse {
	s.Body = v
	return s
}

func (s *RevokeTokenResponse) Validate() error {
	return dara.Validate(s)
}
