// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterAccessWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterAccessWhiteListResponseBody) *DescribeDBClusterAccessWhiteListResponse
	GetBody() *DescribeDBClusterAccessWhiteListResponseBody
}

type DescribeDBClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterAccessWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterAccessWhiteListResponse) GetBody() *DescribeDBClusterAccessWhiteListResponseBody {
	return s.Body
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetStatusCode(v int32) *DescribeDBClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetBody(v *DescribeDBClusterAccessWhiteListResponseBody) *DescribeDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
