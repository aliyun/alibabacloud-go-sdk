// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMetaCollectionRequest
	GetDescription() *string
	SetName(v string) *CreateMetaCollectionRequest
	GetName() *string
	SetParentId(v string) *CreateMetaCollectionRequest
	GetParentId() *string
	SetType(v string) *CreateMetaCollectionRequest
	GetType() *string
}

type CreateMetaCollectionRequest struct {
	// The collection description.
	//
	// example:
	//
	// test comment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_album
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent collection ID.
	//
	// example:
	//
	// category.123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The collection name.
	//
	// 	- Category
	//
	// 	- Album
	//
	// 	- AlbumCategory: Album subcategory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Category
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaCollectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMetaCollectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaCollectionRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateMetaCollectionRequest) GetType() *string {
	return s.Type
}

func (s *CreateMetaCollectionRequest) SetDescription(v string) *CreateMetaCollectionRequest {
	s.Description = &v
	return s
}

func (s *CreateMetaCollectionRequest) SetName(v string) *CreateMetaCollectionRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaCollectionRequest) SetParentId(v string) *CreateMetaCollectionRequest {
	s.ParentId = &v
	return s
}

func (s *CreateMetaCollectionRequest) SetType(v string) *CreateMetaCollectionRequest {
	s.Type = &v
	return s
}

func (s *CreateMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
