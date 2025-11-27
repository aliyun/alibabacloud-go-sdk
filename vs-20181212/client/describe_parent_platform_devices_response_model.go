// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParentPlatformDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParentPlatformDevicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParentPlatformDevicesResponseBody) *DescribeParentPlatformDevicesResponse
	GetBody() *DescribeParentPlatformDevicesResponseBody
}

type DescribeParentPlatformDevicesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParentPlatformDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParentPlatformDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParentPlatformDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParentPlatformDevicesResponse) GetBody() *DescribeParentPlatformDevicesResponseBody {
	return s.Body
}

func (s *DescribeParentPlatformDevicesResponse) SetHeaders(v map[string]*string) *DescribeParentPlatformDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentPlatformDevicesResponse) SetStatusCode(v int32) *DescribeParentPlatformDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponse) SetBody(v *DescribeParentPlatformDevicesResponseBody) *DescribeParentPlatformDevicesResponse {
	s.Body = v
	return s
}

func (s *DescribeParentPlatformDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
