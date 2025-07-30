// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RevokeApplicationFromUsersRequest
	GetApplicationId() *string
	SetInstanceId(v string) *RevokeApplicationFromUsersRequest
	GetInstanceId() *string
	SetUserIds(v []*string) *RevokeApplicationFromUsersRequest
	GetUserIds() []*string
}

type RevokeApplicationFromUsersRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the accounts. You can revoke the access permissions from a maximum of 100 accounts at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RevokeApplicationFromUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromUsersRequest) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromUsersRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RevokeApplicationFromUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeApplicationFromUsersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *RevokeApplicationFromUsersRequest) SetApplicationId(v string) *RevokeApplicationFromUsersRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeApplicationFromUsersRequest) SetInstanceId(v string) *RevokeApplicationFromUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeApplicationFromUsersRequest) SetUserIds(v []*string) *RevokeApplicationFromUsersRequest {
	s.UserIds = v
	return s
}

func (s *RevokeApplicationFromUsersRequest) Validate() error {
	return dara.Validate(s)
}
