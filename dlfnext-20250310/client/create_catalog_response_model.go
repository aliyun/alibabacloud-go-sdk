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
}

type CreateCatalogResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

func (s *CreateCatalogResponse) SetHeaders(v map[string]*string) *CreateCatalogResponse {
	s.Headers = v
	return s
}

func (s *CreateCatalogResponse) SetStatusCode(v int32) *CreateCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCatalogResponse) Validate() error {
	return dara.Validate(s)
}
