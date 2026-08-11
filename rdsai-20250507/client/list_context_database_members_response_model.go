// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseMembersResponseBody) *ListContextDatabaseMembersResponse
	GetBody() *ListContextDatabaseMembersResponseBody
}

type ListContextDatabaseMembersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseMembersResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseMembersResponse) GetBody() *ListContextDatabaseMembersResponseBody {
	return s.Body
}

func (s *ListContextDatabaseMembersResponse) SetHeaders(v map[string]*string) *ListContextDatabaseMembersResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseMembersResponse) SetStatusCode(v int32) *ListContextDatabaseMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseMembersResponse) SetBody(v *ListContextDatabaseMembersResponseBody) *ListContextDatabaseMembersResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
