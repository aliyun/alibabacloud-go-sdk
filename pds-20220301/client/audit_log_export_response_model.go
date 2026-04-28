// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditLogExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuditLogExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuditLogExportResponse
	GetStatusCode() *int32
	SetBody(v *AuditLogExportResponseBody) *AuditLogExportResponse
	GetBody() *AuditLogExportResponseBody
}

type AuditLogExportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditLogExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditLogExportResponse) String() string {
	return dara.Prettify(s)
}

func (s AuditLogExportResponse) GoString() string {
	return s.String()
}

func (s *AuditLogExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuditLogExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuditLogExportResponse) GetBody() *AuditLogExportResponseBody {
	return s.Body
}

func (s *AuditLogExportResponse) SetHeaders(v map[string]*string) *AuditLogExportResponse {
	s.Headers = v
	return s
}

func (s *AuditLogExportResponse) SetStatusCode(v int32) *AuditLogExportResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditLogExportResponse) SetBody(v *AuditLogExportResponseBody) *AuditLogExportResponse {
	s.Body = v
	return s
}

func (s *AuditLogExportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
