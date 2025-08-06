// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultWatermarkConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultWatermarkConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultWatermarkConsoleResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultWatermarkConsoleResponseBody) *SetDefaultWatermarkConsoleResponse
	GetBody() *SetDefaultWatermarkConsoleResponseBody
}

type SetDefaultWatermarkConsoleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultWatermarkConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultWatermarkConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultWatermarkConsoleResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultWatermarkConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultWatermarkConsoleResponse) GetBody() *SetDefaultWatermarkConsoleResponseBody {
	return s.Body
}

func (s *SetDefaultWatermarkConsoleResponse) SetHeaders(v map[string]*string) *SetDefaultWatermarkConsoleResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultWatermarkConsoleResponse) SetStatusCode(v int32) *SetDefaultWatermarkConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultWatermarkConsoleResponse) SetBody(v *SetDefaultWatermarkConsoleResponseBody) *SetDefaultWatermarkConsoleResponse {
	s.Body = v
	return s
}

func (s *SetDefaultWatermarkConsoleResponse) Validate() error {
	return dara.Validate(s)
}
