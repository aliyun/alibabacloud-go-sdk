// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneMessageQrdlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePhoneMessageQrdlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePhoneMessageQrdlResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePhoneMessageQrdlResponseBody) *UpdatePhoneMessageQrdlResponse
	GetBody() *UpdatePhoneMessageQrdlResponseBody
}

type UpdatePhoneMessageQrdlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePhoneMessageQrdlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePhoneMessageQrdlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePhoneMessageQrdlResponse) GetBody() *UpdatePhoneMessageQrdlResponseBody {
	return s.Body
}

func (s *UpdatePhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *UpdatePhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneMessageQrdlResponse) SetStatusCode(v int32) *UpdatePhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponse) SetBody(v *UpdatePhoneMessageQrdlResponseBody) *UpdatePhoneMessageQrdlResponse {
	s.Body = v
	return s
}

func (s *UpdatePhoneMessageQrdlResponse) Validate() error {
	return dara.Validate(s)
}
