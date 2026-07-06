// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertResponseBody) *UpdateAlertResponse
	GetBody() *UpdateAlertResponseBody
}

type UpdateAlertResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertResponse) GetBody() *UpdateAlertResponseBody {
	return s.Body
}

func (s *UpdateAlertResponse) SetHeaders(v map[string]*string) *UpdateAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertResponse) SetStatusCode(v int32) *UpdateAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertResponse) SetBody(v *UpdateAlertResponseBody) *UpdateAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
