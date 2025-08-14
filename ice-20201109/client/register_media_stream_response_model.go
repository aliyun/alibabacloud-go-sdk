// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterMediaStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterMediaStreamResponse
	GetStatusCode() *int32
	SetBody(v *RegisterMediaStreamResponseBody) *RegisterMediaStreamResponse
	GetBody() *RegisterMediaStreamResponseBody
}

type RegisterMediaStreamResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterMediaStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterMediaStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaStreamResponse) GoString() string {
	return s.String()
}

func (s *RegisterMediaStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterMediaStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterMediaStreamResponse) GetBody() *RegisterMediaStreamResponseBody {
	return s.Body
}

func (s *RegisterMediaStreamResponse) SetHeaders(v map[string]*string) *RegisterMediaStreamResponse {
	s.Headers = v
	return s
}

func (s *RegisterMediaStreamResponse) SetStatusCode(v int32) *RegisterMediaStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterMediaStreamResponse) SetBody(v *RegisterMediaStreamResponseBody) *RegisterMediaStreamResponse {
	s.Body = v
	return s
}

func (s *RegisterMediaStreamResponse) Validate() error {
	return dara.Validate(s)
}
