// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJoinedGroupsForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJoinedGroupsForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJoinedGroupsForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListJoinedGroupsForUserResponseBody) *ListJoinedGroupsForUserResponse
	GetBody() *ListJoinedGroupsForUserResponseBody
}

type ListJoinedGroupsForUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJoinedGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJoinedGroupsForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJoinedGroupsForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJoinedGroupsForUserResponse) GetBody() *ListJoinedGroupsForUserResponseBody {
	return s.Body
}

func (s *ListJoinedGroupsForUserResponse) SetHeaders(v map[string]*string) *ListJoinedGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListJoinedGroupsForUserResponse) SetStatusCode(v int32) *ListJoinedGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJoinedGroupsForUserResponse) SetBody(v *ListJoinedGroupsForUserResponseBody) *ListJoinedGroupsForUserResponse {
	s.Body = v
	return s
}

func (s *ListJoinedGroupsForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
