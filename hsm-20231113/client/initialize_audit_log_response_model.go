// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *InitializeAuditLogResponseBody) *InitializeAuditLogResponse
	GetBody() *InitializeAuditLogResponseBody
}

type InitializeAuditLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeAuditLogResponse) GoString() string {
	return s.String()
}

func (s *InitializeAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeAuditLogResponse) GetBody() *InitializeAuditLogResponseBody {
	return s.Body
}

func (s *InitializeAuditLogResponse) SetHeaders(v map[string]*string) *InitializeAuditLogResponse {
	s.Headers = v
	return s
}

func (s *InitializeAuditLogResponse) SetStatusCode(v int32) *InitializeAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeAuditLogResponse) SetBody(v *InitializeAuditLogResponseBody) *InitializeAuditLogResponse {
	s.Body = v
	return s
}

func (s *InitializeAuditLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
