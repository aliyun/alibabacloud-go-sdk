// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMetaCollectionEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollectionQualifiedName(v string) *AddMetaCollectionEntityRequest
	GetCollectionQualifiedName() *string
	SetEntityQualifiedName(v string) *AddMetaCollectionEntityRequest
	GetEntityQualifiedName() *string
	SetRemark(v string) *AddMetaCollectionEntityRequest
	GetRemark() *string
}

type AddMetaCollectionEntityRequest struct {
	// The unique identifier of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// album.2222
	CollectionQualifiedName *string `json:"CollectionQualifiedName,omitempty" xml:"CollectionQualifiedName,omitempty"`
	// The unique identifier of the entity. Example: maxcompute-table.projectA.tableA.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.projectA.tableB
	EntityQualifiedName *string `json:"EntityQualifiedName,omitempty" xml:"EntityQualifiedName,omitempty"`
	// The remarks of the entity. Example: latest product table.
	//
	// example:
	//
	// this is a remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddMetaCollectionEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMetaCollectionEntityRequest) GoString() string {
	return s.String()
}

func (s *AddMetaCollectionEntityRequest) GetCollectionQualifiedName() *string {
	return s.CollectionQualifiedName
}

func (s *AddMetaCollectionEntityRequest) GetEntityQualifiedName() *string {
	return s.EntityQualifiedName
}

func (s *AddMetaCollectionEntityRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddMetaCollectionEntityRequest) SetCollectionQualifiedName(v string) *AddMetaCollectionEntityRequest {
	s.CollectionQualifiedName = &v
	return s
}

func (s *AddMetaCollectionEntityRequest) SetEntityQualifiedName(v string) *AddMetaCollectionEntityRequest {
	s.EntityQualifiedName = &v
	return s
}

func (s *AddMetaCollectionEntityRequest) SetRemark(v string) *AddMetaCollectionEntityRequest {
	s.Remark = &v
	return s
}

func (s *AddMetaCollectionEntityRequest) Validate() error {
	return dara.Validate(s)
}
