// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserResponse
	GetStatusCode() *int32
	SetBody(v *ListUserResponseBody) *ListUserResponse
	GetBody() *ListUserResponseBody
}

type ListUserResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserResponse) GoString() string {
	return s.String()
}

func (s *ListUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserResponse) GetBody() *ListUserResponseBody {
	return s.Body
}

func (s *ListUserResponse) SetHeaders(v map[string]*string) *ListUserResponse {
	s.Headers = v
	return s
}

func (s *ListUserResponse) SetStatusCode(v int32) *ListUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserResponse) SetBody(v *ListUserResponseBody) *ListUserResponse {
	s.Body = v
	return s
}

func (s *ListUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
