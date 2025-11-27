// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateSecurityIPModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateSecurityIPModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateSecurityIPModeResponse
	GetStatusCode() *int32
	SetBody(v *MigrateSecurityIPModeResponseBody) *MigrateSecurityIPModeResponse
	GetBody() *MigrateSecurityIPModeResponseBody
}

type MigrateSecurityIPModeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateSecurityIPModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateSecurityIPModeResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateSecurityIPModeResponse) GoString() string {
	return s.String()
}

func (s *MigrateSecurityIPModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateSecurityIPModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateSecurityIPModeResponse) GetBody() *MigrateSecurityIPModeResponseBody {
	return s.Body
}

func (s *MigrateSecurityIPModeResponse) SetHeaders(v map[string]*string) *MigrateSecurityIPModeResponse {
	s.Headers = v
	return s
}

func (s *MigrateSecurityIPModeResponse) SetStatusCode(v int32) *MigrateSecurityIPModeResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateSecurityIPModeResponse) SetBody(v *MigrateSecurityIPModeResponseBody) *MigrateSecurityIPModeResponse {
	s.Body = v
	return s
}

func (s *MigrateSecurityIPModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
