// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStatDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStatDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStatDeviceInfoResponseBody) *DescribeStatDeviceInfoResponse
	GetBody() *DescribeStatDeviceInfoResponseBody
}

type DescribeStatDeviceInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStatDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStatDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStatDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStatDeviceInfoResponse) GetBody() *DescribeStatDeviceInfoResponseBody {
	return s.Body
}

func (s *DescribeStatDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeStatDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatDeviceInfoResponse) SetStatusCode(v int32) *DescribeStatDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStatDeviceInfoResponse) SetBody(v *DescribeStatDeviceInfoResponseBody) *DescribeStatDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeStatDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
