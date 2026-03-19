// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructDLTableInput interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int32) *OpenStructDLTableInput
	GetCreateTime() *int32
	SetCreatorId(v int64) *OpenStructDLTableInput
	GetCreatorId() *int64
	SetDescription(v string) *OpenStructDLTableInput
	GetDescription() *string
	SetLastAccessTime(v int32) *OpenStructDLTableInput
	GetLastAccessTime() *int32
	SetLocation(v string) *OpenStructDLTableInput
	GetLocation() *string
	SetModifierId(v int64) *OpenStructDLTableInput
	GetModifierId() *int64
	SetName(v string) *OpenStructDLTableInput
	GetName() *string
	SetOwner(v string) *OpenStructDLTableInput
	GetOwner() *string
	SetOwnerType(v string) *OpenStructDLTableInput
	GetOwnerType() *string
	SetParameters(v map[string]*string) *OpenStructDLTableInput
	GetParameters() map[string]*string
	SetPartitionKeys(v []*DLColumn) *OpenStructDLTableInput
	GetPartitionKeys() []*DLColumn
	SetRetention(v int32) *OpenStructDLTableInput
	GetRetention() *int32
	SetStorageDescriptor(v *DLStorageDescriptor) *OpenStructDLTableInput
	GetStorageDescriptor() *DLStorageDescriptor
	SetTableType(v string) *OpenStructDLTableInput
	GetTableType() *string
	SetViewExpandedText(v string) *OpenStructDLTableInput
	GetViewExpandedText() *string
	SetViewOriginalText(v string) *OpenStructDLTableInput
	GetViewOriginalText() *string
}

type OpenStructDLTableInput struct {
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
	// 8***
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the table.
	//
	// example:
	//
	// test
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
	// The ID of the user who last modified the table.
	//
	// example:
	//
	// 1410769
	ModifierId *int64 `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The table name.
	//
	// example:
	//
	// 100g_customer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the table.
	//
	// example:
	//
	// zhangsan
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The type of the owner. Valid values: USER, ROLE, and GROUP.
	//
	// example:
	//
	// USER
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The key-value pairs.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The column attributes of the table.
	PartitionKeys []*DLColumn `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	// The retention period of the table.
	//
	// example:
	//
	// 300
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The description of data storage, including the storage characteristics and format of the table.
	StorageDescriptor *DLStorageDescriptor `json:"StorageDescriptor,omitempty" xml:"StorageDescriptor,omitempty"`
	// The type of the metadata table. Valid values: MANAGED_TABLE, EXTERNAL_TABLE, VIRTUAL_VIEW, INDEX_TABLE, and MATERIALIZED_VIEW.
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

func (s OpenStructDLTableInput) String() string {
	return dara.Prettify(s)
}

func (s OpenStructDLTableInput) GoString() string {
	return s.String()
}

func (s *OpenStructDLTableInput) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *OpenStructDLTableInput) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *OpenStructDLTableInput) GetDescription() *string {
	return s.Description
}

func (s *OpenStructDLTableInput) GetLastAccessTime() *int32 {
	return s.LastAccessTime
}

func (s *OpenStructDLTableInput) GetLocation() *string {
	return s.Location
}

func (s *OpenStructDLTableInput) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *OpenStructDLTableInput) GetName() *string {
	return s.Name
}

func (s *OpenStructDLTableInput) GetOwner() *string {
	return s.Owner
}

func (s *OpenStructDLTableInput) GetOwnerType() *string {
	return s.OwnerType
}

func (s *OpenStructDLTableInput) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *OpenStructDLTableInput) GetPartitionKeys() []*DLColumn {
	return s.PartitionKeys
}

func (s *OpenStructDLTableInput) GetRetention() *int32 {
	return s.Retention
}

func (s *OpenStructDLTableInput) GetStorageDescriptor() *DLStorageDescriptor {
	return s.StorageDescriptor
}

func (s *OpenStructDLTableInput) GetTableType() *string {
	return s.TableType
}

func (s *OpenStructDLTableInput) GetViewExpandedText() *string {
	return s.ViewExpandedText
}

func (s *OpenStructDLTableInput) GetViewOriginalText() *string {
	return s.ViewOriginalText
}

func (s *OpenStructDLTableInput) SetCreateTime(v int32) *OpenStructDLTableInput {
	s.CreateTime = &v
	return s
}

func (s *OpenStructDLTableInput) SetCreatorId(v int64) *OpenStructDLTableInput {
	s.CreatorId = &v
	return s
}

func (s *OpenStructDLTableInput) SetDescription(v string) *OpenStructDLTableInput {
	s.Description = &v
	return s
}

func (s *OpenStructDLTableInput) SetLastAccessTime(v int32) *OpenStructDLTableInput {
	s.LastAccessTime = &v
	return s
}

func (s *OpenStructDLTableInput) SetLocation(v string) *OpenStructDLTableInput {
	s.Location = &v
	return s
}

func (s *OpenStructDLTableInput) SetModifierId(v int64) *OpenStructDLTableInput {
	s.ModifierId = &v
	return s
}

func (s *OpenStructDLTableInput) SetName(v string) *OpenStructDLTableInput {
	s.Name = &v
	return s
}

func (s *OpenStructDLTableInput) SetOwner(v string) *OpenStructDLTableInput {
	s.Owner = &v
	return s
}

func (s *OpenStructDLTableInput) SetOwnerType(v string) *OpenStructDLTableInput {
	s.OwnerType = &v
	return s
}

func (s *OpenStructDLTableInput) SetParameters(v map[string]*string) *OpenStructDLTableInput {
	s.Parameters = v
	return s
}

func (s *OpenStructDLTableInput) SetPartitionKeys(v []*DLColumn) *OpenStructDLTableInput {
	s.PartitionKeys = v
	return s
}

func (s *OpenStructDLTableInput) SetRetention(v int32) *OpenStructDLTableInput {
	s.Retention = &v
	return s
}

func (s *OpenStructDLTableInput) SetStorageDescriptor(v *DLStorageDescriptor) *OpenStructDLTableInput {
	s.StorageDescriptor = v
	return s
}

func (s *OpenStructDLTableInput) SetTableType(v string) *OpenStructDLTableInput {
	s.TableType = &v
	return s
}

func (s *OpenStructDLTableInput) SetViewExpandedText(v string) *OpenStructDLTableInput {
	s.ViewExpandedText = &v
	return s
}

func (s *OpenStructDLTableInput) SetViewOriginalText(v string) *OpenStructDLTableInput {
	s.ViewOriginalText = &v
	return s
}

func (s *OpenStructDLTableInput) Validate() error {
	if s.PartitionKeys != nil {
		for _, item := range s.PartitionKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StorageDescriptor != nil {
		if err := s.StorageDescriptor.Validate(); err != nil {
			return err
		}
	}
	return nil
}
