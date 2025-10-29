// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataAssetTagsResponseBodyPagingInfo) *ListDataAssetTagsResponseBody
	GetPagingInfo() *ListDataAssetTagsResponseBodyPagingInfo
	SetRequestId(v string) *ListDataAssetTagsResponseBody
	GetRequestId() *string
}

type ListDataAssetTagsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDataAssetTagsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataAssetTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAssetTagsResponseBody) GetPagingInfo() *ListDataAssetTagsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataAssetTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAssetTagsResponseBody) SetPagingInfo(v *ListDataAssetTagsResponseBodyPagingInfo) *ListDataAssetTagsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataAssetTagsResponseBody) SetRequestId(v string) *ListDataAssetTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAssetTagsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAssetTagsResponseBodyPagingInfo struct {
	// The tags.
	DataAssetTags []*ListDataAssetTagsResponseBodyPagingInfoDataAssetTags `json:"DataAssetTags,omitempty" xml:"DataAssetTags,omitempty" type:"Repeated"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 2524
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataAssetTagsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetTagsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) GetDataAssetTags() []*ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	return s.DataAssetTags
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) SetDataAssetTags(v []*ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) *ListDataAssetTagsResponseBodyPagingInfo {
	s.DataAssetTags = v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataAssetTagsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) SetPageSize(v int32) *ListDataAssetTagsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataAssetTagsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfo) Validate() error {
	if s.DataAssetTags != nil {
		for _, item := range s.DataAssetTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAssetTagsResponseBodyPagingInfoDataAssetTags struct {
	// The type of the tag.
	//
	// Valid values:
	//
	// 	- Normal
	//
	// 	- System
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the tag was created.
	//
	// example:
	//
	// 1735890003000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the tag.
	//
	// example:
	//
	// 12345
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description of the tag.
	//
	// example:
	//
	// This is a description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag administrators.
	Managers []*string `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Repeated"`
	// The time when the tag was last modified.
	//
	// example:
	//
	// 1735890003000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user who last modified the tag.
	//
	// example:
	//
	// 1234
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The type of the tag value.
	//
	// example:
	//
	// String
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GoString() string {
	return s.String()
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetCategory() *string {
	return s.Category
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetDescription() *string {
	return s.Description
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetKey() *string {
	return s.Key
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetManagers() []*string {
	return s.Managers
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetValueType() *string {
	return s.ValueType
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) GetValues() []*string {
	return s.Values
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetCategory(v string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.Category = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetCreateTime(v int64) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.CreateTime = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetCreateUser(v string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.CreateUser = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetDescription(v string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.Description = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetKey(v string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.Key = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetManagers(v []*string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.Managers = v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetModifyTime(v int64) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.ModifyTime = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetModifyUser(v string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.ModifyUser = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetValueType(v string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.ValueType = &v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) SetValues(v []*string) *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags {
	s.Values = v
	return s
}

func (s *ListDataAssetTagsResponseBodyPagingInfoDataAssetTags) Validate() error {
	return dara.Validate(s)
}
