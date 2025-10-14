// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionIdResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionIdResources(v *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) *DescribeEnsRegionIdResourceResponseBody
	GetEnsRegionIdResources() *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources
	SetPageNumber(v int32) *DescribeEnsRegionIdResourceResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsRegionIdResourceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEnsRegionIdResourceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEnsRegionIdResourceResponseBody
	GetTotalCount() *int32
}

type DescribeEnsRegionIdResourceResponseBody struct {
	// The returned data. For more information, see EnsRegionIdResources in sample JSON responses.
	EnsRegionIdResources *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources `json:"EnsRegionIdResources,omitempty" xml:"EnsRegionIdResources,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 476600B1-C9E2-4245-A26F-DC7EA8071425
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried nodes.
	//
	// example:
	//
	// 58
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnsRegionIdResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseBody) GetEnsRegionIdResources() *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources {
	return s.EnsRegionIdResources
}

func (s *DescribeEnsRegionIdResourceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsRegionIdResourceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsRegionIdResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsRegionIdResourceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetEnsRegionIdResources(v *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) *DescribeEnsRegionIdResourceResponseBody {
	s.EnsRegionIdResources = v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetPageNumber(v int32) *DescribeEnsRegionIdResourceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetPageSize(v int32) *DescribeEnsRegionIdResourceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetRequestId(v string) *DescribeEnsRegionIdResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) SetTotalCount(v int32) *DescribeEnsRegionIdResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBody) Validate() error {
	if s.EnsRegionIdResources != nil {
		if err := s.EnsRegionIdResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources struct {
	EnsRegionIdResource []*DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource `json:"EnsRegionIdResource,omitempty" xml:"EnsRegionIdResource,omitempty" type:"Repeated"`
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) GetEnsRegionIdResource() []*DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	return s.EnsRegionIdResource
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) SetEnsRegionIdResource(v []*DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources {
	s.EnsRegionIdResource = v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResources) Validate() error {
	if s.EnsRegionIdResource != nil {
		for _, item := range s.EnsRegionIdResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource struct {
	// The region. Set the value to West.
	//
	// example:
	//
	// West
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The code of the region.
	//
	// example:
	//
	// 300100
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// The date when the transaction was processed.
	//
	// example:
	//
	// 2019-10-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-hangzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Nanjing Mobile
	EnsRegionIdName *string `json:"EnsRegionIdName,omitempty" xml:"EnsRegionIdName,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 10
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The public bandwidth of the instance. Unit: Bits/s.
	//
	// example:
	//
	// 100
	InternetBandwidth *int64 `json:"InternetBandwidth,omitempty" xml:"InternetBandwidth,omitempty"`
	// The ISP. Valid values:
	//
	// 	- cmcc: China Mobile
	//
	// 	- unicom: China Unicom
	//
	// 	- telecom: China Telecom
	//
	// 	- multiCarrier: multi-line ISP
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 10
	VCpu *int32 `json:"VCpu,omitempty" xml:"VCpu,omitempty"`
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetArea() *string {
	return s.Area
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetAreaCode() *string {
	return s.AreaCode
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetBizDate() *string {
	return s.BizDate
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetEnsRegionIdName() *string {
	return s.EnsRegionIdName
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetInternetBandwidth() *int64 {
	return s.InternetBandwidth
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) GetVCpu() *int32 {
	return s.VCpu
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetArea(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.Area = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetAreaCode(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.AreaCode = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetBizDate(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.BizDate = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetEnsRegionId(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetEnsRegionIdName(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.EnsRegionIdName = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetInstanceCount(v int32) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.InstanceCount = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetInternetBandwidth(v int64) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.InternetBandwidth = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetIsp(v string) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.Isp = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) SetVCpu(v int32) *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource {
	s.VCpu = &v
	return s
}

func (s *DescribeEnsRegionIdResourceResponseBodyEnsRegionIdResourcesEnsRegionIdResource) Validate() error {
	return dara.Validate(s)
}
