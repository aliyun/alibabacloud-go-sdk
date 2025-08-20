// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReminderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReminderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReminderResponse
	GetStatusCode() *int32
	SetBody(v *GetReminderResponseBody) *GetReminderResponse
	GetBody() *GetReminderResponseBody
}

type GetReminderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReminderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReminderResponse) GoString() string {
	return s.String()
}

func (s *GetReminderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReminderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReminderResponse) GetBody() *GetReminderResponseBody {
	return s.Body
}

func (s *GetReminderResponse) SetHeaders(v map[string]*string) *GetReminderResponse {
	s.Headers = v
	return s
}

func (s *GetReminderResponse) SetStatusCode(v int32) *GetReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReminderResponse) SetBody(v *GetReminderResponseBody) *GetReminderResponse {
	s.Body = v
	return s
}

func (s *GetReminderResponse) Validate() error {
	return dara.Validate(s)
}
