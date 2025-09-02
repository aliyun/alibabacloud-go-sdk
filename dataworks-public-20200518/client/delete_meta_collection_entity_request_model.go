// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollectionQualifiedName(v string) *DeleteMetaCollectionEntityRequest
	GetCollectionQualifiedName() *string
	SetEntityQualifiedName(v string) *DeleteMetaCollectionEntityRequest
	GetEntityQualifiedName() *string
}

type DeleteMetaCollectionEntityRequest struct {
	// The unique identifier of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// album.12345
	CollectionQualifiedName *string `json:"CollectionQualifiedName,omitempty" xml:"CollectionQualifiedName,omitempty"`
	// The unique identifier of the entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute.projectA.tableB
	EntityQualifiedName *string `json:"EntityQualifiedName,omitempty" xml:"EntityQualifiedName,omitempty"`
}

func (s DeleteMetaCollectionEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionEntityRequest) GetCollectionQualifiedName() *string {
	return s.CollectionQualifiedName
}

func (s *DeleteMetaCollectionEntityRequest) GetEntityQualifiedName() *string {
	return s.EntityQualifiedName
}

func (s *DeleteMetaCollectionEntityRequest) SetCollectionQualifiedName(v string) *DeleteMetaCollectionEntityRequest {
	s.CollectionQualifiedName = &v
	return s
}

func (s *DeleteMetaCollectionEntityRequest) SetEntityQualifiedName(v string) *DeleteMetaCollectionEntityRequest {
	s.EntityQualifiedName = &v
	return s
}

func (s *DeleteMetaCollectionEntityRequest) Validate() error {
	return dara.Validate(s)
}
