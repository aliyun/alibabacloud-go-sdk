// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAudioAuditConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveAudioAuditConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveAudioAuditConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveAudioAuditConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveAudioAuditConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLiveAudioAuditConfigRequest
	GetStreamName() *string
}

type DescribeLiveAudioAuditConfigRequest struct {
	// The name of the application to which the live stream belongs.
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
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveAudioAuditConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveAudioAuditConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveAudioAuditConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveAudioAuditConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveAudioAuditConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveAudioAuditConfigRequest) SetAppName(v string) *DescribeLiveAudioAuditConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) SetDomainName(v string) *DescribeLiveAudioAuditConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) SetOwnerId(v int64) *DescribeLiveAudioAuditConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) SetRegionId(v string) *DescribeLiveAudioAuditConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) SetStreamName(v string) *DescribeLiveAudioAuditConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigRequest) Validate() error {
	return dara.Validate(s)
}
