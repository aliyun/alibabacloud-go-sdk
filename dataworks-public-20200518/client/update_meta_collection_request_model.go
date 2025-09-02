// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateMetaCollectionRequest
	GetComment() *string
	SetName(v string) *UpdateMetaCollectionRequest
	GetName() *string
	SetQualifiedName(v string) *UpdateMetaCollectionRequest
	GetQualifiedName() *string
}

type UpdateMetaCollectionRequest struct {
	// The comment of the collection. The comment must be 1 to 64 characters in length.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the collection.
	//
	// example:
	//
	// myCollectionName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// album.396397
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s UpdateMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaCollectionRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateMetaCollectionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMetaCollectionRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *UpdateMetaCollectionRequest) SetComment(v string) *UpdateMetaCollectionRequest {
	s.Comment = &v
	return s
}

func (s *UpdateMetaCollectionRequest) SetName(v string) *UpdateMetaCollectionRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetaCollectionRequest) SetQualifiedName(v string) *UpdateMetaCollectionRequest {
	s.QualifiedName = &v
	return s
}

func (s *UpdateMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
