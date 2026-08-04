// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyIdentityRegistrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyIdentityRegistrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyIdentityRegistrationResponse
	GetStatusCode() *int32
	SetBody(v *ApplyIdentityRegistrationResponseBody) *ApplyIdentityRegistrationResponse
	GetBody() *ApplyIdentityRegistrationResponseBody
}

type ApplyIdentityRegistrationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyIdentityRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyIdentityRegistrationResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyIdentityRegistrationResponse) GoString() string {
	return s.String()
}

func (s *ApplyIdentityRegistrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyIdentityRegistrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyIdentityRegistrationResponse) GetBody() *ApplyIdentityRegistrationResponseBody {
	return s.Body
}

func (s *ApplyIdentityRegistrationResponse) SetHeaders(v map[string]*string) *ApplyIdentityRegistrationResponse {
	s.Headers = v
	return s
}

func (s *ApplyIdentityRegistrationResponse) SetStatusCode(v int32) *ApplyIdentityRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyIdentityRegistrationResponse) SetBody(v *ApplyIdentityRegistrationResponseBody) *ApplyIdentityRegistrationResponse {
	s.Body = v
	return s
}

func (s *ApplyIdentityRegistrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
