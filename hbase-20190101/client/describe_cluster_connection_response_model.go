// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterConnectionResponseBody) *DescribeClusterConnectionResponse
	GetBody() *DescribeClusterConnectionResponseBody
}

type DescribeClusterConnectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterConnectionResponse) GetBody() *DescribeClusterConnectionResponseBody {
	return s.Body
}

func (s *DescribeClusterConnectionResponse) SetHeaders(v map[string]*string) *DescribeClusterConnectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterConnectionResponse) SetStatusCode(v int32) *DescribeClusterConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterConnectionResponse) SetBody(v *DescribeClusterConnectionResponseBody) *DescribeClusterConnectionResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
