// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmResponse
	GetStatusCode() *int32
	SetId(v string) *ConfirmResponse
	GetId() *string
	SetEvent(v string) *ConfirmResponse
	GetEvent() *string
	SetBody(v string) *ConfirmResponse
	GetBody() *string
}

type ConfirmResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmResponse) GoString() string {
	return s.String()
}

func (s *ConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmResponse) GetId() *string {
	return s.Id
}

func (s *ConfirmResponse) GetEvent() *string {
	return s.Event
}

func (s *ConfirmResponse) GetBody() *string {
	return s.Body
}

func (s *ConfirmResponse) SetHeaders(v map[string]*string) *ConfirmResponse {
	s.Headers = v
	return s
}

func (s *ConfirmResponse) SetStatusCode(v int32) *ConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmResponse) SetId(v string) *ConfirmResponse {
	s.Id = &v
	return s
}

func (s *ConfirmResponse) SetEvent(v string) *ConfirmResponse {
	s.Event = &v
	return s
}

func (s *ConfirmResponse) SetBody(v string) *ConfirmResponse {
	s.Body = &v
	return s
}

func (s *ConfirmResponse) Validate() error {
	return dara.Validate(s)
}
