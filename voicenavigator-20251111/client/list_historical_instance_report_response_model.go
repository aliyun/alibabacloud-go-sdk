// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalInstanceReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHistoricalInstanceReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHistoricalInstanceReportResponse
	GetStatusCode() *int32
	SetBody(v *ListHistoricalInstanceReportResponseBody) *ListHistoricalInstanceReportResponse
	GetBody() *ListHistoricalInstanceReportResponseBody
}

type ListHistoricalInstanceReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHistoricalInstanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHistoricalInstanceReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalInstanceReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalInstanceReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHistoricalInstanceReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHistoricalInstanceReportResponse) GetBody() *ListHistoricalInstanceReportResponseBody {
	return s.Body
}

func (s *ListHistoricalInstanceReportResponse) SetHeaders(v map[string]*string) *ListHistoricalInstanceReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalInstanceReportResponse) SetStatusCode(v int32) *ListHistoricalInstanceReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoricalInstanceReportResponse) SetBody(v *ListHistoricalInstanceReportResponseBody) *ListHistoricalInstanceReportResponse {
	s.Body = v
	return s
}

func (s *ListHistoricalInstanceReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
