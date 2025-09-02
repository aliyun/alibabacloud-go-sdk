// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *DeleteMetaCollectionRequest
	GetQualifiedName() *string
}

type DeleteMetaCollectionRequest struct {
	// The unique identifier of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// album.12333
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s DeleteMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *DeleteMetaCollectionRequest) SetQualifiedName(v string) *DeleteMetaCollectionRequest {
	s.QualifiedName = &v
	return s
}

func (s *DeleteMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
