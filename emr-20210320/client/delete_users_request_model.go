// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteUsersRequest
	GetClusterId() *string
	SetRegionId(v string) *DeleteUsersRequest
	GetRegionId() *string
	SetUserNames(v []*string) *DeleteUsersRequest
	GetUserNames() []*string
}

type DeleteUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The usernames. Number of elements in the array: 0 to 10.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUsersRequest) GetUserNames() []*string {
	return s.UserNames
}

func (s *DeleteUsersRequest) SetClusterId(v string) *DeleteUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersRequest) SetRegionId(v string) *DeleteUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUsersRequest) SetUserNames(v []*string) *DeleteUsersRequest {
	s.UserNames = v
	return s
}

func (s *DeleteUsersRequest) Validate() error {
	return dara.Validate(s)
}
