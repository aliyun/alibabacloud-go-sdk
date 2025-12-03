// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationMembersResponseBody) *ListApplicationMembersResponse
	GetBody() *ListApplicationMembersResponseBody
}

type ListApplicationMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMembersResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationMembersResponse) GetBody() *ListApplicationMembersResponseBody {
	return s.Body
}

func (s *ListApplicationMembersResponse) SetHeaders(v map[string]*string) *ListApplicationMembersResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationMembersResponse) SetStatusCode(v int32) *ListApplicationMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMembersResponse) SetBody(v *ListApplicationMembersResponseBody) *ListApplicationMembersResponse {
	s.Body = v
	return s
}

func (s *ListApplicationMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
