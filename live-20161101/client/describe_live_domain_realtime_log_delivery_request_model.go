// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainRealtimeLogDeliveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainRealtimeLogDeliveryRequest
	GetRegionId() *string
}

type DescribeLiveDomainRealtimeLogDeliveryRequest struct {
	// The streaming domain.
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

func (s DescribeLiveDomainRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) SetDomainName(v string) *DescribeLiveDomainRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DescribeLiveDomainRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) SetRegionId(v string) *DescribeLiveDomainRealtimeLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
