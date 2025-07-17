// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertContactResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertContactResponseBody) *CreateAlertContactResponse
	GetBody() *CreateAlertContactResponseBody
}

type CreateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertContactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertContactResponse) GetBody() *CreateAlertContactResponseBody {
	return s.Body
}

func (s *CreateAlertContactResponse) SetHeaders(v map[string]*string) *CreateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactResponse) SetStatusCode(v int32) *CreateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactResponse) SetBody(v *CreateAlertContactResponseBody) *CreateAlertContactResponse {
	s.Body = v
	return s
}

func (s *CreateAlertContactResponse) Validate() error {
	return dara.Validate(s)
}
