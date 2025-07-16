// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBenchmarkTaskResponseBody) *DeleteBenchmarkTaskResponse
	GetBody() *DeleteBenchmarkTaskResponseBody
}

type DeleteBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBenchmarkTaskResponse) GetBody() *DeleteBenchmarkTaskResponseBody {
	return s.Body
}

func (s *DeleteBenchmarkTaskResponse) SetHeaders(v map[string]*string) *DeleteBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteBenchmarkTaskResponse) SetStatusCode(v int32) *DeleteBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBenchmarkTaskResponse) SetBody(v *DeleteBenchmarkTaskResponseBody) *DeleteBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteBenchmarkTaskResponse) Validate() error {
	return dara.Validate(s)
}
