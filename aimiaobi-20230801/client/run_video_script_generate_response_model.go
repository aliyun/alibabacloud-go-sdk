// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoScriptGenerateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunVideoScriptGenerateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunVideoScriptGenerateResponse
	GetStatusCode() *int32
	SetId(v string) *RunVideoScriptGenerateResponse
	GetId() *string
	SetEvent(v string) *RunVideoScriptGenerateResponse
	GetEvent() *string
	SetBody(v *RunVideoScriptGenerateResponseBody) *RunVideoScriptGenerateResponse
	GetBody() *RunVideoScriptGenerateResponseBody
}

type RunVideoScriptGenerateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunVideoScriptGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunVideoScriptGenerateResponse) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponse) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunVideoScriptGenerateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunVideoScriptGenerateResponse) GetId() *string {
	return s.Id
}

func (s *RunVideoScriptGenerateResponse) GetEvent() *string {
	return s.Event
}

func (s *RunVideoScriptGenerateResponse) GetBody() *RunVideoScriptGenerateResponseBody {
	return s.Body
}

func (s *RunVideoScriptGenerateResponse) SetHeaders(v map[string]*string) *RunVideoScriptGenerateResponse {
	s.Headers = v
	return s
}

func (s *RunVideoScriptGenerateResponse) SetStatusCode(v int32) *RunVideoScriptGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponse) SetId(v string) *RunVideoScriptGenerateResponse {
	s.Id = &v
	return s
}

func (s *RunVideoScriptGenerateResponse) SetEvent(v string) *RunVideoScriptGenerateResponse {
	s.Event = &v
	return s
}

func (s *RunVideoScriptGenerateResponse) SetBody(v *RunVideoScriptGenerateResponseBody) *RunVideoScriptGenerateResponse {
	s.Body = v
	return s
}

func (s *RunVideoScriptGenerateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
