// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageChannelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteLivePackageChannelGroupRequest
	GetGroupName() *string
}

type DeleteLivePackageChannelGroupRequest struct {
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteLivePackageChannelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageChannelGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageChannelGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteLivePackageChannelGroupRequest) SetGroupName(v string) *DeleteLivePackageChannelGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteLivePackageChannelGroupRequest) Validate() error {
	return dara.Validate(s)
}
