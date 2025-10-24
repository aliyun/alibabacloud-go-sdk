// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImAuditResponse
	GetStatusCode() *int32
	SetBody(v *ImAuditResponseBody) *ImAuditResponse
	GetBody() *ImAuditResponseBody
}

type ImAuditResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s ImAuditResponse) GoString() string {
	return s.String()
}

func (s *ImAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImAuditResponse) GetBody() *ImAuditResponseBody {
	return s.Body
}

func (s *ImAuditResponse) SetHeaders(v map[string]*string) *ImAuditResponse {
	s.Headers = v
	return s
}

func (s *ImAuditResponse) SetStatusCode(v int32) *ImAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *ImAuditResponse) SetBody(v *ImAuditResponseBody) *ImAuditResponse {
	s.Body = v
	return s
}

func (s *ImAuditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
