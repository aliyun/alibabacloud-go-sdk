// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuditLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuditLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAuditLogStatusResponseBody) *GetAuditLogStatusResponse
	GetBody() *GetAuditLogStatusResponseBody
}

type GetAuditLogStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuditLogStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAuditLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuditLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuditLogStatusResponse) GetBody() *GetAuditLogStatusResponseBody {
	return s.Body
}

func (s *GetAuditLogStatusResponse) SetHeaders(v map[string]*string) *GetAuditLogStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAuditLogStatusResponse) SetStatusCode(v int32) *GetAuditLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditLogStatusResponse) SetBody(v *GetAuditLogStatusResponseBody) *GetAuditLogStatusResponse {
	s.Body = v
	return s
}

func (s *GetAuditLogStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
