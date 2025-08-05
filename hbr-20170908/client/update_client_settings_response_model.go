// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClientSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClientSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClientSettingsResponseBody) *UpdateClientSettingsResponse
	GetBody() *UpdateClientSettingsResponseBody
}

type UpdateClientSettingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClientSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClientSettingsResponse) GetBody() *UpdateClientSettingsResponseBody {
	return s.Body
}

func (s *UpdateClientSettingsResponse) SetHeaders(v map[string]*string) *UpdateClientSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientSettingsResponse) SetStatusCode(v int32) *UpdateClientSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientSettingsResponse) SetBody(v *UpdateClientSettingsResponseBody) *UpdateClientSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateClientSettingsResponse) Validate() error {
	return dara.Validate(s)
}
