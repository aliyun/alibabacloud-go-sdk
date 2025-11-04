// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyUserResponse
	GetStatusCode() *int32
	SetBody(v *VerifyUserResponseBody) *VerifyUserResponse
	GetBody() *VerifyUserResponseBody
}

type VerifyUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyUserResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyUserResponse) GoString() string {
	return s.String()
}

func (s *VerifyUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyUserResponse) GetBody() *VerifyUserResponseBody {
	return s.Body
}

func (s *VerifyUserResponse) SetHeaders(v map[string]*string) *VerifyUserResponse {
	s.Headers = v
	return s
}

func (s *VerifyUserResponse) SetStatusCode(v int32) *VerifyUserResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyUserResponse) SetBody(v *VerifyUserResponseBody) *VerifyUserResponse {
	s.Body = v
	return s
}

func (s *VerifyUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
