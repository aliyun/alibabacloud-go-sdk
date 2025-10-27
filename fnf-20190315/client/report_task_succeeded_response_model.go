// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportTaskSucceededResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportTaskSucceededResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportTaskSucceededResponse
	GetStatusCode() *int32
	SetBody(v *ReportTaskSucceededResponseBody) *ReportTaskSucceededResponse
	GetBody() *ReportTaskSucceededResponseBody
}

type ReportTaskSucceededResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportTaskSucceededResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportTaskSucceededResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportTaskSucceededResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportTaskSucceededResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportTaskSucceededResponse) GetBody() *ReportTaskSucceededResponseBody {
	return s.Body
}

func (s *ReportTaskSucceededResponse) SetHeaders(v map[string]*string) *ReportTaskSucceededResponse {
	s.Headers = v
	return s
}

func (s *ReportTaskSucceededResponse) SetStatusCode(v int32) *ReportTaskSucceededResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportTaskSucceededResponse) SetBody(v *ReportTaskSucceededResponseBody) *ReportTaskSucceededResponse {
	s.Body = v
	return s
}

func (s *ReportTaskSucceededResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
