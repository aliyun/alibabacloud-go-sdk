// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuditRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuditRequestResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuditRequestResponseBody) *UpdateAuditRequestResponse
	GetBody() *UpdateAuditRequestResponseBody
}

type UpdateAuditRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuditRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuditRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuditRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuditRequestResponse) GetBody() *UpdateAuditRequestResponseBody {
	return s.Body
}

func (s *UpdateAuditRequestResponse) SetHeaders(v map[string]*string) *UpdateAuditRequestResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuditRequestResponse) SetStatusCode(v int32) *UpdateAuditRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuditRequestResponse) SetBody(v *UpdateAuditRequestResponseBody) *UpdateAuditRequestResponse {
	s.Body = v
	return s
}

func (s *UpdateAuditRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
