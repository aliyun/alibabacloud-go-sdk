// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTagResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagResponseBody
	GetRequestId() *string
	SetTagGroups(v []*ListTagResponseBodyTagGroups) *ListTagResponseBody
	GetTagGroups() []*ListTagResponseBodyTagGroups
	SetTotalCount(v int32) *ListTagResponseBody
	GetTotalCount() *int32
}

type ListTagResponseBody struct {
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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagGroups []*ListTagResponseBodyTagGroups `json:"TagGroups,omitempty" xml:"TagGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResponseBody) GetTagGroups() []*ListTagResponseBodyTagGroups {
	return s.TagGroups
}

func (s *ListTagResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagResponseBody) SetPageNumber(v int32) *ListTagResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagResponseBody) SetPageSize(v int32) *ListTagResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagResponseBody) SetRequestId(v string) *ListTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResponseBody) SetTagGroups(v []*ListTagResponseBodyTagGroups) *ListTagResponseBody {
	s.TagGroups = v
	return s
}

func (s *ListTagResponseBody) SetTotalCount(v int32) *ListTagResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagResponseBody) Validate() error {
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

type ListTagResponseBodyTagGroups struct {
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
	// 54545
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	// example:
	//
	// 标签1
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListTagResponseBodyTagGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponseBodyTagGroups) GoString() string {
	return s.String()
}

func (s *ListTagResponseBodyTagGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTagResponseBodyTagGroups) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *ListTagResponseBodyTagGroups) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListTagResponseBodyTagGroups) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListTagResponseBodyTagGroups) GetId() *int64 {
	return s.Id
}

func (s *ListTagResponseBodyTagGroups) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListTagResponseBodyTagGroups) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *ListTagResponseBodyTagGroups) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListTagResponseBodyTagGroups) GetTagName() *string {
	return s.TagName
}

func (s *ListTagResponseBodyTagGroups) SetCreateTime(v string) *ListTagResponseBodyTagGroups {
	s.CreateTime = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetCreateUserId(v int64) *ListTagResponseBodyTagGroups {
	s.CreateUserId = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetCreateUserName(v string) *ListTagResponseBodyTagGroups {
	s.CreateUserName = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetGroupId(v int64) *ListTagResponseBodyTagGroups {
	s.GroupId = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetId(v int64) *ListTagResponseBodyTagGroups {
	s.Id = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetModifyTime(v string) *ListTagResponseBodyTagGroups {
	s.ModifyTime = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetModifyUserId(v int64) *ListTagResponseBodyTagGroups {
	s.ModifyUserId = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetModifyUserName(v string) *ListTagResponseBodyTagGroups {
	s.ModifyUserName = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) SetTagName(v string) *ListTagResponseBodyTagGroups {
	s.TagName = &v
	return s
}

func (s *ListTagResponseBodyTagGroups) Validate() error {
	return dara.Validate(s)
}
