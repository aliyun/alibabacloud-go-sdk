// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNapalStreamMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendNapalStreamMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendNapalStreamMessageResponse
	GetStatusCode() *int32
	SetId(v string) *SendNapalStreamMessageResponse
	GetId() *string
	SetEvent(v string) *SendNapalStreamMessageResponse
	GetEvent() *string
	SetBody(v *SendNapalStreamMessageResponseBody) *SendNapalStreamMessageResponse
	GetBody() *SendNapalStreamMessageResponseBody
}

type SendNapalStreamMessageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *SendNapalStreamMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendNapalStreamMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponse) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendNapalStreamMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendNapalStreamMessageResponse) GetId() *string {
	return s.Id
}

func (s *SendNapalStreamMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *SendNapalStreamMessageResponse) GetBody() *SendNapalStreamMessageResponseBody {
	return s.Body
}

func (s *SendNapalStreamMessageResponse) SetHeaders(v map[string]*string) *SendNapalStreamMessageResponse {
	s.Headers = v
	return s
}

func (s *SendNapalStreamMessageResponse) SetStatusCode(v int32) *SendNapalStreamMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendNapalStreamMessageResponse) SetId(v string) *SendNapalStreamMessageResponse {
	s.Id = &v
	return s
}

func (s *SendNapalStreamMessageResponse) SetEvent(v string) *SendNapalStreamMessageResponse {
	s.Event = &v
	return s
}

func (s *SendNapalStreamMessageResponse) SetBody(v *SendNapalStreamMessageResponseBody) *SendNapalStreamMessageResponse {
	s.Body = v
	return s
}

func (s *SendNapalStreamMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
