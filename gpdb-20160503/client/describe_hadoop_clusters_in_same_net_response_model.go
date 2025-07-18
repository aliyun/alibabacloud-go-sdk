// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopClustersInSameNetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHadoopClustersInSameNetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHadoopClustersInSameNetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHadoopClustersInSameNetResponseBody) *DescribeHadoopClustersInSameNetResponse
	GetBody() *DescribeHadoopClustersInSameNetResponseBody
}

type DescribeHadoopClustersInSameNetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHadoopClustersInSameNetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHadoopClustersInSameNetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopClustersInSameNetResponse) GoString() string {
	return s.String()
}

func (s *DescribeHadoopClustersInSameNetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHadoopClustersInSameNetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHadoopClustersInSameNetResponse) GetBody() *DescribeHadoopClustersInSameNetResponseBody {
	return s.Body
}

func (s *DescribeHadoopClustersInSameNetResponse) SetHeaders(v map[string]*string) *DescribeHadoopClustersInSameNetResponse {
	s.Headers = v
	return s
}

func (s *DescribeHadoopClustersInSameNetResponse) SetStatusCode(v int32) *DescribeHadoopClustersInSameNetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHadoopClustersInSameNetResponse) SetBody(v *DescribeHadoopClustersInSameNetResponseBody) *DescribeHadoopClustersInSameNetResponse {
	s.Body = v
	return s
}

func (s *DescribeHadoopClustersInSameNetResponse) Validate() error {
	return dara.Validate(s)
}
