// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iA2aResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *A2aResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *A2aResponse
	GetStatusCode() *int32
	SetId(v string) *A2aResponse
	GetId() *string
	SetEvent(v string) *A2aResponse
	GetEvent() *string
	SetBody(v interface{}) *A2aResponse
	GetBody() interface{}
}

type A2aResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s A2aResponse) String() string {
	return dara.Prettify(s)
}

func (s A2aResponse) GoString() string {
	return s.String()
}

func (s *A2aResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *A2aResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *A2aResponse) GetId() *string {
	return s.Id
}

func (s *A2aResponse) GetEvent() *string {
	return s.Event
}

func (s *A2aResponse) GetBody() interface{} {
	return s.Body
}

func (s *A2aResponse) SetHeaders(v map[string]*string) *A2aResponse {
	s.Headers = v
	return s
}

func (s *A2aResponse) SetStatusCode(v int32) *A2aResponse {
	s.StatusCode = &v
	return s
}

func (s *A2aResponse) SetId(v string) *A2aResponse {
	s.Id = &v
	return s
}

func (s *A2aResponse) SetEvent(v string) *A2aResponse {
	s.Event = &v
	return s
}

func (s *A2aResponse) SetBody(v interface{}) *A2aResponse {
	s.Body = v
	return s
}

func (s *A2aResponse) Validate() error {
	return dara.Validate(s)
}
