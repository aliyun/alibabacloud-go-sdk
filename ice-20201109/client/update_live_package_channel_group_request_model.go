// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLivePackageChannelGroupRequest
	GetDescription() *string
	SetGroupName(v string) *UpdateLivePackageChannelGroupRequest
	GetGroupName() *string
}

type UpdateLivePackageChannelGroupRequest struct {
	// The channel group description. It can be up to 1,000 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name. It can contain letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s UpdateLivePackageChannelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageChannelGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageChannelGroupRequest) SetDescription(v string) *UpdateLivePackageChannelGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageChannelGroupRequest) SetGroupName(v string) *UpdateLivePackageChannelGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageChannelGroupRequest) Validate() error {
	return dara.Validate(s)
}
