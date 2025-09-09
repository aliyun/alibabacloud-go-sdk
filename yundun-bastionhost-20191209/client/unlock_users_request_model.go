// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UnlockUsersRequest
	GetInstanceId() *string
	SetRegionId(v string) *UnlockUsersRequest
	GetRegionId() *string
	SetUserIds(v string) *UnlockUsersRequest
	GetUserIds() *string
}

type UnlockUsersRequest struct {
	// The ID of the bastion host to which the users to be unlocked belong.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the users to be unlocked belong.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that you want to unlock. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2"]
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s UnlockUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersRequest) GoString() string {
	return s.String()
}

func (s *UnlockUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnlockUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnlockUsersRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *UnlockUsersRequest) SetInstanceId(v string) *UnlockUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *UnlockUsersRequest) SetRegionId(v string) *UnlockUsersRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockUsersRequest) SetUserIds(v string) *UnlockUsersRequest {
	s.UserIds = &v
	return s
}

func (s *UnlockUsersRequest) Validate() error {
	return dara.Validate(s)
}
