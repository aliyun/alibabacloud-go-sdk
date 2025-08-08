// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogId(v string) *ListShareResourcesResponseBody
	GetCatalogId() *string
	SetNextPageToken(v string) *ListShareResourcesResponseBody
	GetNextPageToken() *string
	SetShareResources(v []*ShareResource) *ListShareResourcesResponseBody
	GetShareResources() []*ShareResource
}

type ListShareResourcesResponseBody struct {
	// example:
	//
	// clg-paimon-xxxx
	CatalogId *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	// example:
	//
	// ""
	NextPageToken  *string          `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	ShareResources []*ShareResource `json:"shareResources,omitempty" xml:"shareResources,omitempty" type:"Repeated"`
}

func (s ListShareResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListShareResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListShareResourcesResponseBody) GetCatalogId() *string {
	return s.CatalogId
}

func (s *ListShareResourcesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListShareResourcesResponseBody) GetShareResources() []*ShareResource {
	return s.ShareResources
}

func (s *ListShareResourcesResponseBody) SetCatalogId(v string) *ListShareResourcesResponseBody {
	s.CatalogId = &v
	return s
}

func (s *ListShareResourcesResponseBody) SetNextPageToken(v string) *ListShareResourcesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListShareResourcesResponseBody) SetShareResources(v []*ShareResource) *ListShareResourcesResponseBody {
	s.ShareResources = v
	return s
}

func (s *ListShareResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
