// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListMembersResponseBody) *ListMembersResponse
	GetBody() *ListMembersResponseBody
}

type ListMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponse) GoString() string {
	return s.String()
}

func (s *ListMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMembersResponse) GetBody() *ListMembersResponseBody {
	return s.Body
}

func (s *ListMembersResponse) SetHeaders(v map[string]*string) *ListMembersResponse {
	s.Headers = v
	return s
}

func (s *ListMembersResponse) SetStatusCode(v int32) *ListMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMembersResponse) SetBody(v *ListMembersResponseBody) *ListMembersResponse {
	s.Body = v
	return s
}

func (s *ListMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
