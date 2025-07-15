// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotDetectPornConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetAppName() *string
	SetDomainName(v string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetDomainName() *string
	SetInterval(v int32) *UpdateLiveSnapshotDetectPornConfigRequest
	GetInterval() *int32
	SetOssBucket(v string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetOssEndpoint() *string
	SetOssObject(v string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetOssObject() *string
	SetOwnerId(v int64) *UpdateLiveSnapshotDetectPornConfigRequest
	GetOwnerId() *int64
	SetScene(v []*string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetScene() []*string
	SetSecurityToken(v string) *UpdateLiveSnapshotDetectPornConfigRequest
	GetSecurityToken() *string
}

type UpdateLiveSnapshotDetectPornConfigRequest struct {
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name in the ingest URL. Otherwise, the configuration does not take effect. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name cannot start with a hyphen (-) or underscore (_). You can also specify an asterisk (\\*) as the value to match all applications.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The interval at which snapshots are captured from the live stream. Valid values: **5 to 3600**. Unit: seconds.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// cn-oss-****.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The name of the snapshot that stores violations such as pornographic content and politically sensitive content.
	//
	// example:
	//
	// {liveApp****}/{liveStream****}/{Date}/{Hour}/{Minute}_{Second}.jpg
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The moderation scenario array.
	//
	// example:
	//
	// porn
	Scene         []*string `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLiveSnapshotDetectPornConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetOssObject() *string {
	return s.OssObject
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetScene() []*string {
	return s.Scene
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetInterval(v int32) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.Interval = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOssBucket(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOssEndpoint(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOssObject(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OssObject = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetOwnerId(v int64) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetScene(v []*string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.Scene = v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *UpdateLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLiveSnapshotDetectPornConfigRequest) Validate() error {
	return dara.Validate(s)
}
