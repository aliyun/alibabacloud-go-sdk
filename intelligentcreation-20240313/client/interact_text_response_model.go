// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInteractTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InteractTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InteractTextResponse
	GetStatusCode() *int32
	SetId(v string) *InteractTextResponse
	GetId() *string
	SetEvent(v string) *InteractTextResponse
	GetEvent() *string
	SetBody(v *InteractTextResponseBody) *InteractTextResponse
	GetBody() *InteractTextResponseBody
}

type InteractTextResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                   `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                   `json:"event,omitempty" xml:"event,omitempty"`
	Body       *InteractTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InteractTextResponse) String() string {
	return dara.Prettify(s)
}

func (s InteractTextResponse) GoString() string {
	return s.String()
}

func (s *InteractTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InteractTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InteractTextResponse) GetId() *string {
	return s.Id
}

func (s *InteractTextResponse) GetEvent() *string {
	return s.Event
}

func (s *InteractTextResponse) GetBody() *InteractTextResponseBody {
	return s.Body
}

func (s *InteractTextResponse) SetHeaders(v map[string]*string) *InteractTextResponse {
	s.Headers = v
	return s
}

func (s *InteractTextResponse) SetStatusCode(v int32) *InteractTextResponse {
	s.StatusCode = &v
	return s
}

func (s *InteractTextResponse) SetId(v string) *InteractTextResponse {
	s.Id = &v
	return s
}

func (s *InteractTextResponse) SetEvent(v string) *InteractTextResponse {
	s.Event = &v
	return s
}

func (s *InteractTextResponse) SetBody(v *InteractTextResponseBody) *InteractTextResponse {
	s.Body = v
	return s
}

func (s *InteractTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
