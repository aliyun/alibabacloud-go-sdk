// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunChatResultGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunChatResultGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunChatResultGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunChatResultGenerationResponseBody) *RunChatResultGenerationResponse
	GetBody() *RunChatResultGenerationResponseBody
}

type RunChatResultGenerationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunChatResultGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunChatResultGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunChatResultGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunChatResultGenerationResponse) GetBody() *RunChatResultGenerationResponseBody {
	return s.Body
}

func (s *RunChatResultGenerationResponse) SetHeaders(v map[string]*string) *RunChatResultGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunChatResultGenerationResponse) SetStatusCode(v int32) *RunChatResultGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunChatResultGenerationResponse) SetBody(v *RunChatResultGenerationResponseBody) *RunChatResultGenerationResponse {
	s.Body = v
	return s
}

func (s *RunChatResultGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
