// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportTaskFailedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportTaskFailedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportTaskFailedResponse
	GetStatusCode() *int32
	SetBody(v *ReportTaskFailedResponseBody) *ReportTaskFailedResponse
	GetBody() *ReportTaskFailedResponseBody
}

type ReportTaskFailedResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportTaskFailedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportTaskFailedResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportTaskFailedResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportTaskFailedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportTaskFailedResponse) GetBody() *ReportTaskFailedResponseBody {
	return s.Body
}

func (s *ReportTaskFailedResponse) SetHeaders(v map[string]*string) *ReportTaskFailedResponse {
	s.Headers = v
	return s
}

func (s *ReportTaskFailedResponse) SetStatusCode(v int32) *ReportTaskFailedResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportTaskFailedResponse) SetBody(v *ReportTaskFailedResponseBody) *ReportTaskFailedResponse {
	s.Body = v
	return s
}

func (s *ReportTaskFailedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
