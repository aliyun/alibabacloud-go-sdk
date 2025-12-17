// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnoseReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiagnoseReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiagnoseReportResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiagnoseReportResponseBody) *DeleteDiagnoseReportResponse
	GetBody() *DeleteDiagnoseReportResponseBody
}

type DeleteDiagnoseReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiagnoseReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiagnoseReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiagnoseReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiagnoseReportResponse) GetBody() *DeleteDiagnoseReportResponseBody {
	return s.Body
}

func (s *DeleteDiagnoseReportResponse) SetHeaders(v map[string]*string) *DeleteDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiagnoseReportResponse) SetStatusCode(v int32) *DeleteDiagnoseReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiagnoseReportResponse) SetBody(v *DeleteDiagnoseReportResponseBody) *DeleteDiagnoseReportResponse {
	s.Body = v
	return s
}

func (s *DeleteDiagnoseReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
