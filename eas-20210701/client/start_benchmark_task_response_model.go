// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartBenchmarkTaskResponseBody) *StartBenchmarkTaskResponse
	GetBody() *StartBenchmarkTaskResponseBody
}

type StartBenchmarkTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartBenchmarkTaskResponse) GetBody() *StartBenchmarkTaskResponseBody {
	return s.Body
}

func (s *StartBenchmarkTaskResponse) SetHeaders(v map[string]*string) *StartBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *StartBenchmarkTaskResponse) SetStatusCode(v int32) *StartBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBenchmarkTaskResponse) SetBody(v *StartBenchmarkTaskResponseBody) *StartBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *StartBenchmarkTaskResponse) Validate() error {
	return dara.Validate(s)
}
