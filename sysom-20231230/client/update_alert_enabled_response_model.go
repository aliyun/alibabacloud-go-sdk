// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertEnabledResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertEnabledResponseBody) *UpdateAlertEnabledResponse
	GetBody() *UpdateAlertEnabledResponseBody
}

type UpdateAlertEnabledResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertEnabledResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertEnabledResponse) GetBody() *UpdateAlertEnabledResponseBody {
	return s.Body
}

func (s *UpdateAlertEnabledResponse) SetHeaders(v map[string]*string) *UpdateAlertEnabledResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertEnabledResponse) SetStatusCode(v int32) *UpdateAlertEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertEnabledResponse) SetBody(v *UpdateAlertEnabledResponseBody) *UpdateAlertEnabledResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertEnabledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
