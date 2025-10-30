// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretAPhoneNoToCustResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecretAPhoneNoToCustResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecretAPhoneNoToCustResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecretAPhoneNoToCustResponseBody) *DeleteSecretAPhoneNoToCustResponse
	GetBody() *DeleteSecretAPhoneNoToCustResponseBody
}

type DeleteSecretAPhoneNoToCustResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretAPhoneNoToCustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretAPhoneNoToCustResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretAPhoneNoToCustResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretAPhoneNoToCustResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecretAPhoneNoToCustResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecretAPhoneNoToCustResponse) GetBody() *DeleteSecretAPhoneNoToCustResponseBody {
	return s.Body
}

func (s *DeleteSecretAPhoneNoToCustResponse) SetHeaders(v map[string]*string) *DeleteSecretAPhoneNoToCustResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponse) SetStatusCode(v int32) *DeleteSecretAPhoneNoToCustResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponse) SetBody(v *DeleteSecretAPhoneNoToCustResponseBody) *DeleteSecretAPhoneNoToCustResponse {
	s.Body = v
	return s
}

func (s *DeleteSecretAPhoneNoToCustResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
