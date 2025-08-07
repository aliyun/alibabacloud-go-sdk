// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIssueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIssueResponse
	GetStatusCode() *int32
	SetBody(v *GetIssueResponseBody) *GetIssueResponse
	GetBody() *GetIssueResponseBody
}

type GetIssueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIssueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIssueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIssueResponse) GoString() string {
	return s.String()
}

func (s *GetIssueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIssueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIssueResponse) GetBody() *GetIssueResponseBody {
	return s.Body
}

func (s *GetIssueResponse) SetHeaders(v map[string]*string) *GetIssueResponse {
	s.Headers = v
	return s
}

func (s *GetIssueResponse) SetStatusCode(v int32) *GetIssueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIssueResponse) SetBody(v *GetIssueResponseBody) *GetIssueResponse {
	s.Body = v
	return s
}

func (s *GetIssueResponse) Validate() error {
	return dara.Validate(s)
}
