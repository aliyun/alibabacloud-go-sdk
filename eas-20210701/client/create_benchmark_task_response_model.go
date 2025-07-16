// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateBenchmarkTaskResponseBody) *CreateBenchmarkTaskResponse
	GetBody() *CreateBenchmarkTaskResponseBody
}

type CreateBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBenchmarkTaskResponse) GetBody() *CreateBenchmarkTaskResponseBody {
	return s.Body
}

func (s *CreateBenchmarkTaskResponse) SetHeaders(v map[string]*string) *CreateBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateBenchmarkTaskResponse) SetStatusCode(v int32) *CreateBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBenchmarkTaskResponse) SetBody(v *CreateBenchmarkTaskResponseBody) *CreateBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *CreateBenchmarkTaskResponse) Validate() error {
	return dara.Validate(s)
}
