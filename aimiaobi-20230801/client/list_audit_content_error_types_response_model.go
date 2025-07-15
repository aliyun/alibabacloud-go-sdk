// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditContentErrorTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuditContentErrorTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuditContentErrorTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListAuditContentErrorTypesResponseBody) *ListAuditContentErrorTypesResponse
	GetBody() *ListAuditContentErrorTypesResponseBody
}

type ListAuditContentErrorTypesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuditContentErrorTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuditContentErrorTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuditContentErrorTypesResponse) GoString() string {
	return s.String()
}

func (s *ListAuditContentErrorTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuditContentErrorTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuditContentErrorTypesResponse) GetBody() *ListAuditContentErrorTypesResponseBody {
	return s.Body
}

func (s *ListAuditContentErrorTypesResponse) SetHeaders(v map[string]*string) *ListAuditContentErrorTypesResponse {
	s.Headers = v
	return s
}

func (s *ListAuditContentErrorTypesResponse) SetStatusCode(v int32) *ListAuditContentErrorTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuditContentErrorTypesResponse) SetBody(v *ListAuditContentErrorTypesResponseBody) *ListAuditContentErrorTypesResponse {
	s.Body = v
	return s
}

func (s *ListAuditContentErrorTypesResponse) Validate() error {
	return dara.Validate(s)
}
