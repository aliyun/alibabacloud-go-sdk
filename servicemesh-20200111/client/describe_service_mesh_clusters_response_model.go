// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshClustersResponseBody) *DescribeServiceMeshClustersResponse
	GetBody() *DescribeServiceMeshClustersResponseBody
}

type DescribeServiceMeshClustersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshClustersResponse) GetBody() *DescribeServiceMeshClustersResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshClustersResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshClustersResponse) SetStatusCode(v int32) *DescribeServiceMeshClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshClustersResponse) SetBody(v *DescribeServiceMeshClustersResponseBody) *DescribeServiceMeshClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
