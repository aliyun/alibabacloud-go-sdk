// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEapDeviceResourceAllocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEapDeviceResourceAllocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEapDeviceResourceAllocationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEapDeviceResourceAllocationResponseBody) *DescribeEapDeviceResourceAllocationResponse
	GetBody() *DescribeEapDeviceResourceAllocationResponseBody
}

type DescribeEapDeviceResourceAllocationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEapDeviceResourceAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEapDeviceResourceAllocationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEapDeviceResourceAllocationResponse) GoString() string {
	return s.String()
}

func (s *DescribeEapDeviceResourceAllocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEapDeviceResourceAllocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEapDeviceResourceAllocationResponse) GetBody() *DescribeEapDeviceResourceAllocationResponseBody {
	return s.Body
}

func (s *DescribeEapDeviceResourceAllocationResponse) SetHeaders(v map[string]*string) *DescribeEapDeviceResourceAllocationResponse {
	s.Headers = v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponse) SetStatusCode(v int32) *DescribeEapDeviceResourceAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponse) SetBody(v *DescribeEapDeviceResourceAllocationResponseBody) *DescribeEapDeviceResourceAllocationResponse {
	s.Body = v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
