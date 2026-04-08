// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeActionResponse
	GetStatusCode() *int32
	SetId(v string) *InvokeActionResponse
	GetId() *string
	SetEvent(v string) *InvokeActionResponse
	GetEvent() *string
	SetBody(v *InvokeActionResponseBody) *InvokeActionResponse
	GetBody() *InvokeActionResponseBody
}

type InvokeActionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                   `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                   `json:"event,omitempty" xml:"event,omitempty"`
	Body       *InvokeActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeActionResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeActionResponse) GoString() string {
	return s.String()
}

func (s *InvokeActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeActionResponse) GetId() *string {
	return s.Id
}

func (s *InvokeActionResponse) GetEvent() *string {
	return s.Event
}

func (s *InvokeActionResponse) GetBody() *InvokeActionResponseBody {
	return s.Body
}

func (s *InvokeActionResponse) SetHeaders(v map[string]*string) *InvokeActionResponse {
	s.Headers = v
	return s
}

func (s *InvokeActionResponse) SetStatusCode(v int32) *InvokeActionResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeActionResponse) SetId(v string) *InvokeActionResponse {
	s.Id = &v
	return s
}

func (s *InvokeActionResponse) SetEvent(v string) *InvokeActionResponse {
	s.Event = &v
	return s
}

func (s *InvokeActionResponse) SetBody(v *InvokeActionResponseBody) *InvokeActionResponse {
	s.Body = v
	return s
}

func (s *InvokeActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
