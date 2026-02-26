// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeSkillResponse
	GetStatusCode() *int32
	SetId(v string) *InvokeSkillResponse
	GetId() *string
	SetEvent(v string) *InvokeSkillResponse
	GetEvent() *string
	SetBody(v *InvokeSkillResponseBody) *InvokeSkillResponse
	GetBody() *InvokeSkillResponseBody
}

type InvokeSkillResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                  `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                  `json:"event,omitempty" xml:"event,omitempty"`
	Body       *InvokeSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillResponse) GoString() string {
	return s.String()
}

func (s *InvokeSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeSkillResponse) GetId() *string {
	return s.Id
}

func (s *InvokeSkillResponse) GetEvent() *string {
	return s.Event
}

func (s *InvokeSkillResponse) GetBody() *InvokeSkillResponseBody {
	return s.Body
}

func (s *InvokeSkillResponse) SetHeaders(v map[string]*string) *InvokeSkillResponse {
	s.Headers = v
	return s
}

func (s *InvokeSkillResponse) SetStatusCode(v int32) *InvokeSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeSkillResponse) SetId(v string) *InvokeSkillResponse {
	s.Id = &v
	return s
}

func (s *InvokeSkillResponse) SetEvent(v string) *InvokeSkillResponse {
	s.Event = &v
	return s
}

func (s *InvokeSkillResponse) SetBody(v *InvokeSkillResponseBody) *InvokeSkillResponse {
	s.Body = v
	return s
}

func (s *InvokeSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
