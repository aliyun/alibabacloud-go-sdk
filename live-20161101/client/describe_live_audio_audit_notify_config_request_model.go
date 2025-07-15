// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAudioAuditNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveAudioAuditNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveAudioAuditNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveAudioAuditNotifyConfigRequest
	GetRegionId() *string
}

type DescribeLiveAudioAuditNotifyConfigRequest struct {
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

func (s DescribeLiveAudioAuditNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) SetDomainName(v string) *DescribeLiveAudioAuditNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) SetOwnerId(v int64) *DescribeLiveAudioAuditNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) SetRegionId(v string) *DescribeLiveAudioAuditNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
