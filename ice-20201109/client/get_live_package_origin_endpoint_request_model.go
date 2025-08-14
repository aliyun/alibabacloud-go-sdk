// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageOriginEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *GetLivePackageOriginEndpointRequest
	GetChannelName() *string
	SetEndpointName(v string) *GetLivePackageOriginEndpointRequest
	GetEndpointName() *string
	SetGroupName(v string) *GetLivePackageOriginEndpointRequest
	GetGroupName() *string
}

type GetLivePackageOriginEndpointRequest struct {
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

func (s GetLivePackageOriginEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageOriginEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetLivePackageOriginEndpointRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetLivePackageOriginEndpointRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetLivePackageOriginEndpointRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLivePackageOriginEndpointRequest) SetChannelName(v string) *GetLivePackageOriginEndpointRequest {
	s.ChannelName = &v
	return s
}

func (s *GetLivePackageOriginEndpointRequest) SetEndpointName(v string) *GetLivePackageOriginEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *GetLivePackageOriginEndpointRequest) SetGroupName(v string) *GetLivePackageOriginEndpointRequest {
	s.GroupName = &v
	return s
}

func (s *GetLivePackageOriginEndpointRequest) Validate() error {
	return dara.Validate(s)
}
