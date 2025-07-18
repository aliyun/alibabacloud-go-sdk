// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupIds(v []*string) *ListPolicesForUserGroupRequest
	GetUserGroupIds() []*string
}

type ListPolicesForUserGroupRequest struct {
	// This parameter is required.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListPolicesForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListPolicesForUserGroupRequest) SetUserGroupIds(v []*string) *ListPolicesForUserGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *ListPolicesForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
