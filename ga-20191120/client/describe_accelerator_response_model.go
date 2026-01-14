// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAcceleratorResponseBody) *DescribeAcceleratorResponse
	GetBody() *DescribeAcceleratorResponseBody
}

type DescribeAcceleratorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAcceleratorResponse) GetBody() *DescribeAcceleratorResponseBody {
	return s.Body
}

func (s *DescribeAcceleratorResponse) SetHeaders(v map[string]*string) *DescribeAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DescribeAcceleratorResponse) SetStatusCode(v int32) *DescribeAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAcceleratorResponse) SetBody(v *DescribeAcceleratorResponseBody) *DescribeAcceleratorResponse {
	s.Body = v
	return s
}

func (s *DescribeAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
