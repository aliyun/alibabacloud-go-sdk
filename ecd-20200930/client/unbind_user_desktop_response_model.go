// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserDesktopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindUserDesktopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindUserDesktopResponse
	GetStatusCode() *int32
	SetBody(v *UnbindUserDesktopResponseBody) *UnbindUserDesktopResponse
	GetBody() *UnbindUserDesktopResponseBody
}

type UnbindUserDesktopResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindUserDesktopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindUserDesktopResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserDesktopResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindUserDesktopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindUserDesktopResponse) GetBody() *UnbindUserDesktopResponseBody {
	return s.Body
}

func (s *UnbindUserDesktopResponse) SetHeaders(v map[string]*string) *UnbindUserDesktopResponse {
	s.Headers = v
	return s
}

func (s *UnbindUserDesktopResponse) SetStatusCode(v int32) *UnbindUserDesktopResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindUserDesktopResponse) SetBody(v *UnbindUserDesktopResponseBody) *UnbindUserDesktopResponse {
	s.Body = v
	return s
}

func (s *UnbindUserDesktopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
