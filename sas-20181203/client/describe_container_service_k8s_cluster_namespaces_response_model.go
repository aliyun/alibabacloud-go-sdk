// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClusterNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerServiceK8sClusterNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerServiceK8sClusterNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerServiceK8sClusterNamespacesResponseBody) *DescribeContainerServiceK8sClusterNamespacesResponse
	GetBody() *DescribeContainerServiceK8sClusterNamespacesResponseBody
}

type DescribeContainerServiceK8sClusterNamespacesResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerServiceK8sClusterNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerServiceK8sClusterNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClusterNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) GetBody() *DescribeContainerServiceK8sClusterNamespacesResponseBody {
	return s.Body
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) SetHeaders(v map[string]*string) *DescribeContainerServiceK8sClusterNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) SetStatusCode(v int32) *DescribeContainerServiceK8sClusterNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) SetBody(v *DescribeContainerServiceK8sClusterNamespacesResponseBody) *DescribeContainerServiceK8sClusterNamespacesResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerServiceK8sClusterNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
