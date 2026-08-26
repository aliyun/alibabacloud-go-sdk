// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveStreamRecordIndexFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetAppName() *string
	SetDomainName(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetDomainName() *string
	SetEndTime(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetEndTime() *string
	SetEndTimeIncluded(v bool) *CreateLiveStreamRecordIndexFilesRequest
	GetEndTimeIncluded() *bool
	SetOssBucket(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetOssEndpoint() *string
	SetOssObject(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetOssObject() *string
	SetOwnerId(v int64) *CreateLiveStreamRecordIndexFilesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetSecurityToken() *string
	SetStartTime(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetStartTime() *string
	SetStreamName(v string) *CreateLiveStreamRecordIndexFilesRequest
	GetStreamName() *string
}

type CreateLiveStreamRecordIndexFilesRequest struct {
	// The name of the application to which the stream belongs. The AppName must match the AppName in the ingest URL for the template to take effect. To match all AppName values, set this parameter to *.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streamer streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the index file. TS files uploaded before this time are included in the index file. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to include the end time. If you set this parameter to true, the system attempts to include one additional TS file so that the created index file fully covers the period between StartTime and EndTime.
	//
	// example:
	//
	// false
	EndTimeIncluded *bool `json:"EndTimeIncluded,omitempty" xml:"EndTimeIncluded,omitempty"`
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-oss-****.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The name of the recording file stored in OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// {AppName}/{StreamName}/{Date}/{Hour}/{Minute}_{Second}.m3u8
	OssObject     *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The start time of the index file. TS files uploaded after this time are included in the index file. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name. The StreamName must match the StreamName in the ingest URL for the template to take effect. To match all StreamName values, set this parameter to *.
	//
	// The stream must have had actual stream ingest activity under the specified DomainName and AppName. Otherwise, the InvalidStream.NotFound error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s CreateLiveStreamRecordIndexFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamRecordIndexFilesRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetEndTimeIncluded() *bool {
	return s.EndTimeIncluded
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetOssObject() *string {
	return s.OssObject
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateLiveStreamRecordIndexFilesRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetAppName(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.AppName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetDomainName(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.DomainName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetEndTime(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetEndTimeIncluded(v bool) *CreateLiveStreamRecordIndexFilesRequest {
	s.EndTimeIncluded = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOssBucket(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOssEndpoint(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOssObject(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.OssObject = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetOwnerId(v int64) *CreateLiveStreamRecordIndexFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetSecurityToken(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetStartTime(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.StartTime = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) SetStreamName(v string) *CreateLiveStreamRecordIndexFilesRequest {
	s.StreamName = &v
	return s
}

func (s *CreateLiveStreamRecordIndexFilesRequest) Validate() error {
	return dara.Validate(s)
}
