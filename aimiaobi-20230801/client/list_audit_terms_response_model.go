// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditTermsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuditTermsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuditTermsResponse
	GetStatusCode() *int32
	SetBody(v *ListAuditTermsResponseBody) *ListAuditTermsResponse
	GetBody() *ListAuditTermsResponseBody
}

type ListAuditTermsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuditTermsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuditTermsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuditTermsResponse) GoString() string {
	return s.String()
}

func (s *ListAuditTermsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuditTermsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuditTermsResponse) GetBody() *ListAuditTermsResponseBody {
	return s.Body
}

func (s *ListAuditTermsResponse) SetHeaders(v map[string]*string) *ListAuditTermsResponse {
	s.Headers = v
	return s
}

func (s *ListAuditTermsResponse) SetStatusCode(v int32) *ListAuditTermsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuditTermsResponse) SetBody(v *ListAuditTermsResponseBody) *ListAuditTermsResponse {
	s.Body = v
	return s
}

func (s *ListAuditTermsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
