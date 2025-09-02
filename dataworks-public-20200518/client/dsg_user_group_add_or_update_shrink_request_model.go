// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupAddOrUpdateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupsShrink(v string) *DsgUserGroupAddOrUpdateShrinkRequest
	GetUserGroupsShrink() *string
}

type DsgUserGroupAddOrUpdateShrinkRequest struct {
	// The information about the user group.
	//
	// This parameter is required.
	UserGroupsShrink *string `json:"UserGroups,omitempty" xml:"UserGroups,omitempty"`
}

func (s DsgUserGroupAddOrUpdateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupAddOrUpdateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgUserGroupAddOrUpdateShrinkRequest) GetUserGroupsShrink() *string {
	return s.UserGroupsShrink
}

func (s *DsgUserGroupAddOrUpdateShrinkRequest) SetUserGroupsShrink(v string) *DsgUserGroupAddOrUpdateShrinkRequest {
	s.UserGroupsShrink = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
