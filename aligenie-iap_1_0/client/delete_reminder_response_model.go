// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReminderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReminderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReminderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReminderResponseBody) *DeleteReminderResponse
	GetBody() *DeleteReminderResponseBody
}

type DeleteReminderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReminderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderResponse) GoString() string {
	return s.String()
}

func (s *DeleteReminderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReminderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReminderResponse) GetBody() *DeleteReminderResponseBody {
	return s.Body
}

func (s *DeleteReminderResponse) SetHeaders(v map[string]*string) *DeleteReminderResponse {
	s.Headers = v
	return s
}

func (s *DeleteReminderResponse) SetStatusCode(v int32) *DeleteReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReminderResponse) SetBody(v *DeleteReminderResponseBody) *DeleteReminderResponse {
	s.Body = v
	return s
}

func (s *DeleteReminderResponse) Validate() error {
	return dara.Validate(s)
}
