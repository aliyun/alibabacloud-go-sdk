// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannels(v *DescribeCasterChannelsResponseBodyChannels) *DescribeCasterChannelsResponseBody
	GetChannels() *DescribeCasterChannelsResponseBodyChannels
	SetRequestId(v string) *DescribeCasterChannelsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCasterChannelsResponseBody
	GetTotal() *int32
}

type DescribeCasterChannelsResponseBody struct {
	// The channels.
	Channels *DescribeCasterChannelsResponseBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 83C52866-281E-4AEA-9582-B124********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of channels.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCasterChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponseBody) GetChannels() *DescribeCasterChannelsResponseBodyChannels {
	return s.Channels
}

func (s *DescribeCasterChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterChannelsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterChannelsResponseBody) SetChannels(v *DescribeCasterChannelsResponseBodyChannels) *DescribeCasterChannelsResponseBody {
	s.Channels = v
	return s
}

func (s *DescribeCasterChannelsResponseBody) SetRequestId(v string) *DescribeCasterChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterChannelsResponseBody) SetTotal(v int32) *DescribeCasterChannelsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterChannelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterChannelsResponseBodyChannels struct {
	Channel []*DescribeCasterChannelsResponseBodyChannelsChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s DescribeCasterChannelsResponseBodyChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterChannelsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponseBodyChannels) GetChannel() []*DescribeCasterChannelsResponseBodyChannelsChannel {
	return s.Channel
}

func (s *DescribeCasterChannelsResponseBodyChannels) SetChannel(v []*DescribeCasterChannelsResponseBodyChannelsChannel) *DescribeCasterChannelsResponseBodyChannels {
	s.Channel = v
	return s
}

func (s *DescribeCasterChannelsResponseBodyChannels) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterChannelsResponseBodyChannelsChannel struct {
	// The ID of the channel.
	//
	// The layout references the channel ID when the channel is enabled. You can specify up to one video resource for the channel. The value of this parameter must be in the RV[Number] format, such as RV01 and RV12.
	//
	// example:
	//
	// RV01
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The face retouching effect. Valid values: 0 (all effects), 1 (skin smoothing), 2 (skin whitening), 3 (dark circles removal), and 4 (nasolabial folds removal).
	//
	// example:
	//
	// 0
	FaceBeauty *string `json:"FaceBeauty,omitempty" xml:"FaceBeauty,omitempty"`
	// The ID of the video resource.
	//
	// example:
	//
	// 87642866-281E-4AEA-9582-B124879****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The URL in the Real-Time Messaging Protocol (RTMP) format.
	//
	// example:
	//
	// rtmp://demo.aliyundoc.com/caster/rtmperf?auth_key=****
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
	// The URL of the output content in the channel.
	//
	// example:
	//
	// http://demo.aliyundoc.com/caster/streamwsx.flv?auth_key=YYYYY
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s DescribeCasterChannelsResponseBodyChannelsChannel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterChannelsResponseBodyChannelsChannel) GoString() string {
	return s.String()
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) GetFaceBeauty() *string {
	return s.FaceBeauty
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) GetRtmpUrl() *string {
	return s.RtmpUrl
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) SetChannelId(v string) *DescribeCasterChannelsResponseBodyChannelsChannel {
	s.ChannelId = &v
	return s
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) SetFaceBeauty(v string) *DescribeCasterChannelsResponseBodyChannelsChannel {
	s.FaceBeauty = &v
	return s
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) SetResourceId(v string) *DescribeCasterChannelsResponseBodyChannelsChannel {
	s.ResourceId = &v
	return s
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) SetRtmpUrl(v string) *DescribeCasterChannelsResponseBodyChannelsChannel {
	s.RtmpUrl = &v
	return s
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) SetStreamUrl(v string) *DescribeCasterChannelsResponseBodyChannelsChannel {
	s.StreamUrl = &v
	return s
}

func (s *DescribeCasterChannelsResponseBodyChannelsChannel) Validate() error {
	return dara.Validate(s)
}
