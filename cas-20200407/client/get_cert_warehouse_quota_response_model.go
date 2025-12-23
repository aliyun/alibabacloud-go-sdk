// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertWarehouseQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCertWarehouseQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCertWarehouseQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetCertWarehouseQuotaResponseBody) *GetCertWarehouseQuotaResponse
	GetBody() *GetCertWarehouseQuotaResponseBody
}

type GetCertWarehouseQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertWarehouseQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertWarehouseQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCertWarehouseQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetCertWarehouseQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCertWarehouseQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCertWarehouseQuotaResponse) GetBody() *GetCertWarehouseQuotaResponseBody {
	return s.Body
}

func (s *GetCertWarehouseQuotaResponse) SetHeaders(v map[string]*string) *GetCertWarehouseQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetCertWarehouseQuotaResponse) SetStatusCode(v int32) *GetCertWarehouseQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertWarehouseQuotaResponse) SetBody(v *GetCertWarehouseQuotaResponseBody) *GetCertWarehouseQuotaResponse {
	s.Body = v
	return s
}

func (s *GetCertWarehouseQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
