// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *ListChannelsRequest
	GetChannelName() *string
	SetChannelTier(v string) *ListChannelsRequest
	GetChannelTier() *string
	SetPageNo(v int32) *ListChannelsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChannelsRequest
	GetPageSize() *int32
	SetPlaybackMode(v string) *ListChannelsRequest
	GetPlaybackMode() *string
	SetSortBy(v string) *ListChannelsRequest
	GetSortBy() *string
	SetSortByModifiedTime(v string) *ListChannelsRequest
	GetSortByModifiedTime() *string
	SetState(v int32) *ListChannelsRequest
	GetState() *int32
}

type ListChannelsRequest struct {
	// The name of the channel.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The tier of the channel. Valid values: basic and standard.
	//
	// example:
	//
	// basic
	ChannelTier *string `json:"ChannelTier,omitempty" xml:"ChannelTier,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The playback mode. Valid values: loop and linear.
	//
	// example:
	//
	// loop
	PlaybackMode *string `json:"PlaybackMode,omitempty" xml:"PlaybackMode,omitempty"`
	// The sorting order by creation time. Valid values: asc and desc.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order by modification time. Valid values: asc and desc.
	//
	// example:
	//
	// desc
	SortByModifiedTime *string `json:"SortByModifiedTime,omitempty" xml:"SortByModifiedTime,omitempty"`
	// The channel status. A value of 0 specifies stopped. A value of 1 specifies started.
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListChannelsRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListChannelsRequest) GetChannelTier() *string {
	return s.ChannelTier
}

func (s *ListChannelsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChannelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChannelsRequest) GetPlaybackMode() *string {
	return s.PlaybackMode
}

func (s *ListChannelsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListChannelsRequest) GetSortByModifiedTime() *string {
	return s.SortByModifiedTime
}

func (s *ListChannelsRequest) GetState() *int32 {
	return s.State
}

func (s *ListChannelsRequest) SetChannelName(v string) *ListChannelsRequest {
	s.ChannelName = &v
	return s
}

func (s *ListChannelsRequest) SetChannelTier(v string) *ListChannelsRequest {
	s.ChannelTier = &v
	return s
}

func (s *ListChannelsRequest) SetPageNo(v int32) *ListChannelsRequest {
	s.PageNo = &v
	return s
}

func (s *ListChannelsRequest) SetPageSize(v int32) *ListChannelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChannelsRequest) SetPlaybackMode(v string) *ListChannelsRequest {
	s.PlaybackMode = &v
	return s
}

func (s *ListChannelsRequest) SetSortBy(v string) *ListChannelsRequest {
	s.SortBy = &v
	return s
}

func (s *ListChannelsRequest) SetSortByModifiedTime(v string) *ListChannelsRequest {
	s.SortByModifiedTime = &v
	return s
}

func (s *ListChannelsRequest) SetState(v int32) *ListChannelsRequest {
	s.State = &v
	return s
}

func (s *ListChannelsRequest) Validate() error {
	return dara.Validate(s)
}
