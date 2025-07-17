// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityFromMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RemoveEntityFromMetaCollectionRequest
	GetId() *string
	SetMetaCollectionId(v string) *RemoveEntityFromMetaCollectionRequest
	GetMetaCollectionId() *string
}

type RemoveEntityFromMetaCollectionRequest struct {
	// The entity ID. Currently, entities can only be tables. You can call the ListTables operation to query the ID.
	//
	// example:
	//
	// dlf-table:123456789:test_catalog:test_database::test_table
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The collection ID. You can call the ListMetaCollections operation to query the ID.
	//
	// example:
	//
	// category.123
	MetaCollectionId *string `json:"MetaCollectionId,omitempty" xml:"MetaCollectionId,omitempty"`
}

func (s RemoveEntityFromMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityFromMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntityFromMetaCollectionRequest) GetId() *string {
	return s.Id
}

func (s *RemoveEntityFromMetaCollectionRequest) GetMetaCollectionId() *string {
	return s.MetaCollectionId
}

func (s *RemoveEntityFromMetaCollectionRequest) SetId(v string) *RemoveEntityFromMetaCollectionRequest {
	s.Id = &v
	return s
}

func (s *RemoveEntityFromMetaCollectionRequest) SetMetaCollectionId(v string) *RemoveEntityFromMetaCollectionRequest {
	s.MetaCollectionId = &v
	return s
}

func (s *RemoveEntityFromMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
