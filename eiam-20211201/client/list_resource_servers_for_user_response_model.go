// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServersForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceServersForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceServersForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceServersForUserResponseBody) *ListResourceServersForUserResponse
	GetBody() *ListResourceServersForUserResponseBody
}

type ListResourceServersForUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceServersForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceServersForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServersForUserResponse) GoString() string {
	return s.String()
}

func (s *ListResourceServersForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceServersForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceServersForUserResponse) GetBody() *ListResourceServersForUserResponseBody {
	return s.Body
}

func (s *ListResourceServersForUserResponse) SetHeaders(v map[string]*string) *ListResourceServersForUserResponse {
	s.Headers = v
	return s
}

func (s *ListResourceServersForUserResponse) SetStatusCode(v int32) *ListResourceServersForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceServersForUserResponse) SetBody(v *ListResourceServersForUserResponseBody) *ListResourceServersForUserResponse {
	s.Body = v
	return s
}

func (s *ListResourceServersForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
