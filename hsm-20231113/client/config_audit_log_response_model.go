// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *ConfigAuditLogResponseBody) *ConfigAuditLogResponse
	GetBody() *ConfigAuditLogResponseBody
}

type ConfigAuditLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ConfigAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigAuditLogResponse) GetBody() *ConfigAuditLogResponseBody {
	return s.Body
}

func (s *ConfigAuditLogResponse) SetHeaders(v map[string]*string) *ConfigAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ConfigAuditLogResponse) SetStatusCode(v int32) *ConfigAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigAuditLogResponse) SetBody(v *ConfigAuditLogResponseBody) *ConfigAuditLogResponse {
	s.Body = v
	return s
}

func (s *ConfigAuditLogResponse) Validate() error {
	return dara.Validate(s)
}
