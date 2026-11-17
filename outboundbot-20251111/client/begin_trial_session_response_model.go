// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginTrialSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BeginTrialSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BeginTrialSessionResponse
	GetStatusCode() *int32
	SetId(v string) *BeginTrialSessionResponse
	GetId() *string
	SetEvent(v string) *BeginTrialSessionResponse
	GetEvent() *string
	SetBody(v *BeginTrialSessionResponseBody) *BeginTrialSessionResponse
	GetBody() *BeginTrialSessionResponseBody
}

type BeginTrialSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
	Body       *BeginTrialSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BeginTrialSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s BeginTrialSessionResponse) GoString() string {
	return s.String()
}

func (s *BeginTrialSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BeginTrialSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BeginTrialSessionResponse) GetId() *string {
	return s.Id
}

func (s *BeginTrialSessionResponse) GetEvent() *string {
	return s.Event
}

func (s *BeginTrialSessionResponse) GetBody() *BeginTrialSessionResponseBody {
	return s.Body
}

func (s *BeginTrialSessionResponse) SetHeaders(v map[string]*string) *BeginTrialSessionResponse {
	s.Headers = v
	return s
}

func (s *BeginTrialSessionResponse) SetStatusCode(v int32) *BeginTrialSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginTrialSessionResponse) SetId(v string) *BeginTrialSessionResponse {
	s.Id = &v
	return s
}

func (s *BeginTrialSessionResponse) SetEvent(v string) *BeginTrialSessionResponse {
	s.Event = &v
	return s
}

func (s *BeginTrialSessionResponse) SetBody(v *BeginTrialSessionResponseBody) *BeginTrialSessionResponse {
	s.Body = v
	return s
}

func (s *BeginTrialSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
