// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContinueContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunContinueContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunContinueContentResponse
	GetStatusCode() *int32
	SetId(v string) *RunContinueContentResponse
	GetId() *string
	SetEvent(v string) *RunContinueContentResponse
	GetEvent() *string
	SetBody(v *RunContinueContentResponseBody) *RunContinueContentResponse
	GetBody() *RunContinueContentResponseBody
}

type RunContinueContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                         `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                         `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunContinueContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContinueContentResponse) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponse) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunContinueContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunContinueContentResponse) GetId() *string {
	return s.Id
}

func (s *RunContinueContentResponse) GetEvent() *string {
	return s.Event
}

func (s *RunContinueContentResponse) GetBody() *RunContinueContentResponseBody {
	return s.Body
}

func (s *RunContinueContentResponse) SetHeaders(v map[string]*string) *RunContinueContentResponse {
	s.Headers = v
	return s
}

func (s *RunContinueContentResponse) SetStatusCode(v int32) *RunContinueContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContinueContentResponse) SetId(v string) *RunContinueContentResponse {
	s.Id = &v
	return s
}

func (s *RunContinueContentResponse) SetEvent(v string) *RunContinueContentResponse {
	s.Event = &v
	return s
}

func (s *RunContinueContentResponse) SetBody(v *RunContinueContentResponseBody) *RunContinueContentResponse {
	s.Body = v
	return s
}

func (s *RunContinueContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
