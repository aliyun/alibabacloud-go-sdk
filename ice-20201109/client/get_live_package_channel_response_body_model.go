// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannel(v *GetLivePackageChannelResponseBodyLivePackageChannel) *GetLivePackageChannelResponseBody
	GetLivePackageChannel() *GetLivePackageChannelResponseBodyLivePackageChannel
	SetRequestId(v string) *GetLivePackageChannelResponseBody
	GetRequestId() *string
}

type GetLivePackageChannelResponseBody struct {
	// Details of the live package channel.
	LivePackageChannel *GetLivePackageChannelResponseBodyLivePackageChannel `json:"LivePackageChannel,omitempty" xml:"LivePackageChannel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// RequestId-12345678
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLivePackageChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelResponseBody) GetLivePackageChannel() *GetLivePackageChannelResponseBodyLivePackageChannel {
	return s.LivePackageChannel
}

func (s *GetLivePackageChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLivePackageChannelResponseBody) SetLivePackageChannel(v *GetLivePackageChannelResponseBodyLivePackageChannel) *GetLivePackageChannelResponseBody {
	s.LivePackageChannel = v
	return s
}

func (s *GetLivePackageChannelResponseBody) SetRequestId(v string) *GetLivePackageChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLivePackageChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLivePackageChannelResponseBodyLivePackageChannel struct {
	// The channel name.
	//
	// example:
	//
	// ch4
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The time when the channel was created.
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
	IngestEndpoints []*GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints `json:"IngestEndpoints,omitempty" xml:"IngestEndpoints,omitempty" type:"Repeated"`
	// The time when the endpoint was last modified.
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

func (s GetLivePackageChannelResponseBodyLivePackageChannel) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelResponseBodyLivePackageChannel) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetDescription() *string {
	return s.Description
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetIngestEndpoints() []*GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	return s.IngestEndpoints
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetLastModified() *string {
	return s.LastModified
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetProtocol() *string {
	return s.Protocol
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetChannelName(v string) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.ChannelName = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetCreateTime(v string) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.CreateTime = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetDescription(v string) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.Description = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetGroupName(v string) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.GroupName = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetIngestEndpoints(v []*GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.IngestEndpoints = v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetLastModified(v string) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.LastModified = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetProtocol(v string) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.Protocol = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetSegmentCount(v int32) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.SegmentCount = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) SetSegmentDuration(v int32) *GetLivePackageChannelResponseBodyLivePackageChannel {
	s.SegmentDuration = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannel) Validate() error {
	return dara.Validate(s)
}

type GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints struct {
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

func (s GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetId() *string {
	return s.Id
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetPassword() *string {
	return s.Password
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetUrl() *string {
	return s.Url
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetUsername() *string {
	return s.Username
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetId(v string) *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Id = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetPassword(v string) *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Password = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetUrl(v string) *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Url = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetUsername(v string) *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Username = &v
	return s
}

func (s *GetLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) Validate() error {
	return dara.Validate(s)
}
