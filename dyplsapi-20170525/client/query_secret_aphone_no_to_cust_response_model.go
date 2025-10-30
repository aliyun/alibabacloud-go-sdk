// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretAPhoneNoToCustResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySecretAPhoneNoToCustResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySecretAPhoneNoToCustResponse
	GetStatusCode() *int32
	SetBody(v *QuerySecretAPhoneNoToCustResponseBody) *QuerySecretAPhoneNoToCustResponse
	GetBody() *QuerySecretAPhoneNoToCustResponseBody
}

type QuerySecretAPhoneNoToCustResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretAPhoneNoToCustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretAPhoneNoToCustResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySecretAPhoneNoToCustResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySecretAPhoneNoToCustResponse) GetBody() *QuerySecretAPhoneNoToCustResponseBody {
	return s.Body
}

func (s *QuerySecretAPhoneNoToCustResponse) SetHeaders(v map[string]*string) *QuerySecretAPhoneNoToCustResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponse) SetStatusCode(v int32) *QuerySecretAPhoneNoToCustResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponse) SetBody(v *QuerySecretAPhoneNoToCustResponseBody) *QuerySecretAPhoneNoToCustResponse {
	s.Body = v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
