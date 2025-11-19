// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAuditSecurityIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAuditSecurityIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAuditSecurityIpResponse
	GetStatusCode() *int32
	SetBody(v *SetAuditSecurityIpResponseBody) *SetAuditSecurityIpResponse
	GetBody() *SetAuditSecurityIpResponseBody
}

type SetAuditSecurityIpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAuditSecurityIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAuditSecurityIpResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAuditSecurityIpResponse) GoString() string {
	return s.String()
}

func (s *SetAuditSecurityIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAuditSecurityIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAuditSecurityIpResponse) GetBody() *SetAuditSecurityIpResponseBody {
	return s.Body
}

func (s *SetAuditSecurityIpResponse) SetHeaders(v map[string]*string) *SetAuditSecurityIpResponse {
	s.Headers = v
	return s
}

func (s *SetAuditSecurityIpResponse) SetStatusCode(v int32) *SetAuditSecurityIpResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAuditSecurityIpResponse) SetBody(v *SetAuditSecurityIpResponseBody) *SetAuditSecurityIpResponse {
	s.Body = v
	return s
}

func (s *SetAuditSecurityIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
