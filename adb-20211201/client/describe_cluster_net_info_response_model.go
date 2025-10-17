// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNetInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterNetInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterNetInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterNetInfoResponseBody) *DescribeClusterNetInfoResponse
	GetBody() *DescribeClusterNetInfoResponseBody
}

type DescribeClusterNetInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterNetInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterNetInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterNetInfoResponse) GetBody() *DescribeClusterNetInfoResponseBody {
	return s.Body
}

func (s *DescribeClusterNetInfoResponse) SetHeaders(v map[string]*string) *DescribeClusterNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNetInfoResponse) SetStatusCode(v int32) *DescribeClusterNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterNetInfoResponse) SetBody(v *DescribeClusterNetInfoResponseBody) *DescribeClusterNetInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterNetInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
