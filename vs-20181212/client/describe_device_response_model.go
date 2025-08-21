// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceResponseBody) *DescribeDeviceResponse
	GetBody() *DescribeDeviceResponseBody
}

type DescribeDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceResponse) GetBody() *DescribeDeviceResponseBody {
	return s.Body
}

func (s *DescribeDeviceResponse) SetHeaders(v map[string]*string) *DescribeDeviceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceResponse) SetStatusCode(v int32) *DescribeDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceResponse) SetBody(v *DescribeDeviceResponseBody) *DescribeDeviceResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceResponse) Validate() error {
	return dara.Validate(s)
}
