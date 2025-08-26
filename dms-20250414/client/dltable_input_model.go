// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLTableInput interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int32) *DLTableInput
	GetCreateTime() *int32
	SetCreatorId(v int64) *DLTableInput
	GetCreatorId() *int64
	SetDescription(v string) *DLTableInput
	GetDescription() *string
	SetLastAccessTime(v int32) *DLTableInput
	GetLastAccessTime() *int32
	SetLocation(v string) *DLTableInput
	GetLocation() *string
	SetModifierId(v int64) *DLTableInput
	GetModifierId() *int64
	SetName(v string) *DLTableInput
	GetName() *string
	SetOwner(v string) *DLTableInput
	GetOwner() *string
	SetOwnerType(v string) *DLTableInput
	GetOwnerType() *string
	SetParameters(v map[string]*string) *DLTableInput
	GetParameters() map[string]*string
	SetPartitionKeys(v []*DLColumn) *DLTableInput
	GetPartitionKeys() []*DLColumn
	SetRetention(v int32) *DLTableInput
	GetRetention() *int32
	SetStorageDescriptor(v *DLStorageDescriptor) *DLTableInput
	GetStorageDescriptor() *DLStorageDescriptor
	SetTableType(v string) *DLTableInput
	GetTableType() *string
	SetViewExpandedText(v string) *DLTableInput
	GetViewExpandedText() *string
	SetViewOriginalText(v string) *DLTableInput
	GetViewOriginalText() *string
}

type DLTableInput struct {
	CreateTime        *int32               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId         *int64               `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description       *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	LastAccessTime    *int32               `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	Location          *string              `json:"Location,omitempty" xml:"Location,omitempty"`
	ModifierId        *int64               `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	Name              *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner             *string              `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerType         *string              `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters        map[string]*string   `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PartitionKeys     []*DLColumn          `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	Retention         *int32               `json:"Retention,omitempty" xml:"Retention,omitempty"`
	StorageDescriptor *DLStorageDescriptor `json:"StorageDescriptor,omitempty" xml:"StorageDescriptor,omitempty"`
	TableType         *string              `json:"TableType,omitempty" xml:"TableType,omitempty"`
	ViewExpandedText  *string              `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	ViewOriginalText  *string              `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
}

func (s DLTableInput) String() string {
	return dara.Prettify(s)
}

func (s DLTableInput) GoString() string {
	return s.String()
}

func (s *DLTableInput) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLTableInput) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *DLTableInput) GetDescription() *string {
	return s.Description
}

func (s *DLTableInput) GetLastAccessTime() *int32 {
	return s.LastAccessTime
}

func (s *DLTableInput) GetLocation() *string {
	return s.Location
}

func (s *DLTableInput) GetModifierId() *int64 {
	return s.ModifierId
}

func (s *DLTableInput) GetName() *string {
	return s.Name
}

func (s *DLTableInput) GetOwner() *string {
	return s.Owner
}

func (s *DLTableInput) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DLTableInput) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *DLTableInput) GetPartitionKeys() []*DLColumn {
	return s.PartitionKeys
}

func (s *DLTableInput) GetRetention() *int32 {
	return s.Retention
}

func (s *DLTableInput) GetStorageDescriptor() *DLStorageDescriptor {
	return s.StorageDescriptor
}

func (s *DLTableInput) GetTableType() *string {
	return s.TableType
}

func (s *DLTableInput) GetViewExpandedText() *string {
	return s.ViewExpandedText
}

func (s *DLTableInput) GetViewOriginalText() *string {
	return s.ViewOriginalText
}

func (s *DLTableInput) SetCreateTime(v int32) *DLTableInput {
	s.CreateTime = &v
	return s
}

func (s *DLTableInput) SetCreatorId(v int64) *DLTableInput {
	s.CreatorId = &v
	return s
}

func (s *DLTableInput) SetDescription(v string) *DLTableInput {
	s.Description = &v
	return s
}

func (s *DLTableInput) SetLastAccessTime(v int32) *DLTableInput {
	s.LastAccessTime = &v
	return s
}

func (s *DLTableInput) SetLocation(v string) *DLTableInput {
	s.Location = &v
	return s
}

func (s *DLTableInput) SetModifierId(v int64) *DLTableInput {
	s.ModifierId = &v
	return s
}

func (s *DLTableInput) SetName(v string) *DLTableInput {
	s.Name = &v
	return s
}

func (s *DLTableInput) SetOwner(v string) *DLTableInput {
	s.Owner = &v
	return s
}

func (s *DLTableInput) SetOwnerType(v string) *DLTableInput {
	s.OwnerType = &v
	return s
}

func (s *DLTableInput) SetParameters(v map[string]*string) *DLTableInput {
	s.Parameters = v
	return s
}

func (s *DLTableInput) SetPartitionKeys(v []*DLColumn) *DLTableInput {
	s.PartitionKeys = v
	return s
}

func (s *DLTableInput) SetRetention(v int32) *DLTableInput {
	s.Retention = &v
	return s
}

func (s *DLTableInput) SetStorageDescriptor(v *DLStorageDescriptor) *DLTableInput {
	s.StorageDescriptor = v
	return s
}

func (s *DLTableInput) SetTableType(v string) *DLTableInput {
	s.TableType = &v
	return s
}

func (s *DLTableInput) SetViewExpandedText(v string) *DLTableInput {
	s.ViewExpandedText = &v
	return s
}

func (s *DLTableInput) SetViewOriginalText(v string) *DLTableInput {
	s.ViewOriginalText = &v
	return s
}

func (s *DLTableInput) Validate() error {
	return dara.Validate(s)
}
