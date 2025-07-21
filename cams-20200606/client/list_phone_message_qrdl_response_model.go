// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneMessageQrdlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPhoneMessageQrdlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPhoneMessageQrdlResponse
	GetStatusCode() *int32
	SetBody(v *ListPhoneMessageQrdlResponseBody) *ListPhoneMessageQrdlResponse
	GetBody() *ListPhoneMessageQrdlResponseBody
}

type ListPhoneMessageQrdlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPhoneMessageQrdlResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPhoneMessageQrdlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPhoneMessageQrdlResponse) GetBody() *ListPhoneMessageQrdlResponseBody {
	return s.Body
}

func (s *ListPhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *ListPhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *ListPhoneMessageQrdlResponse) SetStatusCode(v int32) *ListPhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPhoneMessageQrdlResponse) SetBody(v *ListPhoneMessageQrdlResponseBody) *ListPhoneMessageQrdlResponse {
	s.Body = v
	return s
}

func (s *ListPhoneMessageQrdlResponse) Validate() error {
	return dara.Validate(s)
}
