// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorDataReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorDataReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorDataReportResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorDataReportResponseBody) *CreateMonitorDataReportResponse
	GetBody() *CreateMonitorDataReportResponseBody
}

type CreateMonitorDataReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorDataReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorDataReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorDataReportResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorDataReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorDataReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorDataReportResponse) GetBody() *CreateMonitorDataReportResponseBody {
	return s.Body
}

func (s *CreateMonitorDataReportResponse) SetHeaders(v map[string]*string) *CreateMonitorDataReportResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorDataReportResponse) SetStatusCode(v int32) *CreateMonitorDataReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorDataReportResponse) SetBody(v *CreateMonitorDataReportResponseBody) *CreateMonitorDataReportResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorDataReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
