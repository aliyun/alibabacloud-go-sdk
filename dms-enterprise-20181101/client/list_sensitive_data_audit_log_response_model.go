// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveDataAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSensitiveDataAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSensitiveDataAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *ListSensitiveDataAuditLogResponseBody) *ListSensitiveDataAuditLogResponse
	GetBody() *ListSensitiveDataAuditLogResponseBody
}

type ListSensitiveDataAuditLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSensitiveDataAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSensitiveDataAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveDataAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveDataAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSensitiveDataAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSensitiveDataAuditLogResponse) GetBody() *ListSensitiveDataAuditLogResponseBody {
	return s.Body
}

func (s *ListSensitiveDataAuditLogResponse) SetHeaders(v map[string]*string) *ListSensitiveDataAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveDataAuditLogResponse) SetStatusCode(v int32) *ListSensitiveDataAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitiveDataAuditLogResponse) SetBody(v *ListSensitiveDataAuditLogResponseBody) *ListSensitiveDataAuditLogResponse {
	s.Body = v
	return s
}

func (s *ListSensitiveDataAuditLogResponse) Validate() error {
	return dara.Validate(s)
}
