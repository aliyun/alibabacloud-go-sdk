// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenPrivateZoneRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenPrivateZoneRoutesResponseBody
	GetCenId() *string
	SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenPrivateZoneRoutesResponseBody
	GetPageSize() *int32
	SetPrivateZoneDnsServers(v string) *DescribeCenPrivateZoneRoutesResponseBody
	GetPrivateZoneDnsServers() *string
	SetPrivateZoneInfos(v *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) *DescribeCenPrivateZoneRoutesResponseBody
	GetPrivateZoneInfos() *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos
	SetRequestId(v string) *DescribeCenPrivateZoneRoutesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenPrivateZoneRoutesResponseBody
	GetTotalCount() *int32
}

type DescribeCenPrivateZoneRoutesResponseBody struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
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
	// The IP address of the DNS server used by PrivateZone.
	//
	// example:
	//
	// 100.100.XX.XX/32,100.100.XX.XX/32
	PrivateZoneDnsServers *string `json:"PrivateZoneDnsServers,omitempty" xml:"PrivateZoneDnsServers,omitempty"`
	// The detailed configuration of PrivateZone.
	PrivateZoneInfos *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos `json:"PrivateZoneInfos,omitempty" xml:"PrivateZoneInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 461EC1B5-04A8-4706-8764-8F5BCEF48A6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetPrivateZoneDnsServers() *string {
	return s.PrivateZoneDnsServers
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetPrivateZoneInfos() *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	return s.PrivateZoneInfos
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetCenId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPageNumber(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPageSize(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneDnsServers(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneDnsServers = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetPrivateZoneInfos(v *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) *DescribeCenPrivateZoneRoutesResponseBody {
	s.PrivateZoneInfos = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetRequestId(v string) *DescribeCenPrivateZoneRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) SetTotalCount(v int32) *DescribeCenPrivateZoneRoutesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBody) Validate() error {
	if s.PrivateZoneInfos != nil {
		if err := s.PrivateZoneInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos struct {
	PrivateZoneInfo []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo `json:"PrivateZoneInfo,omitempty" xml:"PrivateZoneInfo,omitempty" type:"Repeated"`
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) GetPrivateZoneInfo() []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	return s.PrivateZoneInfo
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) SetPrivateZoneInfo(v []*DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos {
	s.PrivateZoneInfo = v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfos) Validate() error {
	if s.PrivateZoneInfo != nil {
		for _, item := range s.PrivateZoneInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo struct {
	// The ID of the region where PrivateZone is accessed.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	// The ID of the region where PrivateZone is deployed.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	// The ID of the VPC that is associated with PrivateZone.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnvo****
	HostVpcId *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	// The status of PrivateZone. Valid values:
	//
	// 	- **Creating**: being created
	//
	// 	- **Active**: available
	//
	// 	- **Deleting**: being deleted
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GoString() string {
	return s.String()
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GetHostVpcId() *string {
	return s.HostVpcId
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetAccessRegionId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetHostRegionId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.HostRegionId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetHostVpcId(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.HostVpcId = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) SetStatus(v string) *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo {
	s.Status = &v
	return s
}

func (s *DescribeCenPrivateZoneRoutesResponseBodyPrivateZoneInfosPrivateZoneInfo) Validate() error {
	return dara.Validate(s)
}
