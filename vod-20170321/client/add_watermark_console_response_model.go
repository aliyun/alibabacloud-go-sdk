// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWatermarkConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWatermarkConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWatermarkConsoleResponse
	GetStatusCode() *int32
	SetBody(v *AddWatermarkConsoleResponseBody) *AddWatermarkConsoleResponse
	GetBody() *AddWatermarkConsoleResponseBody
}

type AddWatermarkConsoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWatermarkConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWatermarkConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkConsoleResponse) GoString() string {
	return s.String()
}

func (s *AddWatermarkConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWatermarkConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWatermarkConsoleResponse) GetBody() *AddWatermarkConsoleResponseBody {
	return s.Body
}

func (s *AddWatermarkConsoleResponse) SetHeaders(v map[string]*string) *AddWatermarkConsoleResponse {
	s.Headers = v
	return s
}

func (s *AddWatermarkConsoleResponse) SetStatusCode(v int32) *AddWatermarkConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWatermarkConsoleResponse) SetBody(v *AddWatermarkConsoleResponseBody) *AddWatermarkConsoleResponse {
	s.Body = v
	return s
}

func (s *AddWatermarkConsoleResponse) Validate() error {
	return dara.Validate(s)
}
