// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuditLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAuditLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAuditLogResponse
	GetStatusCode() *int32
	SetBody(v *QueryAuditLogResponseBody) *QueryAuditLogResponse
	GetBody() *QueryAuditLogResponseBody
}

type QueryAuditLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuditLogResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditLogResponse) GoString() string {
	return s.String()
}

func (s *QueryAuditLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAuditLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAuditLogResponse) GetBody() *QueryAuditLogResponseBody {
	return s.Body
}

func (s *QueryAuditLogResponse) SetHeaders(v map[string]*string) *QueryAuditLogResponse {
	s.Headers = v
	return s
}

func (s *QueryAuditLogResponse) SetStatusCode(v int32) *QueryAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuditLogResponse) SetBody(v *QueryAuditLogResponseBody) *QueryAuditLogResponse {
	s.Body = v
	return s
}

func (s *QueryAuditLogResponse) Validate() error {
	return dara.Validate(s)
}
