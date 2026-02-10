// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCompletionMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCompletionMessageResponse
	GetStatusCode() *int32
	SetId(v string) *RunCompletionMessageResponse
	GetId() *string
	SetEvent(v string) *RunCompletionMessageResponse
	GetEvent() *string
	SetBody(v *RunCompletionMessageResponseBody) *RunCompletionMessageResponse
	GetBody() *RunCompletionMessageResponseBody
}

type RunCompletionMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                           `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                           `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunCompletionMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCompletionMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageResponse) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCompletionMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCompletionMessageResponse) GetId() *string {
	return s.Id
}

func (s *RunCompletionMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *RunCompletionMessageResponse) GetBody() *RunCompletionMessageResponseBody {
	return s.Body
}

func (s *RunCompletionMessageResponse) SetHeaders(v map[string]*string) *RunCompletionMessageResponse {
	s.Headers = v
	return s
}

func (s *RunCompletionMessageResponse) SetStatusCode(v int32) *RunCompletionMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCompletionMessageResponse) SetId(v string) *RunCompletionMessageResponse {
	s.Id = &v
	return s
}

func (s *RunCompletionMessageResponse) SetEvent(v string) *RunCompletionMessageResponse {
	s.Event = &v
	return s
}

func (s *RunCompletionMessageResponse) SetBody(v *RunCompletionMessageResponseBody) *RunCompletionMessageResponse {
	s.Body = v
	return s
}

func (s *RunCompletionMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
