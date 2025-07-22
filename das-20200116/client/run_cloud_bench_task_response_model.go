// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCloudBenchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCloudBenchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCloudBenchTaskResponse
	GetStatusCode() *int32
	SetBody(v *RunCloudBenchTaskResponseBody) *RunCloudBenchTaskResponse
	GetBody() *RunCloudBenchTaskResponseBody
}

type RunCloudBenchTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCloudBenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCloudBenchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCloudBenchTaskResponse) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCloudBenchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCloudBenchTaskResponse) GetBody() *RunCloudBenchTaskResponseBody {
	return s.Body
}

func (s *RunCloudBenchTaskResponse) SetHeaders(v map[string]*string) *RunCloudBenchTaskResponse {
	s.Headers = v
	return s
}

func (s *RunCloudBenchTaskResponse) SetStatusCode(v int32) *RunCloudBenchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCloudBenchTaskResponse) SetBody(v *RunCloudBenchTaskResponseBody) *RunCloudBenchTaskResponse {
	s.Body = v
	return s
}

func (s *RunCloudBenchTaskResponse) Validate() error {
	return dara.Validate(s)
}
