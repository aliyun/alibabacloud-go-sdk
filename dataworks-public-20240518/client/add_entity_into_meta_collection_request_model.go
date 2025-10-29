// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntityIntoMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *AddEntityIntoMetaCollectionRequest
	GetId() *string
	SetMetaCollectionId(v string) *AddEntityIntoMetaCollectionRequest
	GetMetaCollectionId() *string
	SetRemark(v string) *AddEntityIntoMetaCollectionRequest
	GetRemark() *string
}

type AddEntityIntoMetaCollectionRequest struct {
	// The entity ID. Currently, only table entities are supported. You can call the ListTables operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The collection ID. You can refer to the return result of the ListMetaCollections operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// category.123
	MetaCollectionId *string `json:"MetaCollectionId,omitempty" xml:"MetaCollectionId,omitempty"`
	// Remarks added when adding the entity to a collection. This parameter is currently valid only for album collections.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddEntityIntoMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEntityIntoMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *AddEntityIntoMetaCollectionRequest) GetId() *string {
	return s.Id
}

func (s *AddEntityIntoMetaCollectionRequest) GetMetaCollectionId() *string {
	return s.MetaCollectionId
}

func (s *AddEntityIntoMetaCollectionRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddEntityIntoMetaCollectionRequest) SetId(v string) *AddEntityIntoMetaCollectionRequest {
	s.Id = &v
	return s
}

func (s *AddEntityIntoMetaCollectionRequest) SetMetaCollectionId(v string) *AddEntityIntoMetaCollectionRequest {
	s.MetaCollectionId = &v
	return s
}

func (s *AddEntityIntoMetaCollectionRequest) SetRemark(v string) *AddEntityIntoMetaCollectionRequest {
	s.Remark = &v
	return s
}

func (s *AddEntityIntoMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
