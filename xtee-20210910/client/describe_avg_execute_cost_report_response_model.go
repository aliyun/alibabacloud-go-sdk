// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvgExecuteCostReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvgExecuteCostReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvgExecuteCostReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvgExecuteCostReportResponseBody) *DescribeAvgExecuteCostReportResponse
	GetBody() *DescribeAvgExecuteCostReportResponseBody
}

type DescribeAvgExecuteCostReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvgExecuteCostReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvgExecuteCostReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvgExecuteCostReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvgExecuteCostReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvgExecuteCostReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvgExecuteCostReportResponse) GetBody() *DescribeAvgExecuteCostReportResponseBody {
	return s.Body
}

func (s *DescribeAvgExecuteCostReportResponse) SetHeaders(v map[string]*string) *DescribeAvgExecuteCostReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvgExecuteCostReportResponse) SetStatusCode(v int32) *DescribeAvgExecuteCostReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponse) SetBody(v *DescribeAvgExecuteCostReportResponseBody) *DescribeAvgExecuteCostReportResponse {
	s.Body = v
	return s
}

func (s *DescribeAvgExecuteCostReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
