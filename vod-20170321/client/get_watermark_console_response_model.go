// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWatermarkConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWatermarkConsoleResponse
	GetStatusCode() *int32
	SetBody(v *GetWatermarkConsoleResponseBody) *GetWatermarkConsoleResponse
	GetBody() *GetWatermarkConsoleResponseBody
}

type GetWatermarkConsoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWatermarkConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWatermarkConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkConsoleResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarkConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWatermarkConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWatermarkConsoleResponse) GetBody() *GetWatermarkConsoleResponseBody {
	return s.Body
}

func (s *GetWatermarkConsoleResponse) SetHeaders(v map[string]*string) *GetWatermarkConsoleResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarkConsoleResponse) SetStatusCode(v int32) *GetWatermarkConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWatermarkConsoleResponse) SetBody(v *GetWatermarkConsoleResponseBody) *GetWatermarkConsoleResponse {
	s.Body = v
	return s
}

func (s *GetWatermarkConsoleResponse) Validate() error {
	return dara.Validate(s)
}
