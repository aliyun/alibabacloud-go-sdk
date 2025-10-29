// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMetaCollectionsResponseBodyData) *ListMetaCollectionsResponseBody
	GetData() *ListMetaCollectionsResponseBodyData
	SetRequestId(v string) *ListMetaCollectionsResponseBody
	GetRequestId() *string
}

type ListMetaCollectionsResponseBody struct {
	// Pagination information.
	Data *ListMetaCollectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A0DE367DDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMetaCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsResponseBody) GetData() *ListMetaCollectionsResponseBodyData {
	return s.Data
}

func (s *ListMetaCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaCollectionsResponseBody) SetData(v *ListMetaCollectionsResponseBodyData) *ListMetaCollectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaCollectionsResponseBody) SetRequestId(v string) *ListMetaCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaCollectionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMetaCollectionsResponseBodyData struct {
	// The list of collections.
	MetaCollections []*ListMetaCollectionsResponseBodyDataMetaCollections `json:"MetaCollections,omitempty" xml:"MetaCollections,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetaCollectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsResponseBodyData) GetMetaCollections() []*ListMetaCollectionsResponseBodyDataMetaCollections {
	return s.MetaCollections
}

func (s *ListMetaCollectionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetaCollectionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaCollectionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMetaCollectionsResponseBodyData) SetMetaCollections(v []*ListMetaCollectionsResponseBodyDataMetaCollections) *ListMetaCollectionsResponseBodyData {
	s.MetaCollections = v
	return s
}

func (s *ListMetaCollectionsResponseBodyData) SetPageNumber(v int32) *ListMetaCollectionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyData) SetPageSize(v int32) *ListMetaCollectionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyData) SetTotalCount(v int32) *ListMetaCollectionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyData) Validate() error {
	if s.MetaCollections != nil {
		for _, item := range s.MetaCollections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetaCollectionsResponseBodyDataMetaCollections struct {
	// The list of administrator IDs. Supported only for album types. Administrators must be users within the same tenant. Multiple administrators can be specified.
	Administrators []*string `json:"Administrators,omitempty" xml:"Administrators,omitempty" type:"Repeated"`
	// The creation time in milliseconds (timestamp).
	//
	// example:
	//
	// 1668568601000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator user ID.
	//
	// example:
	//
	// 456789
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The collection description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The collection name.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time in milliseconds (timestamp).
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
	// The ID of the parent collection. Can be empty.
	//
	// example:
	//
	// category.1
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The collection type. Valid values:
	//
	// 	- Category
	//
	// 	- Album
	//
	// 	- AlbumCategory: Album subcategory
	//
	// example:
	//
	// Category
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMetaCollectionsResponseBodyDataMetaCollections) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsResponseBodyDataMetaCollections) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetAdministrators() []*string {
	return s.Administrators
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetDescription() *string {
	return s.Description
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetId() *string {
	return s.Id
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetName() *string {
	return s.Name
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetParentId() *string {
	return s.ParentId
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) GetType() *string {
	return s.Type
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetAdministrators(v []*string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.Administrators = v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetCreateTime(v int64) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.CreateTime = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetCreateUser(v string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.CreateUser = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetDescription(v string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.Description = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetId(v string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.Id = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetModifyTime(v int64) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.ModifyTime = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetName(v string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.Name = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetParentId(v string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.ParentId = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) SetType(v string) *ListMetaCollectionsResponseBodyDataMetaCollections {
	s.Type = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyDataMetaCollections) Validate() error {
	return dara.Validate(s)
}
