// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetCatalogsRequest
	GetCatalogName() *string
}

type GetCatalogsRequest struct {
	// example:
	//
	// paimon
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
}

func (s GetCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogsRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogsRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetCatalogsRequest) SetCatalogName(v string) *GetCatalogsRequest {
	s.CatalogName = &v
	return s
}

func (s *GetCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
