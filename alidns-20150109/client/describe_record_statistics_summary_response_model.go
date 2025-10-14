// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordStatisticsSummaryResponseBody) *DescribeRecordStatisticsSummaryResponse
	GetBody() *DescribeRecordStatisticsSummaryResponseBody
}

type DescribeRecordStatisticsSummaryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordStatisticsSummaryResponse) GetBody() *DescribeRecordStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeRecordStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeRecordStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeRecordStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponse) SetBody(v *DescribeRecordStatisticsSummaryResponseBody) *DescribeRecordStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordStatisticsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
