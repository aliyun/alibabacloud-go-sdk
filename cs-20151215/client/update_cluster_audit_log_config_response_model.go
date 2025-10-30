// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAuditLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterAuditLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterAuditLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterAuditLogConfigResponseBody) *UpdateClusterAuditLogConfigResponse
	GetBody() *UpdateClusterAuditLogConfigResponseBody
}

type UpdateClusterAuditLogConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterAuditLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterAuditLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterAuditLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterAuditLogConfigResponse) GetBody() *UpdateClusterAuditLogConfigResponseBody {
	return s.Body
}

func (s *UpdateClusterAuditLogConfigResponse) SetHeaders(v map[string]*string) *UpdateClusterAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterAuditLogConfigResponse) SetStatusCode(v int32) *UpdateClusterAuditLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterAuditLogConfigResponse) SetBody(v *UpdateClusterAuditLogConfigResponseBody) *UpdateClusterAuditLogConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterAuditLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
