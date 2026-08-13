// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetPasswordResponseBody) *ResetPasswordResponse
	GetBody() *ResetPasswordResponseBody
}

type ResetPasswordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetPasswordResponse) GetBody() *ResetPasswordResponseBody {
	return s.Body
}

func (s *ResetPasswordResponse) SetHeaders(v map[string]*string) *ResetPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetPasswordResponse) SetStatusCode(v int32) *ResetPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetPasswordResponse) SetBody(v *ResetPasswordResponseBody) *ResetPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
