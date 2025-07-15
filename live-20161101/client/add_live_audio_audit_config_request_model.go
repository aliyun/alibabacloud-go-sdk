// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAudioAuditConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddLiveAudioAuditConfigRequest
	GetAppName() *string
	SetBizType(v string) *AddLiveAudioAuditConfigRequest
	GetBizType() *string
	SetDomainName(v string) *AddLiveAudioAuditConfigRequest
	GetDomainName() *string
	SetOssBucket(v string) *AddLiveAudioAuditConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *AddLiveAudioAuditConfigRequest
	GetOssEndpoint() *string
	SetOssObject(v string) *AddLiveAudioAuditConfigRequest
	GetOssObject() *string
	SetOwnerId(v int64) *AddLiveAudioAuditConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveAudioAuditConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *AddLiveAudioAuditConfigRequest
	GetStreamName() *string
}

type AddLiveAudioAuditConfigRequest struct {
	// The name of the application to which the live stream belongs. The value of this parameter must be the same as the application name in the ingest URL. Otherwise, the configuration does not take effect. The application name is case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business type. You can specify a model. The default value is the domain name.
	//
	// example:
	//
	// example.edu
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the OSS bucket in which the recording is stored.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of OSS.
	//
	// example:
	//
	// cn-oss-****.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The name of the recording stored in OSS.
	//
	// example:
	//
	// liveObject****
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. The value of this parameter must be the same as the stream name in the ingest URL. Otherwise, the configuration does not take effect. The stream name is case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s AddLiveAudioAuditConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLiveAudioAuditConfigRequest) GetBizType() *string {
	return s.BizType
}

func (s *AddLiveAudioAuditConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveAudioAuditConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *AddLiveAudioAuditConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *AddLiveAudioAuditConfigRequest) GetOssObject() *string {
	return s.OssObject
}

func (s *AddLiveAudioAuditConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveAudioAuditConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveAudioAuditConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddLiveAudioAuditConfigRequest) SetAppName(v string) *AddLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetBizType(v string) *AddLiveAudioAuditConfigRequest {
	s.BizType = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetDomainName(v string) *AddLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOssBucket(v string) *AddLiveAudioAuditConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOssEndpoint(v string) *AddLiveAudioAuditConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOssObject(v string) *AddLiveAudioAuditConfigRequest {
	s.OssObject = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetOwnerId(v int64) *AddLiveAudioAuditConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetRegionId(v string) *AddLiveAudioAuditConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) SetStreamName(v string) *AddLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

func (s *AddLiveAudioAuditConfigRequest) Validate() error {
	return dara.Validate(s)
}
