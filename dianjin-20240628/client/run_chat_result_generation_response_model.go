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
	SetId(v string) *RunChatResultGenerationResponse
	GetId() *string
	SetEvent(v string) *RunChatResultGenerationResponse
	GetEvent() *string
	SetBody(v *RunChatResultGenerationResponseBody) *RunChatResultGenerationResponse
	GetBody() *RunChatResultGenerationResponseBody
}

type RunChatResultGenerationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                              `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                              `json:"event,omitempty" xml:"event,omitempty"`
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

func (s *RunChatResultGenerationResponse) GetId() *string {
	return s.Id
}

func (s *RunChatResultGenerationResponse) GetEvent() *string {
	return s.Event
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

func (s *RunChatResultGenerationResponse) SetId(v string) *RunChatResultGenerationResponse {
	s.Id = &v
	return s
}

func (s *RunChatResultGenerationResponse) SetEvent(v string) *RunChatResultGenerationResponse {
	s.Event = &v
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
