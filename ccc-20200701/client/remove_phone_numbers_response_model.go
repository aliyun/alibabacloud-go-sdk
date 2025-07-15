// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePhoneNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePhoneNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePhoneNumbersResponse
	GetStatusCode() *int32
	SetBody(v *RemovePhoneNumbersResponseBody) *RemovePhoneNumbersResponse
	GetBody() *RemovePhoneNumbersResponseBody
}

type RemovePhoneNumbersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePhoneNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePhoneNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePhoneNumbersResponse) GetBody() *RemovePhoneNumbersResponseBody {
	return s.Body
}

func (s *RemovePhoneNumbersResponse) SetHeaders(v map[string]*string) *RemovePhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *RemovePhoneNumbersResponse) SetStatusCode(v int32) *RemovePhoneNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePhoneNumbersResponse) SetBody(v *RemovePhoneNumbersResponseBody) *RemovePhoneNumbersResponse {
	s.Body = v
	return s
}

func (s *RemovePhoneNumbersResponse) Validate() error {
	return dara.Validate(s)
}
