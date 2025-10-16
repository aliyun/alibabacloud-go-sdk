// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAckClusterNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAckClusterNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAckClusterNamespacesResponseBody) *DescribeAckClusterNamespacesResponse
	GetBody() *DescribeAckClusterNamespacesResponseBody
}

type DescribeAckClusterNamespacesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAckClusterNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAckClusterNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAckClusterNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAckClusterNamespacesResponse) GetBody() *DescribeAckClusterNamespacesResponseBody {
	return s.Body
}

func (s *DescribeAckClusterNamespacesResponse) SetHeaders(v map[string]*string) *DescribeAckClusterNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckClusterNamespacesResponse) SetStatusCode(v int32) *DescribeAckClusterNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAckClusterNamespacesResponse) SetBody(v *DescribeAckClusterNamespacesResponseBody) *DescribeAckClusterNamespacesResponse {
	s.Body = v
	return s
}

func (s *DescribeAckClusterNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
