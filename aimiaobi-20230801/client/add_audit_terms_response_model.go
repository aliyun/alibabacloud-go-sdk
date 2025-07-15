// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditTermsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAuditTermsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAuditTermsResponse
	GetStatusCode() *int32
	SetBody(v *AddAuditTermsResponseBody) *AddAuditTermsResponse
	GetBody() *AddAuditTermsResponseBody
}

type AddAuditTermsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAuditTermsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAuditTermsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAuditTermsResponse) GoString() string {
	return s.String()
}

func (s *AddAuditTermsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAuditTermsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAuditTermsResponse) GetBody() *AddAuditTermsResponseBody {
	return s.Body
}

func (s *AddAuditTermsResponse) SetHeaders(v map[string]*string) *AddAuditTermsResponse {
	s.Headers = v
	return s
}

func (s *AddAuditTermsResponse) SetStatusCode(v int32) *AddAuditTermsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAuditTermsResponse) SetBody(v *AddAuditTermsResponseBody) *AddAuditTermsResponse {
	s.Body = v
	return s
}

func (s *AddAuditTermsResponse) Validate() error {
	return dara.Validate(s)
}
