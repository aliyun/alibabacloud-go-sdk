// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAckClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAckClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAckClustersResponseBody) *DescribeAckClustersResponse
	GetBody() *DescribeAckClustersResponseBody
}

type DescribeAckClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAckClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAckClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeAckClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAckClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAckClustersResponse) GetBody() *DescribeAckClustersResponseBody {
	return s.Body
}

func (s *DescribeAckClustersResponse) SetHeaders(v map[string]*string) *DescribeAckClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeAckClustersResponse) SetStatusCode(v int32) *DescribeAckClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAckClustersResponse) SetBody(v *DescribeAckClustersResponseBody) *DescribeAckClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeAckClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
