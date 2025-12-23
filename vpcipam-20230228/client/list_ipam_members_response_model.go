// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamMembersResponseBody) *ListIpamMembersResponse
	GetBody() *ListIpamMembersResponseBody
}

type ListIpamMembersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamMembersResponse) GoString() string {
	return s.String()
}

func (s *ListIpamMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamMembersResponse) GetBody() *ListIpamMembersResponseBody {
	return s.Body
}

func (s *ListIpamMembersResponse) SetHeaders(v map[string]*string) *ListIpamMembersResponse {
	s.Headers = v
	return s
}

func (s *ListIpamMembersResponse) SetStatusCode(v int32) *ListIpamMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamMembersResponse) SetBody(v *ListIpamMembersResponseBody) *ListIpamMembersResponse {
	s.Body = v
	return s
}

func (s *ListIpamMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
