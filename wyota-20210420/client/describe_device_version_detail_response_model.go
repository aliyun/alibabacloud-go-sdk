// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceVersionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceVersionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceVersionDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceVersionDetailResponseBody) *DescribeDeviceVersionDetailResponse
	GetBody() *DescribeDeviceVersionDetailResponseBody
}

type DescribeDeviceVersionDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceVersionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceVersionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceVersionDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceVersionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceVersionDetailResponse) GetBody() *DescribeDeviceVersionDetailResponseBody {
	return s.Body
}

func (s *DescribeDeviceVersionDetailResponse) SetHeaders(v map[string]*string) *DescribeDeviceVersionDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceVersionDetailResponse) SetStatusCode(v int32) *DescribeDeviceVersionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponse) SetBody(v *DescribeDeviceVersionDetailResponseBody) *DescribeDeviceVersionDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceVersionDetailResponse) Validate() error {
	return dara.Validate(s)
}
