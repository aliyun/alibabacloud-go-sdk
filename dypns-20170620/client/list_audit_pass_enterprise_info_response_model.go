// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditPassEnterpriseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuditPassEnterpriseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuditPassEnterpriseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListAuditPassEnterpriseInfoResponseBody) *ListAuditPassEnterpriseInfoResponse
	GetBody() *ListAuditPassEnterpriseInfoResponseBody
}

type ListAuditPassEnterpriseInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuditPassEnterpriseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuditPassEnterpriseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuditPassEnterpriseInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAuditPassEnterpriseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuditPassEnterpriseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuditPassEnterpriseInfoResponse) GetBody() *ListAuditPassEnterpriseInfoResponseBody {
	return s.Body
}

func (s *ListAuditPassEnterpriseInfoResponse) SetHeaders(v map[string]*string) *ListAuditPassEnterpriseInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponse) SetStatusCode(v int32) *ListAuditPassEnterpriseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponse) SetBody(v *ListAuditPassEnterpriseInfoResponseBody) *ListAuditPassEnterpriseInfoResponse {
	s.Body = v
	return s
}

func (s *ListAuditPassEnterpriseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
