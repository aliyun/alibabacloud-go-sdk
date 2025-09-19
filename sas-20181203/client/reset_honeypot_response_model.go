// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *ResetHoneypotResponseBody) *ResetHoneypotResponse
	GetBody() *ResetHoneypotResponseBody
}

type ResetHoneypotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetHoneypotResponse) GoString() string {
	return s.String()
}

func (s *ResetHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetHoneypotResponse) GetBody() *ResetHoneypotResponseBody {
	return s.Body
}

func (s *ResetHoneypotResponse) SetHeaders(v map[string]*string) *ResetHoneypotResponse {
	s.Headers = v
	return s
}

func (s *ResetHoneypotResponse) SetStatusCode(v int32) *ResetHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetHoneypotResponse) SetBody(v *ResetHoneypotResponseBody) *ResetHoneypotResponse {
	s.Body = v
	return s
}

func (s *ResetHoneypotResponse) Validate() error {
	return dara.Validate(s)
}
