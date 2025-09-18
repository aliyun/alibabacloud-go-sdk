// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIcebergTableMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentSnapshot(v *IcebergSnapshot) *IcebergTableMetadata
	GetCurrentSnapshot() *IcebergSnapshot
	SetFields(v []*IcebergNestedField) *IcebergTableMetadata
	GetFields() []*IcebergNestedField
	SetIdentifierFieldIds(v []*int32) *IcebergTableMetadata
	GetIdentifierFieldIds() []*int32
	SetPartitionFields(v []*IcebergPartitionField) *IcebergTableMetadata
	GetPartitionFields() []*IcebergPartitionField
	SetProperties(v map[string]*string) *IcebergTableMetadata
	GetProperties() map[string]*string
}

type IcebergTableMetadata struct {
	CurrentSnapshot    *IcebergSnapshot         `json:"currentSnapshot,omitempty" xml:"currentSnapshot,omitempty"`
	Fields             []*IcebergNestedField    `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	IdentifierFieldIds []*int32                 `json:"identifierFieldIds,omitempty" xml:"identifierFieldIds,omitempty" type:"Repeated"`
	PartitionFields    []*IcebergPartitionField `json:"partitionFields,omitempty" xml:"partitionFields,omitempty" type:"Repeated"`
	Properties         map[string]*string       `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s IcebergTableMetadata) String() string {
	return dara.Prettify(s)
}

func (s IcebergTableMetadata) GoString() string {
	return s.String()
}

func (s *IcebergTableMetadata) GetCurrentSnapshot() *IcebergSnapshot {
	return s.CurrentSnapshot
}

func (s *IcebergTableMetadata) GetFields() []*IcebergNestedField {
	return s.Fields
}

func (s *IcebergTableMetadata) GetIdentifierFieldIds() []*int32 {
	return s.IdentifierFieldIds
}

func (s *IcebergTableMetadata) GetPartitionFields() []*IcebergPartitionField {
	return s.PartitionFields
}

func (s *IcebergTableMetadata) GetProperties() map[string]*string {
	return s.Properties
}

func (s *IcebergTableMetadata) SetCurrentSnapshot(v *IcebergSnapshot) *IcebergTableMetadata {
	s.CurrentSnapshot = v
	return s
}

func (s *IcebergTableMetadata) SetFields(v []*IcebergNestedField) *IcebergTableMetadata {
	s.Fields = v
	return s
}

func (s *IcebergTableMetadata) SetIdentifierFieldIds(v []*int32) *IcebergTableMetadata {
	s.IdentifierFieldIds = v
	return s
}

func (s *IcebergTableMetadata) SetPartitionFields(v []*IcebergPartitionField) *IcebergTableMetadata {
	s.PartitionFields = v
	return s
}

func (s *IcebergTableMetadata) SetProperties(v map[string]*string) *IcebergTableMetadata {
	s.Properties = v
	return s
}

func (s *IcebergTableMetadata) Validate() error {
	return dara.Validate(s)
}
