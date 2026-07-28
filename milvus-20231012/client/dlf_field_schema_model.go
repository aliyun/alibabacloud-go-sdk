// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDlfFieldSchema interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DlfFieldSchema
	GetComment() *string
	SetDimension(v int32) *DlfFieldSchema
	GetDimension() *int32
	SetDlfFieldType(v string) *DlfFieldSchema
	GetDlfFieldType() *string
	SetFieldName(v string) *DlfFieldSchema
	GetFieldName() *string
	SetIsPrimaryKey(v bool) *DlfFieldSchema
	GetIsPrimaryKey() *bool
	SetIsSupported(v bool) *DlfFieldSchema
	GetIsSupported() *bool
	SetIsVectorField(v bool) *DlfFieldSchema
	GetIsVectorField() *bool
	SetMilvusFieldType(v string) *DlfFieldSchema
	GetMilvusFieldType() *string
	SetNullable(v bool) *DlfFieldSchema
	GetNullable() *bool
	SetUnsupportedReason(v string) *DlfFieldSchema
	GetUnsupportedReason() *string
}

type DlfFieldSchema struct {
	// A comment for the field.
	//
	// example:
	//
	// Primary key field
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The dimension of the vector field. This parameter applies only when `isVectorField` is `true`.
	//
	// example:
	//
	// 128
	Dimension *int32 `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// The DLF field type.
	//
	// example:
	//
	// BIGINT
	DlfFieldType *string `json:"dlfFieldType,omitempty" xml:"dlfFieldType,omitempty"`
	// The field name.
	//
	// example:
	//
	// id
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Indicates whether the field is a primary key.
	//
	// example:
	//
	// true
	IsPrimaryKey *bool `json:"isPrimaryKey,omitempty" xml:"isPrimaryKey,omitempty"`
	// Indicates whether the DLF field type can be mapped to a Milvus field type.
	//
	// example:
	//
	// true
	IsSupported *bool `json:"isSupported,omitempty" xml:"isSupported,omitempty"`
	// Indicates whether the field is a vector field.
	//
	// example:
	//
	// false
	IsVectorField *bool `json:"isVectorField,omitempty" xml:"isVectorField,omitempty"`
	// The corresponding Milvus field type.
	//
	// example:
	//
	// Int64
	MilvusFieldType *string `json:"milvusFieldType,omitempty" xml:"milvusFieldType,omitempty"`
	// Indicates whether the field can be null.
	//
	// example:
	//
	// true
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// The reason the DLF field type is unsupported. This field is present only when `isSupported` is `false`.
	//
	// example:
	//
	// Unsupported type
	UnsupportedReason *string `json:"unsupportedReason,omitempty" xml:"unsupportedReason,omitempty"`
}

func (s DlfFieldSchema) String() string {
	return dara.Prettify(s)
}

func (s DlfFieldSchema) GoString() string {
	return s.String()
}

func (s *DlfFieldSchema) GetComment() *string {
	return s.Comment
}

func (s *DlfFieldSchema) GetDimension() *int32 {
	return s.Dimension
}

func (s *DlfFieldSchema) GetDlfFieldType() *string {
	return s.DlfFieldType
}

func (s *DlfFieldSchema) GetFieldName() *string {
	return s.FieldName
}

func (s *DlfFieldSchema) GetIsPrimaryKey() *bool {
	return s.IsPrimaryKey
}

func (s *DlfFieldSchema) GetIsSupported() *bool {
	return s.IsSupported
}

func (s *DlfFieldSchema) GetIsVectorField() *bool {
	return s.IsVectorField
}

func (s *DlfFieldSchema) GetMilvusFieldType() *string {
	return s.MilvusFieldType
}

func (s *DlfFieldSchema) GetNullable() *bool {
	return s.Nullable
}

func (s *DlfFieldSchema) GetUnsupportedReason() *string {
	return s.UnsupportedReason
}

func (s *DlfFieldSchema) SetComment(v string) *DlfFieldSchema {
	s.Comment = &v
	return s
}

func (s *DlfFieldSchema) SetDimension(v int32) *DlfFieldSchema {
	s.Dimension = &v
	return s
}

func (s *DlfFieldSchema) SetDlfFieldType(v string) *DlfFieldSchema {
	s.DlfFieldType = &v
	return s
}

func (s *DlfFieldSchema) SetFieldName(v string) *DlfFieldSchema {
	s.FieldName = &v
	return s
}

func (s *DlfFieldSchema) SetIsPrimaryKey(v bool) *DlfFieldSchema {
	s.IsPrimaryKey = &v
	return s
}

func (s *DlfFieldSchema) SetIsSupported(v bool) *DlfFieldSchema {
	s.IsSupported = &v
	return s
}

func (s *DlfFieldSchema) SetIsVectorField(v bool) *DlfFieldSchema {
	s.IsVectorField = &v
	return s
}

func (s *DlfFieldSchema) SetMilvusFieldType(v string) *DlfFieldSchema {
	s.MilvusFieldType = &v
	return s
}

func (s *DlfFieldSchema) SetNullable(v bool) *DlfFieldSchema {
	s.Nullable = &v
	return s
}

func (s *DlfFieldSchema) SetUnsupportedReason(v string) *DlfFieldSchema {
	s.UnsupportedReason = &v
	return s
}

func (s *DlfFieldSchema) Validate() error {
	return dara.Validate(s)
}
