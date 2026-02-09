// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDeepWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDeepWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDeepWritingResponse
	GetStatusCode() *int32
	SetId(v string) *RunDeepWritingResponse
	GetId() *string
	SetEvent(v string) *RunDeepWritingResponse
	GetEvent() *string
	SetBody(v *RunDeepWritingResponseBody) *RunDeepWritingResponse
	GetBody() *RunDeepWritingResponseBody
}

type RunDeepWritingResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                     `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                     `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDeepWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDeepWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponse) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDeepWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDeepWritingResponse) GetId() *string {
	return s.Id
}

func (s *RunDeepWritingResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDeepWritingResponse) GetBody() *RunDeepWritingResponseBody {
	return s.Body
}

func (s *RunDeepWritingResponse) SetHeaders(v map[string]*string) *RunDeepWritingResponse {
	s.Headers = v
	return s
}

func (s *RunDeepWritingResponse) SetStatusCode(v int32) *RunDeepWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDeepWritingResponse) SetId(v string) *RunDeepWritingResponse {
	s.Id = &v
	return s
}

func (s *RunDeepWritingResponse) SetEvent(v string) *RunDeepWritingResponse {
	s.Event = &v
	return s
}

func (s *RunDeepWritingResponse) SetBody(v *RunDeepWritingResponseBody) *RunDeepWritingResponse {
	s.Body = v
	return s
}

func (s *RunDeepWritingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
