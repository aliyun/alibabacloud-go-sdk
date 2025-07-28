// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserVpcAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserVpcAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserVpcAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserVpcAuthorizationResponseBody) *DeleteUserVpcAuthorizationResponse
	GetBody() *DeleteUserVpcAuthorizationResponseBody
}

type DeleteUserVpcAuthorizationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserVpcAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserVpcAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserVpcAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserVpcAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserVpcAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserVpcAuthorizationResponse) GetBody() *DeleteUserVpcAuthorizationResponseBody {
	return s.Body
}

func (s *DeleteUserVpcAuthorizationResponse) SetHeaders(v map[string]*string) *DeleteUserVpcAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserVpcAuthorizationResponse) SetStatusCode(v int32) *DeleteUserVpcAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserVpcAuthorizationResponse) SetBody(v *DeleteUserVpcAuthorizationResponseBody) *DeleteUserVpcAuthorizationResponse {
	s.Body = v
	return s
}

func (s *DeleteUserVpcAuthorizationResponse) Validate() error {
	return dara.Validate(s)
}
