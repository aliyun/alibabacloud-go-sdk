// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCVClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCVClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCVClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCVClusterResponseBody) *DescribeRCVClusterResponse
	GetBody() *DescribeRCVClusterResponseBody
}

type DescribeRCVClusterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCVClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCVClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCVClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCVClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCVClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCVClusterResponse) GetBody() *DescribeRCVClusterResponseBody {
	return s.Body
}

func (s *DescribeRCVClusterResponse) SetHeaders(v map[string]*string) *DescribeRCVClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCVClusterResponse) SetStatusCode(v int32) *DescribeRCVClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCVClusterResponse) SetBody(v *DescribeRCVClusterResponseBody) *DescribeRCVClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeRCVClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
