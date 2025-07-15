// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordIndexFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordIndexInfo(v *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) *DescribeLiveStreamRecordIndexFileResponseBody
	GetRecordIndexInfo() *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo
	SetRequestId(v string) *DescribeLiveStreamRecordIndexFileResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamRecordIndexFileResponseBody struct {
	// The information about the index file.
	RecordIndexInfo *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo `json:"RecordIndexInfo,omitempty" xml:"RecordIndexInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5EBF2AC3-4B73-40A5-8B32-83F49D5F035E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileResponseBody) GetRecordIndexInfo() *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	return s.RecordIndexInfo
}

func (s *DescribeLiveStreamRecordIndexFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamRecordIndexFileResponseBody) SetRecordIndexInfo(v *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) *DescribeLiveStreamRecordIndexFileResponseBody {
	s.RecordIndexInfo = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBody) SetRequestId(v string) *DescribeLiveStreamRecordIndexFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the index file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-05-27T09:40:56Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The recording length. Unit: seconds.
	//
	// example:
	//
	// 588.849
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time of the index file. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-05-25T05:47:11Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The video format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The video height.
	//
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// example:
	//
	// cn-oss-****.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The name of the storage file in OSS.
	//
	// example:
	//
	// liveObject****
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	// The ID of the index file.
	//
	// example:
	//
	// c4d7f0a4-b506-43f9-8de3-07732c3f****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The URL of the index file.
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// The start time of the index file. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T05:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The video width.
	//
	// example:
	//
	// 640
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetDuration() *float32 {
	return s.Duration
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetFormat() *string {
	return s.Format
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetOssObject() *string {
	return s.OssObject
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetAppName(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetCreateTime(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetDomainName(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetDuration(v float32) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.Duration = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetEndTime(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetFormat(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.Format = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetHeight(v int32) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetOssBucket(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetOssEndpoint(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetOssObject(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetRecordId(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.RecordId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetRecordUrl(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.RecordUrl = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetStartTime(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetStreamName(v string) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) SetWidth(v int32) *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFileResponseBodyRecordIndexInfo) Validate() error {
	return dara.Validate(s)
}
