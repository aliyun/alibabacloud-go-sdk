// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterNodesResponseBody) *DescribeClusterNodesResponse
	GetBody() *DescribeClusterNodesResponseBody
}

type DescribeClusterNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterNodesResponse) GetBody() *DescribeClusterNodesResponseBody {
	return s.Body
}

func (s *DescribeClusterNodesResponse) SetHeaders(v map[string]*string) *DescribeClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodesResponse) SetStatusCode(v int32) *DescribeClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterNodesResponse) SetBody(v *DescribeClusterNodesResponseBody) *DescribeClusterNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
