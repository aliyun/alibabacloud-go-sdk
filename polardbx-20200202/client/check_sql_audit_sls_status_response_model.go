// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlAuditSlsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSqlAuditSlsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSqlAuditSlsStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckSqlAuditSlsStatusResponseBody) *CheckSqlAuditSlsStatusResponse
	GetBody() *CheckSqlAuditSlsStatusResponseBody
}

type CheckSqlAuditSlsStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSqlAuditSlsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSqlAuditSlsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlAuditSlsStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditSlsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSqlAuditSlsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSqlAuditSlsStatusResponse) GetBody() *CheckSqlAuditSlsStatusResponseBody {
	return s.Body
}

func (s *CheckSqlAuditSlsStatusResponse) SetHeaders(v map[string]*string) *CheckSqlAuditSlsStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSqlAuditSlsStatusResponse) SetStatusCode(v int32) *CheckSqlAuditSlsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSqlAuditSlsStatusResponse) SetBody(v *CheckSqlAuditSlsStatusResponseBody) *CheckSqlAuditSlsStatusResponse {
	s.Body = v
	return s
}

func (s *CheckSqlAuditSlsStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
