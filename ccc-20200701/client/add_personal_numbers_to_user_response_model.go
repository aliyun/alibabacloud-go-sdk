// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPersonalNumbersToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPersonalNumbersToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPersonalNumbersToUserResponse
	GetStatusCode() *int32
	SetBody(v *AddPersonalNumbersToUserResponseBody) *AddPersonalNumbersToUserResponse
	GetBody() *AddPersonalNumbersToUserResponseBody
}

type AddPersonalNumbersToUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPersonalNumbersToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPersonalNumbersToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPersonalNumbersToUserResponse) GoString() string {
	return s.String()
}

func (s *AddPersonalNumbersToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPersonalNumbersToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPersonalNumbersToUserResponse) GetBody() *AddPersonalNumbersToUserResponseBody {
	return s.Body
}

func (s *AddPersonalNumbersToUserResponse) SetHeaders(v map[string]*string) *AddPersonalNumbersToUserResponse {
	s.Headers = v
	return s
}

func (s *AddPersonalNumbersToUserResponse) SetStatusCode(v int32) *AddPersonalNumbersToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPersonalNumbersToUserResponse) SetBody(v *AddPersonalNumbersToUserResponseBody) *AddPersonalNumbersToUserResponse {
	s.Body = v
	return s
}

func (s *AddPersonalNumbersToUserResponse) Validate() error {
	return dara.Validate(s)
}
