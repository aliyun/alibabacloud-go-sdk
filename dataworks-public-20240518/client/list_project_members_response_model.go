// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectMembersResponseBody) *ListProjectMembersResponse
	GetBody() *ListProjectMembersResponseBody
}

type ListProjectMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectMembersResponse) GetBody() *ListProjectMembersResponseBody {
	return s.Body
}

func (s *ListProjectMembersResponse) SetHeaders(v map[string]*string) *ListProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *ListProjectMembersResponse) SetStatusCode(v int32) *ListProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectMembersResponse) SetBody(v *ListProjectMembersResponseBody) *ListProjectMembersResponse {
	s.Body = v
	return s
}

func (s *ListProjectMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
