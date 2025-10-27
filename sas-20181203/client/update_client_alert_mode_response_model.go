// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientAlertModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClientAlertModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClientAlertModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClientAlertModeResponseBody) *UpdateClientAlertModeResponse
	GetBody() *UpdateClientAlertModeResponseBody
}

type UpdateClientAlertModeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientAlertModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientAlertModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientAlertModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientAlertModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClientAlertModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClientAlertModeResponse) GetBody() *UpdateClientAlertModeResponseBody {
	return s.Body
}

func (s *UpdateClientAlertModeResponse) SetHeaders(v map[string]*string) *UpdateClientAlertModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientAlertModeResponse) SetStatusCode(v int32) *UpdateClientAlertModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientAlertModeResponse) SetBody(v *UpdateClientAlertModeResponseBody) *UpdateClientAlertModeResponse {
	s.Body = v
	return s
}

func (s *UpdateClientAlertModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
