// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserVpcAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserVpcAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserVpcAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *AddUserVpcAuthorizationResponseBody) *AddUserVpcAuthorizationResponse
	GetBody() *AddUserVpcAuthorizationResponseBody
}

type AddUserVpcAuthorizationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserVpcAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserVpcAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserVpcAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *AddUserVpcAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserVpcAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserVpcAuthorizationResponse) GetBody() *AddUserVpcAuthorizationResponseBody {
	return s.Body
}

func (s *AddUserVpcAuthorizationResponse) SetHeaders(v map[string]*string) *AddUserVpcAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *AddUserVpcAuthorizationResponse) SetStatusCode(v int32) *AddUserVpcAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserVpcAuthorizationResponse) SetBody(v *AddUserVpcAuthorizationResponseBody) *AddUserVpcAuthorizationResponse {
	s.Body = v
	return s
}

func (s *AddUserVpcAuthorizationResponse) Validate() error {
	return dara.Validate(s)
}
