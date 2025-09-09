// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteUserRequest
	GetRegionId() *string
	SetUserId(v string) *DeleteUserRequest
	GetUserId() *string
}

type DeleteUserRequest struct {
	// The ID of the bastion host to which the user to be deleted belongs.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the user to be deleted belongs.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to be deleted.
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserRequest) SetInstanceId(v string) *DeleteUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserRequest) SetRegionId(v string) *DeleteUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
