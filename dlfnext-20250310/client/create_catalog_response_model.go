// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCatalogResponse
	GetStatusCode() *int32
	SetBody(v *CreateCatalogResponseBody) *CreateCatalogResponse
	GetBody() *CreateCatalogResponseBody
}

type CreateCatalogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCatalogResponse) GoString() string {
	return s.String()
}

func (s *CreateCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCatalogResponse) GetBody() *CreateCatalogResponseBody {
	return s.Body
}

func (s *CreateCatalogResponse) SetHeaders(v map[string]*string) *CreateCatalogResponse {
	s.Headers = v
	return s
}

func (s *CreateCatalogResponse) SetStatusCode(v int32) *CreateCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCatalogResponse) SetBody(v *CreateCatalogResponseBody) *CreateCatalogResponse {
	s.Body = v
	return s
}

func (s *CreateCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
