// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlAuditEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSqlAuditEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSqlAuditEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckSqlAuditEnableStatusResponseBody) *CheckSqlAuditEnableStatusResponse
	GetBody() *CheckSqlAuditEnableStatusResponseBody
}

type CheckSqlAuditEnableStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSqlAuditEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSqlAuditEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSqlAuditEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSqlAuditEnableStatusResponse) GetBody() *CheckSqlAuditEnableStatusResponseBody {
	return s.Body
}

func (s *CheckSqlAuditEnableStatusResponse) SetHeaders(v map[string]*string) *CheckSqlAuditEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) SetStatusCode(v int32) *CheckSqlAuditEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) SetBody(v *CheckSqlAuditEnableStatusResponseBody) *CheckSqlAuditEnableStatusResponse {
	s.Body = v
	return s
}

func (s *CheckSqlAuditEnableStatusResponse) Validate() error {
	return dara.Validate(s)
}
