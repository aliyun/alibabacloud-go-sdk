// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookSmartCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunBookSmartCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunBookSmartCardResponse
	GetStatusCode() *int32
	SetId(v string) *RunBookSmartCardResponse
	GetId() *string
	SetEvent(v string) *RunBookSmartCardResponse
	GetEvent() *string
	SetBody(v *RunBookSmartCardResponseBody) *RunBookSmartCardResponse
	GetBody() *RunBookSmartCardResponseBody
}

type RunBookSmartCardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunBookSmartCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunBookSmartCardResponse) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardResponse) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunBookSmartCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunBookSmartCardResponse) GetId() *string {
	return s.Id
}

func (s *RunBookSmartCardResponse) GetEvent() *string {
	return s.Event
}

func (s *RunBookSmartCardResponse) GetBody() *RunBookSmartCardResponseBody {
	return s.Body
}

func (s *RunBookSmartCardResponse) SetHeaders(v map[string]*string) *RunBookSmartCardResponse {
	s.Headers = v
	return s
}

func (s *RunBookSmartCardResponse) SetStatusCode(v int32) *RunBookSmartCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RunBookSmartCardResponse) SetId(v string) *RunBookSmartCardResponse {
	s.Id = &v
	return s
}

func (s *RunBookSmartCardResponse) SetEvent(v string) *RunBookSmartCardResponse {
	s.Event = &v
	return s
}

func (s *RunBookSmartCardResponse) SetBody(v *RunBookSmartCardResponseBody) *RunBookSmartCardResponse {
	s.Body = v
	return s
}

func (s *RunBookSmartCardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
