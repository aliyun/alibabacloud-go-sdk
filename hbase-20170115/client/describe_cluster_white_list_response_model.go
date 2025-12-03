// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterWhiteListResponseBody) *DescribeClusterWhiteListResponse
	GetBody() *DescribeClusterWhiteListResponseBody
}

type DescribeClusterWhiteListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterWhiteListResponse) GetBody() *DescribeClusterWhiteListResponseBody {
	return s.Body
}

func (s *DescribeClusterWhiteListResponse) SetHeaders(v map[string]*string) *DescribeClusterWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterWhiteListResponse) SetStatusCode(v int32) *DescribeClusterWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterWhiteListResponse) SetBody(v *DescribeClusterWhiteListResponseBody) *DescribeClusterWhiteListResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
