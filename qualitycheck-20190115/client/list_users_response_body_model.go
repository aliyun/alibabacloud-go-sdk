// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUsersResponseBody
	GetCode() *string
	SetCount(v int32) *ListUsersResponseBody
	GetCount() *int32
	SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody
	GetData() *ListUsersResponseBodyData
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListUsersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUsersResponseBody
	GetSuccess() *bool
}

type ListUsersResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 12
	Count *int32                     `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUsersResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListUsersResponseBody) GetData() *ListUsersResponseBodyData {
	return s.Data
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetCount(v int32) *ListUsersResponseBody {
	s.Count = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyData struct {
	User []*ListUsersResponseBodyDataUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) GetUser() []*ListUsersResponseBodyDataUser {
	return s.User
}

func (s *ListUsersResponseBodyData) SetUser(v []*ListUsersResponseBodyDataUser) *ListUsersResponseBodyData {
	s.User = v
	return s
}

func (s *ListUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyDataUser struct {
	// example:
	//
	// 2951869706989****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 2020-03-11T16:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// XXX
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// xxx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2
	LoginUserType *int32 `json:"LoginUserType,omitempty" xml:"LoginUserType,omitempty"`
	// example:
	//
	// AGENT
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 2020-03-11T16:54Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// xxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyDataUser) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUser) GetAliUid() *string {
	return s.AliUid
}

func (s *ListUsersResponseBodyDataUser) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUsersResponseBodyDataUser) GetDescription() *string {
	return s.Description
}

func (s *ListUsersResponseBodyDataUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyDataUser) GetId() *int64 {
	return s.Id
}

func (s *ListUsersResponseBodyDataUser) GetLoginUserType() *int32 {
	return s.LoginUserType
}

func (s *ListUsersResponseBodyDataUser) GetRoleName() *string {
	return s.RoleName
}

func (s *ListUsersResponseBodyDataUser) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUsersResponseBodyDataUser) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersResponseBodyDataUser) SetAliUid(v string) *ListUsersResponseBodyDataUser {
	s.AliUid = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetCreateTime(v string) *ListUsersResponseBodyDataUser {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetDescription(v string) *ListUsersResponseBodyDataUser {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetDisplayName(v string) *ListUsersResponseBodyDataUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetId(v int64) *ListUsersResponseBodyDataUser {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetLoginUserType(v int32) *ListUsersResponseBodyDataUser {
	s.LoginUserType = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetRoleName(v string) *ListUsersResponseBodyDataUser {
	s.RoleName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetUpdateTime(v string) *ListUsersResponseBodyDataUser {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) SetUserName(v string) *ListUsersResponseBodyDataUser {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyDataUser) Validate() error {
	return dara.Validate(s)
}
