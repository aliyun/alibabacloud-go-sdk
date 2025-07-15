// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalInstanceReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHistoricalInstanceReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHistoricalInstanceReportResponse
	GetStatusCode() *int32
	SetBody(v *GetHistoricalInstanceReportResponseBody) *GetHistoricalInstanceReportResponse
	GetBody() *GetHistoricalInstanceReportResponseBody
}

type GetHistoricalInstanceReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoricalInstanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoricalInstanceReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponse) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHistoricalInstanceReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHistoricalInstanceReportResponse) GetBody() *GetHistoricalInstanceReportResponseBody {
	return s.Body
}

func (s *GetHistoricalInstanceReportResponse) SetHeaders(v map[string]*string) *GetHistoricalInstanceReportResponse {
	s.Headers = v
	return s
}

func (s *GetHistoricalInstanceReportResponse) SetStatusCode(v int32) *GetHistoricalInstanceReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoricalInstanceReportResponse) SetBody(v *GetHistoricalInstanceReportResponseBody) *GetHistoricalInstanceReportResponse {
	s.Body = v
	return s
}

func (s *GetHistoricalInstanceReportResponse) Validate() error {
	return dara.Validate(s)
}
