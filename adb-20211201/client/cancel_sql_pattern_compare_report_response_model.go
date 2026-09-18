// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPatternCompareReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSqlPatternCompareReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSqlPatternCompareReportResponse
	GetStatusCode() *int32
	SetBody(v *CancelSqlPatternCompareReportResponseBody) *CancelSqlPatternCompareReportResponse
	GetBody() *CancelSqlPatternCompareReportResponseBody
}

type CancelSqlPatternCompareReportResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSqlPatternCompareReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSqlPatternCompareReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPatternCompareReportResponse) GoString() string {
	return s.String()
}

func (s *CancelSqlPatternCompareReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSqlPatternCompareReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSqlPatternCompareReportResponse) GetBody() *CancelSqlPatternCompareReportResponseBody {
	return s.Body
}

func (s *CancelSqlPatternCompareReportResponse) SetHeaders(v map[string]*string) *CancelSqlPatternCompareReportResponse {
	s.Headers = v
	return s
}

func (s *CancelSqlPatternCompareReportResponse) SetStatusCode(v int32) *CancelSqlPatternCompareReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSqlPatternCompareReportResponse) SetBody(v *CancelSqlPatternCompareReportResponseBody) *CancelSqlPatternCompareReportResponse {
	s.Body = v
	return s
}

func (s *CancelSqlPatternCompareReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
