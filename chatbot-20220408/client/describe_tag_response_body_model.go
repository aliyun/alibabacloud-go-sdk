// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeTagResponseBody
	GetCreateTime() *string
	SetCreateUserId(v int64) *DescribeTagResponseBody
	GetCreateUserId() *int64
	SetCreateUserName(v string) *DescribeTagResponseBody
	GetCreateUserName() *string
	SetGroupId(v int64) *DescribeTagResponseBody
	GetGroupId() *int64
	SetId(v int64) *DescribeTagResponseBody
	GetId() *int64
	SetModifyTime(v string) *DescribeTagResponseBody
	GetModifyTime() *string
	SetModifyUserId(v int64) *DescribeTagResponseBody
	GetModifyUserId() *int64
	SetModifyUserName(v string) *DescribeTagResponseBody
	GetModifyUserName() *string
	SetRequestId(v string) *DescribeTagResponseBody
	GetRequestId() *string
	SetTagName(v string) *DescribeTagResponseBody
	GetTagName() *string
}

type DescribeTagResponseBody struct {
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
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	// example:
	//
	// 标签1
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTagResponseBody) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *DescribeTagResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeTagResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeTagResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeTagResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeTagResponseBody) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *DescribeTagResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagResponseBody) GetTagName() *string {
	return s.TagName
}

func (s *DescribeTagResponseBody) SetCreateTime(v string) *DescribeTagResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTagResponseBody) SetCreateUserId(v int64) *DescribeTagResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeTagResponseBody) SetCreateUserName(v string) *DescribeTagResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeTagResponseBody) SetGroupId(v int64) *DescribeTagResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeTagResponseBody) SetId(v int64) *DescribeTagResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeTagResponseBody) SetModifyTime(v string) *DescribeTagResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeTagResponseBody) SetModifyUserId(v int64) *DescribeTagResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeTagResponseBody) SetModifyUserName(v string) *DescribeTagResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeTagResponseBody) SetRequestId(v string) *DescribeTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagResponseBody) SetTagName(v string) *DescribeTagResponseBody {
	s.TagName = &v
	return s
}

func (s *DescribeTagResponseBody) Validate() error {
	return dara.Validate(s)
}
