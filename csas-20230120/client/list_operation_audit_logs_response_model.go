// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationAuditLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationAuditLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationAuditLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationAuditLogsResponseBody) *ListOperationAuditLogsResponse
	GetBody() *ListOperationAuditLogsResponseBody
}

type ListOperationAuditLogsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationAuditLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationAuditLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationAuditLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationAuditLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationAuditLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationAuditLogsResponse) GetBody() *ListOperationAuditLogsResponseBody {
	return s.Body
}

func (s *ListOperationAuditLogsResponse) SetHeaders(v map[string]*string) *ListOperationAuditLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationAuditLogsResponse) SetStatusCode(v int32) *ListOperationAuditLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationAuditLogsResponse) SetBody(v *ListOperationAuditLogsResponseBody) *ListOperationAuditLogsResponse {
	s.Body = v
	return s
}

func (s *ListOperationAuditLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
