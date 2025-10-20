// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterCatalogResponse
	GetStatusCode() *int32
	SetBody(v *AlterCatalogResponseBody) *AlterCatalogResponse
	GetBody() *AlterCatalogResponseBody
}

type AlterCatalogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlterCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlterCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterCatalogResponse) GoString() string {
	return s.String()
}

func (s *AlterCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterCatalogResponse) GetBody() *AlterCatalogResponseBody {
	return s.Body
}

func (s *AlterCatalogResponse) SetHeaders(v map[string]*string) *AlterCatalogResponse {
	s.Headers = v
	return s
}

func (s *AlterCatalogResponse) SetStatusCode(v int32) *AlterCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterCatalogResponse) SetBody(v *AlterCatalogResponseBody) *AlterCatalogResponse {
	s.Body = v
	return s
}

func (s *AlterCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
