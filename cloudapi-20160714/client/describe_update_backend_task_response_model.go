// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpdateBackendTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpdateBackendTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpdateBackendTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpdateBackendTaskResponseBody) *DescribeUpdateBackendTaskResponse
	GetBody() *DescribeUpdateBackendTaskResponseBody
}

type DescribeUpdateBackendTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpdateBackendTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpdateBackendTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateBackendTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpdateBackendTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpdateBackendTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpdateBackendTaskResponse) GetBody() *DescribeUpdateBackendTaskResponseBody {
	return s.Body
}

func (s *DescribeUpdateBackendTaskResponse) SetHeaders(v map[string]*string) *DescribeUpdateBackendTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpdateBackendTaskResponse) SetStatusCode(v int32) *DescribeUpdateBackendTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpdateBackendTaskResponse) SetBody(v *DescribeUpdateBackendTaskResponseBody) *DescribeUpdateBackendTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeUpdateBackendTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
