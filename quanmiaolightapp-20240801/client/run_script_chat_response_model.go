// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunScriptChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunScriptChatResponse
	GetStatusCode() *int32
	SetBody(v *RunScriptChatResponseBody) *RunScriptChatResponse
	GetBody() *RunScriptChatResponseBody
}

type RunScriptChatResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunScriptChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunScriptChatResponse) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatResponse) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunScriptChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunScriptChatResponse) GetBody() *RunScriptChatResponseBody {
	return s.Body
}

func (s *RunScriptChatResponse) SetHeaders(v map[string]*string) *RunScriptChatResponse {
	s.Headers = v
	return s
}

func (s *RunScriptChatResponse) SetStatusCode(v int32) *RunScriptChatResponse {
	s.StatusCode = &v
	return s
}

func (s *RunScriptChatResponse) SetBody(v *RunScriptChatResponseBody) *RunScriptChatResponse {
	s.Body = v
	return s
}

func (s *RunScriptChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
