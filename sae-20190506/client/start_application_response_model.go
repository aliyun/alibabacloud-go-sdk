// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartApplicationResponse
	GetStatusCode() *int32
	SetBody(v *StartApplicationResponseBody) *StartApplicationResponse
	GetBody() *StartApplicationResponseBody
}

type StartApplicationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartApplicationResponse) GoString() string {
	return s.String()
}

func (s *StartApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartApplicationResponse) GetBody() *StartApplicationResponseBody {
	return s.Body
}

func (s *StartApplicationResponse) SetHeaders(v map[string]*string) *StartApplicationResponse {
	s.Headers = v
	return s
}

func (s *StartApplicationResponse) SetStatusCode(v int32) *StartApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartApplicationResponse) SetBody(v *StartApplicationResponseBody) *StartApplicationResponse {
	s.Body = v
	return s
}

func (s *StartApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
