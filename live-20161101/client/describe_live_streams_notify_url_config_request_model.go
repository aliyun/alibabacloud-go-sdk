// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsNotifyUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamsNotifyUrlConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamsNotifyUrlConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamsNotifyUrlConfigRequest
	GetRegionId() *string
}

type DescribeLiveStreamsNotifyUrlConfigRequest struct {
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

func (s DescribeLiveStreamsNotifyUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) SetDomainName(v string) *DescribeLiveStreamsNotifyUrlConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) SetOwnerId(v int64) *DescribeLiveStreamsNotifyUrlConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) SetRegionId(v string) *DescribeLiveStreamsNotifyUrlConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
