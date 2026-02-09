// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunQuickWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunQuickWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunQuickWritingResponse
	GetStatusCode() *int32
	SetId(v string) *RunQuickWritingResponse
	GetId() *string
	SetEvent(v string) *RunQuickWritingResponse
	GetEvent() *string
	SetBody(v *RunQuickWritingResponseBody) *RunQuickWritingResponse
	GetBody() *RunQuickWritingResponseBody
}

type RunQuickWritingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                      `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                      `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunQuickWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunQuickWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingResponse) GoString() string {
	return s.String()
}

func (s *RunQuickWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunQuickWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunQuickWritingResponse) GetId() *string {
	return s.Id
}

func (s *RunQuickWritingResponse) GetEvent() *string {
	return s.Event
}

func (s *RunQuickWritingResponse) GetBody() *RunQuickWritingResponseBody {
	return s.Body
}

func (s *RunQuickWritingResponse) SetHeaders(v map[string]*string) *RunQuickWritingResponse {
	s.Headers = v
	return s
}

func (s *RunQuickWritingResponse) SetStatusCode(v int32) *RunQuickWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunQuickWritingResponse) SetId(v string) *RunQuickWritingResponse {
	s.Id = &v
	return s
}

func (s *RunQuickWritingResponse) SetEvent(v string) *RunQuickWritingResponse {
	s.Event = &v
	return s
}

func (s *RunQuickWritingResponse) SetBody(v *RunQuickWritingResponseBody) *RunQuickWritingResponse {
	s.Body = v
	return s
}

func (s *RunQuickWritingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
