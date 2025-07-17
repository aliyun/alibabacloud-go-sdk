// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceUserPermissionsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListInstanceUserPermissionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceUserPermissionsRequest
	GetPageSize() *int32
	SetTid(v int64) *ListInstanceUserPermissionsRequest
	GetTid() *int64
	SetUserName(v string) *ListInstanceUserPermissionsRequest
	GetUserName() *string
}

type ListInstanceUserPermissionsRequest struct {
	// The ID of the instance. You can call the [ListInstances](https://help.aliyun.com/document_detail/141936.html) or [GetInstance](https://help.aliyun.com/document_detail/141567.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The nickname of the user. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to query the nickname of the user.
	//
	// >  The value of the NickName parameter is that of the UserName parameter.
	//
	// example:
	//
	// test_nick_name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListInstanceUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceUserPermissionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceUserPermissionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceUserPermissionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListInstanceUserPermissionsRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListInstanceUserPermissionsRequest) SetInstanceId(v string) *ListInstanceUserPermissionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceUserPermissionsRequest) SetPageNumber(v int32) *ListInstanceUserPermissionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceUserPermissionsRequest) SetPageSize(v int32) *ListInstanceUserPermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceUserPermissionsRequest) SetTid(v int64) *ListInstanceUserPermissionsRequest {
	s.Tid = &v
	return s
}

func (s *ListInstanceUserPermissionsRequest) SetUserName(v string) *ListInstanceUserPermissionsRequest {
	s.UserName = &v
	return s
}

func (s *ListInstanceUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
