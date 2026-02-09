// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStepByStepWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunStepByStepWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunStepByStepWritingResponse
	GetStatusCode() *int32
	SetId(v string) *RunStepByStepWritingResponse
	GetId() *string
	SetEvent(v string) *RunStepByStepWritingResponse
	GetEvent() *string
	SetBody(v *RunStepByStepWritingResponseBody) *RunStepByStepWritingResponse
	GetBody() *RunStepByStepWritingResponseBody
}

type RunStepByStepWritingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                           `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                           `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunStepByStepWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStepByStepWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponse) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunStepByStepWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunStepByStepWritingResponse) GetId() *string {
	return s.Id
}

func (s *RunStepByStepWritingResponse) GetEvent() *string {
	return s.Event
}

func (s *RunStepByStepWritingResponse) GetBody() *RunStepByStepWritingResponseBody {
	return s.Body
}

func (s *RunStepByStepWritingResponse) SetHeaders(v map[string]*string) *RunStepByStepWritingResponse {
	s.Headers = v
	return s
}

func (s *RunStepByStepWritingResponse) SetStatusCode(v int32) *RunStepByStepWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStepByStepWritingResponse) SetId(v string) *RunStepByStepWritingResponse {
	s.Id = &v
	return s
}

func (s *RunStepByStepWritingResponse) SetEvent(v string) *RunStepByStepWritingResponse {
	s.Event = &v
	return s
}

func (s *RunStepByStepWritingResponse) SetBody(v *RunStepByStepWritingResponseBody) *RunStepByStepWritingResponse {
	s.Body = v
	return s
}

func (s *RunStepByStepWritingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
