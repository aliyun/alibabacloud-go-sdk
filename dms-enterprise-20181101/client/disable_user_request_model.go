// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *DisableUserRequest
	GetTid() *int64
	SetUid(v string) *DisableUserRequest
	GetUid() *string
}

type DisableUserRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The UID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DisableUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableUserRequest) GoString() string {
	return s.String()
}

func (s *DisableUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DisableUserRequest) GetUid() *string {
	return s.Uid
}

func (s *DisableUserRequest) SetTid(v int64) *DisableUserRequest {
	s.Tid = &v
	return s
}

func (s *DisableUserRequest) SetUid(v string) *DisableUserRequest {
	s.Uid = &v
	return s
}

func (s *DisableUserRequest) Validate() error {
	return dara.Validate(s)
}
