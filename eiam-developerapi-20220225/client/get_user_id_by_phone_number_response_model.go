// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdByPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdByPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdByPhoneNumberResponseBody) *GetUserIdByPhoneNumberResponse
	GetBody() *GetUserIdByPhoneNumberResponseBody
}

type GetUserIdByPhoneNumberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdByPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdByPhoneNumberResponse) GetBody() *GetUserIdByPhoneNumberResponseBody {
	return s.Body
}

func (s *GetUserIdByPhoneNumberResponse) SetHeaders(v map[string]*string) *GetUserIdByPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) SetStatusCode(v int32) *GetUserIdByPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) SetBody(v *GetUserIdByPhoneNumberResponseBody) *GetUserIdByPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *GetUserIdByPhoneNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
