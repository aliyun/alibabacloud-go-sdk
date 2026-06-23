// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizationServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizationServersResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizationServersResponseBody) *ListAuthorizationServersResponse
	GetBody() *ListAuthorizationServersResponseBody
}

type ListAuthorizationServersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationServersResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizationServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizationServersResponse) GetBody() *ListAuthorizationServersResponseBody {
	return s.Body
}

func (s *ListAuthorizationServersResponse) SetHeaders(v map[string]*string) *ListAuthorizationServersResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationServersResponse) SetStatusCode(v int32) *ListAuthorizationServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationServersResponse) SetBody(v *ListAuthorizationServersResponseBody) *ListAuthorizationServersResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizationServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
