// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLTablebaseInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DLTablebaseInfo
	GetCatalogName() *string
	SetCreateTime(v int32) *DLTablebaseInfo
	GetCreateTime() *int32
	SetCreatorId(v int64) *DLTablebaseInfo
	GetCreatorId() *int64
	SetDbId(v int64) *DLTablebaseInfo
	GetDbId() *int64
	SetDbName(v string) *DLTablebaseInfo
	GetDbName() *string
	SetDescription(v string) *DLTablebaseInfo
	GetDescription() *string
	SetLastAccessTime(v int32) *DLTablebaseInfo
	GetLastAccessTime() *int32
	SetLocation(v string) *DLTablebaseInfo
	GetLocation() *string
	SetModifierId(v int64) *DLTablebaseInfo
	GetModifierId() *int64
	SetName(v string) *DLTablebaseInfo
	GetName() *string
	SetOwner(v string) *DLTablebaseInfo
	GetOwner() *string
	SetOwnerType(v string) *DLTablebaseInfo
	GetOwnerType() *string
	SetParameters(v map[string]interface{}) *DLTablebaseInfo
	GetParameters() map[string]interface{}
	SetPartitionKeys(v []*DLColumn) *DLTablebaseInfo
	GetPartitionKeys() []*DLColumn
	SetRetention(v int32) *DLTablebaseInfo
	GetRetention() *int32
	SetTableType(v string) *DLTablebaseInfo
	GetTableType() *string
	SetViewExpandedText(v string) *DLTablebaseInfo
	GetViewExpandedText() *string
	SetViewOriginalText(v string) *DLTablebaseInfo
	GetViewOriginalText() *string
}

type DLTablebaseInfo struct {
	// The catalog to which the table belongs.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The time when the table was created.
	//
	// example:
	//
	// 1731586286
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the table.
	//
	// example:
	//
	// 141****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The ID of the database in which the table is stored.
	//
	// example:
	//
	// 19
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the database in which the table is stored.
	//
	// example:
	//
	// 100g_customer
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The comment of the table.
	//
	// example:
	//
	// "table"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the table was last accessed.
	//
	// example:
	//
	// 1608707407
	LastAccessTime *int32 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// The storage path of the table.
	//
	// example:
	//
	// oss://xxx
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the table was modified.
	//
	// example:
	//
	// 1410769
	ModifierId *int64 `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// 100g_customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the table.
	//
	// example:
	//
	// test
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The type of the table owner. Valid values: USER, ROLE, and GROUP.
	//
	// example:
	//
	// ROLE
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// Additional parameters for the table.
	//
	// example:
	//
	// "EXTERNAL": "TRUE"
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The column attributes of the table.
	PartitionKeys []*DLColumn `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	// The retention period of the table. Unit: days.
	//
	// example:
	//
	// 30000
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The type of the table. Valid values: MANAGED_TABLE, EXTERNAL_TABLE, VIRTUAL_VIEW, INDEX_TABLE, and MATERIALIZED_VIEW.
	//
	// example:
	//
	// EXTERNAL_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// The expanded text of the view if the table type is view.
	//
	// example:
	//
	// ""
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	// The original text of the view if the table type is view.
	//
	// example:
	//
	// ""
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
}

func (s DLTablebaseInfo) String() string {
	return dara.Prettify(s)
}

func (s DLTablebaseInfo) GoString() string {
	return s.String()
}

func (s *DLTablebaseInfo) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DLTablebaseInfo) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLTablebaseInfo) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *DLTablebaseInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *DLTablebaseInfo) GetDbName() *string {
	return s.DbName
}

func (s *DLTablebaseInfo) GetDescription() *string {
	return s.Description
}

func (s *DLTablebaseInfo) GetLastAccessTime() *int32 {
	return s.LastAccessTime
}

func (s *DLTablebaseInfo) GetLocation() *string {
	return s.Location
}

func (s *DLTablebaseInfo) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *DLTablebaseInfo) GetName() *string {
	return s.Name
}

func (s *DLTablebaseInfo) GetOwner() *string {
	return s.Owner
}

func (s *DLTablebaseInfo) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DLTablebaseInfo) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DLTablebaseInfo) GetPartitionKeys() []*DLColumn {
	return s.PartitionKeys
}

func (s *DLTablebaseInfo) GetRetention() *int32 {
	return s.Retention
}

func (s *DLTablebaseInfo) GetTableType() *string {
	return s.TableType
}

func (s *DLTablebaseInfo) GetViewExpandedText() *string {
	return s.ViewExpandedText
}

func (s *DLTablebaseInfo) GetViewOriginalText() *string {
	return s.ViewOriginalText
}

func (s *DLTablebaseInfo) SetCatalogName(v string) *DLTablebaseInfo {
	s.CatalogName = &v
	return s
}

func (s *DLTablebaseInfo) SetCreateTime(v int32) *DLTablebaseInfo {
	s.CreateTime = &v
	return s
}

func (s *DLTablebaseInfo) SetCreatorId(v int64) *DLTablebaseInfo {
	s.CreatorId = &v
	return s
}

func (s *DLTablebaseInfo) SetDbId(v int64) *DLTablebaseInfo {
	s.DbId = &v
	return s
}

func (s *DLTablebaseInfo) SetDbName(v string) *DLTablebaseInfo {
	s.DbName = &v
	return s
}

func (s *DLTablebaseInfo) SetDescription(v string) *DLTablebaseInfo {
	s.Description = &v
	return s
}

func (s *DLTablebaseInfo) SetLastAccessTime(v int32) *DLTablebaseInfo {
	s.LastAccessTime = &v
	return s
}

func (s *DLTablebaseInfo) SetLocation(v string) *DLTablebaseInfo {
	s.Location = &v
	return s
}

func (s *DLTablebaseInfo) SetModifierId(v int64) *DLTablebaseInfo {
	s.ModifierId = &v
	return s
}

func (s *DLTablebaseInfo) SetName(v string) *DLTablebaseInfo {
	s.Name = &v
	return s
}

func (s *DLTablebaseInfo) SetOwner(v string) *DLTablebaseInfo {
	s.Owner = &v
	return s
}

func (s *DLTablebaseInfo) SetOwnerType(v string) *DLTablebaseInfo {
	s.OwnerType = &v
	return s
}

func (s *DLTablebaseInfo) SetParameters(v map[string]interface{}) *DLTablebaseInfo {
	s.Parameters = v
	return s
}

func (s *DLTablebaseInfo) SetPartitionKeys(v []*DLColumn) *DLTablebaseInfo {
	s.PartitionKeys = v
	return s
}

func (s *DLTablebaseInfo) SetRetention(v int32) *DLTablebaseInfo {
	s.Retention = &v
	return s
}

func (s *DLTablebaseInfo) SetTableType(v string) *DLTablebaseInfo {
	s.TableType = &v
	return s
}

func (s *DLTablebaseInfo) SetViewExpandedText(v string) *DLTablebaseInfo {
	s.ViewExpandedText = &v
	return s
}

func (s *DLTablebaseInfo) SetViewOriginalText(v string) *DLTablebaseInfo {
	s.ViewOriginalText = &v
	return s
}

func (s *DLTablebaseInfo) Validate() error {
	if s.PartitionKeys != nil {
		for _, item := range s.PartitionKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
