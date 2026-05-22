// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQaChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QaChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QaChatResponse
	GetStatusCode() *int32
	SetId(v string) *QaChatResponse
	GetId() *string
	SetEvent(v string) *QaChatResponse
	GetEvent() *string
	SetBody(v *QaChatResponseBody) *QaChatResponse
	GetBody() *QaChatResponseBody
}

type QaChatResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *QaChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QaChatResponse) String() string {
	return dara.Prettify(s)
}

func (s QaChatResponse) GoString() string {
	return s.String()
}

func (s *QaChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QaChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QaChatResponse) GetId() *string {
	return s.Id
}

func (s *QaChatResponse) GetEvent() *string {
	return s.Event
}

func (s *QaChatResponse) GetBody() *QaChatResponseBody {
	return s.Body
}

func (s *QaChatResponse) SetHeaders(v map[string]*string) *QaChatResponse {
	s.Headers = v
	return s
}

func (s *QaChatResponse) SetStatusCode(v int32) *QaChatResponse {
	s.StatusCode = &v
	return s
}

func (s *QaChatResponse) SetId(v string) *QaChatResponse {
	s.Id = &v
	return s
}

func (s *QaChatResponse) SetEvent(v string) *QaChatResponse {
	s.Event = &v
	return s
}

func (s *QaChatResponse) SetBody(v *QaChatResponseBody) *QaChatResponse {
	s.Body = v
	return s
}

func (s *QaChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
