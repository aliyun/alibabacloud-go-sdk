// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RevokeApplicationFromGroupsRequest
	GetApplicationId() *string
	SetGroupIds(v []*string) *RevokeApplicationFromGroupsRequest
	GetGroupIds() []*string
	SetInstanceId(v string) *RevokeApplicationFromGroupsRequest
	GetInstanceId() *string
}

type RevokeApplicationFromGroupsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The group IDs. You can specify up to 100 group IDs at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RevokeApplicationFromGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromGroupsRequest) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromGroupsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RevokeApplicationFromGroupsRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *RevokeApplicationFromGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeApplicationFromGroupsRequest) SetApplicationId(v string) *RevokeApplicationFromGroupsRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeApplicationFromGroupsRequest) SetGroupIds(v []*string) *RevokeApplicationFromGroupsRequest {
	s.GroupIds = v
	return s
}

func (s *RevokeApplicationFromGroupsRequest) SetInstanceId(v string) *RevokeApplicationFromGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeApplicationFromGroupsRequest) Validate() error {
	return dara.Validate(s)
}
