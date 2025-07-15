// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveStreamRecordIndexFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordInfo(v *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) *CreateLiveStreamRecordIndexFilesResponseBody
	GetRecordInfo() *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo
	SetRequestId(v string) *CreateLiveStreamRecordIndexFilesResponseBody
	GetRequestId() *string
}

type CreateLiveStreamRecordIndexFilesResponseBody struct {
	// The recording configuration.
	RecordInfo *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo `json:"RecordInfo,omitempty" xml:"RecordInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 550439A3-F8EC-4CA2-BB62-B9DB43EEEF30
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLiveStreamRecordIndexFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesResponseBody) GetRecordInfo() *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	return s.RecordInfo
}

func (s *CreateLiveStreamRecordIndexFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveStreamRecordIndexFilesResponseBody) SetRecordInfo(v *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) *CreateLiveStreamRecordIndexFilesResponseBody {
	s.RecordInfo = v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBody) SetRequestId(v string) *CreateLiveStreamRecordIndexFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the index file was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
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
	// 20
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time of the index file. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T07:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The height of the video.
	//
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the OSS bucket.
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
	// The name of the recording that is stored in OSS.
	//
	// example:
	//
	// liveObject****.m3u8
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	// The ID of the index file.
	//
	// example:
	//
	// c4d7f0a4-b506-43f9-8de3-07732c3f****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The URL of the M3U8 index file.
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// The start time of the index file. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T07:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The width of the video.
	//
	// example:
	//
	// 640
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetAppName() *string {
	return s.AppName
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetDuration() *float32 {
	return s.Duration
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetHeight() *int32 {
	return s.Height
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetOssBucket() *string {
	return s.OssBucket
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetOssObject() *string {
	return s.OssObject
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetRecordId() *string {
	return s.RecordId
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) GetWidth() *int32 {
	return s.Width
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetAppName(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.AppName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetCreateTime(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.CreateTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetDomainName(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.DomainName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetDuration(v float32) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.Duration = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetEndTime(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.EndTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetHeight(v int32) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.Height = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetOssBucket(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.OssBucket = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetOssEndpoint(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.OssEndpoint = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetOssObject(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.OssObject = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetRecordId(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.RecordId = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetRecordUrl(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.RecordUrl = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetStartTime(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.StartTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetStreamName(v string) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.StreamName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) SetWidth(v int32) *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo {
	s.Width = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesResponseBodyRecordInfo) Validate() error {
	return dara.Validate(s)
}
