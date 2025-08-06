// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditSecurityIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuditSecurityIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuditSecurityIpResponse
	GetStatusCode() *int32
	SetBody(v *ListAuditSecurityIpResponseBody) *ListAuditSecurityIpResponse
	GetBody() *ListAuditSecurityIpResponseBody
}

type ListAuditSecurityIpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuditSecurityIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuditSecurityIpResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuditSecurityIpResponse) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuditSecurityIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuditSecurityIpResponse) GetBody() *ListAuditSecurityIpResponseBody {
	return s.Body
}

func (s *ListAuditSecurityIpResponse) SetHeaders(v map[string]*string) *ListAuditSecurityIpResponse {
	s.Headers = v
	return s
}

func (s *ListAuditSecurityIpResponse) SetStatusCode(v int32) *ListAuditSecurityIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuditSecurityIpResponse) SetBody(v *ListAuditSecurityIpResponseBody) *ListAuditSecurityIpResponse {
	s.Body = v
	return s
}

func (s *ListAuditSecurityIpResponse) Validate() error {
	return dara.Validate(s)
}
