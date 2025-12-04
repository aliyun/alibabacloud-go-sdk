// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMonitorAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMonitorAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMonitorAlertResponse
	GetStatusCode() *int32
	SetBody(v *CheckMonitorAlertResponseBody) *CheckMonitorAlertResponse
	GetBody() *CheckMonitorAlertResponseBody
}

type CheckMonitorAlertResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMonitorAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMonitorAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMonitorAlertResponse) GoString() string {
	return s.String()
}

func (s *CheckMonitorAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMonitorAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMonitorAlertResponse) GetBody() *CheckMonitorAlertResponseBody {
	return s.Body
}

func (s *CheckMonitorAlertResponse) SetHeaders(v map[string]*string) *CheckMonitorAlertResponse {
	s.Headers = v
	return s
}

func (s *CheckMonitorAlertResponse) SetStatusCode(v int32) *CheckMonitorAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMonitorAlertResponse) SetBody(v *CheckMonitorAlertResponseBody) *CheckMonitorAlertResponse {
	s.Body = v
	return s
}

func (s *CheckMonitorAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
