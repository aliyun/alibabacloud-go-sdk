// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollection interface {
	dara.Model
	String() string
	GoString() string
	SetCollectionType(v string) *Collection
	GetCollectionType() *string
	SetComment(v string) *Collection
	GetComment() *string
	SetCreateTime(v int64) *Collection
	GetCreateTime() *int64
	SetLevel(v int32) *Collection
	GetLevel() *int32
	SetName(v string) *Collection
	GetName() *string
	SetOwnerId(v string) *Collection
	GetOwnerId() *string
	SetOwnerName(v string) *Collection
	GetOwnerName() *string
	SetQualifiedName(v string) *Collection
	GetQualifiedName() *string
	SetUpdateTime(v int64) *Collection
	GetUpdateTime() *int64
}

type Collection struct {
	// example:
	//
	// album
	CollectionType *string `json:"CollectionType,omitempty" xml:"CollectionType,omitempty"`
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1668600147617
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// collectionName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1234444
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// owner
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// album.12334
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// example:
	//
	// 1668600148617
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Collection) String() string {
	return dara.Prettify(s)
}

func (s Collection) GoString() string {
	return s.String()
}

func (s *Collection) GetCollectionType() *string {
	return s.CollectionType
}

func (s *Collection) GetComment() *string {
	return s.Comment
}

func (s *Collection) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Collection) GetLevel() *int32 {
	return s.Level
}

func (s *Collection) GetName() *string {
	return s.Name
}

func (s *Collection) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Collection) GetOwnerName() *string {
	return s.OwnerName
}

func (s *Collection) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *Collection) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *Collection) SetCollectionType(v string) *Collection {
	s.CollectionType = &v
	return s
}

func (s *Collection) SetComment(v string) *Collection {
	s.Comment = &v
	return s
}

func (s *Collection) SetCreateTime(v int64) *Collection {
	s.CreateTime = &v
	return s
}

func (s *Collection) SetLevel(v int32) *Collection {
	s.Level = &v
	return s
}

func (s *Collection) SetName(v string) *Collection {
	s.Name = &v
	return s
}

func (s *Collection) SetOwnerId(v string) *Collection {
	s.OwnerId = &v
	return s
}

func (s *Collection) SetOwnerName(v string) *Collection {
	s.OwnerName = &v
	return s
}

func (s *Collection) SetQualifiedName(v string) *Collection {
	s.QualifiedName = &v
	return s
}

func (s *Collection) SetUpdateTime(v int64) *Collection {
	s.UpdateTime = &v
	return s
}

func (s *Collection) Validate() error {
	return dara.Validate(s)
}
