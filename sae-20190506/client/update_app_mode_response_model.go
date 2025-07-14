// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppModeResponseBody) *UpdateAppModeResponse
	GetBody() *UpdateAppModeResponseBody
}

type UpdateAppModeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppModeResponse) GetBody() *UpdateAppModeResponseBody {
	return s.Body
}

func (s *UpdateAppModeResponse) SetHeaders(v map[string]*string) *UpdateAppModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppModeResponse) SetStatusCode(v int32) *UpdateAppModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppModeResponse) SetBody(v *UpdateAppModeResponseBody) *UpdateAppModeResponse {
	s.Body = v
	return s
}

func (s *UpdateAppModeResponse) Validate() error {
	return dara.Validate(s)
}
