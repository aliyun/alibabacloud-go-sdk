// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWatermarkConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWatermarkConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWatermarkConsoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWatermarkConsoleResponseBody) *UpdateWatermarkConsoleResponse
	GetBody() *UpdateWatermarkConsoleResponseBody
}

type UpdateWatermarkConsoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWatermarkConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWatermarkConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkConsoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWatermarkConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWatermarkConsoleResponse) GetBody() *UpdateWatermarkConsoleResponseBody {
	return s.Body
}

func (s *UpdateWatermarkConsoleResponse) SetHeaders(v map[string]*string) *UpdateWatermarkConsoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWatermarkConsoleResponse) SetStatusCode(v int32) *UpdateWatermarkConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWatermarkConsoleResponse) SetBody(v *UpdateWatermarkConsoleResponseBody) *UpdateWatermarkConsoleResponse {
	s.Body = v
	return s
}

func (s *UpdateWatermarkConsoleResponse) Validate() error {
	return dara.Validate(s)
}
