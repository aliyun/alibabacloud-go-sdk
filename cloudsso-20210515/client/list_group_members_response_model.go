// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupMembersResponseBody) *ListGroupMembersResponse
	GetBody() *ListGroupMembersResponseBody
}

type ListGroupMembersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *ListGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupMembersResponse) GetBody() *ListGroupMembersResponseBody {
	return s.Body
}

func (s *ListGroupMembersResponse) SetHeaders(v map[string]*string) *ListGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *ListGroupMembersResponse) SetStatusCode(v int32) *ListGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupMembersResponse) SetBody(v *ListGroupMembersResponseBody) *ListGroupMembersResponse {
	s.Body = v
	return s
}

func (s *ListGroupMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
