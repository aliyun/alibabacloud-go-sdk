// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *GetLivePackageChannelRequest
	GetChannelName() *string
	SetGroupName(v string) *GetLivePackageChannelRequest
	GetGroupName() *string
}

type GetLivePackageChannelRequest struct {
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

func (s GetLivePackageChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelRequest) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetLivePackageChannelRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLivePackageChannelRequest) SetChannelName(v string) *GetLivePackageChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *GetLivePackageChannelRequest) SetGroupName(v string) *GetLivePackageChannelRequest {
	s.GroupName = &v
	return s
}

func (s *GetLivePackageChannelRequest) Validate() error {
	return dara.Validate(s)
}
