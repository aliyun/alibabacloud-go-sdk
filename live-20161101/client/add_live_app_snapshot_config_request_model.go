// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAppSnapshotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveAppSnapshotConfigRequest
	GetAppName() *string
	SetCallback(v string) *AddLiveAppSnapshotConfigRequest
	GetCallback() *string
	SetDomainName(v string) *AddLiveAppSnapshotConfigRequest
	GetDomainName() *string
	SetOssBucket(v string) *AddLiveAppSnapshotConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *AddLiveAppSnapshotConfigRequest
	GetOssEndpoint() *string
	SetOverwriteOssObject(v string) *AddLiveAppSnapshotConfigRequest
	GetOverwriteOssObject() *string
	SetOwnerId(v int64) *AddLiveAppSnapshotConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *AddLiveAppSnapshotConfigRequest
	GetSecurityToken() *string
	SetSequenceOssObject(v string) *AddLiveAppSnapshotConfigRequest
	GetSequenceOssObject() *string
	SetTimeInterval(v int32) *AddLiveAppSnapshotConfigRequest
	GetTimeInterval() *int32
}

type AddLiveAppSnapshotConfigRequest struct {
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name in the ingest URL. Otherwise, the configuration does not take effect. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name cannot start with a hyphen (-) or underscore (_). You can also specify an asterisk (\\*) as the value to match all applications.
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
	// After the configuration is complete, you can search for specific snapshots in the OSS console based on the callback information. You must create the OSS bucket in advance. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// After the configuration is complete, you can search for specific snapshots in the OSS console based on the callback information. You must configure the OSS endpoint in advance. For more information, see [Configure OSS](https://help.aliyun.com/document_detail/84932.html).
	//
	// This parameter is required.
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
	// example:
	//
	// {AppName}/{StreamName}.jpg
	OverwriteOssObject *string `json:"OverwriteOssObject,omitempty" xml:"OverwriteOssObject,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The naming format of snapshots that are stored in sequence,
	//
	// which means that a new snapshot does not overwrite the previous snapshot. You can call the [DescribeLiveStreamSnapshotInfo](https://help.aliyun.com/document_detail/2847902.html) operation to query the snapshots that were captured within a specific time period.
	//
	// 	- The name must be less than 256 bytes in length.
	//
	// 	- Only JPG images are supported.
	//
	// 	- The name can contain variables such as {AppName}, {StreamName}, {UnixTimestamp}, and {Sequence}. The name must contain at least one of the {UnixTimestamp} and {Sequence} variables.
	//
	// example:
	//
	// snapshot/{AppName}/{StreamName}/{UnixTimestamp}.jpg
	SequenceOssObject *string `json:"SequenceOssObject,omitempty" xml:"SequenceOssObject,omitempty"`
	// The interval at which snapshots are captured. Unit: seconds. Valid values: **5 to 3600**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s AddLiveAppSnapshotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppSnapshotConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAppSnapshotConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveAppSnapshotConfigRequest) GetCallback() *string {
	return s.Callback
}

func (s *AddLiveAppSnapshotConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveAppSnapshotConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *AddLiveAppSnapshotConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *AddLiveAppSnapshotConfigRequest) GetOverwriteOssObject() *string {
	return s.OverwriteOssObject
}

func (s *AddLiveAppSnapshotConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveAppSnapshotConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveAppSnapshotConfigRequest) GetSequenceOssObject() *string {
	return s.SequenceOssObject
}

func (s *AddLiveAppSnapshotConfigRequest) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *AddLiveAppSnapshotConfigRequest) SetAppName(v string) *AddLiveAppSnapshotConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetCallback(v string) *AddLiveAppSnapshotConfigRequest {
	s.Callback = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetDomainName(v string) *AddLiveAppSnapshotConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOssBucket(v string) *AddLiveAppSnapshotConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOssEndpoint(v string) *AddLiveAppSnapshotConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOverwriteOssObject(v string) *AddLiveAppSnapshotConfigRequest {
	s.OverwriteOssObject = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetOwnerId(v int64) *AddLiveAppSnapshotConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetSecurityToken(v string) *AddLiveAppSnapshotConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetSequenceOssObject(v string) *AddLiveAppSnapshotConfigRequest {
	s.SequenceOssObject = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) SetTimeInterval(v int32) *AddLiveAppSnapshotConfigRequest {
	s.TimeInterval = &v
	return s
}

func (s *AddLiveAppSnapshotConfigRequest) Validate() error {
	return dara.Validate(s)
}
