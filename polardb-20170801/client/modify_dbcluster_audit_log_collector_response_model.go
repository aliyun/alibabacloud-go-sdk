// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAuditLogCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterAuditLogCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterAuditLogCollectorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterAuditLogCollectorResponseBody) *ModifyDBClusterAuditLogCollectorResponse
	GetBody() *ModifyDBClusterAuditLogCollectorResponseBody
}

type ModifyDBClusterAuditLogCollectorResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterAuditLogCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterAuditLogCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterAuditLogCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterAuditLogCollectorResponse) GetBody() *ModifyDBClusterAuditLogCollectorResponseBody {
	return s.Body
}

func (s *ModifyDBClusterAuditLogCollectorResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAuditLogCollectorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorResponse) SetStatusCode(v int32) *ModifyDBClusterAuditLogCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorResponse) SetBody(v *ModifyDBClusterAuditLogCollectorResponseBody) *ModifyDBClusterAuditLogCollectorResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorResponse) Validate() error {
	return dara.Validate(s)
}
