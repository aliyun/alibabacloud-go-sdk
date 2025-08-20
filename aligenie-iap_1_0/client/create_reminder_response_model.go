// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReminderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReminderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReminderResponse
	GetStatusCode() *int32
	SetBody(v *CreateReminderResponseBody) *CreateReminderResponse
	GetBody() *CreateReminderResponseBody
}

type CreateReminderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReminderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderResponse) GoString() string {
	return s.String()
}

func (s *CreateReminderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReminderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReminderResponse) GetBody() *CreateReminderResponseBody {
	return s.Body
}

func (s *CreateReminderResponse) SetHeaders(v map[string]*string) *CreateReminderResponse {
	s.Headers = v
	return s
}

func (s *CreateReminderResponse) SetStatusCode(v int32) *CreateReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReminderResponse) SetBody(v *CreateReminderResponseBody) *CreateReminderResponse {
	s.Body = v
	return s
}

func (s *CreateReminderResponse) Validate() error {
	return dara.Validate(s)
}
