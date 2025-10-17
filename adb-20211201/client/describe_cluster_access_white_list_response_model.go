// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAccessWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAccessWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAccessWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterAccessWhiteListResponseBody) *DescribeClusterAccessWhiteListResponse
	GetBody() *DescribeClusterAccessWhiteListResponseBody
}

type DescribeClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAccessWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAccessWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAccessWhiteListResponse) GetBody() *DescribeClusterAccessWhiteListResponseBody {
	return s.Body
}

func (s *DescribeClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *DescribeClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAccessWhiteListResponse) SetStatusCode(v int32) *DescribeClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponse) SetBody(v *DescribeClusterAccessWhiteListResponseBody) *DescribeClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAccessWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
