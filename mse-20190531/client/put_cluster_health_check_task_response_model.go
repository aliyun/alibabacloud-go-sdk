// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutClusterHealthCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutClusterHealthCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutClusterHealthCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *PutClusterHealthCheckTaskResponseBody) *PutClusterHealthCheckTaskResponse
	GetBody() *PutClusterHealthCheckTaskResponseBody
}

type PutClusterHealthCheckTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutClusterHealthCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutClusterHealthCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PutClusterHealthCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *PutClusterHealthCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutClusterHealthCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutClusterHealthCheckTaskResponse) GetBody() *PutClusterHealthCheckTaskResponseBody {
	return s.Body
}

func (s *PutClusterHealthCheckTaskResponse) SetHeaders(v map[string]*string) *PutClusterHealthCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *PutClusterHealthCheckTaskResponse) SetStatusCode(v int32) *PutClusterHealthCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponse) SetBody(v *PutClusterHealthCheckTaskResponseBody) *PutClusterHealthCheckTaskResponse {
	s.Body = v
	return s
}

func (s *PutClusterHealthCheckTaskResponse) Validate() error {
	return dara.Validate(s)
}
