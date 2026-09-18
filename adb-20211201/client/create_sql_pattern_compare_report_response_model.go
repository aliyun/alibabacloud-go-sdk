// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlPatternCompareReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSqlPatternCompareReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSqlPatternCompareReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateSqlPatternCompareReportResponseBody) *CreateSqlPatternCompareReportResponse
	GetBody() *CreateSqlPatternCompareReportResponseBody
}

type CreateSqlPatternCompareReportResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlPatternCompareReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlPatternCompareReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlPatternCompareReportResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlPatternCompareReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSqlPatternCompareReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSqlPatternCompareReportResponse) GetBody() *CreateSqlPatternCompareReportResponseBody {
	return s.Body
}

func (s *CreateSqlPatternCompareReportResponse) SetHeaders(v map[string]*string) *CreateSqlPatternCompareReportResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlPatternCompareReportResponse) SetStatusCode(v int32) *CreateSqlPatternCompareReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlPatternCompareReportResponse) SetBody(v *CreateSqlPatternCompareReportResponseBody) *CreateSqlPatternCompareReportResponse {
	s.Body = v
	return s
}

func (s *CreateSqlPatternCompareReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
