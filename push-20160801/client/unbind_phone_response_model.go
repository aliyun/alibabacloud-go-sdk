// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindPhoneResponse
	GetStatusCode() *int32
	SetBody(v *UnbindPhoneResponseBody) *UnbindPhoneResponse
	GetBody() *UnbindPhoneResponseBody
}

type UnbindPhoneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindPhoneResponse) GoString() string {
	return s.String()
}

func (s *UnbindPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindPhoneResponse) GetBody() *UnbindPhoneResponseBody {
	return s.Body
}

func (s *UnbindPhoneResponse) SetHeaders(v map[string]*string) *UnbindPhoneResponse {
	s.Headers = v
	return s
}

func (s *UnbindPhoneResponse) SetStatusCode(v int32) *UnbindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPhoneResponse) SetBody(v *UnbindPhoneResponseBody) *UnbindPhoneResponse {
	s.Body = v
	return s
}

func (s *UnbindPhoneResponse) Validate() error {
	return dara.Validate(s)
}
