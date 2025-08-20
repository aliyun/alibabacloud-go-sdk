// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSpaceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterSpaceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterSpaceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterSpaceSummaryResponseBody) *DescribeDBClusterSpaceSummaryResponse
	GetBody() *DescribeDBClusterSpaceSummaryResponseBody
}

type DescribeDBClusterSpaceSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterSpaceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterSpaceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSpaceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSpaceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterSpaceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterSpaceSummaryResponse) GetBody() *DescribeDBClusterSpaceSummaryResponseBody {
	return s.Body
}

func (s *DescribeDBClusterSpaceSummaryResponse) SetHeaders(v map[string]*string) *DescribeDBClusterSpaceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponse) SetStatusCode(v int32) *DescribeDBClusterSpaceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponse) SetBody(v *DescribeDBClusterSpaceSummaryResponseBody) *DescribeDBClusterSpaceSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterSpaceSummaryResponse) Validate() error {
	return dara.Validate(s)
}
