// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalNumbersOfUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPersonalNumbersOfUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPersonalNumbersOfUserResponse
	GetStatusCode() *int32
	SetBody(v *ListPersonalNumbersOfUserResponseBody) *ListPersonalNumbersOfUserResponse
	GetBody() *ListPersonalNumbersOfUserResponseBody
}

type ListPersonalNumbersOfUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPersonalNumbersOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPersonalNumbersOfUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPersonalNumbersOfUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPersonalNumbersOfUserResponse) GetBody() *ListPersonalNumbersOfUserResponseBody {
	return s.Body
}

func (s *ListPersonalNumbersOfUserResponse) SetHeaders(v map[string]*string) *ListPersonalNumbersOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListPersonalNumbersOfUserResponse) SetStatusCode(v int32) *ListPersonalNumbersOfUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponse) SetBody(v *ListPersonalNumbersOfUserResponseBody) *ListPersonalNumbersOfUserResponse {
	s.Body = v
	return s
}

func (s *ListPersonalNumbersOfUserResponse) Validate() error {
	return dara.Validate(s)
}
