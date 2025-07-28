// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUnassignedMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudUnassignedMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudUnassignedMachinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudUnassignedMachinesResponseBody) *DescribeHybridCloudUnassignedMachinesResponse
	GetBody() *DescribeHybridCloudUnassignedMachinesResponseBody
}

type DescribeHybridCloudUnassignedMachinesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudUnassignedMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudUnassignedMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnassignedMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) GetBody() *DescribeHybridCloudUnassignedMachinesResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudUnassignedMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) SetStatusCode(v int32) *DescribeHybridCloudUnassignedMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) SetBody(v *DescribeHybridCloudUnassignedMachinesResponseBody) *DescribeHybridCloudUnassignedMachinesResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponse) Validate() error {
	return dara.Validate(s)
}
