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
	CenId                   *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChildInstanceAttachTime *string `json:"ChildInstanceAttachTime,omitempty" xml:"ChildInstanceAttachTime,omitempty"`
	ChildInstanceId         *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId    *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId   *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType       *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ManagedService          *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
