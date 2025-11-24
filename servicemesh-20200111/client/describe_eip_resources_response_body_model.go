// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipList(v []*DescribeEipResourcesResponseBodyEipList) *DescribeEipResourcesResponseBody
	GetEipList() []*DescribeEipResourcesResponseBodyEipList
	SetPageResult(v *DescribeEipResourcesResponseBodyPageResult) *DescribeEipResourcesResponseBody
	GetPageResult() *DescribeEipResourcesResponseBodyPageResult
	SetRequestId(v string) *DescribeEipResourcesResponseBody
	GetRequestId() *string
}

type DescribeEipResourcesResponseBody struct {
	// The EIPs.
	EipList []*DescribeEipResourcesResponseBodyEipList `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	// The pagination information.
	PageResult *DescribeEipResourcesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEipResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponseBody) GetEipList() []*DescribeEipResourcesResponseBodyEipList {
	return s.EipList
}

func (s *DescribeEipResourcesResponseBody) GetPageResult() *DescribeEipResourcesResponseBodyPageResult {
	return s.PageResult
}

func (s *DescribeEipResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipResourcesResponseBody) SetEipList(v []*DescribeEipResourcesResponseBodyEipList) *DescribeEipResourcesResponseBody {
	s.EipList = v
	return s
}

func (s *DescribeEipResourcesResponseBody) SetPageResult(v *DescribeEipResourcesResponseBodyPageResult) *DescribeEipResourcesResponseBody {
	s.PageResult = v
	return s
}

func (s *DescribeEipResourcesResponseBody) SetRequestId(v string) *DescribeEipResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipResourcesResponseBody) Validate() error {
	if s.EipList != nil {
		for _, item := range s.EipList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEipResourcesResponseBodyEipList struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-2zeerraiwb7ujxscd****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The type of the resource that is associated with the EIP. Valid values:
	//
	// 	- `EcsInstance`: an ECS instance in a VPC
	//
	// 	- `SlbInstance`: a Server Load Balancer (SLB) instance in a VPC
	//
	// 	- `Nat`: a NAT gateway
	//
	// 	- `HaVip`: a high-availability virtual IP address (HAVIP)
	//
	// 	- `NetworkInterface`: a secondary elastic network interface (ENI)
	//
	// Default value: `EcsInstance`.
	//
	// example:
	//
	// SlbInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IP address of the EIP.
	//
	// example:
	//
	// 120.xx.xx.112
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The status of the EIP. Valid values:
	//
	// 	- `Associating`: The EIP is being associated with a resource.
	//
	// 	- `Unassociating`: The EIP is being disassociated from a resource.
	//
	// 	- `InUse`: The EIP is associated with a resource.
	//
	// 	- `Available`: The EIP is available.
	//
	// example:
	//
	// InUse
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEipResourcesResponseBodyEipList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipResourcesResponseBodyEipList) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponseBodyEipList) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEipResourcesResponseBodyEipList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEipResourcesResponseBodyEipList) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeEipResourcesResponseBodyEipList) GetStatus() *string {
	return s.Status
}

func (s *DescribeEipResourcesResponseBodyEipList) SetAllocationId(v string) *DescribeEipResourcesResponseBodyEipList {
	s.AllocationId = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) SetInstanceType(v string) *DescribeEipResourcesResponseBodyEipList {
	s.InstanceType = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) SetIpAddress(v string) *DescribeEipResourcesResponseBodyEipList {
	s.IpAddress = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) SetStatus(v string) *DescribeEipResourcesResponseBodyEipList {
	s.Status = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyEipList) Validate() error {
	return dara.Validate(s)
}

type DescribeEipResourcesResponseBodyPageResult struct {
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
	// The number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEipResourcesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipResourcesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponseBodyPageResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEipResourcesResponseBodyPageResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipResourcesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEipResourcesResponseBodyPageResult) SetPageNumber(v int32) *DescribeEipResourcesResponseBodyPageResult {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyPageResult) SetPageSize(v int32) *DescribeEipResourcesResponseBodyPageResult {
	s.PageSize = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyPageResult) SetTotalCount(v int32) *DescribeEipResourcesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *DescribeEipResourcesResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}
