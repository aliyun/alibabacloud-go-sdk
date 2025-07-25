// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordStatisticsResponseBody) *DescribeRecordStatisticsResponse
	GetBody() *DescribeRecordStatisticsResponseBody
}

type DescribeRecordStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordStatisticsResponse) GetBody() *DescribeRecordStatisticsResponseBody {
	return s.Body
}

func (s *DescribeRecordStatisticsResponse) SetHeaders(v map[string]*string) *DescribeRecordStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordStatisticsResponse) SetStatusCode(v int32) *DescribeRecordStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordStatisticsResponse) SetBody(v *DescribeRecordStatisticsResponseBody) *DescribeRecordStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
