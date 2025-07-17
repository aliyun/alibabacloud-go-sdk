// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxySQLExecAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProxySQLExecAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProxySQLExecAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *ListProxySQLExecAuditLogResponseBody) *ListProxySQLExecAuditLogResponse
	GetBody() *ListProxySQLExecAuditLogResponseBody
}

type ListProxySQLExecAuditLogResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProxySQLExecAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProxySQLExecAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProxySQLExecAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ListProxySQLExecAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProxySQLExecAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProxySQLExecAuditLogResponse) GetBody() *ListProxySQLExecAuditLogResponseBody {
	return s.Body
}

func (s *ListProxySQLExecAuditLogResponse) SetHeaders(v map[string]*string) *ListProxySQLExecAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ListProxySQLExecAuditLogResponse) SetStatusCode(v int32) *ListProxySQLExecAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponse) SetBody(v *ListProxySQLExecAuditLogResponseBody) *ListProxySQLExecAuditLogResponse {
	s.Body = v
	return s
}

func (s *ListProxySQLExecAuditLogResponse) Validate() error {
	return dara.Validate(s)
}
