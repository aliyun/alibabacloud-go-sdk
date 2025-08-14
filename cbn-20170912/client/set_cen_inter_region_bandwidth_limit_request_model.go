// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCenInterRegionBandwidthLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthLimit(v int64) *SetCenInterRegionBandwidthLimitRequest
	GetBandwidthLimit() *int64
	SetBandwidthType(v string) *SetCenInterRegionBandwidthLimitRequest
	GetBandwidthType() *string
	SetCenId(v string) *SetCenInterRegionBandwidthLimitRequest
	GetCenId() *string
	SetLocalRegionId(v string) *SetCenInterRegionBandwidthLimitRequest
	GetLocalRegionId() *string
	SetOppositeRegionId(v string) *SetCenInterRegionBandwidthLimitRequest
	GetOppositeRegionId() *string
	SetOwnerAccount(v string) *SetCenInterRegionBandwidthLimitRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetCenInterRegionBandwidthLimitRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetCenInterRegionBandwidthLimitRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetCenInterRegionBandwidthLimitRequest
	GetResourceOwnerId() *int64
}

type SetCenInterRegionBandwidthLimitRequest struct {
	// The maximum bandwidth value of the inter-region connection. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	BandwidthLimit *int64 `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The bandwidth allocation method. Valid values:
	//
	// **BandwidthPackage**: allocates bandwidth from a bandwidth plan.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmx****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the local region.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query regions where you can attach network instances to a CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	LocalRegionId *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	// The ID of the peer region.
	//
	// This parameter is required.
	//
	// example:
	//
	// us-west-1
	OppositeRegionId     *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetCenInterRegionBandwidthLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitRequest) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetBandwidthLimit() *int64 {
	return s.BandwidthLimit
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetCenId() *string {
	return s.CenId
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetLocalRegionId() *string {
	return s.LocalRegionId
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetCenInterRegionBandwidthLimitRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetBandwidthLimit(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.BandwidthLimit = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetBandwidthType(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.BandwidthType = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetCenId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.CenId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetLocalRegionId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.LocalRegionId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetOppositeRegionId(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.OppositeRegionId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetOwnerAccount(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetOwnerId(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetResourceOwnerAccount(v string) *SetCenInterRegionBandwidthLimitRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) SetResourceOwnerId(v int64) *SetCenInterRegionBandwidthLimitRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitRequest) Validate() error {
	return dara.Validate(s)
}
