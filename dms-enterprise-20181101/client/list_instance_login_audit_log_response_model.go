// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceLoginAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceLoginAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceLoginAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceLoginAuditLogResponseBody) *ListInstanceLoginAuditLogResponse
	GetBody() *ListInstanceLoginAuditLogResponseBody
}

type ListInstanceLoginAuditLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceLoginAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceLoginAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLoginAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceLoginAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceLoginAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceLoginAuditLogResponse) GetBody() *ListInstanceLoginAuditLogResponseBody {
	return s.Body
}

func (s *ListInstanceLoginAuditLogResponse) SetHeaders(v map[string]*string) *ListInstanceLoginAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceLoginAuditLogResponse) SetStatusCode(v int32) *ListInstanceLoginAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceLoginAuditLogResponse) SetBody(v *ListInstanceLoginAuditLogResponseBody) *ListInstanceLoginAuditLogResponse {
	s.Body = v
	return s
}

func (s *ListInstanceLoginAuditLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
