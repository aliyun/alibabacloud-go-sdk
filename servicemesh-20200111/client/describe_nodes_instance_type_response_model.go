// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodesInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodesInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodesInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodesInstanceTypeResponseBody) *DescribeNodesInstanceTypeResponse
	GetBody() *DescribeNodesInstanceTypeResponseBody
}

type DescribeNodesInstanceTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodesInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodesInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodesInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodesInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodesInstanceTypeResponse) GetBody() *DescribeNodesInstanceTypeResponseBody {
	return s.Body
}

func (s *DescribeNodesInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeNodesInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodesInstanceTypeResponse) SetStatusCode(v int32) *DescribeNodesInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodesInstanceTypeResponse) SetBody(v *DescribeNodesInstanceTypeResponseBody) *DescribeNodesInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeNodesInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
