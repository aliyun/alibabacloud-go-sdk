// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIcebergPartitionField interface {
	dara.Model
	String() string
	GoString() string
	SetFieldId(v int64) *IcebergPartitionField
	GetFieldId() *int64
	SetName(v string) *IcebergPartitionField
	GetName() *string
	SetSourceId(v int64) *IcebergPartitionField
	GetSourceId() *int64
	SetTransform(v string) *IcebergPartitionField
	GetTransform() *string
}

type IcebergPartitionField struct {
	FieldId   *int64  `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceId  *int64  `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Transform *string `json:"transform,omitempty" xml:"transform,omitempty"`
}

func (s IcebergPartitionField) String() string {
	return dara.Prettify(s)
}

func (s IcebergPartitionField) GoString() string {
	return s.String()
}

func (s *IcebergPartitionField) GetFieldId() *int64 {
	return s.FieldId
}

func (s *IcebergPartitionField) GetName() *string {
	return s.Name
}

func (s *IcebergPartitionField) GetSourceId() *int64 {
	return s.SourceId
}

func (s *IcebergPartitionField) GetTransform() *string {
	return s.Transform
}

func (s *IcebergPartitionField) SetFieldId(v int64) *IcebergPartitionField {
	s.FieldId = &v
	return s
}

func (s *IcebergPartitionField) SetName(v string) *IcebergPartitionField {
	s.Name = &v
	return s
}

func (s *IcebergPartitionField) SetSourceId(v int64) *IcebergPartitionField {
	s.SourceId = &v
	return s
}

func (s *IcebergPartitionField) SetTransform(v string) *IcebergPartitionField {
	s.Transform = &v
	return s
}

func (s *IcebergPartitionField) Validate() error {
	return dara.Validate(s)
}
