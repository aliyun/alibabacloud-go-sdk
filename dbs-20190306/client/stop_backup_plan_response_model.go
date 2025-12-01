// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *StopBackupPlanResponseBody) *StopBackupPlanResponse
	GetBody() *StopBackupPlanResponseBody
}

type StopBackupPlanResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s StopBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *StopBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopBackupPlanResponse) GetBody() *StopBackupPlanResponseBody {
	return s.Body
}

func (s *StopBackupPlanResponse) SetHeaders(v map[string]*string) *StopBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *StopBackupPlanResponse) SetStatusCode(v int32) *StopBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBackupPlanResponse) SetBody(v *StopBackupPlanResponseBody) *StopBackupPlanResponse {
	s.Body = v
	return s
}

func (s *StopBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
