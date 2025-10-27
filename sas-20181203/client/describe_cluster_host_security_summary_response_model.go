// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterHostSecuritySummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterHostSecuritySummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterHostSecuritySummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterHostSecuritySummaryResponseBody) *DescribeClusterHostSecuritySummaryResponse
	GetBody() *DescribeClusterHostSecuritySummaryResponseBody
}

type DescribeClusterHostSecuritySummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterHostSecuritySummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterHostSecuritySummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterHostSecuritySummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostSecuritySummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterHostSecuritySummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterHostSecuritySummaryResponse) GetBody() *DescribeClusterHostSecuritySummaryResponseBody {
	return s.Body
}

func (s *DescribeClusterHostSecuritySummaryResponse) SetHeaders(v map[string]*string) *DescribeClusterHostSecuritySummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponse) SetStatusCode(v int32) *DescribeClusterHostSecuritySummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponse) SetBody(v *DescribeClusterHostSecuritySummaryResponseBody) *DescribeClusterHostSecuritySummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterHostSecuritySummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
