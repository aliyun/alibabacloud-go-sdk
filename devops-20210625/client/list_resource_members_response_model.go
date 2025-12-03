// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceMembersResponseBody) *ListResourceMembersResponse
	GetBody() *ListResourceMembersResponseBody
}

type ListResourceMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceMembersResponse) GoString() string {
	return s.String()
}

func (s *ListResourceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceMembersResponse) GetBody() *ListResourceMembersResponseBody {
	return s.Body
}

func (s *ListResourceMembersResponse) SetHeaders(v map[string]*string) *ListResourceMembersResponse {
	s.Headers = v
	return s
}

func (s *ListResourceMembersResponse) SetStatusCode(v int32) *ListResourceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceMembersResponse) SetBody(v *ListResourceMembersResponseBody) *ListResourceMembersResponse {
	s.Body = v
	return s
}

func (s *ListResourceMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
