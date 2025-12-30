// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatisticSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStatisticSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStatisticSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStatisticSummaryResponseBody) *DescribeStatisticSummaryResponse
	GetBody() *DescribeStatisticSummaryResponseBody
}

type DescribeStatisticSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStatisticSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStatisticSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStatisticSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStatisticSummaryResponse) GetBody() *DescribeStatisticSummaryResponseBody {
	return s.Body
}

func (s *DescribeStatisticSummaryResponse) SetHeaders(v map[string]*string) *DescribeStatisticSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatisticSummaryResponse) SetStatusCode(v int32) *DescribeStatisticSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStatisticSummaryResponse) SetBody(v *DescribeStatisticSummaryResponseBody) *DescribeStatisticSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeStatisticSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
