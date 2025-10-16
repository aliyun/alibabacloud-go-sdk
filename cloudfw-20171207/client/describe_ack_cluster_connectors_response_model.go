// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterConnectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAckClusterConnectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAckClusterConnectorsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAckClusterConnectorsResponseBody) *DescribeAckClusterConnectorsResponse
	GetBody() *DescribeAckClusterConnectorsResponseBody
}

type DescribeAckClusterConnectorsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAckClusterConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAckClusterConnectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAckClusterConnectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAckClusterConnectorsResponse) GetBody() *DescribeAckClusterConnectorsResponseBody {
	return s.Body
}

func (s *DescribeAckClusterConnectorsResponse) SetHeaders(v map[string]*string) *DescribeAckClusterConnectorsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckClusterConnectorsResponse) SetStatusCode(v int32) *DescribeAckClusterConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponse) SetBody(v *DescribeAckClusterConnectorsResponseBody) *DescribeAckClusterConnectorsResponse {
	s.Body = v
	return s
}

func (s *DescribeAckClusterConnectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
