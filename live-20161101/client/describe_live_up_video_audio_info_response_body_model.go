// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUpVideoAudioInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveUpVideoAudioInfoResponseBody
	GetRequestId() *string
	SetUpItems(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItems) *DescribeLiveUpVideoAudioInfoResponseBody
	GetUpItems() *DescribeLiveUpVideoAudioInfoResponseBodyUpItems
}

type DescribeLiveUpVideoAudioInfoResponseBody struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request ID.
	UpItems *DescribeLiveUpVideoAudioInfoResponseBodyUpItems `json:"UpItems,omitempty" xml:"UpItems,omitempty" type:"Struct"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveUpVideoAudioInfoResponseBody) GetUpItems() *DescribeLiveUpVideoAudioInfoResponseBodyUpItems {
	return s.UpItems
}

func (s *DescribeLiveUpVideoAudioInfoResponseBody) SetRequestId(v string) *DescribeLiveUpVideoAudioInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBody) SetUpItems(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItems) *DescribeLiveUpVideoAudioInfoResponseBody {
	s.UpItems = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItems struct {
	PublishItem []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem `json:"PublishItem,omitempty" xml:"PublishItem,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItems) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItems) GetPublishItem() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	return s.PublishItem
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItems) SetPublishItem(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) *DescribeLiveUpVideoAudioInfoResponseBodyUpItems {
	s.PublishItem = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItems) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem struct {
	// The details about the audio and video data of the stream ingest occurrences.
	AacHeaders *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders `json:"AacHeaders,omitempty" xml:"AacHeaders,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the application to which the ingested stream belongs.
	AudioBitRate *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate `json:"AudioBitRate,omitempty" xml:"AudioBitRate,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	AudioFrames *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames `json:"AudioFrames,omitempty" xml:"AudioFrames,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	AudioInterval *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval `json:"AudioInterval,omitempty" xml:"AudioInterval,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	AudioStamps *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps `json:"AudioStamps,omitempty" xml:"AudioStamps,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	AvcHeaders *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders `json:"AvcHeaders,omitempty" xml:"AvcHeaders,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	//
	// example:
	//
	// H264/AAC
	CodecInfo *string `json:"CodecInfo,omitempty" xml:"CodecInfo,omitempty"`
	// The audio and video encoding information.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ingest domain.
	ErrorFlags *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags `json:"ErrorFlags,omitempty" xml:"ErrorFlags,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	//
	// example:
	//
	// -
	PublishInterval *string `json:"PublishInterval,omitempty" xml:"PublishInterval,omitempty"`
	// The stream ingest duration. Unit: seconds. A hyphen (-) indicates that the stream is being ingested and the duration cannot be returned.
	//
	// example:
	//
	// cn397
	PublishIp *string `json:"PublishIp,omitempty" xml:"PublishIp,omitempty"`
	// The IP address of the stream ingest client.
	//
	// example:
	//
	// 1
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The stream ingest status. A value of 1 indicates that the stream is being ingested. A value of 0 indicates that the stream was ingested.
	//
	// example:
	//
	// 2015-12-10T15:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The start time of stream ingest. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T15:10:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The end time of stream ingest. The time is displayed in UTC.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The name of the stream.
	//
	// example:
	//
	// 2.-395_37261_9848098_1538080899396
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// The unique ID of each stream ingest occurrence.
	VideoAndAudioStamp *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp `json:"VideoAndAudioStamp,omitempty" xml:"VideoAndAudioStamp,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	VideoBitRate *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate `json:"VideoBitRate,omitempty" xml:"VideoBitRate,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	VideoFrames *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames `json:"VideoFrames,omitempty" xml:"VideoFrames,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	VideoInterval *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval `json:"VideoInterval,omitempty" xml:"VideoInterval,omitempty" type:"Struct"`
	// The metric value at a granularity of seconds at the query time.
	VideoStamps *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps `json:"VideoStamps,omitempty" xml:"VideoStamps,omitempty" type:"Struct"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAacHeaders() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders {
	return s.AacHeaders
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAudioBitRate() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate {
	return s.AudioBitRate
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAudioFrames() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames {
	return s.AudioFrames
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAudioInterval() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval {
	return s.AudioInterval
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAudioStamps() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps {
	return s.AudioStamps
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetAvcHeaders() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders {
	return s.AvcHeaders
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetCodecInfo() *string {
	return s.CodecInfo
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetErrorFlags() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags {
	return s.ErrorFlags
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetPublishInterval() *string {
	return s.PublishInterval
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetPublishIp() *string {
	return s.PublishIp
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetPublishTime() *string {
	return s.PublishTime
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetUniqueId() *string {
	return s.UniqueId
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetVideoAndAudioStamp() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp {
	return s.VideoAndAudioStamp
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetVideoBitRate() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate {
	return s.VideoBitRate
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetVideoFrames() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames {
	return s.VideoFrames
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetVideoInterval() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval {
	return s.VideoInterval
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) GetVideoStamps() *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps {
	return s.VideoStamps
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAacHeaders(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AacHeaders = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAppName(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AppName = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAudioBitRate(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AudioBitRate = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAudioFrames(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AudioFrames = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAudioInterval(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AudioInterval = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAudioStamps(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AudioStamps = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetAvcHeaders(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.AvcHeaders = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetCodecInfo(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.CodecInfo = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetDomainName(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetErrorFlags(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.ErrorFlags = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetPublishInterval(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.PublishInterval = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetPublishIp(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.PublishIp = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetPublishStatus(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.PublishStatus = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetPublishTime(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.PublishTime = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetStopTime(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.StopTime = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetStreamName(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetUniqueId(v string) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.UniqueId = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetVideoAndAudioStamp(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.VideoAndAudioStamp = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetVideoBitRate(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.VideoBitRate = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetVideoFrames(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.VideoFrames = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetVideoInterval(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.VideoInterval = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) SetVideoStamps(v *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem {
	s.VideoStamps = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItem) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders struct {
	AacHeaders []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders `json:"AacHeaders,omitempty" xml:"AacHeaders,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders) GetAacHeaders() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders {
	return s.AacHeaders
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders) SetAacHeaders(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders {
	s.AacHeaders = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders struct {
	// The number of AAC headers in the audio.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 20
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAacHeadersAacHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate struct {
	AudioBitRate []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate `json:"AudioBitRate,omitempty" xml:"AudioBitRate,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate) GetAudioBitRate() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate {
	return s.AudioBitRate
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate) SetAudioBitRate(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate {
	s.AudioBitRate = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRate) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate struct {
	// The bitrate of the audio. Unit: bit/s.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 24552
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioBitRateAudioBitRate) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames struct {
	AudioFrames []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames `json:"AudioFrames,omitempty" xml:"AudioFrames,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames) GetAudioFrames() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames {
	return s.AudioFrames
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames) SetAudioFrames(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames {
	s.AudioFrames = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFrames) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames struct {
	// The frame rate of the audio. Unit: frames.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 23
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioFramesAudioFrames) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval struct {
	AudioInterval []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval `json:"AudioInterval,omitempty" xml:"AudioInterval,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval) GetAudioInterval() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval {
	return s.AudioInterval
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval) SetAudioInterval(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval {
	s.AudioInterval = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval struct {
	// The maximum audio frame interval. Unit: milliseconds.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 254
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioIntervalAudioInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps struct {
	AudioStamps []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps `json:"AudioStamps,omitempty" xml:"AudioStamps,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps) GetAudioStamps() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps {
	return s.AudioStamps
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps) SetAudioStamps(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps {
	s.AudioStamps = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStamps) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps struct {
	// The audio timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 725053422
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAudioStampsAudioStamps) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders struct {
	AvcHeaders []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders `json:"AvcHeaders,omitempty" xml:"AvcHeaders,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders) GetAvcHeaders() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders {
	return s.AvcHeaders
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders) SetAvcHeaders(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders {
	s.AvcHeaders = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders struct {
	// The number of AVC headers in the audio.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 11
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemAvcHeadersAvcHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags struct {
	ErrorFlags []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags `json:"ErrorFlags,omitempty" xml:"ErrorFlags,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags) GetErrorFlags() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags {
	return s.ErrorFlags
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags) SetErrorFlags(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags {
	s.ErrorFlags = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlags) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags struct {
	// The number of times the error code that indicates interrupted stream ingest was returned.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 0
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemErrorFlagsErrorFlags) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp struct {
	VAStamp []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp `json:"V_AStamp,omitempty" xml:"V_AStamp,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp) GetVAStamp() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp {
	return s.VAStamp
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp) SetVAStamp(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp {
	s.VAStamp = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStamp) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp struct {
	// The difference between the audio and video timestamps. Unit: milliseconds.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 359
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoAndAudioStampVAStamp) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate struct {
	VideoBitRate []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate `json:"VideoBitRate,omitempty" xml:"VideoBitRate,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate) GetVideoBitRate() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate {
	return s.VideoBitRate
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate) SetVideoBitRate(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate {
	s.VideoBitRate = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRate) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate struct {
	// The bitrate of the video. Unit: bit/s.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 3970160
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoBitRateVideoBitRate) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames struct {
	VideoFrames []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames `json:"VideoFrames,omitempty" xml:"VideoFrames,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames) GetVideoFrames() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames {
	return s.VideoFrames
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames) SetVideoFrames(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames {
	s.VideoFrames = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFrames) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames struct {
	// The frame rate of the video. Unit: frames.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 29
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoFramesVideoFrames) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval struct {
	VideoInterval []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval `json:"VideoInterval,omitempty" xml:"VideoInterval,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval) GetVideoInterval() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval {
	return s.VideoInterval
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval) SetVideoInterval(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval {
	s.VideoInterval = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval struct {
	// The maximum video frame interval. Unit: milliseconds.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 278
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoIntervalVideoInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps struct {
	VideoStamps []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps `json:"VideoStamps,omitempty" xml:"VideoStamps,omitempty" type:"Repeated"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps) GetVideoStamps() []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps {
	return s.VideoStamps
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps) SetVideoStamps(v []*DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps {
	s.VideoStamps = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStamps) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps struct {
	// The video timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1538134750408
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The query time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 725053781
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) GetTime() *int64 {
	return s.Time
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) GetValue() *int32 {
	return s.Value
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) SetTime(v int64) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps {
	s.Time = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) SetValue(v int32) *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps {
	s.Value = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponseBodyUpItemsPublishItemVideoStampsVideoStamps) Validate() error {
	return dara.Validate(s)
}
