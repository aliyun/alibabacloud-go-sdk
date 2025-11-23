// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedInstancesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedInstancesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedInstancesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedInstancesForUserResponseBody) *ListAuthorizedInstancesForUserResponse
	GetBody() *ListAuthorizedInstancesForUserResponseBody
}

type ListAuthorizedInstancesForUserResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedInstancesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedInstancesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedInstancesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedInstancesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedInstancesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedInstancesForUserResponse) GetBody() *ListAuthorizedInstancesForUserResponseBody {
	return s.Body
}

func (s *ListAuthorizedInstancesForUserResponse) SetHeaders(v map[string]*string) *ListAuthorizedInstancesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedInstancesForUserResponse) SetStatusCode(v int32) *ListAuthorizedInstancesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponse) SetBody(v *ListAuthorizedInstancesForUserResponseBody) *ListAuthorizedInstancesForUserResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedInstancesForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
