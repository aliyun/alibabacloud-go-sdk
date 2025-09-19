// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSessionResponse
	GetStatusCode() *int32
	SetBody(v *Session) *GetSessionResponse
	GetBody() *Session
}

type GetSessionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Session           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSessionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSessionResponse) GetBody() *Session {
	return s.Body
}

func (s *GetSessionResponse) SetHeaders(v map[string]*string) *GetSessionResponse {
	s.Headers = v
	return s
}

func (s *GetSessionResponse) SetStatusCode(v int32) *GetSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSessionResponse) SetBody(v *Session) *GetSessionResponse {
	s.Body = v
	return s
}

func (s *GetSessionResponse) Validate() error {
	return dara.Validate(s)
}
