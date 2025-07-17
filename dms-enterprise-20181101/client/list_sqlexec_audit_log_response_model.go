// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLExecAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSQLExecAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSQLExecAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *ListSQLExecAuditLogResponseBody) *ListSQLExecAuditLogResponse
	GetBody() *ListSQLExecAuditLogResponseBody
}

type ListSQLExecAuditLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSQLExecAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSQLExecAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSQLExecAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSQLExecAuditLogResponse) GetBody() *ListSQLExecAuditLogResponseBody {
	return s.Body
}

func (s *ListSQLExecAuditLogResponse) SetHeaders(v map[string]*string) *ListSQLExecAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ListSQLExecAuditLogResponse) SetStatusCode(v int32) *ListSQLExecAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSQLExecAuditLogResponse) SetBody(v *ListSQLExecAuditLogResponseBody) *ListSQLExecAuditLogResponse {
	s.Body = v
	return s
}

func (s *ListSQLExecAuditLogResponse) Validate() error {
	return dara.Validate(s)
}
