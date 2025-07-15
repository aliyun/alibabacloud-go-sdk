// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseChatResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseChatResponseBody) *ReleaseChatResponse
	GetBody() *ReleaseChatResponseBody
}

type ReleaseChatResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseChatResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseChatResponse) GoString() string {
	return s.String()
}

func (s *ReleaseChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseChatResponse) GetBody() *ReleaseChatResponseBody {
	return s.Body
}

func (s *ReleaseChatResponse) SetHeaders(v map[string]*string) *ReleaseChatResponse {
	s.Headers = v
	return s
}

func (s *ReleaseChatResponse) SetStatusCode(v int32) *ReleaseChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseChatResponse) SetBody(v *ReleaseChatResponseBody) *ReleaseChatResponse {
	s.Body = v
	return s
}

func (s *ReleaseChatResponse) Validate() error {
	return dara.Validate(s)
}
