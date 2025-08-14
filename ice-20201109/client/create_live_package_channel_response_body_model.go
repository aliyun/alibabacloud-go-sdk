// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannel(v *CreateLivePackageChannelResponseBodyLivePackageChannel) *CreateLivePackageChannelResponseBody
	GetLivePackageChannel() *CreateLivePackageChannelResponseBodyLivePackageChannel
	SetRequestId(v string) *CreateLivePackageChannelResponseBody
	GetRequestId() *string
}

type CreateLivePackageChannelResponseBody struct {
	// The information about the live package channel.
	LivePackageChannel *CreateLivePackageChannelResponseBodyLivePackageChannel `json:"LivePackageChannel,omitempty" xml:"LivePackageChannel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLivePackageChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelResponseBody) GetLivePackageChannel() *CreateLivePackageChannelResponseBodyLivePackageChannel {
	return s.LivePackageChannel
}

func (s *CreateLivePackageChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivePackageChannelResponseBody) SetLivePackageChannel(v *CreateLivePackageChannelResponseBodyLivePackageChannel) *CreateLivePackageChannelResponseBody {
	s.LivePackageChannel = v
	return s
}

func (s *CreateLivePackageChannelResponseBody) SetRequestId(v string) *CreateLivePackageChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivePackageChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLivePackageChannelResponseBodyLivePackageChannel struct {
	// The channel name.
	//
	// example:
	//
	// example-channel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The time when the channel was created. It is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The channel description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ingest endpoints.
	IngestEndpoints []*CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints `json:"IngestEndpoints,omitempty" xml:"IngestEndpoints,omitempty" type:"Repeated"`
	// The time when the channel was last modified. It is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The ingest protocol. Only HLS is supported.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of M3U8 segments.
	//
	// example:
	//
	// 3
	SegmentCount *int32 `json:"SegmentCount,omitempty" xml:"SegmentCount,omitempty"`
	// The segment duration.
	//
	// example:
	//
	// 5
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
}

func (s CreateLivePackageChannelResponseBodyLivePackageChannel) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelResponseBodyLivePackageChannel) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetIngestEndpoints() []*CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	return s.IngestEndpoints
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetLastModified() *string {
	return s.LastModified
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetChannelName(v string) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.ChannelName = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetCreateTime(v string) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.CreateTime = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetDescription(v string) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.Description = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetGroupName(v string) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetIngestEndpoints(v []*CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.IngestEndpoints = v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetLastModified(v string) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.LastModified = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetProtocol(v string) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.Protocol = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetSegmentCount(v int32) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.SegmentCount = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) SetSegmentDuration(v int32) *CreateLivePackageChannelResponseBodyLivePackageChannel {
	s.SegmentDuration = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannel) Validate() error {
	return dara.Validate(s)
}

type CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints struct {
	// The ingest endpoint ID.
	//
	// example:
	//
	// ingest1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The password.
	//
	// example:
	//
	// 2F9e******b569c8
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ingest endpoint URL.
	//
	// example:
	//
	// http://xxx-1.packagepush-abcxxx.ap-southeast-1.aliyuncsiceintl.com/v1/group01/1/ch01/manifest
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The username.
	//
	// example:
	//
	// us12******das
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetId() *string {
	return s.Id
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetPassword() *string {
	return s.Password
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetUrl() *string {
	return s.Url
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetUsername() *string {
	return s.Username
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetId(v string) *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Id = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetPassword(v string) *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Password = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetUrl(v string) *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Url = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetUsername(v string) *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Username = &v
	return s
}

func (s *CreateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) Validate() error {
	return dara.Validate(s)
}
