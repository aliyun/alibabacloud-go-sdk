// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalCallerReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHistoricalCallerReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHistoricalCallerReportResponse
	GetStatusCode() *int32
	SetBody(v *GetHistoricalCallerReportResponseBody) *GetHistoricalCallerReportResponse
	GetBody() *GetHistoricalCallerReportResponseBody
}

type GetHistoricalCallerReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoricalCallerReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoricalCallerReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCallerReportResponse) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHistoricalCallerReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHistoricalCallerReportResponse) GetBody() *GetHistoricalCallerReportResponseBody {
	return s.Body
}

func (s *GetHistoricalCallerReportResponse) SetHeaders(v map[string]*string) *GetHistoricalCallerReportResponse {
	s.Headers = v
	return s
}

func (s *GetHistoricalCallerReportResponse) SetStatusCode(v int32) *GetHistoricalCallerReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoricalCallerReportResponse) SetBody(v *GetHistoricalCallerReportResponseBody) *GetHistoricalCallerReportResponse {
	s.Body = v
	return s
}

func (s *GetHistoricalCallerReportResponse) Validate() error {
	return dara.Validate(s)
}
