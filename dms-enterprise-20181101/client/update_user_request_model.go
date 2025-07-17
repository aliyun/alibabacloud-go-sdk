// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxExecuteCount(v int64) *UpdateUserRequest
	GetMaxExecuteCount() *int64
	SetMaxResultCount(v int64) *UpdateUserRequest
	GetMaxResultCount() *int64
	SetMobile(v string) *UpdateUserRequest
	GetMobile() *string
	SetRoleNames(v string) *UpdateUserRequest
	GetRoleNames() *string
	SetTid(v int64) *UpdateUserRequest
	GetTid() *int64
	SetUid(v int64) *UpdateUserRequest
	GetUid() *int64
	SetUidString(v string) *UpdateUserRequest
	GetUidString() *string
	SetUserNick(v string) *UpdateUserRequest
	GetUserNick() *string
}

type UpdateUserRequest struct {
	// The maximum number of queries that can be performed each day.
	//
	// example:
	//
	// 1000
	MaxExecuteCount *int64 `json:"MaxExecuteCount,omitempty" xml:"MaxExecuteCount,omitempty"`
	// The maximum number of rows that can be queried each day.
	//
	// example:
	//
	// 1000
	MaxResultCount *int64 `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
	// The DingTalk ID or mobile number of the user.
	//
	// example:
	//
	// 188xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The roles that the user assumes. For more information about the valid values, see the Request parameters section in the [UpdateUser](https://help.aliyun.com/document_detail/465812.html) topic.
	//
	// example:
	//
	// ADMIN,DBA
	RoleNames *string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The Alibaba Cloud unique ID (UID) of the user to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The UID of the String type. If you specify this parameter, the UID of the Long type is replaced.
	//
	// example:
	//
	// 322824****:dmstest.wu@A201***
	UidString *string `json:"UidString,omitempty" xml:"UidString,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// test
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetMaxExecuteCount() *int64 {
	return s.MaxExecuteCount
}

func (s *UpdateUserRequest) GetMaxResultCount() *int64 {
	return s.MaxResultCount
}

func (s *UpdateUserRequest) GetMobile() *string {
	return s.Mobile
}

func (s *UpdateUserRequest) GetRoleNames() *string {
	return s.RoleNames
}

func (s *UpdateUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateUserRequest) GetUid() *int64 {
	return s.Uid
}

func (s *UpdateUserRequest) GetUidString() *string {
	return s.UidString
}

func (s *UpdateUserRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *UpdateUserRequest) SetMaxExecuteCount(v int64) *UpdateUserRequest {
	s.MaxExecuteCount = &v
	return s
}

func (s *UpdateUserRequest) SetMaxResultCount(v int64) *UpdateUserRequest {
	s.MaxResultCount = &v
	return s
}

func (s *UpdateUserRequest) SetMobile(v string) *UpdateUserRequest {
	s.Mobile = &v
	return s
}

func (s *UpdateUserRequest) SetRoleNames(v string) *UpdateUserRequest {
	s.RoleNames = &v
	return s
}

func (s *UpdateUserRequest) SetTid(v int64) *UpdateUserRequest {
	s.Tid = &v
	return s
}

func (s *UpdateUserRequest) SetUid(v int64) *UpdateUserRequest {
	s.Uid = &v
	return s
}

func (s *UpdateUserRequest) SetUidString(v string) *UpdateUserRequest {
	s.UidString = &v
	return s
}

func (s *UpdateUserRequest) SetUserNick(v string) *UpdateUserRequest {
	s.UserNick = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
