// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReplicationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReplicationJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReplicationJobsResponseBody) *DescribeReplicationJobsResponse
	GetBody() *DescribeReplicationJobsResponseBody
}

type DescribeReplicationJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReplicationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReplicationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReplicationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReplicationJobsResponse) GetBody() *DescribeReplicationJobsResponseBody {
	return s.Body
}

func (s *DescribeReplicationJobsResponse) SetHeaders(v map[string]*string) *DescribeReplicationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicationJobsResponse) SetStatusCode(v int32) *DescribeReplicationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicationJobsResponse) SetBody(v *DescribeReplicationJobsResponseBody) *DescribeReplicationJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeReplicationJobsResponse) Validate() error {
	return dara.Validate(s)
}
