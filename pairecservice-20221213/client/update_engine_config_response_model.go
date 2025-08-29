// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEngineConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEngineConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEngineConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEngineConfigResponseBody) *UpdateEngineConfigResponse
	GetBody() *UpdateEngineConfigResponseBody
}

type UpdateEngineConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEngineConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEngineConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateEngineConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEngineConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEngineConfigResponse) GetBody() *UpdateEngineConfigResponseBody {
	return s.Body
}

func (s *UpdateEngineConfigResponse) SetHeaders(v map[string]*string) *UpdateEngineConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateEngineConfigResponse) SetStatusCode(v int32) *UpdateEngineConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEngineConfigResponse) SetBody(v *UpdateEngineConfigResponseBody) *UpdateEngineConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateEngineConfigResponse) Validate() error {
	return dara.Validate(s)
}
