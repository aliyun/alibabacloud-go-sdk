// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveSnapshotDetectPornConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveSnapshotDetectPornConfigRequest
	GetAppName() *string
	SetDomainName(v string) *AddLiveSnapshotDetectPornConfigRequest
	GetDomainName() *string
	SetInterval(v int32) *AddLiveSnapshotDetectPornConfigRequest
	GetInterval() *int32
	SetOssBucket(v string) *AddLiveSnapshotDetectPornConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *AddLiveSnapshotDetectPornConfigRequest
	GetOssEndpoint() *string
	SetOssObject(v string) *AddLiveSnapshotDetectPornConfigRequest
	GetOssObject() *string
	SetOwnerId(v int64) *AddLiveSnapshotDetectPornConfigRequest
	GetOwnerId() *int64
	SetScene(v []*string) *AddLiveSnapshotDetectPornConfigRequest
	GetScene() []*string
	SetSecurityToken(v string) *AddLiveSnapshotDetectPornConfigRequest
	GetSecurityToken() *string
}

type AddLiveSnapshotDetectPornConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// testApp
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
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the OSS bucket.
	//
	// After the review is complete, you can search for specific violations in the OSS console based on the callback information. You must create the OSS bucket in advance. For more information, see [Configure content moderation](https://help.aliyun.com/document_detail/199449.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// livebucket
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the Object Storage Service (OSS) bucket.
	//
	// After the review is complete, you can search for specific violations in the OSS console based on the callback information. You must configure the OSS endpoint in advance. For more information, see [Configure content moderation](https://help.aliyun.com/document_detail/199449.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-oss-****.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The name of the snapshot that stores violations such as pornographic content and politically sensitive content.
	//
	// example:
	//
	// record/{AppName}/{StreamName}/{Sequence}.jpg
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Scene list detection.
	//
	// example:
	//
	// live
	Scene         []*string `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddLiveSnapshotDetectPornConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveSnapshotDetectPornConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetOssObject() *string {
	return s.OssObject
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetScene() []*string {
	return s.Scene
}

func (s *AddLiveSnapshotDetectPornConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetAppName(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetDomainName(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetInterval(v int32) *AddLiveSnapshotDetectPornConfigRequest {
	s.Interval = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOssBucket(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOssEndpoint(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOssObject(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.OssObject = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetOwnerId(v int64) *AddLiveSnapshotDetectPornConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetScene(v []*string) *AddLiveSnapshotDetectPornConfigRequest {
	s.Scene = v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) SetSecurityToken(v string) *AddLiveSnapshotDetectPornConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddLiveSnapshotDetectPornConfigRequest) Validate() error {
	return dara.Validate(s)
}
