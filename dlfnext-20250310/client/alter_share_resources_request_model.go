// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterShareResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogId(v string) *AlterShareResourcesRequest
	GetCatalogId() *string
	SetShareResourceList(v []*ShareResource) *AlterShareResourcesRequest
	GetShareResourceList() []*ShareResource
}

type AlterShareResourcesRequest struct {
	// example:
	//
	// clg-paimon-xxxx
	CatalogId         *string          `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	ShareResourceList []*ShareResource `json:"shareResourceList,omitempty" xml:"shareResourceList,omitempty" type:"Repeated"`
}

func (s AlterShareResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterShareResourcesRequest) GoString() string {
	return s.String()
}

func (s *AlterShareResourcesRequest) GetCatalogId() *string {
	return s.CatalogId
}

func (s *AlterShareResourcesRequest) GetShareResourceList() []*ShareResource {
	return s.ShareResourceList
}

func (s *AlterShareResourcesRequest) SetCatalogId(v string) *AlterShareResourcesRequest {
	s.CatalogId = &v
	return s
}

func (s *AlterShareResourcesRequest) SetShareResourceList(v []*ShareResource) *AlterShareResourcesRequest {
	s.ShareResourceList = v
	return s
}

func (s *AlterShareResourcesRequest) Validate() error {
	if s.ShareResourceList != nil {
		for _, item := range s.ShareResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
