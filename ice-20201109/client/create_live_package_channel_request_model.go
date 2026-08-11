// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *CreateLivePackageChannelRequest
	GetChannelName() *string
	SetClientToken(v string) *CreateLivePackageChannelRequest
	GetClientToken() *string
	SetDescription(v string) *CreateLivePackageChannelRequest
	GetDescription() *string
	SetGroupName(v string) *CreateLivePackageChannelRequest
	GetGroupName() *string
	SetProtocol(v string) *CreateLivePackageChannelRequest
	GetProtocol() *string
	SetSegmentCount(v int32) *CreateLivePackageChannelRequest
	GetSegmentCount() *int32
	SetSegmentDuration(v int32) *CreateLivePackageChannelRequest
	GetSegmentDuration() *int32
}

type CreateLivePackageChannelRequest struct {
	// The channel name. The name can contain only uppercase and lowercase letters, digits, hyphens, and underscores. The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description. The description can be up to 1000 characters in length.
	//
	// example:
	//
	// This is a test channel.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name. The name can contain only uppercase and lowercase letters, digits, hyphens, and underscores. The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The input protocol. Currently, only HLS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of m3u8 segments for the input stream. Valid values: 2 to 100.
	//
	// example:
	//
	// 3
	SegmentCount *int32 `json:"SegmentCount,omitempty" xml:"SegmentCount,omitempty"`
	// The segment duration of the input stream, in seconds. Valid values: 1 to 30.
	//
	// example:
	//
	// 6
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
}

func (s CreateLivePackageChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateLivePackageChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLivePackageChannelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageChannelRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageChannelRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateLivePackageChannelRequest) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *CreateLivePackageChannelRequest) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *CreateLivePackageChannelRequest) SetChannelName(v string) *CreateLivePackageChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateLivePackageChannelRequest) SetClientToken(v string) *CreateLivePackageChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLivePackageChannelRequest) SetDescription(v string) *CreateLivePackageChannelRequest {
	s.Description = &v
	return s
}

func (s *CreateLivePackageChannelRequest) SetGroupName(v string) *CreateLivePackageChannelRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageChannelRequest) SetProtocol(v string) *CreateLivePackageChannelRequest {
	s.Protocol = &v
	return s
}

func (s *CreateLivePackageChannelRequest) SetSegmentCount(v int32) *CreateLivePackageChannelRequest {
	s.SegmentCount = &v
	return s
}

func (s *CreateLivePackageChannelRequest) SetSegmentDuration(v int32) *CreateLivePackageChannelRequest {
	s.SegmentDuration = &v
	return s
}

func (s *CreateLivePackageChannelRequest) Validate() error {
	return dara.Validate(s)
}
