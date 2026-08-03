// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionScheduleReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionScheduleReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionScheduleReportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionScheduleReportsResponseBody) *DescribeInspectionScheduleReportsResponse
	GetBody() *DescribeInspectionScheduleReportsResponseBody
}

type DescribeInspectionScheduleReportsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionScheduleReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionScheduleReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionScheduleReportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionScheduleReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionScheduleReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionScheduleReportsResponse) GetBody() *DescribeInspectionScheduleReportsResponseBody {
	return s.Body
}

func (s *DescribeInspectionScheduleReportsResponse) SetHeaders(v map[string]*string) *DescribeInspectionScheduleReportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionScheduleReportsResponse) SetStatusCode(v int32) *DescribeInspectionScheduleReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionScheduleReportsResponse) SetBody(v *DescribeInspectionScheduleReportsResponseBody) *DescribeInspectionScheduleReportsResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionScheduleReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
