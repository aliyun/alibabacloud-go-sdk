// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCopilotTurnStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeCopilotTurnStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeCopilotTurnStreamResponse
	GetStatusCode() *int32
	SetId(v string) *ResumeCopilotTurnStreamResponse
	GetId() *string
	SetEvent(v string) *ResumeCopilotTurnStreamResponse
	GetEvent() *string
	SetBody(v string) *ResumeCopilotTurnStreamResponse
	GetBody() *string
}

type ResumeCopilotTurnStreamResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeCopilotTurnStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeCopilotTurnStreamResponse) GoString() string {
	return s.String()
}

func (s *ResumeCopilotTurnStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeCopilotTurnStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeCopilotTurnStreamResponse) GetId() *string {
	return s.Id
}

func (s *ResumeCopilotTurnStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *ResumeCopilotTurnStreamResponse) GetBody() *string {
	return s.Body
}

func (s *ResumeCopilotTurnStreamResponse) SetHeaders(v map[string]*string) *ResumeCopilotTurnStreamResponse {
	s.Headers = v
	return s
}

func (s *ResumeCopilotTurnStreamResponse) SetStatusCode(v int32) *ResumeCopilotTurnStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeCopilotTurnStreamResponse) SetId(v string) *ResumeCopilotTurnStreamResponse {
	s.Id = &v
	return s
}

func (s *ResumeCopilotTurnStreamResponse) SetEvent(v string) *ResumeCopilotTurnStreamResponse {
	s.Event = &v
	return s
}

func (s *ResumeCopilotTurnStreamResponse) SetBody(v string) *ResumeCopilotTurnStreamResponse {
	s.Body = &v
	return s
}

func (s *ResumeCopilotTurnStreamResponse) Validate() error {
	return dara.Validate(s)
}
