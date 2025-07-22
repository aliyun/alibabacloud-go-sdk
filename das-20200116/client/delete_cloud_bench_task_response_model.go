// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudBenchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudBenchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudBenchTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudBenchTaskResponseBody) *DeleteCloudBenchTaskResponse
	GetBody() *DeleteCloudBenchTaskResponseBody
}

type DeleteCloudBenchTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudBenchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudBenchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudBenchTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudBenchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudBenchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudBenchTaskResponse) GetBody() *DeleteCloudBenchTaskResponseBody {
	return s.Body
}

func (s *DeleteCloudBenchTaskResponse) SetHeaders(v map[string]*string) *DeleteCloudBenchTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudBenchTaskResponse) SetStatusCode(v int32) *DeleteCloudBenchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudBenchTaskResponse) SetBody(v *DeleteCloudBenchTaskResponseBody) *DeleteCloudBenchTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudBenchTaskResponse) Validate() error {
	return dara.Validate(s)
}
