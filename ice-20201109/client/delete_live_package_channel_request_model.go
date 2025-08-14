// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DeleteLivePackageChannelRequest
	GetChannelName() *string
	SetGroupName(v string) *DeleteLivePackageChannelRequest
	GetGroupName() *string
}

type DeleteLivePackageChannelRequest struct {
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteLivePackageChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DeleteLivePackageChannelRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteLivePackageChannelRequest) SetChannelName(v string) *DeleteLivePackageChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *DeleteLivePackageChannelRequest) SetGroupName(v string) *DeleteLivePackageChannelRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteLivePackageChannelRequest) Validate() error {
	return dara.Validate(s)
}
