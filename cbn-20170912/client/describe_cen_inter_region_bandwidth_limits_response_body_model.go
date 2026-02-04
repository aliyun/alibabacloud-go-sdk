// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenInterRegionBandwidthLimitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenInterRegionBandwidthLimits(v *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) *DescribeCenInterRegionBandwidthLimitsResponseBody
	GetCenInterRegionBandwidthLimits() *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits
	SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody
	GetTotalCount() *int32
}

type DescribeCenInterRegionBandwidthLimitsResponseBody struct {
	// A list of inter-region connections.
	CenInterRegionBandwidthLimits *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits `json:"CenInterRegionBandwidthLimits,omitempty" xml:"CenInterRegionBandwidthLimits,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A30C665-8766-5AAA-9274-C97380E2D850
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) GetCenInterRegionBandwidthLimits() *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	return s.CenInterRegionBandwidthLimits
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetCenInterRegionBandwidthLimits(v *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.CenInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetPageNumber(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetPageSize(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetRequestId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) SetTotalCount(v int32) *DescribeCenInterRegionBandwidthLimitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBody) Validate() error {
	if s.CenInterRegionBandwidthLimits != nil {
		if err := s.CenInterRegionBandwidthLimits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits struct {
	CenInterRegionBandwidthLimit []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit `json:"CenInterRegionBandwidthLimit,omitempty" xml:"CenInterRegionBandwidthLimit,omitempty" type:"Repeated"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) GetCenInterRegionBandwidthLimit() []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	return s.CenInterRegionBandwidthLimit
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) SetCenInterRegionBandwidthLimit(v []*DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits {
	s.CenInterRegionBandwidthLimit = v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimits) Validate() error {
	if s.CenInterRegionBandwidthLimit != nil {
		for _, item := range s.CenInterRegionBandwidthLimit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit struct {
	// The maximum bandwidth of the inter-region connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	BandwidthLimit *int64 `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// cenbwp-uenczwb592fnvv****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The bandwidth allocation method. Valid values:
	//
	// 	- **BandwidthPackage**: allocates bandwidth from a bandwidth plan.
	//
	// 	- **DataTransfer**: bandwidth is billed based on the pay-by-data-transfer metering method.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-pfa6ugf3xl0qsd****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The IDs of the local and peer regions.
	//
	// example:
	//
	// china_china
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	// The ID of the local region.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// ccn-cn-shanghai
	LocalRegionId *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	// The ID of the peer region.
	//
	// example:
	//
	// cn-hangzhou
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	// The bandwidth status of the inter-region connection. Valid values:
	//
	// 	- **Active**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GoString() string {
	return s.String()
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetBandwidthLimit() *int64 {
	return s.BandwidthLimit
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetGeographicSpanId() *string {
	return s.GeographicSpanId
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetLocalRegionId() *string {
	return s.LocalRegionId
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthLimit(v int64) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthPackageId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetBandwidthType(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.BandwidthType = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetCenId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.CenId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetGeographicSpanId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetLocalRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetOppositeRegionId(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) SetStatus(v string) *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit {
	s.Status = &v
	return s
}

func (s *DescribeCenInterRegionBandwidthLimitsResponseBodyCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit) Validate() error {
	return dara.Validate(s)
}
