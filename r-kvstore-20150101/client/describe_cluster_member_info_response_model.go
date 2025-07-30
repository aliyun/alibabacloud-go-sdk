// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterMemberInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterMemberInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterMemberInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterMemberInfoResponseBody) *DescribeClusterMemberInfoResponse
	GetBody() *DescribeClusterMemberInfoResponseBody
}

type DescribeClusterMemberInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterMemberInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterMemberInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterMemberInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterMemberInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterMemberInfoResponse) GetBody() *DescribeClusterMemberInfoResponseBody {
	return s.Body
}

func (s *DescribeClusterMemberInfoResponse) SetHeaders(v map[string]*string) *DescribeClusterMemberInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterMemberInfoResponse) SetStatusCode(v int32) *DescribeClusterMemberInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterMemberInfoResponse) SetBody(v *DescribeClusterMemberInfoResponseBody) *DescribeClusterMemberInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterMemberInfoResponse) Validate() error {
	return dara.Validate(s)
}
