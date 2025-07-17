// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertContactResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertContactResponseBody) *UpdateAlertContactResponse
	GetBody() *UpdateAlertContactResponseBody
}

type UpdateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertContactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertContactResponse) GetBody() *UpdateAlertContactResponseBody {
	return s.Body
}

func (s *UpdateAlertContactResponse) SetHeaders(v map[string]*string) *UpdateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertContactResponse) SetStatusCode(v int32) *UpdateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertContactResponse) SetBody(v *UpdateAlertContactResponseBody) *UpdateAlertContactResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertContactResponse) Validate() error {
	return dara.Validate(s)
}
