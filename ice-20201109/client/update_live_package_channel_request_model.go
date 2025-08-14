// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *UpdateLivePackageChannelRequest
	GetChannelName() *string
	SetDescription(v string) *UpdateLivePackageChannelRequest
	GetDescription() *string
	SetGroupName(v string) *UpdateLivePackageChannelRequest
	GetGroupName() *string
	SetProtocol(v string) *UpdateLivePackageChannelRequest
	GetProtocol() *string
	SetSegmentCount(v int32) *UpdateLivePackageChannelRequest
	GetSegmentCount() *int32
	SetSegmentDuration(v int32) *UpdateLivePackageChannelRequest
	GetSegmentDuration() *int32
}

type UpdateLivePackageChannelRequest struct {
	// The channel name. It can contain letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel description. It can be up to 1,000 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name. It can contain letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ingest protocol. Only HLS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of M3U8 segments. Valid values: 2 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SegmentCount *int32 `json:"SegmentCount,omitempty" xml:"SegmentCount,omitempty"`
	// The segment duration. Valid values: 1 to 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
}

func (s UpdateLivePackageChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateLivePackageChannelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageChannelRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageChannelRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateLivePackageChannelRequest) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *UpdateLivePackageChannelRequest) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *UpdateLivePackageChannelRequest) SetChannelName(v string) *UpdateLivePackageChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateLivePackageChannelRequest) SetDescription(v string) *UpdateLivePackageChannelRequest {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageChannelRequest) SetGroupName(v string) *UpdateLivePackageChannelRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageChannelRequest) SetProtocol(v string) *UpdateLivePackageChannelRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateLivePackageChannelRequest) SetSegmentCount(v int32) *UpdateLivePackageChannelRequest {
	s.SegmentCount = &v
	return s
}

func (s *UpdateLivePackageChannelRequest) SetSegmentDuration(v int32) *UpdateLivePackageChannelRequest {
	s.SegmentDuration = &v
	return s
}

func (s *UpdateLivePackageChannelRequest) Validate() error {
	return dara.Validate(s)
}
