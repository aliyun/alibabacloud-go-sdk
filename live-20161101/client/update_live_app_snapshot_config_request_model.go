// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAppSnapshotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLiveAppSnapshotConfigRequest
	GetAppName() *string
	SetCallback(v string) *UpdateLiveAppSnapshotConfigRequest
	GetCallback() *string
	SetDomainName(v string) *UpdateLiveAppSnapshotConfigRequest
	GetDomainName() *string
	SetOssBucket(v string) *UpdateLiveAppSnapshotConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *UpdateLiveAppSnapshotConfigRequest
	GetOssEndpoint() *string
	SetOverwriteOssObject(v string) *UpdateLiveAppSnapshotConfigRequest
	GetOverwriteOssObject() *string
	SetOwnerId(v int64) *UpdateLiveAppSnapshotConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *UpdateLiveAppSnapshotConfigRequest
	GetSecurityToken() *string
	SetSequenceOssObject(v string) *UpdateLiveAppSnapshotConfigRequest
	GetSequenceOssObject() *string
	SetTimeInterval(v int32) *UpdateLiveAppSnapshotConfigRequest
	GetTimeInterval() *int32
}

type UpdateLiveAppSnapshotConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The callback URL that is used to receive notifications about snapshot capture.
	//
	// example:
	//
	// https://learn.aliyundoc.com
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
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
	// The naming format of snapshots that are stored in the overwrite mode, which means that a new snapshot overwrites the previous snapshot.
	//
	// 	- The name must be less than 256 bytes in length.
	//
	// 	- Only JPG images are supported.
	//
	// 	- The name can contain variables such as {AppName} and {StreamName}.
	//
	// 	- A value of hyphen (-) indicates that this parameter is deleted.
	//
	// example:
	//
	// {liveApp****}/{liveStream****}.jpg
	OverwriteOssObject *string `json:"OverwriteOssObject,omitempty" xml:"OverwriteOssObject,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The naming format of snapshots that are stored in sequence, which means that a new snapshot does not overwrite the previous snapshot. You can call the [DescribeLiveStreamSnapshotInfo](https://help.aliyun.com/document_detail/2847902.html) operation to query the snapshots that were captured within a specific time period.
	//
	// 	- The name must be less than 256 bytes in length.
	//
	// 	- Only JPG images are supported.
	//
	// 	- The name can contain variables such as {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}. The name must contain at least one of the {UnixTimestamp} and {Sequence} variables.
	//
	// 	- A value of hyphen (-) indicates that this parameter is deleted.
	//
	// example:
	//
	// snapshot/{liveApp****}/{liveStream****}/{UnixTimestamp****}.jpg
	SequenceOssObject *string `json:"SequenceOssObject,omitempty" xml:"SequenceOssObject,omitempty"`
	// The interval at which snapshots are captured. Valid values: **5 to 3600**. Unit: seconds.
	//
	// example:
	//
	// 5
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s UpdateLiveAppSnapshotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetCallback() *string {
	return s.Callback
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetOverwriteOssObject() *string {
	return s.OverwriteOssObject
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetSequenceOssObject() *string {
	return s.SequenceOssObject
}

func (s *UpdateLiveAppSnapshotConfigRequest) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetAppName(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetCallback(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.Callback = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetDomainName(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOssBucket(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOssEndpoint(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOverwriteOssObject(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.OverwriteOssObject = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetOwnerId(v int64) *UpdateLiveAppSnapshotConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetSecurityToken(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetSequenceOssObject(v string) *UpdateLiveAppSnapshotConfigRequest {
	s.SequenceOssObject = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) SetTimeInterval(v int32) *UpdateLiveAppSnapshotConfigRequest {
	s.TimeInterval = &v
	return s
}

func (s *UpdateLiveAppSnapshotConfigRequest) Validate() error {
	return dara.Validate(s)
}
