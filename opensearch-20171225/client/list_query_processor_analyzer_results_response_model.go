// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorAnalyzerResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueryProcessorAnalyzerResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueryProcessorAnalyzerResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListQueryProcessorAnalyzerResultsResponseBody) *ListQueryProcessorAnalyzerResultsResponse
	GetBody() *ListQueryProcessorAnalyzerResultsResponseBody
}

type ListQueryProcessorAnalyzerResultsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryProcessorAnalyzerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueryProcessorAnalyzerResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueryProcessorAnalyzerResultsResponse) GetBody() *ListQueryProcessorAnalyzerResultsResponseBody {
	return s.Body
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetHeaders(v map[string]*string) *ListQueryProcessorAnalyzerResultsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetStatusCode(v int32) *ListQueryProcessorAnalyzerResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetBody(v *ListQueryProcessorAnalyzerResultsResponseBody) *ListQueryProcessorAnalyzerResultsResponse {
	s.Body = v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
