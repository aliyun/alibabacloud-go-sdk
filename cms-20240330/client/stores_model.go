// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStores interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v string) *Stores
	GetProject() *string
	SetRegionId(v string) *Stores
	GetRegionId() *string
	SetStore(v string) *Stores
	GetStore() *string
	SetStoreType(v string) *Stores
	GetStoreType() *string
}

type Stores struct {
	// Deprecated
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// Deprecated
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Store     *string `json:"store,omitempty" xml:"store,omitempty"`
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s Stores) String() string {
	return dara.Prettify(s)
}

func (s Stores) GoString() string {
	return s.String()
}

func (s *Stores) GetProject() *string {
	return s.Project
}

func (s *Stores) GetRegionId() *string {
	return s.RegionId
}

func (s *Stores) GetStore() *string {
	return s.Store
}

func (s *Stores) GetStoreType() *string {
	return s.StoreType
}

func (s *Stores) SetProject(v string) *Stores {
	s.Project = &v
	return s
}

func (s *Stores) SetRegionId(v string) *Stores {
	s.RegionId = &v
	return s
}

func (s *Stores) SetStore(v string) *Stores {
	s.Store = &v
	return s
}

func (s *Stores) SetStoreType(v string) *Stores {
	s.StoreType = &v
	return s
}

func (s *Stores) Validate() error {
	return dara.Validate(s)
}
