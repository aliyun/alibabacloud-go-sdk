// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAuditResponse
	GetStatusCode() *int32
	SetBody(v *CreateAuditResponseBody) *CreateAuditResponse
	GetBody() *CreateAuditResponseBody
}

type CreateAuditResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAuditResponse) GoString() string {
	return s.String()
}

func (s *CreateAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAuditResponse) GetBody() *CreateAuditResponseBody {
	return s.Body
}

func (s *CreateAuditResponse) SetHeaders(v map[string]*string) *CreateAuditResponse {
	s.Headers = v
	return s
}

func (s *CreateAuditResponse) SetStatusCode(v int32) *CreateAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuditResponse) SetBody(v *CreateAuditResponseBody) *CreateAuditResponse {
	s.Body = v
	return s
}

func (s *CreateAuditResponse) Validate() error {
	return dara.Validate(s)
}
