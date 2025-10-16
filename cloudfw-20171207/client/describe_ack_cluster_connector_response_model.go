// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAckClusterConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAckClusterConnectorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAckClusterConnectorResponseBody) *DescribeAckClusterConnectorResponse
	GetBody() *DescribeAckClusterConnectorResponseBody
}

type DescribeAckClusterConnectorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAckClusterConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAckClusterConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAckClusterConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAckClusterConnectorResponse) GetBody() *DescribeAckClusterConnectorResponseBody {
	return s.Body
}

func (s *DescribeAckClusterConnectorResponse) SetHeaders(v map[string]*string) *DescribeAckClusterConnectorResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckClusterConnectorResponse) SetStatusCode(v int32) *DescribeAckClusterConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAckClusterConnectorResponse) SetBody(v *DescribeAckClusterConnectorResponseBody) *DescribeAckClusterConnectorResponse {
	s.Body = v
	return s
}

func (s *DescribeAckClusterConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
