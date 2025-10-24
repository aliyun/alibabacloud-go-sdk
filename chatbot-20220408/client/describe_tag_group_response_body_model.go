// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeTagGroupResponseBody
	GetCreateTime() *string
	SetCreateUserId(v int64) *DescribeTagGroupResponseBody
	GetCreateUserId() *int64
	SetCreateUserName(v string) *DescribeTagGroupResponseBody
	GetCreateUserName() *string
	SetDefaultTagId(v int64) *DescribeTagGroupResponseBody
	GetDefaultTagId() *int64
	SetGroupName(v string) *DescribeTagGroupResponseBody
	GetGroupName() *string
	SetId(v int64) *DescribeTagGroupResponseBody
	GetId() *int64
	SetModifyTime(v string) *DescribeTagGroupResponseBody
	GetModifyTime() *string
	SetModifyUserId(v int64) *DescribeTagGroupResponseBody
	GetModifyUserId() *int64
	SetModifyUserName(v string) *DescribeTagGroupResponseBody
	GetModifyUserName() *string
	SetRequestId(v string) *DescribeTagGroupResponseBody
	GetRequestId() *string
}

type DescribeTagGroupResponseBody struct {
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
	// 7393472
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
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagGroupResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTagGroupResponseBody) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *DescribeTagGroupResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeTagGroupResponseBody) GetDefaultTagId() *int64 {
	return s.DefaultTagId
}

func (s *DescribeTagGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeTagGroupResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeTagGroupResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeTagGroupResponseBody) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *DescribeTagGroupResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagGroupResponseBody) SetCreateTime(v string) *DescribeTagGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetCreateUserId(v int64) *DescribeTagGroupResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetCreateUserName(v string) *DescribeTagGroupResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetDefaultTagId(v int64) *DescribeTagGroupResponseBody {
	s.DefaultTagId = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetGroupName(v string) *DescribeTagGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetId(v int64) *DescribeTagGroupResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetModifyTime(v string) *DescribeTagGroupResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetModifyUserId(v int64) *DescribeTagGroupResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetModifyUserName(v string) *DescribeTagGroupResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeTagGroupResponseBody) SetRequestId(v string) *DescribeTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
