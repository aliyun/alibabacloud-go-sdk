// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDevicesDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDevicesDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDevicesDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDevicesDataResponseBody) *DescribeVsDevicesDataResponse
	GetBody() *DescribeVsDevicesDataResponseBody
}

type DescribeVsDevicesDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDevicesDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDevicesDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDevicesDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDevicesDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDevicesDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDevicesDataResponse) GetBody() *DescribeVsDevicesDataResponseBody {
	return s.Body
}

func (s *DescribeVsDevicesDataResponse) SetHeaders(v map[string]*string) *DescribeVsDevicesDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDevicesDataResponse) SetStatusCode(v int32) *DescribeVsDevicesDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDevicesDataResponse) SetBody(v *DescribeVsDevicesDataResponseBody) *DescribeVsDevicesDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDevicesDataResponse) Validate() error {
	return dara.Validate(s)
}
