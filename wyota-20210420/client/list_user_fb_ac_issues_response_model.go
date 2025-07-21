// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserFbAcIssuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserFbAcIssuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserFbAcIssuesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserFbAcIssuesResponseBody) *ListUserFbAcIssuesResponse
	GetBody() *ListUserFbAcIssuesResponseBody
}

type ListUserFbAcIssuesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserFbAcIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserFbAcIssuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserFbAcIssuesResponse) GoString() string {
	return s.String()
}

func (s *ListUserFbAcIssuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserFbAcIssuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserFbAcIssuesResponse) GetBody() *ListUserFbAcIssuesResponseBody {
	return s.Body
}

func (s *ListUserFbAcIssuesResponse) SetHeaders(v map[string]*string) *ListUserFbAcIssuesResponse {
	s.Headers = v
	return s
}

func (s *ListUserFbAcIssuesResponse) SetStatusCode(v int32) *ListUserFbAcIssuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserFbAcIssuesResponse) SetBody(v *ListUserFbAcIssuesResponseBody) *ListUserFbAcIssuesResponse {
	s.Body = v
	return s
}

func (s *ListUserFbAcIssuesResponse) Validate() error {
	return dara.Validate(s)
}
