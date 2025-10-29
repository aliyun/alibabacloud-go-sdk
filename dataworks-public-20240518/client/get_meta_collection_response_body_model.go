// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaCollection(v *GetMetaCollectionResponseBodyMetaCollection) *GetMetaCollectionResponseBody
	GetMetaCollection() *GetMetaCollectionResponseBodyMetaCollection
	SetRequestId(v string) *GetMetaCollectionResponseBody
	GetRequestId() *string
}

type GetMetaCollectionResponseBody struct {
	// The collection details.
	MetaCollection *GetMetaCollectionResponseBodyMetaCollection `json:"MetaCollection,omitempty" xml:"MetaCollection,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionResponseBody) GetMetaCollection() *GetMetaCollectionResponseBodyMetaCollection {
	return s.MetaCollection
}

func (s *GetMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaCollectionResponseBody) SetMetaCollection(v *GetMetaCollectionResponseBodyMetaCollection) *GetMetaCollectionResponseBody {
	s.MetaCollection = v
	return s
}

func (s *GetMetaCollectionResponseBody) SetRequestId(v string) *GetMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaCollectionResponseBody) Validate() error {
	if s.MetaCollection != nil {
		if err := s.MetaCollection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetaCollectionResponseBodyMetaCollection struct {
	// The list of administrator IDs. Valid only for the album type. The IDs must belong to users in the same tenant. Multiple IDs can be specified.
	Administrators []*int64 `json:"Administrators,omitempty" xml:"Administrators,omitempty" type:"Repeated"`
	// The creation time in milliseconds.
	//
	// example:
	//
	// 1668568601000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 456789
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The collection description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The collection ID.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modified time in milliseconds.
	//
	// example:
	//
	// 1668568601000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The collection name.
	//
	// example:
	//
	// test_category
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent collection ID. This parameter can be empty.
	//
	// example:
	//
	// category.12
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The collection type. Valid values:
	//
	// 	- Category
	//
	// 	- Album
	//
	// 	- AlbumCategory: Album subcategory.
	//
	// example:
	//
	// Category
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMetaCollectionResponseBodyMetaCollection) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionResponseBodyMetaCollection) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetAdministrators() []*int64 {
	return s.Administrators
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetDescription() *string {
	return s.Description
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetId() *string {
	return s.Id
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetName() *string {
	return s.Name
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetParentId() *string {
	return s.ParentId
}

func (s *GetMetaCollectionResponseBodyMetaCollection) GetType() *string {
	return s.Type
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetAdministrators(v []*int64) *GetMetaCollectionResponseBodyMetaCollection {
	s.Administrators = v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetCreateTime(v int64) *GetMetaCollectionResponseBodyMetaCollection {
	s.CreateTime = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetCreateUser(v string) *GetMetaCollectionResponseBodyMetaCollection {
	s.CreateUser = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetDescription(v string) *GetMetaCollectionResponseBodyMetaCollection {
	s.Description = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetId(v string) *GetMetaCollectionResponseBodyMetaCollection {
	s.Id = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetModifyTime(v int64) *GetMetaCollectionResponseBodyMetaCollection {
	s.ModifyTime = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetName(v string) *GetMetaCollectionResponseBodyMetaCollection {
	s.Name = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetParentId(v string) *GetMetaCollectionResponseBodyMetaCollection {
	s.ParentId = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) SetType(v string) *GetMetaCollectionResponseBodyMetaCollection {
	s.Type = &v
	return s
}

func (s *GetMetaCollectionResponseBodyMetaCollection) Validate() error {
	return dara.Validate(s)
}
