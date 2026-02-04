// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenAttachedChildInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstances(v *DescribeCenAttachedChildInstancesResponseBodyChildInstances) *DescribeCenAttachedChildInstancesResponseBody
	GetChildInstances() *DescribeCenAttachedChildInstancesResponseBodyChildInstances
	SetPageNumber(v int32) *DescribeCenAttachedChildInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenAttachedChildInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenAttachedChildInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenAttachedChildInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeCenAttachedChildInstancesResponseBody struct {
	// The information about the network instances.
	ChildInstances *DescribeCenAttachedChildInstancesResponseBodyChildInstances `json:"ChildInstances,omitempty" xml:"ChildInstances,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// B2063B16-852B-5B66-B73D-4ED4D1A5E5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBody) GetChildInstances() *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	return s.ChildInstances
}

func (s *DescribeCenAttachedChildInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenAttachedChildInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenAttachedChildInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenAttachedChildInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetChildInstances(v *DescribeCenAttachedChildInstancesResponseBodyChildInstances) *DescribeCenAttachedChildInstancesResponseBody {
	s.ChildInstances = v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetPageNumber(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetPageSize(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetRequestId(v string) *DescribeCenAttachedChildInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) SetTotalCount(v int32) *DescribeCenAttachedChildInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBody) Validate() error {
	if s.ChildInstances != nil {
		if err := s.ChildInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenAttachedChildInstancesResponseBodyChildInstances struct {
	ChildInstance []*DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance `json:"ChildInstance,omitempty" xml:"ChildInstance,omitempty" type:"Repeated"`
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstances) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) GetChildInstance() []*DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	return s.ChildInstance
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) SetChildInstance(v []*DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) *DescribeCenAttachedChildInstancesResponseBodyChildInstances {
	s.ChildInstance = v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstances) Validate() error {
	if s.ChildInstance != nil {
		for _, item := range s.ChildInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the network instance was attached to the CEN instance.
	//
	// The time follows the ISO8601 standard in the YYYY-MM-DDThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-10T06:27Z
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-8vb1lu55yt9rlwgxl****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the network instance belongs.
	//
	// example:
	//
	// 1688000000000000
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// example:
	//
	// cn-zhangjiakou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// example:
	//
	// VPC
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ManagedService    *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	// The status of the network instance. Valid values:
	//
	// 	- **Attaching**: The network instance is being created on the transit router.
	//
	// 	- **Attached**: The network instance has been created on the transit router.
	//
	// 	- **Detaching**: The network instance is being deleted from the transit router.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetChildInstanceAttachTime() *string {
	return s.ChildInstanceAttachTime
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetManagedService() *string {
	return s.ManagedService
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetCenId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.CenId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceAttachTime(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceAttachTime = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceOwnerId(v int64) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceRegionId(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetChildInstanceType(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetManagedService(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.ManagedService = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) SetStatus(v string) *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance {
	s.Status = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponseBodyChildInstancesChildInstance) Validate() error {
	return dara.Validate(s)
}
