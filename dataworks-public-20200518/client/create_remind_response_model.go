// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRemindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRemindResponse
	GetStatusCode() *int32
	SetBody(v *CreateRemindResponseBody) *CreateRemindResponse
	GetBody() *CreateRemindResponseBody
}

type CreateRemindResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRemindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRemindResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRemindResponse) GoString() string {
	return s.String()
}

func (s *CreateRemindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRemindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRemindResponse) GetBody() *CreateRemindResponseBody {
	return s.Body
}

func (s *CreateRemindResponse) SetHeaders(v map[string]*string) *CreateRemindResponse {
	s.Headers = v
	return s
}

func (s *CreateRemindResponse) SetStatusCode(v int32) *CreateRemindResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRemindResponse) SetBody(v *CreateRemindResponseBody) *CreateRemindResponse {
	s.Body = v
	return s
}

func (s *CreateRemindResponse) Validate() error {
	return dara.Validate(s)
}
