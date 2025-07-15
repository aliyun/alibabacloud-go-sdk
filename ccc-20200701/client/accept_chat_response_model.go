// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptChatResponse
	GetStatusCode() *int32
	SetBody(v *AcceptChatResponseBody) *AcceptChatResponse
	GetBody() *AcceptChatResponseBody
}

type AcceptChatResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptChatResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptChatResponse) GoString() string {
	return s.String()
}

func (s *AcceptChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptChatResponse) GetBody() *AcceptChatResponseBody {
	return s.Body
}

func (s *AcceptChatResponse) SetHeaders(v map[string]*string) *AcceptChatResponse {
	s.Headers = v
	return s
}

func (s *AcceptChatResponse) SetStatusCode(v int32) *AcceptChatResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptChatResponse) SetBody(v *AcceptChatResponseBody) *AcceptChatResponse {
	s.Body = v
	return s
}

func (s *AcceptChatResponse) Validate() error {
	return dara.Validate(s)
}
