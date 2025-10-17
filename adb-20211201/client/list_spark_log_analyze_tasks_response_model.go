// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkLogAnalyzeTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSparkLogAnalyzeTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSparkLogAnalyzeTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListSparkLogAnalyzeTasksResponseBody) *ListSparkLogAnalyzeTasksResponse
	GetBody() *ListSparkLogAnalyzeTasksResponseBody
}

type ListSparkLogAnalyzeTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkLogAnalyzeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkLogAnalyzeTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSparkLogAnalyzeTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSparkLogAnalyzeTasksResponse) GetBody() *ListSparkLogAnalyzeTasksResponseBody {
	return s.Body
}

func (s *ListSparkLogAnalyzeTasksResponse) SetHeaders(v map[string]*string) *ListSparkLogAnalyzeTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponse) SetStatusCode(v int32) *ListSparkLogAnalyzeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponse) SetBody(v *ListSparkLogAnalyzeTasksResponseBody) *ListSparkLogAnalyzeTasksResponse {
	s.Body = v
	return s
}

func (s *ListSparkLogAnalyzeTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
