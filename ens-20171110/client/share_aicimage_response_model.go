// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareAICImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShareAICImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShareAICImageResponse
	GetStatusCode() *int32
	SetBody(v *ShareAICImageResponseBody) *ShareAICImageResponse
	GetBody() *ShareAICImageResponseBody
}

type ShareAICImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareAICImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShareAICImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ShareAICImageResponse) GoString() string {
	return s.String()
}

func (s *ShareAICImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShareAICImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShareAICImageResponse) GetBody() *ShareAICImageResponseBody {
	return s.Body
}

func (s *ShareAICImageResponse) SetHeaders(v map[string]*string) *ShareAICImageResponse {
	s.Headers = v
	return s
}

func (s *ShareAICImageResponse) SetStatusCode(v int32) *ShareAICImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ShareAICImageResponse) SetBody(v *ShareAICImageResponseBody) *ShareAICImageResponse {
	s.Body = v
	return s
}

func (s *ShareAICImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
