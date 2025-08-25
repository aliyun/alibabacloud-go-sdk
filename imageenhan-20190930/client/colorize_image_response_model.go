// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColorizeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ColorizeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ColorizeImageResponse
	GetStatusCode() *int32
	SetBody(v *ColorizeImageResponseBody) *ColorizeImageResponse
	GetBody() *ColorizeImageResponseBody
}

type ColorizeImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ColorizeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ColorizeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ColorizeImageResponse) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ColorizeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ColorizeImageResponse) GetBody() *ColorizeImageResponseBody {
	return s.Body
}

func (s *ColorizeImageResponse) SetHeaders(v map[string]*string) *ColorizeImageResponse {
	s.Headers = v
	return s
}

func (s *ColorizeImageResponse) SetStatusCode(v int32) *ColorizeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ColorizeImageResponse) SetBody(v *ColorizeImageResponseBody) *ColorizeImageResponse {
	s.Body = v
	return s
}

func (s *ColorizeImageResponse) Validate() error {
	return dara.Validate(s)
}
