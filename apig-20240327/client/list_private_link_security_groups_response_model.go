// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateLinkSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateLinkSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateLinkSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateLinkSecurityGroupsResponseBody) *ListPrivateLinkSecurityGroupsResponse
	GetBody() *ListPrivateLinkSecurityGroupsResponseBody
}

type ListPrivateLinkSecurityGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateLinkSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateLinkSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateLinkSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateLinkSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateLinkSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateLinkSecurityGroupsResponse) GetBody() *ListPrivateLinkSecurityGroupsResponseBody {
	return s.Body
}

func (s *ListPrivateLinkSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListPrivateLinkSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponse) SetStatusCode(v int32) *ListPrivateLinkSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponse) SetBody(v *ListPrivateLinkSecurityGroupsResponseBody) *ListPrivateLinkSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ListPrivateLinkSecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
