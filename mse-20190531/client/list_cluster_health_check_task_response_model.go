// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterHealthCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterHealthCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterHealthCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterHealthCheckTaskResponseBody) *ListClusterHealthCheckTaskResponse
	GetBody() *ListClusterHealthCheckTaskResponseBody
}

type ListClusterHealthCheckTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterHealthCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterHealthCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHealthCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *ListClusterHealthCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterHealthCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterHealthCheckTaskResponse) GetBody() *ListClusterHealthCheckTaskResponseBody {
	return s.Body
}

func (s *ListClusterHealthCheckTaskResponse) SetHeaders(v map[string]*string) *ListClusterHealthCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *ListClusterHealthCheckTaskResponse) SetStatusCode(v int32) *ListClusterHealthCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponse) SetBody(v *ListClusterHealthCheckTaskResponseBody) *ListClusterHealthCheckTaskResponse {
	s.Body = v
	return s
}

func (s *ListClusterHealthCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
