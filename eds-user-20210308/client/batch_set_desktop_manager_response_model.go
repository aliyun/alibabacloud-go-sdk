// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDesktopManagerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetDesktopManagerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetDesktopManagerResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetDesktopManagerResponseBody) *BatchSetDesktopManagerResponse
	GetBody() *BatchSetDesktopManagerResponseBody
}

type BatchSetDesktopManagerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetDesktopManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetDesktopManagerResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDesktopManagerResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDesktopManagerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetDesktopManagerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetDesktopManagerResponse) GetBody() *BatchSetDesktopManagerResponseBody {
	return s.Body
}

func (s *BatchSetDesktopManagerResponse) SetHeaders(v map[string]*string) *BatchSetDesktopManagerResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDesktopManagerResponse) SetStatusCode(v int32) *BatchSetDesktopManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetDesktopManagerResponse) SetBody(v *BatchSetDesktopManagerResponseBody) *BatchSetDesktopManagerResponse {
	s.Body = v
	return s
}

func (s *BatchSetDesktopManagerResponse) Validate() error {
	return dara.Validate(s)
}
