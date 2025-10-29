// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogResponse
	GetStatusCode() *int32
	SetBody(v *GetCatalogResponseBody) *GetCatalogResponse
	GetBody() *GetCatalogResponseBody
}

type GetCatalogResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogResponse) GetBody() *GetCatalogResponseBody {
	return s.Body
}

func (s *GetCatalogResponse) SetHeaders(v map[string]*string) *GetCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogResponse) SetStatusCode(v int32) *GetCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogResponse) SetBody(v *GetCatalogResponseBody) *GetCatalogResponse {
	s.Body = v
	return s
}

func (s *GetCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
