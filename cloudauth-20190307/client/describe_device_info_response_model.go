// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse
	GetBody() *DescribeDeviceInfoResponseBody
}

type DescribeDeviceInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceInfoResponse) GetBody() *DescribeDeviceInfoResponseBody {
	return s.Body
}

func (s *DescribeDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceInfoResponse) SetStatusCode(v int32) *DescribeDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceInfoResponse) SetBody(v *DescribeDeviceInfoResponseBody) *DescribeDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
