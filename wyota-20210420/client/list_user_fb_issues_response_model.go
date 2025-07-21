// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserFbIssuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserFbIssuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserFbIssuesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserFbIssuesResponseBody) *ListUserFbIssuesResponse
	GetBody() *ListUserFbIssuesResponseBody
}

type ListUserFbIssuesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserFbIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserFbIssuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbIssuesResponse) GoString() string {
	return s.String()
}

func (s *ListUserFbIssuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserFbIssuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserFbIssuesResponse) GetBody() *ListUserFbIssuesResponseBody {
	return s.Body
}

func (s *ListUserFbIssuesResponse) SetHeaders(v map[string]*string) *ListUserFbIssuesResponse {
	s.Headers = v
	return s
}

func (s *ListUserFbIssuesResponse) SetStatusCode(v int32) *ListUserFbIssuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserFbIssuesResponse) SetBody(v *ListUserFbIssuesResponseBody) *ListUserFbIssuesResponse {
	s.Body = v
	return s
}

func (s *ListUserFbIssuesResponse) Validate() error {
	return dara.Validate(s)
}
