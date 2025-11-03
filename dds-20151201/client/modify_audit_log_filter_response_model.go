// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAuditLogFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAuditLogFilterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAuditLogFilterResponseBody) *ModifyAuditLogFilterResponse
	GetBody() *ModifyAuditLogFilterResponseBody
}

type ModifyAuditLogFilterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAuditLogFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAuditLogFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogFilterResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAuditLogFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAuditLogFilterResponse) GetBody() *ModifyAuditLogFilterResponseBody {
	return s.Body
}

func (s *ModifyAuditLogFilterResponse) SetHeaders(v map[string]*string) *ModifyAuditLogFilterResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditLogFilterResponse) SetStatusCode(v int32) *ModifyAuditLogFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAuditLogFilterResponse) SetBody(v *ModifyAuditLogFilterResponseBody) *ModifyAuditLogFilterResponse {
	s.Body = v
	return s
}

func (s *ModifyAuditLogFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
