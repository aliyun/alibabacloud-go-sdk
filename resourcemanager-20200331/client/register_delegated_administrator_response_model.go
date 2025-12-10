// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDelegatedAdministratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterDelegatedAdministratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterDelegatedAdministratorResponse
	GetStatusCode() *int32
	SetBody(v *RegisterDelegatedAdministratorResponseBody) *RegisterDelegatedAdministratorResponse
	GetBody() *RegisterDelegatedAdministratorResponseBody
}

type RegisterDelegatedAdministratorResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDelegatedAdministratorResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *RegisterDelegatedAdministratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterDelegatedAdministratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterDelegatedAdministratorResponse) GetBody() *RegisterDelegatedAdministratorResponseBody {
	return s.Body
}

func (s *RegisterDelegatedAdministratorResponse) SetHeaders(v map[string]*string) *RegisterDelegatedAdministratorResponse {
	s.Headers = v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) SetStatusCode(v int32) *RegisterDelegatedAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) SetBody(v *RegisterDelegatedAdministratorResponseBody) *RegisterDelegatedAdministratorResponse {
	s.Body = v
	return s
}

func (s *RegisterDelegatedAdministratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
