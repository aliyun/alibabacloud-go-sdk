// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v *Catalog) *GetCatalogResponseBody
	GetCatalog() *Catalog
	SetRequestId(v string) *GetCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCatalogResponseBody
	GetSuccess() *bool
}

type GetCatalogResponseBody struct {
	// Catalog information.
	Catalog *Catalog `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogResponseBody) GetCatalog() *Catalog {
	return s.Catalog
}

func (s *GetCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCatalogResponseBody) SetCatalog(v *Catalog) *GetCatalogResponseBody {
	s.Catalog = v
	return s
}

func (s *GetCatalogResponseBody) SetRequestId(v string) *GetCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogResponseBody) SetSuccess(v bool) *GetCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *GetCatalogResponseBody) Validate() error {
	if s.Catalog != nil {
		if err := s.Catalog.Validate(); err != nil {
			return err
		}
	}
	return nil
}
