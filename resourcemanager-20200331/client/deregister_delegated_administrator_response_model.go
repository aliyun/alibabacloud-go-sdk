// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterDelegatedAdministratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeregisterDelegatedAdministratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeregisterDelegatedAdministratorResponse
	GetStatusCode() *int32
	SetBody(v *DeregisterDelegatedAdministratorResponseBody) *DeregisterDelegatedAdministratorResponse
	GetBody() *DeregisterDelegatedAdministratorResponseBody
}

type DeregisterDelegatedAdministratorResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterDelegatedAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterDelegatedAdministratorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeregisterDelegatedAdministratorResponse) GoString() string {
	return s.String()
}

func (s *DeregisterDelegatedAdministratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeregisterDelegatedAdministratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeregisterDelegatedAdministratorResponse) GetBody() *DeregisterDelegatedAdministratorResponseBody {
	return s.Body
}

func (s *DeregisterDelegatedAdministratorResponse) SetHeaders(v map[string]*string) *DeregisterDelegatedAdministratorResponse {
	s.Headers = v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) SetStatusCode(v int32) *DeregisterDelegatedAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) SetBody(v *DeregisterDelegatedAdministratorResponseBody) *DeregisterDelegatedAdministratorResponse {
	s.Body = v
	return s
}

func (s *DeregisterDelegatedAdministratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
