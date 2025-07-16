// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropCatalogResponse
	GetStatusCode() *int32
}

type DropCatalogResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DropCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s DropCatalogResponse) GoString() string {
	return s.String()
}

func (s *DropCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropCatalogResponse) SetHeaders(v map[string]*string) *DropCatalogResponse {
	s.Headers = v
	return s
}

func (s *DropCatalogResponse) SetStatusCode(v int32) *DropCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *DropCatalogResponse) Validate() error {
	return dara.Validate(s)
}
