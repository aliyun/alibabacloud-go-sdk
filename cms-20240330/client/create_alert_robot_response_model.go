// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRobotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertRobotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertRobotResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertRobotResponseBody) *CreateAlertRobotResponse
	GetBody() *CreateAlertRobotResponseBody
}

type CreateAlertRobotResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertRobotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRobotResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertRobotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertRobotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertRobotResponse) GetBody() *CreateAlertRobotResponseBody {
	return s.Body
}

func (s *CreateAlertRobotResponse) SetHeaders(v map[string]*string) *CreateAlertRobotResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertRobotResponse) SetStatusCode(v int32) *CreateAlertRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertRobotResponse) SetBody(v *CreateAlertRobotResponseBody) *CreateAlertRobotResponse {
	s.Body = v
	return s
}

func (s *CreateAlertRobotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
