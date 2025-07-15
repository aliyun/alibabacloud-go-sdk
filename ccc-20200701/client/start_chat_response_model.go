// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartChatResponse
	GetStatusCode() *int32
	SetBody(v *StartChatResponseBody) *StartChatResponse
	GetBody() *StartChatResponseBody
}

type StartChatResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartChatResponse) String() string {
	return dara.Prettify(s)
}

func (s StartChatResponse) GoString() string {
	return s.String()
}

func (s *StartChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartChatResponse) GetBody() *StartChatResponseBody {
	return s.Body
}

func (s *StartChatResponse) SetHeaders(v map[string]*string) *StartChatResponse {
	s.Headers = v
	return s
}

func (s *StartChatResponse) SetStatusCode(v int32) *StartChatResponse {
	s.StatusCode = &v
	return s
}

func (s *StartChatResponse) SetBody(v *StartChatResponseBody) *StartChatResponse {
	s.Body = v
	return s
}

func (s *StartChatResponse) Validate() error {
	return dara.Validate(s)
}
