// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogByIdResponse
	GetStatusCode() *int32
	SetBody(v *Catalog) *GetCatalogByIdResponse
	GetBody() *Catalog
}

type GetCatalogByIdResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Catalog           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogByIdResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogByIdResponse) GetBody() *Catalog {
	return s.Body
}

func (s *GetCatalogByIdResponse) SetHeaders(v map[string]*string) *GetCatalogByIdResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogByIdResponse) SetStatusCode(v int32) *GetCatalogByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogByIdResponse) SetBody(v *Catalog) *GetCatalogByIdResponse {
	s.Body = v
	return s
}

func (s *GetCatalogByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
