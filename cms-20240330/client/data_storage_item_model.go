// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataStorageItem interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *DataStorageItem
	GetDataType() *string
	SetProject(v string) *DataStorageItem
	GetProject() *string
	SetRegionId(v string) *DataStorageItem
	GetRegionId() *string
	SetStoreName(v string) *DataStorageItem
	GetStoreName() *string
	SetStoreType(v string) *DataStorageItem
	GetStoreType() *string
}

type DataStorageItem struct {
	DataType  *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	Project   *string `json:"project,omitempty" xml:"project,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	StoreName *string `json:"storeName,omitempty" xml:"storeName,omitempty"`
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s DataStorageItem) String() string {
	return dara.Prettify(s)
}

func (s DataStorageItem) GoString() string {
	return s.String()
}

func (s *DataStorageItem) GetDataType() *string {
	return s.DataType
}

func (s *DataStorageItem) GetProject() *string {
	return s.Project
}

func (s *DataStorageItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DataStorageItem) GetStoreName() *string {
	return s.StoreName
}

func (s *DataStorageItem) GetStoreType() *string {
	return s.StoreType
}

func (s *DataStorageItem) SetDataType(v string) *DataStorageItem {
	s.DataType = &v
	return s
}

func (s *DataStorageItem) SetProject(v string) *DataStorageItem {
	s.Project = &v
	return s
}

func (s *DataStorageItem) SetRegionId(v string) *DataStorageItem {
	s.RegionId = &v
	return s
}

func (s *DataStorageItem) SetStoreName(v string) *DataStorageItem {
	s.StoreName = &v
	return s
}

func (s *DataStorageItem) SetStoreType(v string) *DataStorageItem {
	s.StoreType = &v
	return s
}

func (s *DataStorageItem) Validate() error {
	return dara.Validate(s)
}
