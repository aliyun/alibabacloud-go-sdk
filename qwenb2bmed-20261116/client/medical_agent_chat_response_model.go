// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAgentChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MedicalAgentChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MedicalAgentChatResponse
	GetStatusCode() *int32
	SetId(v string) *MedicalAgentChatResponse
	GetId() *string
	SetEvent(v string) *MedicalAgentChatResponse
	GetEvent() *string
	SetBody(v *MedicalAgentChatResponseBody) *MedicalAgentChatResponse
	GetBody() *MedicalAgentChatResponseBody
}

type MedicalAgentChatResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *MedicalAgentChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MedicalAgentChatResponse) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatResponse) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MedicalAgentChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MedicalAgentChatResponse) GetId() *string {
	return s.Id
}

func (s *MedicalAgentChatResponse) GetEvent() *string {
	return s.Event
}

func (s *MedicalAgentChatResponse) GetBody() *MedicalAgentChatResponseBody {
	return s.Body
}

func (s *MedicalAgentChatResponse) SetHeaders(v map[string]*string) *MedicalAgentChatResponse {
	s.Headers = v
	return s
}

func (s *MedicalAgentChatResponse) SetStatusCode(v int32) *MedicalAgentChatResponse {
	s.StatusCode = &v
	return s
}

func (s *MedicalAgentChatResponse) SetId(v string) *MedicalAgentChatResponse {
	s.Id = &v
	return s
}

func (s *MedicalAgentChatResponse) SetEvent(v string) *MedicalAgentChatResponse {
	s.Event = &v
	return s
}

func (s *MedicalAgentChatResponse) SetBody(v *MedicalAgentChatResponseBody) *MedicalAgentChatResponse {
	s.Body = v
	return s
}

func (s *MedicalAgentChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
