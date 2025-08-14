// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuditResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuditResponseBody) *UpdateAuditResponse
	GetBody() *UpdateAuditResponseBody
}

type UpdateAuditResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuditResponse) GetBody() *UpdateAuditResponseBody {
	return s.Body
}

func (s *UpdateAuditResponse) SetHeaders(v map[string]*string) *UpdateAuditResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuditResponse) SetStatusCode(v int32) *UpdateAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuditResponse) SetBody(v *UpdateAuditResponseBody) *UpdateAuditResponse {
	s.Body = v
	return s
}

func (s *UpdateAuditResponse) Validate() error {
	return dara.Validate(s)
}
