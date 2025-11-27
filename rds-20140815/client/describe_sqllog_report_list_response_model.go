// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLLogReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLLogReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLLogReportListResponseBody) *DescribeSQLLogReportListResponse
	GetBody() *DescribeSQLLogReportListResponseBody
}

type DescribeSQLLogReportListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLLogReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLLogReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLLogReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLLogReportListResponse) GetBody() *DescribeSQLLogReportListResponseBody {
	return s.Body
}

func (s *DescribeSQLLogReportListResponse) SetHeaders(v map[string]*string) *DescribeSQLLogReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogReportListResponse) SetStatusCode(v int32) *DescribeSQLLogReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogReportListResponse) SetBody(v *DescribeSQLLogReportListResponseBody) *DescribeSQLLogReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLLogReportListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
