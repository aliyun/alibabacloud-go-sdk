// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannel(v *UpdateLivePackageChannelResponseBodyLivePackageChannel) *UpdateLivePackageChannelResponseBody
	GetLivePackageChannel() *UpdateLivePackageChannelResponseBodyLivePackageChannel
	SetRequestId(v string) *UpdateLivePackageChannelResponseBody
	GetRequestId() *string
}

type UpdateLivePackageChannelResponseBody struct {
	// The information about the live package channel.
	LivePackageChannel *UpdateLivePackageChannelResponseBodyLivePackageChannel `json:"LivePackageChannel,omitempty" xml:"LivePackageChannel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 771A1414-27BF-53E6-AB73-EFCB*****ACF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLivePackageChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelResponseBody) GetLivePackageChannel() *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	return s.LivePackageChannel
}

func (s *UpdateLivePackageChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePackageChannelResponseBody) SetLivePackageChannel(v *UpdateLivePackageChannelResponseBodyLivePackageChannel) *UpdateLivePackageChannelResponseBody {
	s.LivePackageChannel = v
	return s
}

func (s *UpdateLivePackageChannelResponseBody) SetRequestId(v string) *UpdateLivePackageChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateLivePackageChannelResponseBodyLivePackageChannel struct {
	// The channel name.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The time when the channel was created.
	//
	// example:
	//
	// 2024-07-16T02:24:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The channel description. It can be up to 1,000 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ingest endpoints.
	IngestEndpoints []*UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints `json:"IngestEndpoints,omitempty" xml:"IngestEndpoints,omitempty" type:"Repeated"`
	// The time when the channel was last modified.
	//
	// example:
	//
	// 2024-07-16T02:24:42Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The ingest protocol. Only HLS is supported.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of segments.
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

func (s UpdateLivePackageChannelResponseBodyLivePackageChannel) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelResponseBodyLivePackageChannel) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetIngestEndpoints() []*UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	return s.IngestEndpoints
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetLastModified() *string {
	return s.LastModified
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetChannelName(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.ChannelName = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetCreateTime(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.CreateTime = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetDescription(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetGroupName(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetIngestEndpoints(v []*UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.IngestEndpoints = v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetLastModified(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.LastModified = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetProtocol(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.Protocol = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetSegmentCount(v int32) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.SegmentCount = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) SetSegmentDuration(v int32) *UpdateLivePackageChannelResponseBodyLivePackageChannel {
	s.SegmentDuration = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannel) Validate() error {
	return dara.Validate(s)
}

type UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints struct {
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

func (s UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetId() *string {
	return s.Id
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetPassword() *string {
	return s.Password
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetUrl() *string {
	return s.Url
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) GetUsername() *string {
	return s.Username
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetId(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Id = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetPassword(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Password = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetUrl(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Url = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) SetUsername(v string) *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints {
	s.Username = &v
	return s
}

func (s *UpdateLivePackageChannelResponseBodyLivePackageChannelIngestEndpoints) Validate() error {
	return dara.Validate(s)
}
