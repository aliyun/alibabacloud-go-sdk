// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannels(v []*ListLivePackageChannelsResponseBodyLivePackageChannels) *ListLivePackageChannelsResponseBody
	GetLivePackageChannels() []*ListLivePackageChannelsResponseBodyLivePackageChannels
	SetPageNo(v int64) *ListLivePackageChannelsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListLivePackageChannelsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListLivePackageChannelsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLivePackageChannelsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListLivePackageChannelsResponseBody
	GetTotalCount() *int64
}

type ListLivePackageChannelsResponseBody struct {
	// The live package channels.
	LivePackageChannels []*ListLivePackageChannelsResponseBodyLivePackageChannels `json:"LivePackageChannels,omitempty" xml:"LivePackageChannels,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sort order. Valid values: asc and desc (default).
	//
	// example:
	//
	// asc/desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLivePackageChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelsResponseBody) GetLivePackageChannels() []*ListLivePackageChannelsResponseBodyLivePackageChannels {
	return s.LivePackageChannels
}

func (s *ListLivePackageChannelsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLivePackageChannelsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLivePackageChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLivePackageChannelsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLivePackageChannelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLivePackageChannelsResponseBody) SetLivePackageChannels(v []*ListLivePackageChannelsResponseBodyLivePackageChannels) *ListLivePackageChannelsResponseBody {
	s.LivePackageChannels = v
	return s
}

func (s *ListLivePackageChannelsResponseBody) SetPageNo(v int64) *ListLivePackageChannelsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLivePackageChannelsResponseBody) SetPageSize(v int64) *ListLivePackageChannelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLivePackageChannelsResponseBody) SetRequestId(v string) *ListLivePackageChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivePackageChannelsResponseBody) SetSortBy(v string) *ListLivePackageChannelsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLivePackageChannelsResponseBody) SetTotalCount(v int64) *ListLivePackageChannelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLivePackageChannelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLivePackageChannelsResponseBodyLivePackageChannels struct {
	// The channel name.
	//
	// example:
	//
	// ch3
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
	IngestEndpoints []*ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints `json:"IngestEndpoints,omitempty" xml:"IngestEndpoints,omitempty" type:"Repeated"`
	// The time when the channel was last modified.
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

func (s ListLivePackageChannelsResponseBodyLivePackageChannels) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelsResponseBodyLivePackageChannels) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetDescription() *string {
	return s.Description
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetIngestEndpoints() []*ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints {
	return s.IngestEndpoints
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetLastModified() *string {
	return s.LastModified
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetProtocol() *string {
	return s.Protocol
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetSegmentCount() *int32 {
	return s.SegmentCount
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetChannelName(v string) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.ChannelName = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetCreateTime(v string) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.CreateTime = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetDescription(v string) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.Description = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetGroupName(v string) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.GroupName = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetIngestEndpoints(v []*ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.IngestEndpoints = v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetLastModified(v string) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.LastModified = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetProtocol(v string) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.Protocol = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetSegmentCount(v int32) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.SegmentCount = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) SetSegmentDuration(v int32) *ListLivePackageChannelsResponseBodyLivePackageChannels {
	s.SegmentDuration = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannels) Validate() error {
	return dara.Validate(s)
}

type ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints struct {
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
	// 2F9e9******18b569c8
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

func (s ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) GetId() *string {
	return s.Id
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) GetPassword() *string {
	return s.Password
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) GetUrl() *string {
	return s.Url
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) GetUsername() *string {
	return s.Username
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) SetId(v string) *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints {
	s.Id = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) SetPassword(v string) *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints {
	s.Password = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) SetUrl(v string) *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints {
	s.Url = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) SetUsername(v string) *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints {
	s.Username = &v
	return s
}

func (s *ListLivePackageChannelsResponseBodyLivePackageChannelsIngestEndpoints) Validate() error {
	return dara.Validate(s)
}
