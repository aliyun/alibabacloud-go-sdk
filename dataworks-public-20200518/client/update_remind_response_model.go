// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRemindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRemindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRemindResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRemindResponseBody) *UpdateRemindResponse
	GetBody() *UpdateRemindResponseBody
}

type UpdateRemindResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRemindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRemindResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRemindResponse) GoString() string {
	return s.String()
}

func (s *UpdateRemindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRemindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRemindResponse) GetBody() *UpdateRemindResponseBody {
	return s.Body
}

func (s *UpdateRemindResponse) SetHeaders(v map[string]*string) *UpdateRemindResponse {
	s.Headers = v
	return s
}

func (s *UpdateRemindResponse) SetStatusCode(v int32) *UpdateRemindResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRemindResponse) SetBody(v *UpdateRemindResponseBody) *UpdateRemindResponse {
	s.Body = v
	return s
}

func (s *UpdateRemindResponse) Validate() error {
	return dara.Validate(s)
}
