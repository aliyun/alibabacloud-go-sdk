// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseSessionResponse
	GetStatusCode() *int32
	SetBody(v *Session) *PauseSessionResponse
	GetBody() *Session
}

type PauseSessionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Session           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseSessionResponse) GoString() string {
	return s.String()
}

func (s *PauseSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseSessionResponse) GetBody() *Session {
	return s.Body
}

func (s *PauseSessionResponse) SetHeaders(v map[string]*string) *PauseSessionResponse {
	s.Headers = v
	return s
}

func (s *PauseSessionResponse) SetStatusCode(v int32) *PauseSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseSessionResponse) SetBody(v *Session) *PauseSessionResponse {
	s.Body = v
	return s
}

func (s *PauseSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
