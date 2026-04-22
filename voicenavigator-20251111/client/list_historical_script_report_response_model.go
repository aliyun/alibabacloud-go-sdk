// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalScriptReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHistoricalScriptReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHistoricalScriptReportResponse
	GetStatusCode() *int32
	SetBody(v *ListHistoricalScriptReportResponseBody) *ListHistoricalScriptReportResponse
	GetBody() *ListHistoricalScriptReportResponseBody
}

type ListHistoricalScriptReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHistoricalScriptReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHistoricalScriptReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalScriptReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalScriptReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHistoricalScriptReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHistoricalScriptReportResponse) GetBody() *ListHistoricalScriptReportResponseBody {
	return s.Body
}

func (s *ListHistoricalScriptReportResponse) SetHeaders(v map[string]*string) *ListHistoricalScriptReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalScriptReportResponse) SetStatusCode(v int32) *ListHistoricalScriptReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoricalScriptReportResponse) SetBody(v *ListHistoricalScriptReportResponseBody) *ListHistoricalScriptReportResponse {
	s.Body = v
	return s
}

func (s *ListHistoricalScriptReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
