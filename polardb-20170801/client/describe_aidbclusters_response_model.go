// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClustersResponseBody) *DescribeAIDBClustersResponse
	GetBody() *DescribeAIDBClustersResponseBody
}

type DescribeAIDBClustersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClustersResponse) GetBody() *DescribeAIDBClustersResponseBody {
	return s.Body
}

func (s *DescribeAIDBClustersResponse) SetHeaders(v map[string]*string) *DescribeAIDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClustersResponse) SetStatusCode(v int32) *DescribeAIDBClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClustersResponse) SetBody(v *DescribeAIDBClustersResponseBody) *DescribeAIDBClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClustersResponse) Validate() error {
	return dara.Validate(s)
}
