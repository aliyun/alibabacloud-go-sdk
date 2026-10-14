// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *GetAgenticCatalogRequest
	GetCatalogUuid() *string
}

type GetAgenticCatalogRequest struct {
	// The catalog UUID. You can obtain it by calling ListAgenticCatalog or SearchAgenticCatalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
}

func (s GetAgenticCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticCatalogRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *GetAgenticCatalogRequest) SetCatalogUuid(v string) *GetAgenticCatalogRequest {
	s.CatalogUuid = &v
	return s
}

func (s *GetAgenticCatalogRequest) Validate() error {
	return dara.Validate(s)
}
