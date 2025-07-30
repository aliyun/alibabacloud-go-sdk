// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAuditLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAuditLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAuditLogConfigResponseBody) *ModifyAuditLogConfigResponse
	GetBody() *ModifyAuditLogConfigResponseBody
}

type ModifyAuditLogConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAuditLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAuditLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAuditLogConfigResponse) GetBody() *ModifyAuditLogConfigResponseBody {
	return s.Body
}

func (s *ModifyAuditLogConfigResponse) SetHeaders(v map[string]*string) *ModifyAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetStatusCode(v int32) *ModifyAuditLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetBody(v *ModifyAuditLogConfigResponseBody) *ModifyAuditLogConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAuditLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
