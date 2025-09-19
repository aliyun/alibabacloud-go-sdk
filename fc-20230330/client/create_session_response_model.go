// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSessionResponse
	GetStatusCode() *int32
	SetBody(v *Session) *CreateSessionResponse
	GetBody() *Session
}

type CreateSessionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Session           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSessionResponse) GetBody() *Session {
	return s.Body
}

func (s *CreateSessionResponse) SetHeaders(v map[string]*string) *CreateSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionResponse) SetStatusCode(v int32) *CreateSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSessionResponse) SetBody(v *Session) *CreateSessionResponse {
	s.Body = v
	return s
}

func (s *CreateSessionResponse) Validate() error {
	return dara.Validate(s)
}
