// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamsNotifyUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveStreamsNotifyUrlConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveStreamsNotifyUrlConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveStreamsNotifyUrlConfigRequest
	GetRegionId() *string
}

type DeleteLiveStreamsNotifyUrlConfigRequest struct {
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLiveStreamsNotifyUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DeleteLiveStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *DeleteLiveStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) SetRegionId(v string) *DeleteLiveStreamsNotifyUrlConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
