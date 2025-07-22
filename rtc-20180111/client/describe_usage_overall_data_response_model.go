// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageOverallDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsageOverallDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsageOverallDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsageOverallDataResponseBody) *DescribeUsageOverallDataResponse
	GetBody() *DescribeUsageOverallDataResponseBody
}

type DescribeUsageOverallDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageOverallDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsageOverallDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsageOverallDataResponse) GetBody() *DescribeUsageOverallDataResponseBody {
	return s.Body
}

func (s *DescribeUsageOverallDataResponse) SetHeaders(v map[string]*string) *DescribeUsageOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageOverallDataResponse) SetStatusCode(v int32) *DescribeUsageOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageOverallDataResponse) SetBody(v *DescribeUsageOverallDataResponseBody) *DescribeUsageOverallDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUsageOverallDataResponse) Validate() error {
	return dara.Validate(s)
}
