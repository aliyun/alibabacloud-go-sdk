// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrialDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrialDialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrialDialogueResponse
	GetStatusCode() *int32
	SetId(v string) *TrialDialogueResponse
	GetId() *string
	SetEvent(v string) *TrialDialogueResponse
	GetEvent() *string
	SetBody(v *TrialDialogueResponseBody) *TrialDialogueResponse
	GetBody() *TrialDialogueResponseBody
}

type TrialDialogueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                    `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                    `json:"event,omitempty" xml:"event,omitempty"`
	Body       *TrialDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrialDialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s TrialDialogueResponse) GoString() string {
	return s.String()
}

func (s *TrialDialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrialDialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrialDialogueResponse) GetId() *string {
	return s.Id
}

func (s *TrialDialogueResponse) GetEvent() *string {
	return s.Event
}

func (s *TrialDialogueResponse) GetBody() *TrialDialogueResponseBody {
	return s.Body
}

func (s *TrialDialogueResponse) SetHeaders(v map[string]*string) *TrialDialogueResponse {
	s.Headers = v
	return s
}

func (s *TrialDialogueResponse) SetStatusCode(v int32) *TrialDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *TrialDialogueResponse) SetId(v string) *TrialDialogueResponse {
	s.Id = &v
	return s
}

func (s *TrialDialogueResponse) SetEvent(v string) *TrialDialogueResponse {
	s.Event = &v
	return s
}

func (s *TrialDialogueResponse) SetBody(v *TrialDialogueResponseBody) *TrialDialogueResponse {
	s.Body = v
	return s
}

func (s *TrialDialogueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
