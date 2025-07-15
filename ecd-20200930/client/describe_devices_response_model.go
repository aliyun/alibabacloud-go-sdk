// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDevicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDevicesResponseBody) *DescribeDevicesResponse
	GetBody() *DescribeDevicesResponseBody
}

type DescribeDevicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDevicesResponse) GetBody() *DescribeDevicesResponseBody {
	return s.Body
}

func (s *DescribeDevicesResponse) SetHeaders(v map[string]*string) *DescribeDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDevicesResponse) SetStatusCode(v int32) *DescribeDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDevicesResponse) SetBody(v *DescribeDevicesResponseBody) *DescribeDevicesResponse {
	s.Body = v
	return s
}

func (s *DescribeDevicesResponse) Validate() error {
	return dara.Validate(s)
}
