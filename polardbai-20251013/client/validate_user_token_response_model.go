// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateUserTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateUserTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateUserTokenResponse
	GetStatusCode() *int32
	SetBody(v *ValidateUserTokenResponseBody) *ValidateUserTokenResponse
	GetBody() *ValidateUserTokenResponseBody
}

type ValidateUserTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateUserTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateUserTokenResponse) GoString() string {
	return s.String()
}

func (s *ValidateUserTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateUserTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateUserTokenResponse) GetBody() *ValidateUserTokenResponseBody {
	return s.Body
}

func (s *ValidateUserTokenResponse) SetHeaders(v map[string]*string) *ValidateUserTokenResponse {
	s.Headers = v
	return s
}

func (s *ValidateUserTokenResponse) SetStatusCode(v int32) *ValidateUserTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateUserTokenResponse) SetBody(v *ValidateUserTokenResponseBody) *ValidateUserTokenResponse {
	s.Body = v
	return s
}

func (s *ValidateUserTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
