// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *GetUserRequest
	GetTid() *int64
	SetUid(v string) *GetUserRequest
	GetUid() *string
	SetUserId(v string) *GetUserRequest
	GetUserId() *string
}

type GetUserRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The UID of the Alibaba Cloud account. You can view your UID by moving the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console.
	//
	// example:
	//
	// 22973492647626****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The ID of the user. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) operation to query the ID of the user.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetUserRequest) GetUid() *string {
	return s.Uid
}

func (s *GetUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUserRequest) SetTid(v int64) *GetUserRequest {
	s.Tid = &v
	return s
}

func (s *GetUserRequest) SetUid(v string) *GetUserRequest {
	s.Uid = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
