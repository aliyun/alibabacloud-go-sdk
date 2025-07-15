// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAudioAuditConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateLiveAudioAuditConfigRequest
	GetAppName() *string
	SetBizType(v string) *UpdateLiveAudioAuditConfigRequest
	GetBizType() *string
	SetDomainName(v string) *UpdateLiveAudioAuditConfigRequest
	GetDomainName() *string
	SetOssBucket(v string) *UpdateLiveAudioAuditConfigRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *UpdateLiveAudioAuditConfigRequest
	GetOssEndpoint() *string
	SetOssObject(v string) *UpdateLiveAudioAuditConfigRequest
	GetOssObject() *string
	SetOwnerId(v int64) *UpdateLiveAudioAuditConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveAudioAuditConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *UpdateLiveAudioAuditConfigRequest
	GetStreamName() *string
}

type UpdateLiveAudioAuditConfigRequest struct {
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
	// The name of the recording that is stored in OSS.
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

func (s UpdateLiveAudioAuditConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateLiveAudioAuditConfigRequest) GetBizType() *string {
	return s.BizType
}

func (s *UpdateLiveAudioAuditConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveAudioAuditConfigRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *UpdateLiveAudioAuditConfigRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *UpdateLiveAudioAuditConfigRequest) GetOssObject() *string {
	return s.OssObject
}

func (s *UpdateLiveAudioAuditConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveAudioAuditConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveAudioAuditConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateLiveAudioAuditConfigRequest) SetAppName(v string) *UpdateLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetBizType(v string) *UpdateLiveAudioAuditConfigRequest {
	s.BizType = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetDomainName(v string) *UpdateLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOssBucket(v string) *UpdateLiveAudioAuditConfigRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOssEndpoint(v string) *UpdateLiveAudioAuditConfigRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOssObject(v string) *UpdateLiveAudioAuditConfigRequest {
	s.OssObject = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetOwnerId(v int64) *UpdateLiveAudioAuditConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetRegionId(v string) *UpdateLiveAudioAuditConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) SetStreamName(v string) *UpdateLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigRequest) Validate() error {
	return dara.Validate(s)
}
