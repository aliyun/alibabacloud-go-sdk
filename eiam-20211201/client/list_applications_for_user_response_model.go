// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForUserResponseBody) *ListApplicationsForUserResponse
	GetBody() *ListApplicationsForUserResponseBody
}

type ListApplicationsForUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForUserResponse) GetBody() *ListApplicationsForUserResponseBody {
	return s.Body
}

func (s *ListApplicationsForUserResponse) SetHeaders(v map[string]*string) *ListApplicationsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForUserResponse) SetStatusCode(v int32) *ListApplicationsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForUserResponse) SetBody(v *ListApplicationsForUserResponseBody) *ListApplicationsForUserResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
