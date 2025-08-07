// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIssuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIssuesResponse
	GetStatusCode() *int32
	SetBody(v *GetIssuesResponseBody) *GetIssuesResponse
	GetBody() *GetIssuesResponseBody
}

type GetIssuesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIssuesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesResponse) GoString() string {
	return s.String()
}

func (s *GetIssuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIssuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIssuesResponse) GetBody() *GetIssuesResponseBody {
	return s.Body
}

func (s *GetIssuesResponse) SetHeaders(v map[string]*string) *GetIssuesResponse {
	s.Headers = v
	return s
}

func (s *GetIssuesResponse) SetStatusCode(v int32) *GetIssuesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIssuesResponse) SetBody(v *GetIssuesResponseBody) *GetIssuesResponse {
	s.Body = v
	return s
}

func (s *GetIssuesResponse) Validate() error {
	return dara.Validate(s)
}
