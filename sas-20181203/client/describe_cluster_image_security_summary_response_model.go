// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterImageSecuritySummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterImageSecuritySummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterImageSecuritySummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterImageSecuritySummaryResponseBody) *DescribeClusterImageSecuritySummaryResponse
	GetBody() *DescribeClusterImageSecuritySummaryResponseBody
}

type DescribeClusterImageSecuritySummaryResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterImageSecuritySummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterImageSecuritySummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterImageSecuritySummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterImageSecuritySummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterImageSecuritySummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterImageSecuritySummaryResponse) GetBody() *DescribeClusterImageSecuritySummaryResponseBody {
	return s.Body
}

func (s *DescribeClusterImageSecuritySummaryResponse) SetHeaders(v map[string]*string) *DescribeClusterImageSecuritySummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponse) SetStatusCode(v int32) *DescribeClusterImageSecuritySummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponse) SetBody(v *DescribeClusterImageSecuritySummaryResponseBody) *DescribeClusterImageSecuritySummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterImageSecuritySummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
