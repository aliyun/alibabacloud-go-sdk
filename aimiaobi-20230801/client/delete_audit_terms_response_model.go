// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuditTermsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuditTermsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuditTermsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuditTermsResponseBody) *DeleteAuditTermsResponse
	GetBody() *DeleteAuditTermsResponseBody
}

type DeleteAuditTermsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuditTermsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuditTermsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuditTermsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuditTermsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuditTermsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuditTermsResponse) GetBody() *DeleteAuditTermsResponseBody {
	return s.Body
}

func (s *DeleteAuditTermsResponse) SetHeaders(v map[string]*string) *DeleteAuditTermsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuditTermsResponse) SetStatusCode(v int32) *DeleteAuditTermsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuditTermsResponse) SetBody(v *DeleteAuditTermsResponseBody) *DeleteAuditTermsResponse {
	s.Body = v
	return s
}

func (s *DeleteAuditTermsResponse) Validate() error {
	return dara.Validate(s)
}
