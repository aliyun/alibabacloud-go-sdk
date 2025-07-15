// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRealtimeLogAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveOpenapiReserve(v string) *DescribeLiveRealtimeLogAuthorizedRequest
	GetLiveOpenapiReserve() *string
	SetOwnerId(v int64) *DescribeLiveRealtimeLogAuthorizedRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveRealtimeLogAuthorizedRequest
	GetRegionId() *string
}

type DescribeLiveRealtimeLogAuthorizedRequest struct {
	// This parameter has no practical significance.
	//
	// example:
	//
	// none
	LiveOpenapiReserve *string `json:"LiveOpenapiReserve,omitempty" xml:"LiveOpenapiReserve,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveRealtimeLogAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeLogAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) GetLiveOpenapiReserve() *string {
	return s.LiveOpenapiReserve
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) SetLiveOpenapiReserve(v string) *DescribeLiveRealtimeLogAuthorizedRequest {
	s.LiveOpenapiReserve = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) SetOwnerId(v int64) *DescribeLiveRealtimeLogAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) SetRegionId(v string) *DescribeLiveRealtimeLogAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
