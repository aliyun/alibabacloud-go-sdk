// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFbIssueLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFbIssueLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFbIssueLabelsResponse
	GetStatusCode() *int32
	SetBody(v *ListFbIssueLabelsResponseBody) *ListFbIssueLabelsResponse
	GetBody() *ListFbIssueLabelsResponseBody
}

type ListFbIssueLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFbIssueLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFbIssueLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFbIssueLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListFbIssueLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFbIssueLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFbIssueLabelsResponse) GetBody() *ListFbIssueLabelsResponseBody {
	return s.Body
}

func (s *ListFbIssueLabelsResponse) SetHeaders(v map[string]*string) *ListFbIssueLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListFbIssueLabelsResponse) SetStatusCode(v int32) *ListFbIssueLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFbIssueLabelsResponse) SetBody(v *ListFbIssueLabelsResponseBody) *ListFbIssueLabelsResponse {
	s.Body = v
	return s
}

func (s *ListFbIssueLabelsResponse) Validate() error {
	return dara.Validate(s)
}
