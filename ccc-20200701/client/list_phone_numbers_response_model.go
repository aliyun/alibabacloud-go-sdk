// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPhoneNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPhoneNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ListPhoneNumbersResponseBody) *ListPhoneNumbersResponse
	GetBody() *ListPhoneNumbersResponseBody
}

type ListPhoneNumbersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPhoneNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPhoneNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPhoneNumbersResponse) GetBody() *ListPhoneNumbersResponseBody {
	return s.Body
}

func (s *ListPhoneNumbersResponse) SetHeaders(v map[string]*string) *ListPhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListPhoneNumbersResponse) SetStatusCode(v int32) *ListPhoneNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPhoneNumbersResponse) SetBody(v *ListPhoneNumbersResponseBody) *ListPhoneNumbersResponse {
	s.Body = v
	return s
}

func (s *ListPhoneNumbersResponse) Validate() error {
	return dara.Validate(s)
}
