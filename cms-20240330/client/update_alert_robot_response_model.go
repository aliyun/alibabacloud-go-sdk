// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRobotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertRobotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertRobotResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertRobotResponseBody) *UpdateAlertRobotResponse
	GetBody() *UpdateAlertRobotResponseBody
}

type UpdateAlertRobotResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertRobotResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRobotResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertRobotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertRobotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertRobotResponse) GetBody() *UpdateAlertRobotResponseBody {
	return s.Body
}

func (s *UpdateAlertRobotResponse) SetHeaders(v map[string]*string) *UpdateAlertRobotResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertRobotResponse) SetStatusCode(v int32) *UpdateAlertRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertRobotResponse) SetBody(v *UpdateAlertRobotResponseBody) *UpdateAlertRobotResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertRobotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
