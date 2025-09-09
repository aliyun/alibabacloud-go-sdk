// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *LockUsersRequest
	GetInstanceId() *string
	SetRegionId(v string) *LockUsersRequest
	GetRegionId() *string
	SetUserIds(v string) *LockUsersRequest
	GetUserIds() *string
}

type LockUsersRequest struct {
	// The ID of the bastion host to which the users to be locked belong.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the users to be locked belong.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to be locked. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2","3"]
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s LockUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s LockUsersRequest) GoString() string {
	return s.String()
}

func (s *LockUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LockUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LockUsersRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *LockUsersRequest) SetInstanceId(v string) *LockUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *LockUsersRequest) SetRegionId(v string) *LockUsersRequest {
	s.RegionId = &v
	return s
}

func (s *LockUsersRequest) SetUserIds(v string) *LockUsersRequest {
	s.UserIds = &v
	return s
}

func (s *LockUsersRequest) Validate() error {
	return dara.Validate(s)
}
