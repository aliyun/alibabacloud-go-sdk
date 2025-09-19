// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClusterKritisStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerServiceK8sClusterKritisStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerServiceK8sClusterKritisStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerServiceK8sClusterKritisStatusResponseBody) *DescribeContainerServiceK8sClusterKritisStatusResponse
	GetBody() *DescribeContainerServiceK8sClusterKritisStatusResponseBody
}

type DescribeContainerServiceK8sClusterKritisStatusResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerServiceK8sClusterKritisStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerServiceK8sClusterKritisStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterKritisStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) GetBody() *DescribeContainerServiceK8sClusterKritisStatusResponseBody {
	return s.Body
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) SetHeaders(v map[string]*string) *DescribeContainerServiceK8sClusterKritisStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) SetStatusCode(v int32) *DescribeContainerServiceK8sClusterKritisStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) SetBody(v *DescribeContainerServiceK8sClusterKritisStatusResponseBody) *DescribeContainerServiceK8sClusterKritisStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerServiceK8sClusterKritisStatusResponse) Validate() error {
	return dara.Validate(s)
}
