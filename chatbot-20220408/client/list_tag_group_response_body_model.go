// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTagGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagGroupResponseBody
	GetRequestId() *string
	SetTagGroups(v []*ListTagGroupResponseBodyTagGroups) *ListTagGroupResponseBody
	GetTagGroups() []*ListTagGroupResponseBodyTagGroups
	SetTotalCount(v int32) *ListTagGroupResponseBody
	GetTotalCount() *int32
}

type ListTagGroupResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxXxxxxx
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagGroups []*ListTagGroupResponseBodyTagGroups `json:"TagGroups,omitempty" xml:"TagGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagGroupResponseBody) GetTagGroups() []*ListTagGroupResponseBodyTagGroups {
	return s.TagGroups
}

func (s *ListTagGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagGroupResponseBody) SetPageNumber(v int32) *ListTagGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagGroupResponseBody) SetPageSize(v int32) *ListTagGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagGroupResponseBody) SetRequestId(v string) *ListTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagGroupResponseBody) SetTagGroups(v []*ListTagGroupResponseBodyTagGroups) *ListTagGroupResponseBody {
	s.TagGroups = v
	return s
}

func (s *ListTagGroupResponseBody) SetTotalCount(v int32) *ListTagGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagGroupResponseBody) Validate() error {
	if s.TagGroups != nil {
		for _, item := range s.TagGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagGroupResponseBodyTagGroups struct {
	// example:
	//
	// xxxx
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 32323
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 232323
	DefaultTagId *int64 `json:"DefaultTagId,omitempty" xml:"DefaultTagId,omitempty"`
	// example:
	//
	// 标签1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 7562321
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// xxx
	ModifyUserId *int64 `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxx
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
}

func (s ListTagGroupResponseBodyTagGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTagGroupResponseBodyTagGroups) GoString() string {
	return s.String()
}

func (s *ListTagGroupResponseBodyTagGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTagGroupResponseBodyTagGroups) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *ListTagGroupResponseBodyTagGroups) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListTagGroupResponseBodyTagGroups) GetDefaultTagId() *int64 {
	return s.DefaultTagId
}

func (s *ListTagGroupResponseBodyTagGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListTagGroupResponseBodyTagGroups) GetId() *int64 {
	return s.Id
}

func (s *ListTagGroupResponseBodyTagGroups) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListTagGroupResponseBodyTagGroups) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *ListTagGroupResponseBodyTagGroups) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListTagGroupResponseBodyTagGroups) SetCreateTime(v string) *ListTagGroupResponseBodyTagGroups {
	s.CreateTime = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetCreateUserId(v int64) *ListTagGroupResponseBodyTagGroups {
	s.CreateUserId = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetCreateUserName(v string) *ListTagGroupResponseBodyTagGroups {
	s.CreateUserName = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetDefaultTagId(v int64) *ListTagGroupResponseBodyTagGroups {
	s.DefaultTagId = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetGroupName(v string) *ListTagGroupResponseBodyTagGroups {
	s.GroupName = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetId(v int64) *ListTagGroupResponseBodyTagGroups {
	s.Id = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetModifyTime(v string) *ListTagGroupResponseBodyTagGroups {
	s.ModifyTime = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetModifyUserId(v int64) *ListTagGroupResponseBodyTagGroups {
	s.ModifyUserId = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) SetModifyUserName(v string) *ListTagGroupResponseBodyTagGroups {
	s.ModifyUserName = &v
	return s
}

func (s *ListTagGroupResponseBodyTagGroups) Validate() error {
	return dara.Validate(s)
}
