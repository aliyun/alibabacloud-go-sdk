// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineCheckSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageBaselineCheckSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageBaselineCheckSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageBaselineCheckSummaryResponseBody) *DescribeImageBaselineCheckSummaryResponse
	GetBody() *DescribeImageBaselineCheckSummaryResponseBody
}

type DescribeImageBaselineCheckSummaryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageBaselineCheckSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageBaselineCheckSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageBaselineCheckSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageBaselineCheckSummaryResponse) GetBody() *DescribeImageBaselineCheckSummaryResponseBody {
	return s.Body
}

func (s *DescribeImageBaselineCheckSummaryResponse) SetHeaders(v map[string]*string) *DescribeImageBaselineCheckSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponse) SetStatusCode(v int32) *DescribeImageBaselineCheckSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponse) SetBody(v *DescribeImageBaselineCheckSummaryResponseBody) *DescribeImageBaselineCheckSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponse) Validate() error {
	return dara.Validate(s)
}
