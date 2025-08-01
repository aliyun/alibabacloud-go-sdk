// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineEvalTaskResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOnlineEvalTaskResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOnlineEvalTaskResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListOnlineEvalTaskResultsResponseBody) *ListOnlineEvalTaskResultsResponse
	GetBody() *ListOnlineEvalTaskResultsResponseBody
}

type ListOnlineEvalTaskResultsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnlineEvalTaskResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnlineEvalTaskResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineEvalTaskResultsResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOnlineEvalTaskResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOnlineEvalTaskResultsResponse) GetBody() *ListOnlineEvalTaskResultsResponseBody {
	return s.Body
}

func (s *ListOnlineEvalTaskResultsResponse) SetHeaders(v map[string]*string) *ListOnlineEvalTaskResultsResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineEvalTaskResultsResponse) SetStatusCode(v int32) *ListOnlineEvalTaskResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponse) SetBody(v *ListOnlineEvalTaskResultsResponseBody) *ListOnlineEvalTaskResultsResponse {
	s.Body = v
	return s
}

func (s *ListOnlineEvalTaskResultsResponse) Validate() error {
	return dara.Validate(s)
}
