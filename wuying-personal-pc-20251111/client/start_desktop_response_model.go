// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDesktopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDesktopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDesktopResponse
	GetStatusCode() *int32
	SetBody(v *StartDesktopResponseBody) *StartDesktopResponse
	GetBody() *StartDesktopResponseBody
}

type StartDesktopResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDesktopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDesktopResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDesktopResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDesktopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDesktopResponse) GetBody() *StartDesktopResponseBody {
	return s.Body
}

func (s *StartDesktopResponse) SetHeaders(v map[string]*string) *StartDesktopResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopResponse) SetStatusCode(v int32) *StartDesktopResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDesktopResponse) SetBody(v *StartDesktopResponseBody) *StartDesktopResponse {
	s.Body = v
	return s
}

func (s *StartDesktopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
