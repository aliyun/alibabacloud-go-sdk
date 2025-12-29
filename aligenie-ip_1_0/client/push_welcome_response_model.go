// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushWelcomeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushWelcomeResponse
	GetStatusCode() *int32
	SetBody(v *PushWelcomeResponseBody) *PushWelcomeResponse
	GetBody() *PushWelcomeResponseBody
}

type PushWelcomeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushWelcomeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushWelcomeResponse) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeResponse) GoString() string {
	return s.String()
}

func (s *PushWelcomeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushWelcomeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushWelcomeResponse) GetBody() *PushWelcomeResponseBody {
	return s.Body
}

func (s *PushWelcomeResponse) SetHeaders(v map[string]*string) *PushWelcomeResponse {
	s.Headers = v
	return s
}

func (s *PushWelcomeResponse) SetStatusCode(v int32) *PushWelcomeResponse {
	s.StatusCode = &v
	return s
}

func (s *PushWelcomeResponse) SetBody(v *PushWelcomeResponseBody) *PushWelcomeResponse {
	s.Body = v
	return s
}

func (s *PushWelcomeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
