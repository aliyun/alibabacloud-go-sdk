// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRectifyImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RectifyImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RectifyImageResponse
	GetStatusCode() *int32
	SetBody(v *RectifyImageResponseBody) *RectifyImageResponse
	GetBody() *RectifyImageResponseBody
}

type RectifyImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RectifyImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RectifyImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RectifyImageResponse) GoString() string {
	return s.String()
}

func (s *RectifyImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RectifyImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RectifyImageResponse) GetBody() *RectifyImageResponseBody {
	return s.Body
}

func (s *RectifyImageResponse) SetHeaders(v map[string]*string) *RectifyImageResponse {
	s.Headers = v
	return s
}

func (s *RectifyImageResponse) SetStatusCode(v int32) *RectifyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RectifyImageResponse) SetBody(v *RectifyImageResponseBody) *RectifyImageResponse {
	s.Body = v
	return s
}

func (s *RectifyImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
