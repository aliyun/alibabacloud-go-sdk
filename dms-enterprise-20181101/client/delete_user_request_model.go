// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *DeleteUserRequest
	GetTid() *int64
	SetUid(v string) *DeleteUserRequest
	GetUid() *string
}

type DeleteUserRequest struct {
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://www.alibabacloud.com/help/en/data-management-service/latest/manage-dms-tenants) topic.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The unique ID (UID) of Alibaba Cloud account to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteUserRequest) GetUid() *string {
	return s.Uid
}

func (s *DeleteUserRequest) SetTid(v int64) *DeleteUserRequest {
	s.Tid = &v
	return s
}

func (s *DeleteUserRequest) SetUid(v string) *DeleteUserRequest {
	s.Uid = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
