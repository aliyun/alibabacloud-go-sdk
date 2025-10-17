// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportSessionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportSessionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportSessionStatusResponse
	GetStatusCode() *int32
	SetBody(v *ReportSessionStatusResponseBody) *ReportSessionStatusResponse
	GetBody() *ReportSessionStatusResponseBody
}

type ReportSessionStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportSessionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportSessionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportSessionStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportSessionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportSessionStatusResponse) GetBody() *ReportSessionStatusResponseBody {
	return s.Body
}

func (s *ReportSessionStatusResponse) SetHeaders(v map[string]*string) *ReportSessionStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportSessionStatusResponse) SetStatusCode(v int32) *ReportSessionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportSessionStatusResponse) SetBody(v *ReportSessionStatusResponseBody) *ReportSessionStatusResponse {
	s.Body = v
	return s
}

func (s *ReportSessionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
