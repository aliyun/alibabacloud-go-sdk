// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocWashingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocWashingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocWashingResponse
	GetStatusCode() *int32
	SetId(v string) *RunDocWashingResponse
	GetId() *string
	SetEvent(v string) *RunDocWashingResponse
	GetEvent() *string
	SetBody(v *RunDocWashingResponseBody) *RunDocWashingResponse
	GetBody() *RunDocWashingResponseBody
}

type RunDocWashingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                    `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                    `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDocWashingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocWashingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingResponse) GoString() string {
	return s.String()
}

func (s *RunDocWashingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocWashingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocWashingResponse) GetId() *string {
	return s.Id
}

func (s *RunDocWashingResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDocWashingResponse) GetBody() *RunDocWashingResponseBody {
	return s.Body
}

func (s *RunDocWashingResponse) SetHeaders(v map[string]*string) *RunDocWashingResponse {
	s.Headers = v
	return s
}

func (s *RunDocWashingResponse) SetStatusCode(v int32) *RunDocWashingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocWashingResponse) SetId(v string) *RunDocWashingResponse {
	s.Id = &v
	return s
}

func (s *RunDocWashingResponse) SetEvent(v string) *RunDocWashingResponse {
	s.Event = &v
	return s
}

func (s *RunDocWashingResponse) SetBody(v *RunDocWashingResponseBody) *RunDocWashingResponse {
	s.Body = v
	return s
}

func (s *RunDocWashingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
