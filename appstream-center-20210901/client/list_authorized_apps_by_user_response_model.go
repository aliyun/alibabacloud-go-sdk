// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAppsByUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedAppsByUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedAppsByUserResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedAppsByUserResponseBody) *ListAuthorizedAppsByUserResponse
	GetBody() *ListAuthorizedAppsByUserResponseBody
}

type ListAuthorizedAppsByUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedAppsByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedAppsByUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAppsByUserResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAppsByUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedAppsByUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedAppsByUserResponse) GetBody() *ListAuthorizedAppsByUserResponseBody {
	return s.Body
}

func (s *ListAuthorizedAppsByUserResponse) SetHeaders(v map[string]*string) *ListAuthorizedAppsByUserResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedAppsByUserResponse) SetStatusCode(v int32) *ListAuthorizedAppsByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedAppsByUserResponse) SetBody(v *ListAuthorizedAppsByUserResponseBody) *ListAuthorizedAppsByUserResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedAppsByUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
