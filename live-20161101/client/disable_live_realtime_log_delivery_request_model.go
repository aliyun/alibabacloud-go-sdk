// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLiveRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DisableLiveRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DisableLiveRealtimeLogDeliveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableLiveRealtimeLogDeliveryRequest
	GetRegionId() *string
}

type DisableLiveRealtimeLogDeliveryRequest struct {
	// The streaming domain for which you want to suspend real-time log delivery. Separate multiple streaming domains with commas (,).
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

func (s DisableLiveRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DisableLiveRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DisableLiveRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableLiveRealtimeLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *DisableLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DisableLiveRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryRequest) SetRegionId(v string) *DisableLiveRealtimeLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
