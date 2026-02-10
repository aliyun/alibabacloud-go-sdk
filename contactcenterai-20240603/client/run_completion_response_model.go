// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCompletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCompletionResponse
	GetStatusCode() *int32
	SetId(v string) *RunCompletionResponse
	GetId() *string
	SetEvent(v string) *RunCompletionResponse
	GetEvent() *string
	SetBody(v *RunCompletionResponseBody) *RunCompletionResponse
	GetBody() *RunCompletionResponseBody
}

type RunCompletionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                    `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                    `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunCompletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCompletionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponse) GoString() string {
	return s.String()
}

func (s *RunCompletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCompletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCompletionResponse) GetId() *string {
	return s.Id
}

func (s *RunCompletionResponse) GetEvent() *string {
	return s.Event
}

func (s *RunCompletionResponse) GetBody() *RunCompletionResponseBody {
	return s.Body
}

func (s *RunCompletionResponse) SetHeaders(v map[string]*string) *RunCompletionResponse {
	s.Headers = v
	return s
}

func (s *RunCompletionResponse) SetStatusCode(v int32) *RunCompletionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCompletionResponse) SetId(v string) *RunCompletionResponse {
	s.Id = &v
	return s
}

func (s *RunCompletionResponse) SetEvent(v string) *RunCompletionResponse {
	s.Event = &v
	return s
}

func (s *RunCompletionResponse) SetBody(v *RunCompletionResponseBody) *RunCompletionResponse {
	s.Body = v
	return s
}

func (s *RunCompletionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
