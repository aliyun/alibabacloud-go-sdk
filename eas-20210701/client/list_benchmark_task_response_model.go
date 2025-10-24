// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListBenchmarkTaskResponseBody) *ListBenchmarkTaskResponse
	GetBody() *ListBenchmarkTaskResponseBody
}

type ListBenchmarkTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBenchmarkTaskResponse) GetBody() *ListBenchmarkTaskResponseBody {
	return s.Body
}

func (s *ListBenchmarkTaskResponse) SetHeaders(v map[string]*string) *ListBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *ListBenchmarkTaskResponse) SetStatusCode(v int32) *ListBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBenchmarkTaskResponse) SetBody(v *ListBenchmarkTaskResponseBody) *ListBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *ListBenchmarkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
