// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReminderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateReminderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateReminderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateReminderResponseBody) *UpdateReminderResponse
	GetBody() *UpdateReminderResponseBody
}

type UpdateReminderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateReminderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderResponse) GoString() string {
	return s.String()
}

func (s *UpdateReminderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateReminderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateReminderResponse) GetBody() *UpdateReminderResponseBody {
	return s.Body
}

func (s *UpdateReminderResponse) SetHeaders(v map[string]*string) *UpdateReminderResponse {
	s.Headers = v
	return s
}

func (s *UpdateReminderResponse) SetStatusCode(v int32) *UpdateReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReminderResponse) SetBody(v *UpdateReminderResponseBody) *UpdateReminderResponse {
	s.Body = v
	return s
}

func (s *UpdateReminderResponse) Validate() error {
	return dara.Validate(s)
}
