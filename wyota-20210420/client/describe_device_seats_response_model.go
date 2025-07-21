// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceSeatsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceSeatsResponseBody) *DescribeDeviceSeatsResponse
	GetBody() *DescribeDeviceSeatsResponseBody
}

type DescribeDeviceSeatsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceSeatsResponse) GetBody() *DescribeDeviceSeatsResponseBody {
	return s.Body
}

func (s *DescribeDeviceSeatsResponse) SetHeaders(v map[string]*string) *DescribeDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceSeatsResponse) SetStatusCode(v int32) *DescribeDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceSeatsResponse) SetBody(v *DescribeDeviceSeatsResponseBody) *DescribeDeviceSeatsResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceSeatsResponse) Validate() error {
	return dara.Validate(s)
}
