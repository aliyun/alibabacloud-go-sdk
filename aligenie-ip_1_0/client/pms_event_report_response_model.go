// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPmsEventReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PmsEventReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PmsEventReportResponse
	GetStatusCode() *int32
	SetBody(v *PmsEventReportResponseBody) *PmsEventReportResponse
	GetBody() *PmsEventReportResponseBody
}

type PmsEventReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PmsEventReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PmsEventReportResponse) String() string {
	return dara.Prettify(s)
}

func (s PmsEventReportResponse) GoString() string {
	return s.String()
}

func (s *PmsEventReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PmsEventReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PmsEventReportResponse) GetBody() *PmsEventReportResponseBody {
	return s.Body
}

func (s *PmsEventReportResponse) SetHeaders(v map[string]*string) *PmsEventReportResponse {
	s.Headers = v
	return s
}

func (s *PmsEventReportResponse) SetStatusCode(v int32) *PmsEventReportResponse {
	s.StatusCode = &v
	return s
}

func (s *PmsEventReportResponse) SetBody(v *PmsEventReportResponseBody) *PmsEventReportResponse {
	s.Body = v
	return s
}

func (s *PmsEventReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
