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
	// The ID of the entity. Currently, only the table type is supported. You can obtain the ID from the response of the ListTables operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table:::project_name:[schema_name]:table_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the collection object. You can obtain the ID from the response of the ListMetaCollections operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// as78d756asd
	MetaCollectionId *string `json:"MetaCollectionId,omitempty" xml:"MetaCollectionId,omitempty"`
	// The remarks when adding the entity to the collection. Currently, this parameter takes effect only for the album type.
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
