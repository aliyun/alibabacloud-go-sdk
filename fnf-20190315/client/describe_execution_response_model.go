// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExecutionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExecutionResponseBody) *DescribeExecutionResponse
	GetBody() *DescribeExecutionResponseBody
}

type DescribeExecutionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExecutionResponse) GetBody() *DescribeExecutionResponseBody {
	return s.Body
}

func (s *DescribeExecutionResponse) SetHeaders(v map[string]*string) *DescribeExecutionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExecutionResponse) SetStatusCode(v int32) *DescribeExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExecutionResponse) SetBody(v *DescribeExecutionResponseBody) *DescribeExecutionResponse {
	s.Body = v
	return s
}

func (s *DescribeExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
