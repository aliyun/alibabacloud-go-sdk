// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceServiceResponseBody) *DescribeDeviceServiceResponse
	GetBody() *DescribeDeviceServiceResponseBody
}

type DescribeDeviceServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceServiceResponse) GetBody() *DescribeDeviceServiceResponseBody {
	return s.Body
}

func (s *DescribeDeviceServiceResponse) SetHeaders(v map[string]*string) *DescribeDeviceServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceServiceResponse) SetStatusCode(v int32) *DescribeDeviceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceServiceResponse) SetBody(v *DescribeDeviceServiceResponseBody) *DescribeDeviceServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceServiceResponse) Validate() error {
	return dara.Validate(s)
}
