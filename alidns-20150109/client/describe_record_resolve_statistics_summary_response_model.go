// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordResolveStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordResolveStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordResolveStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordResolveStatisticsSummaryResponseBody) *DescribeRecordResolveStatisticsSummaryResponse
	GetBody() *DescribeRecordResolveStatisticsSummaryResponseBody
}

type DescribeRecordResolveStatisticsSummaryResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordResolveStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordResolveStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordResolveStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) GetBody() *DescribeRecordResolveStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeRecordResolveStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeRecordResolveStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) SetBody(v *DescribeRecordResolveStatisticsSummaryResponseBody) *DescribeRecordResolveStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryResponse) Validate() error {
	return dara.Validate(s)
}
