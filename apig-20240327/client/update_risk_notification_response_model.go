// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRiskNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRiskNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRiskNotificationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRiskNotificationResponseBody) *UpdateRiskNotificationResponse
	GetBody() *UpdateRiskNotificationResponseBody
}

type UpdateRiskNotificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRiskNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRiskNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRiskNotificationResponse) GoString() string {
	return s.String()
}

func (s *UpdateRiskNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRiskNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRiskNotificationResponse) GetBody() *UpdateRiskNotificationResponseBody {
	return s.Body
}

func (s *UpdateRiskNotificationResponse) SetHeaders(v map[string]*string) *UpdateRiskNotificationResponse {
	s.Headers = v
	return s
}

func (s *UpdateRiskNotificationResponse) SetStatusCode(v int32) *UpdateRiskNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRiskNotificationResponse) SetBody(v *UpdateRiskNotificationResponseBody) *UpdateRiskNotificationResponse {
	s.Body = v
	return s
}

func (s *UpdateRiskNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
