// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageOriginEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DeleteLivePackageOriginEndpointRequest
	GetChannelName() *string
	SetEndpointName(v string) *DeleteLivePackageOriginEndpointRequest
	GetEndpointName() *string
	SetGroupName(v string) *DeleteLivePackageOriginEndpointRequest
	GetGroupName() *string
}

type DeleteLivePackageOriginEndpointRequest struct {
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The endpoint name.
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint-1
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteLivePackageOriginEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageOriginEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageOriginEndpointRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DeleteLivePackageOriginEndpointRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *DeleteLivePackageOriginEndpointRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteLivePackageOriginEndpointRequest) SetChannelName(v string) *DeleteLivePackageOriginEndpointRequest {
	s.ChannelName = &v
	return s
}

func (s *DeleteLivePackageOriginEndpointRequest) SetEndpointName(v string) *DeleteLivePackageOriginEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *DeleteLivePackageOriginEndpointRequest) SetGroupName(v string) *DeleteLivePackageOriginEndpointRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteLivePackageOriginEndpointRequest) Validate() error {
	return dara.Validate(s)
}
