// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAudioAuditNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveAudioAuditNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveAudioAuditNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveAudioAuditNotifyConfigRequest
	GetRegionId() *string
}

type DeleteLiveAudioAuditNotifyConfigRequest struct {
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
}

func (s DeleteLiveAudioAuditNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *DeleteLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) SetOwnerId(v int64) *DeleteLiveAudioAuditNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) SetRegionId(v string) *DeleteLiveAudioAuditNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
