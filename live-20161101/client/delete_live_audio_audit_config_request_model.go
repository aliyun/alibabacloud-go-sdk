// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAudioAuditConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveAudioAuditConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveAudioAuditConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveAudioAuditConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveAudioAuditConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLiveAudioAuditConfigRequest
	GetStreamName() *string
}

type DeleteLiveAudioAuditConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// >  A value of asterisk (\\*) specifies all applications under the domain name.
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
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream.
	//
	// >  A value of asterisk (\\*) specifies all live streams in the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveAudioAuditConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveAudioAuditConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveAudioAuditConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAudioAuditConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveAudioAuditConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveAudioAuditConfigRequest) SetAppName(v string) *DeleteLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) SetDomainName(v string) *DeleteLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) SetOwnerId(v int64) *DeleteLiveAudioAuditConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) SetRegionId(v string) *DeleteLiveAudioAuditConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) SetStreamName(v string) *DeleteLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigRequest) Validate() error {
	return dara.Validate(s)
}
