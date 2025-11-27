// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceURLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceURLResponseBody) *DescribeDeviceURLResponse
	GetBody() *DescribeDeviceURLResponseBody
}

type DescribeDeviceURLResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceURLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceURLResponse) GetBody() *DescribeDeviceURLResponseBody {
	return s.Body
}

func (s *DescribeDeviceURLResponse) SetHeaders(v map[string]*string) *DescribeDeviceURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceURLResponse) SetStatusCode(v int32) *DescribeDeviceURLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceURLResponse) SetBody(v *DescribeDeviceURLResponseBody) *DescribeDeviceURLResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceURLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
