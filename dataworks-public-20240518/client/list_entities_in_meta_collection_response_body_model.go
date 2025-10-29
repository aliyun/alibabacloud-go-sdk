// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesInMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListEntitiesInMetaCollectionResponseBodyPagingInfo) *ListEntitiesInMetaCollectionResponseBody
	GetPagingInfo() *ListEntitiesInMetaCollectionResponseBodyPagingInfo
	SetRequestId(v string) *ListEntitiesInMetaCollectionResponseBody
	GetRequestId() *string
}

type ListEntitiesInMetaCollectionResponseBody struct {
	// The pagination result.
	PagingInfo *ListEntitiesInMetaCollectionResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F05080B0-CCE6-5D22-B284-34A51C5D4E28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEntitiesInMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesInMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesInMetaCollectionResponseBody) GetPagingInfo() *ListEntitiesInMetaCollectionResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListEntitiesInMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEntitiesInMetaCollectionResponseBody) SetPagingInfo(v *ListEntitiesInMetaCollectionResponseBodyPagingInfo) *ListEntitiesInMetaCollectionResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBody) SetRequestId(v string) *ListEntitiesInMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEntitiesInMetaCollectionResponseBodyPagingInfo struct {
	// The list of entities in the collection.
	Entities []*ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEntitiesInMetaCollectionResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesInMetaCollectionResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) GetEntities() []*ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	return s.Entities
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) SetEntities(v []*ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) *ListEntitiesInMetaCollectionResponseBodyPagingInfo {
	s.Entities = v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) SetPageNumber(v int32) *ListEntitiesInMetaCollectionResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) SetPageSize(v int32) *ListEntitiesInMetaCollectionResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) SetTotalCount(v int32) *ListEntitiesInMetaCollectionResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfo) Validate() error {
	if s.Entities != nil {
		for _, item := range s.Entities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities struct {
	// The entity comment.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time in milliseconds.
	//
	// example:
	//
	// 1737078994080
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description specified when the entity was added to the collection. Valid only for albums.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the entity. Currently, only the Table type is supported. If the entity is deleted, this field is empty.
	//
	// example:
	//
	// dlf-table:123456789:test_catalog:test_database::test_table
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modified time in milliseconds.
	//
	// example:
	//
	// 1737078994080
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The entity name.
	//
	// example:
	//
	// test_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The entity type.
	//
	// example:
	//
	// dlf-table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GoString() string {
	return s.String()
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetComment() *string {
	return s.Comment
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetDescription() *string {
	return s.Description
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetId() *string {
	return s.Id
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetName() *string {
	return s.Name
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) GetType() *string {
	return s.Type
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetComment(v string) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.Comment = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetCreateTime(v int64) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.CreateTime = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetDescription(v string) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.Description = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetId(v string) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.Id = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetModifyTime(v int64) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.ModifyTime = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetName(v string) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.Name = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) SetType(v string) *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities {
	s.Type = &v
	return s
}

func (s *ListEntitiesInMetaCollectionResponseBodyPagingInfoEntities) Validate() error {
	return dara.Validate(s)
}
