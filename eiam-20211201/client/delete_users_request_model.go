// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUsersRequest
	GetInstanceId() *string
	SetUserIds(v []*string) *DeleteUsersRequest
	GetUserIds() []*string
}

type DeleteUsersRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账号ID列表
	//
	// This parameter is required.
	//
	// example:
	//
	// u_001
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUsersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *DeleteUsersRequest) SetInstanceId(v string) *DeleteUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUsersRequest) SetUserIds(v []*string) *DeleteUsersRequest {
	s.UserIds = v
	return s
}

func (s *DeleteUsersRequest) Validate() error {
	return dara.Validate(s)
}
