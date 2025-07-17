// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLTable interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DLTable
	GetCatalogName() *string
	SetCreateTime(v int32) *DLTable
	GetCreateTime() *int32
	SetCreatorId(v int64) *DLTable
	GetCreatorId() *int64
	SetDbId(v int64) *DLTable
	GetDbId() *int64
	SetDbName(v string) *DLTable
	GetDbName() *string
	SetDescription(v string) *DLTable
	GetDescription() *string
	SetLastAccessTime(v int32) *DLTable
	GetLastAccessTime() *int32
	SetLocation(v string) *DLTable
	GetLocation() *string
	SetModifierId(v int64) *DLTable
	GetModifierId() *int64
	SetName(v string) *DLTable
	GetName() *string
	SetOwner(v string) *DLTable
	GetOwner() *string
	SetOwnerType(v string) *DLTable
	GetOwnerType() *string
	SetParameters(v map[string]interface{}) *DLTable
	GetParameters() map[string]interface{}
	SetPartitionKeys(v []*DLColumn) *DLTable
	GetPartitionKeys() []*DLColumn
	SetRetention(v int32) *DLTable
	GetRetention() *int32
	SetStorageDescriptor(v *DLStorageDescriptor) *DLTable
	GetStorageDescriptor() *DLStorageDescriptor
	SetTableType(v string) *DLTable
	GetTableType() *string
	SetViewExpandedText(v string) *DLTable
	GetViewExpandedText() *string
	SetViewOriginalText(v string) *DLTable
	GetViewOriginalText() *string
}

type DLTable struct {
	CatalogName       *string                `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	CreateTime        *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId         *int64                 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DbId              *int64                 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName            *string                `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Description       *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	LastAccessTime    *int32                 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	Location          *string                `json:"Location,omitempty" xml:"Location,omitempty"`
	ModifierId        *int64                 `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	Name              *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner             *string                `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerType         *string                `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters        map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PartitionKeys     []*DLColumn            `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	Retention         *int32                 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	StorageDescriptor *DLStorageDescriptor   `json:"StorageDescriptor,omitempty" xml:"StorageDescriptor,omitempty"`
	TableType         *string                `json:"TableType,omitempty" xml:"TableType,omitempty"`
	ViewExpandedText  *string                `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	ViewOriginalText  *string                `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
}

func (s DLTable) String() string {
	return dara.Prettify(s)
}

func (s DLTable) GoString() string {
	return s.String()
}

func (s *DLTable) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DLTable) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLTable) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *DLTable) GetDbId() *int64 {
	return s.DbId
}

func (s *DLTable) GetDbName() *string {
	return s.DbName
}

func (s *DLTable) GetDescription() *string {
	return s.Description
}

func (s *DLTable) GetLastAccessTime() *int32 {
	return s.LastAccessTime
}

func (s *DLTable) GetLocation() *string {
	return s.Location
}

func (s *DLTable) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *DLTable) GetName() *string {
	return s.Name
}

func (s *DLTable) GetOwner() *string {
	return s.Owner
}

func (s *DLTable) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DLTable) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DLTable) GetPartitionKeys() []*DLColumn {
	return s.PartitionKeys
}

func (s *DLTable) GetRetention() *int32 {
	return s.Retention
}

func (s *DLTable) GetStorageDescriptor() *DLStorageDescriptor {
	return s.StorageDescriptor
}

func (s *DLTable) GetTableType() *string {
	return s.TableType
}

func (s *DLTable) GetViewExpandedText() *string {
	return s.ViewExpandedText
}

func (s *DLTable) GetViewOriginalText() *string {
	return s.ViewOriginalText
}

func (s *DLTable) SetCatalogName(v string) *DLTable {
	s.CatalogName = &v
	return s
}

func (s *DLTable) SetCreateTime(v int32) *DLTable {
	s.CreateTime = &v
	return s
}

func (s *DLTable) SetCreatorId(v int64) *DLTable {
	s.CreatorId = &v
	return s
}

func (s *DLTable) SetDbId(v int64) *DLTable {
	s.DbId = &v
	return s
}

func (s *DLTable) SetDbName(v string) *DLTable {
	s.DbName = &v
	return s
}

func (s *DLTable) SetDescription(v string) *DLTable {
	s.Description = &v
	return s
}

func (s *DLTable) SetLastAccessTime(v int32) *DLTable {
	s.LastAccessTime = &v
	return s
}

func (s *DLTable) SetLocation(v string) *DLTable {
	s.Location = &v
	return s
}

func (s *DLTable) SetModifierId(v int64) *DLTable {
	s.ModifierId = &v
	return s
}

func (s *DLTable) SetName(v string) *DLTable {
	s.Name = &v
	return s
}

func (s *DLTable) SetOwner(v string) *DLTable {
	s.Owner = &v
	return s
}

func (s *DLTable) SetOwnerType(v string) *DLTable {
	s.OwnerType = &v
	return s
}

func (s *DLTable) SetParameters(v map[string]interface{}) *DLTable {
	s.Parameters = v
	return s
}

func (s *DLTable) SetPartitionKeys(v []*DLColumn) *DLTable {
	s.PartitionKeys = v
	return s
}

func (s *DLTable) SetRetention(v int32) *DLTable {
	s.Retention = &v
	return s
}

func (s *DLTable) SetStorageDescriptor(v *DLStorageDescriptor) *DLTable {
	s.StorageDescriptor = v
	return s
}

func (s *DLTable) SetTableType(v string) *DLTable {
	s.TableType = &v
	return s
}

func (s *DLTable) SetViewExpandedText(v string) *DLTable {
	s.ViewExpandedText = &v
	return s
}

func (s *DLTable) SetViewOriginalText(v string) *DLTable {
	s.ViewOriginalText = &v
	return s
}

func (s *DLTable) Validate() error {
	return dara.Validate(s)
}
