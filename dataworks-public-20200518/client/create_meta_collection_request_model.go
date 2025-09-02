// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollectionType(v string) *CreateMetaCollectionRequest
	GetCollectionType() *string
	SetComment(v string) *CreateMetaCollectionRequest
	GetComment() *string
	SetName(v string) *CreateMetaCollectionRequest
	GetName() *string
	SetParentQualifiedName(v string) *CreateMetaCollectionRequest
	GetParentQualifiedName() *string
}

type CreateMetaCollectionRequest struct {
	// The type of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALBUM
	CollectionType *string `json:"CollectionType,omitempty" xml:"CollectionType,omitempty"`
	// The comment of the collection. The comment must be 1 to 64 characters in length.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the collection. The name must be 1 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// collection_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the parent collection.
	//
	// example:
	//
	// album.333508
	ParentQualifiedName *string `json:"ParentQualifiedName,omitempty" xml:"ParentQualifiedName,omitempty"`
}

func (s CreateMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaCollectionRequest) GetCollectionType() *string {
	return s.CollectionType
}

func (s *CreateMetaCollectionRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateMetaCollectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaCollectionRequest) GetParentQualifiedName() *string {
	return s.ParentQualifiedName
}

func (s *CreateMetaCollectionRequest) SetCollectionType(v string) *CreateMetaCollectionRequest {
	s.CollectionType = &v
	return s
}

func (s *CreateMetaCollectionRequest) SetComment(v string) *CreateMetaCollectionRequest {
	s.Comment = &v
	return s
}

func (s *CreateMetaCollectionRequest) SetName(v string) *CreateMetaCollectionRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaCollectionRequest) SetParentQualifiedName(v string) *CreateMetaCollectionRequest {
	s.ParentQualifiedName = &v
	return s
}

func (s *CreateMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
