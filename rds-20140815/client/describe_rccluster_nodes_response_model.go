// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCClusterNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCClusterNodesResponseBody) *DescribeRCClusterNodesResponse
	GetBody() *DescribeRCClusterNodesResponseBody
}

type DescribeRCClusterNodesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCClusterNodesResponse) GetBody() *DescribeRCClusterNodesResponseBody {
	return s.Body
}

func (s *DescribeRCClusterNodesResponse) SetHeaders(v map[string]*string) *DescribeRCClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCClusterNodesResponse) SetStatusCode(v int32) *DescribeRCClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCClusterNodesResponse) SetBody(v *DescribeRCClusterNodesResponseBody) *DescribeRCClusterNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeRCClusterNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
