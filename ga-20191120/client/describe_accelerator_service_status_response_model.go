// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAcceleratorServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAcceleratorServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAcceleratorServiceStatusResponseBody) *DescribeAcceleratorServiceStatusResponse
	GetBody() *DescribeAcceleratorServiceStatusResponseBody
}

type DescribeAcceleratorServiceStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAcceleratorServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAcceleratorServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAcceleratorServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAcceleratorServiceStatusResponse) GetBody() *DescribeAcceleratorServiceStatusResponseBody {
	return s.Body
}

func (s *DescribeAcceleratorServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeAcceleratorServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponse) SetStatusCode(v int32) *DescribeAcceleratorServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponse) SetBody(v *DescribeAcceleratorServiceStatusResponseBody) *DescribeAcceleratorServiceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
