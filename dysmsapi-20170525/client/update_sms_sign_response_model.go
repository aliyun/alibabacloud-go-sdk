// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmsSignResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmsSignResponseBody) *UpdateSmsSignResponse
	GetBody() *UpdateSmsSignResponseBody
}

type UpdateSmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmsSignResponse) GetBody() *UpdateSmsSignResponseBody {
	return s.Body
}

func (s *UpdateSmsSignResponse) SetHeaders(v map[string]*string) *UpdateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsSignResponse) SetStatusCode(v int32) *UpdateSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmsSignResponse) SetBody(v *UpdateSmsSignResponseBody) *UpdateSmsSignResponse {
	s.Body = v
	return s
}

func (s *UpdateSmsSignResponse) Validate() error {
	return dara.Validate(s)
}
