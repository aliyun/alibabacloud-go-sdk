// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketSummaryReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTicketSummaryReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTicketSummaryReportResponse
	GetStatusCode() *int32
	SetBody(v *GetTicketSummaryReportResponseBody) *GetTicketSummaryReportResponse
	GetBody() *GetTicketSummaryReportResponseBody
}

type GetTicketSummaryReportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTicketSummaryReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTicketSummaryReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTicketSummaryReportResponse) GoString() string {
	return s.String()
}

func (s *GetTicketSummaryReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTicketSummaryReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTicketSummaryReportResponse) GetBody() *GetTicketSummaryReportResponseBody {
	return s.Body
}

func (s *GetTicketSummaryReportResponse) SetHeaders(v map[string]*string) *GetTicketSummaryReportResponse {
	s.Headers = v
	return s
}

func (s *GetTicketSummaryReportResponse) SetStatusCode(v int32) *GetTicketSummaryReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketSummaryReportResponse) SetBody(v *GetTicketSummaryReportResponseBody) *GetTicketSummaryReportResponse {
	s.Body = v
	return s
}

func (s *GetTicketSummaryReportResponse) Validate() error {
	return dara.Validate(s)
}
