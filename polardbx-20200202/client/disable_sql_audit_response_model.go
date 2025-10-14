// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSqlAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSqlAuditResponse
	GetStatusCode() *int32
	SetBody(v *DisableSqlAuditResponseBody) *DisableSqlAuditResponse
	GetBody() *DisableSqlAuditResponseBody
}

type DisableSqlAuditResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSqlAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSqlAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlAuditResponse) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSqlAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSqlAuditResponse) GetBody() *DisableSqlAuditResponseBody {
	return s.Body
}

func (s *DisableSqlAuditResponse) SetHeaders(v map[string]*string) *DisableSqlAuditResponse {
	s.Headers = v
	return s
}

func (s *DisableSqlAuditResponse) SetStatusCode(v int32) *DisableSqlAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSqlAuditResponse) SetBody(v *DisableSqlAuditResponseBody) *DisableSqlAuditResponse {
	s.Body = v
	return s
}

func (s *DisableSqlAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
