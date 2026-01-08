// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreValidatePhoneIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPreValidatePhoneIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPreValidatePhoneIdResponse
	GetStatusCode() *int32
	SetBody(v *GetPreValidatePhoneIdResponseBody) *GetPreValidatePhoneIdResponse
	GetBody() *GetPreValidatePhoneIdResponseBody
}

type GetPreValidatePhoneIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPreValidatePhoneIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPreValidatePhoneIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPreValidatePhoneIdResponse) GoString() string {
	return s.String()
}

func (s *GetPreValidatePhoneIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPreValidatePhoneIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPreValidatePhoneIdResponse) GetBody() *GetPreValidatePhoneIdResponseBody {
	return s.Body
}

func (s *GetPreValidatePhoneIdResponse) SetHeaders(v map[string]*string) *GetPreValidatePhoneIdResponse {
	s.Headers = v
	return s
}

func (s *GetPreValidatePhoneIdResponse) SetStatusCode(v int32) *GetPreValidatePhoneIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPreValidatePhoneIdResponse) SetBody(v *GetPreValidatePhoneIdResponseBody) *GetPreValidatePhoneIdResponse {
	s.Body = v
	return s
}

func (s *GetPreValidatePhoneIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
