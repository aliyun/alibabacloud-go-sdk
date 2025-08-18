// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevelopmentModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDevelopmentModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDevelopmentModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDevelopmentModeResponseBody) *UpdateDevelopmentModeResponse
	GetBody() *UpdateDevelopmentModeResponseBody
}

type UpdateDevelopmentModeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDevelopmentModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDevelopmentModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevelopmentModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevelopmentModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDevelopmentModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDevelopmentModeResponse) GetBody() *UpdateDevelopmentModeResponseBody {
	return s.Body
}

func (s *UpdateDevelopmentModeResponse) SetHeaders(v map[string]*string) *UpdateDevelopmentModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevelopmentModeResponse) SetStatusCode(v int32) *UpdateDevelopmentModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDevelopmentModeResponse) SetBody(v *UpdateDevelopmentModeResponseBody) *UpdateDevelopmentModeResponse {
	s.Body = v
	return s
}

func (s *UpdateDevelopmentModeResponse) Validate() error {
	return dara.Validate(s)
}
