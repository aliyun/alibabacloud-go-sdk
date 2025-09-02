// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCollectionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifiedName(v string) *GetMetaCollectionDetailRequest
	GetQualifiedName() *string
}

type GetMetaCollectionDetailRequest struct {
	// The unique identifier of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// album.12345
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s GetMetaCollectionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionDetailRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *GetMetaCollectionDetailRequest) SetQualifiedName(v string) *GetMetaCollectionDetailRequest {
	s.QualifiedName = &v
	return s
}

func (s *GetMetaCollectionDetailRequest) Validate() error {
	return dara.Validate(s)
}
