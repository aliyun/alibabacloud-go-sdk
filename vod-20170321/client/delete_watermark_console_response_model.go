// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWatermarkConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWatermarkConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWatermarkConsoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWatermarkConsoleResponseBody) *DeleteWatermarkConsoleResponse
	GetBody() *DeleteWatermarkConsoleResponseBody
}

type DeleteWatermarkConsoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWatermarkConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWatermarkConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWatermarkConsoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWatermarkConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWatermarkConsoleResponse) GetBody() *DeleteWatermarkConsoleResponseBody {
	return s.Body
}

func (s *DeleteWatermarkConsoleResponse) SetHeaders(v map[string]*string) *DeleteWatermarkConsoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWatermarkConsoleResponse) SetStatusCode(v int32) *DeleteWatermarkConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWatermarkConsoleResponse) SetBody(v *DeleteWatermarkConsoleResponseBody) *DeleteWatermarkConsoleResponse {
	s.Body = v
	return s
}

func (s *DeleteWatermarkConsoleResponse) Validate() error {
	return dara.Validate(s)
}
