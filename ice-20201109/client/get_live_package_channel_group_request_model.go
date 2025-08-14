// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageChannelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetLivePackageChannelGroupRequest
	GetGroupName() *string
}

type GetLivePackageChannelGroupRequest struct {
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetLivePackageChannelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelGroupRequest) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLivePackageChannelGroupRequest) SetGroupName(v string) *GetLivePackageChannelGroupRequest {
	s.GroupName = &v
	return s
}

func (s *GetLivePackageChannelGroupRequest) Validate() error {
	return dara.Validate(s)
}
