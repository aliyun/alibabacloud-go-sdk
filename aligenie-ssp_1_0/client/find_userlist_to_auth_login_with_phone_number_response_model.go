// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindUserlistToAuthLoginWithPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindUserlistToAuthLoginWithPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindUserlistToAuthLoginWithPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *FindUserlistToAuthLoginWithPhoneNumberResponseBody) *FindUserlistToAuthLoginWithPhoneNumberResponse
	GetBody() *FindUserlistToAuthLoginWithPhoneNumberResponseBody
}

type FindUserlistToAuthLoginWithPhoneNumberResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindUserlistToAuthLoginWithPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) GetBody() *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	return s.Body
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) SetHeaders(v map[string]*string) *FindUserlistToAuthLoginWithPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) SetStatusCode(v int32) *FindUserlistToAuthLoginWithPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) SetBody(v *FindUserlistToAuthLoginWithPhoneNumberResponseBody) *FindUserlistToAuthLoginWithPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
