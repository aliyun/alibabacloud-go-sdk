// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappConnectionCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWhatsappConnectionCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWhatsappConnectionCatalogResponse
	GetStatusCode() *int32
	SetBody(v *GetWhatsappConnectionCatalogResponseBody) *GetWhatsappConnectionCatalogResponse
	GetBody() *GetWhatsappConnectionCatalogResponseBody
}

type GetWhatsappConnectionCatalogResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhatsappConnectionCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhatsappConnectionCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConnectionCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetWhatsappConnectionCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWhatsappConnectionCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWhatsappConnectionCatalogResponse) GetBody() *GetWhatsappConnectionCatalogResponseBody {
	return s.Body
}

func (s *GetWhatsappConnectionCatalogResponse) SetHeaders(v map[string]*string) *GetWhatsappConnectionCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetWhatsappConnectionCatalogResponse) SetStatusCode(v int32) *GetWhatsappConnectionCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhatsappConnectionCatalogResponse) SetBody(v *GetWhatsappConnectionCatalogResponseBody) *GetWhatsappConnectionCatalogResponse {
	s.Body = v
	return s
}

func (s *GetWhatsappConnectionCatalogResponse) Validate() error {
	return dara.Validate(s)
}
