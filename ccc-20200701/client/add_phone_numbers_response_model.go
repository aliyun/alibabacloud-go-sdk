// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPhoneNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPhoneNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPhoneNumbersResponse
	GetStatusCode() *int32
	SetBody(v *AddPhoneNumbersResponseBody) *AddPhoneNumbersResponse
	GetBody() *AddPhoneNumbersResponseBody
}

type AddPhoneNumbersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPhoneNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *AddPhoneNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPhoneNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPhoneNumbersResponse) GetBody() *AddPhoneNumbersResponseBody {
	return s.Body
}

func (s *AddPhoneNumbersResponse) SetHeaders(v map[string]*string) *AddPhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *AddPhoneNumbersResponse) SetStatusCode(v int32) *AddPhoneNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPhoneNumbersResponse) SetBody(v *AddPhoneNumbersResponseBody) *AddPhoneNumbersResponse {
	s.Body = v
	return s
}

func (s *AddPhoneNumbersResponse) Validate() error {
	return dara.Validate(s)
}
