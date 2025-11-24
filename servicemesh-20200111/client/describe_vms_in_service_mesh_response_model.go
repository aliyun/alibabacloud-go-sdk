// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVMsInServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVMsInServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVMsInServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVMsInServiceMeshResponseBody) *DescribeVMsInServiceMeshResponse
	GetBody() *DescribeVMsInServiceMeshResponseBody
}

type DescribeVMsInServiceMeshResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVMsInServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVMsInServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVMsInServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVMsInServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVMsInServiceMeshResponse) GetBody() *DescribeVMsInServiceMeshResponseBody {
	return s.Body
}

func (s *DescribeVMsInServiceMeshResponse) SetHeaders(v map[string]*string) *DescribeVMsInServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *DescribeVMsInServiceMeshResponse) SetStatusCode(v int32) *DescribeVMsInServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVMsInServiceMeshResponse) SetBody(v *DescribeVMsInServiceMeshResponseBody) *DescribeVMsInServiceMeshResponse {
	s.Body = v
	return s
}

func (s *DescribeVMsInServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
