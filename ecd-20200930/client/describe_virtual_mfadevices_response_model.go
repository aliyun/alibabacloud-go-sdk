// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualMFADevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVirtualMFADevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVirtualMFADevicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVirtualMFADevicesResponseBody) *DescribeVirtualMFADevicesResponse
	GetBody() *DescribeVirtualMFADevicesResponseBody
}

type DescribeVirtualMFADevicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVirtualMFADevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVirtualMFADevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVirtualMFADevicesResponse) GetBody() *DescribeVirtualMFADevicesResponseBody {
	return s.Body
}

func (s *DescribeVirtualMFADevicesResponse) SetHeaders(v map[string]*string) *DescribeVirtualMFADevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetStatusCode(v int32) *DescribeVirtualMFADevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetBody(v *DescribeVirtualMFADevicesResponseBody) *DescribeVirtualMFADevicesResponse {
	s.Body = v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
