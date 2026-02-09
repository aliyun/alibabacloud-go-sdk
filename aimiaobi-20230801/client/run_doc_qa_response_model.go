// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocQaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocQaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocQaResponse
	GetStatusCode() *int32
	SetId(v string) *RunDocQaResponse
	GetId() *string
	SetEvent(v string) *RunDocQaResponse
	GetEvent() *string
	SetBody(v *RunDocQaResponseBody) *RunDocQaResponse
	GetBody() *RunDocQaResponseBody
}

type RunDocQaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string               `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string               `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDocQaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocQaResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponse) GoString() string {
	return s.String()
}

func (s *RunDocQaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocQaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocQaResponse) GetId() *string {
	return s.Id
}

func (s *RunDocQaResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDocQaResponse) GetBody() *RunDocQaResponseBody {
	return s.Body
}

func (s *RunDocQaResponse) SetHeaders(v map[string]*string) *RunDocQaResponse {
	s.Headers = v
	return s
}

func (s *RunDocQaResponse) SetStatusCode(v int32) *RunDocQaResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocQaResponse) SetId(v string) *RunDocQaResponse {
	s.Id = &v
	return s
}

func (s *RunDocQaResponse) SetEvent(v string) *RunDocQaResponse {
	s.Event = &v
	return s
}

func (s *RunDocQaResponse) SetBody(v *RunDocQaResponseBody) *RunDocQaResponse {
	s.Body = v
	return s
}

func (s *RunDocQaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
