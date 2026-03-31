// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupsForUserResponseBody) *ListGroupsForUserResponse
	GetBody() *ListGroupsForUserResponseBody
}

type ListGroupsForUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupsForUserResponse) GetBody() *ListGroupsForUserResponseBody {
	return s.Body
}

func (s *ListGroupsForUserResponse) SetHeaders(v map[string]*string) *ListGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForUserResponse) SetStatusCode(v int32) *ListGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForUserResponse) SetBody(v *ListGroupsForUserResponseBody) *ListGroupsForUserResponse {
	s.Body = v
	return s
}

func (s *ListGroupsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
