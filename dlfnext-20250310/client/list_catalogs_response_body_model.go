// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogs(v []*Catalog) *ListCatalogsResponseBody
	GetCatalogs() []*Catalog
	SetNextPageToken(v string) *ListCatalogsResponseBody
	GetNextPageToken() *string
}

type ListCatalogsResponseBody struct {
	Catalogs []*Catalog `json:"catalogs,omitempty" xml:"catalogs,omitempty" type:"Repeated"`
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
}

func (s ListCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBody) GetCatalogs() []*Catalog {
	return s.Catalogs
}

func (s *ListCatalogsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListCatalogsResponseBody) SetCatalogs(v []*Catalog) *ListCatalogsResponseBody {
	s.Catalogs = v
	return s
}

func (s *ListCatalogsResponseBody) SetNextPageToken(v string) *ListCatalogsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListCatalogsResponseBody) Validate() error {
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
