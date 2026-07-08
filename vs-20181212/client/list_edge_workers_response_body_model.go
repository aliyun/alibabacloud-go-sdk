// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeWorkersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListEdgeWorkersResponseBodyInstances) *ListEdgeWorkersResponseBody
	GetInstances() []*ListEdgeWorkersResponseBodyInstances
	SetPageNumber(v int32) *ListEdgeWorkersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeWorkersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEdgeWorkersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListEdgeWorkersResponseBody
	GetTotalCount() *int64
}

type ListEdgeWorkersResponseBody struct {
	// The list of payloads.
	Instances []*ListEdgeWorkersResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number of the query. The value starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 100. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEdgeWorkersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeWorkersResponseBody) GetInstances() []*ListEdgeWorkersResponseBodyInstances {
	return s.Instances
}

func (s *ListEdgeWorkersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeWorkersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeWorkersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeWorkersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEdgeWorkersResponseBody) SetInstances(v []*ListEdgeWorkersResponseBodyInstances) *ListEdgeWorkersResponseBody {
	s.Instances = v
	return s
}

func (s *ListEdgeWorkersResponseBody) SetPageNumber(v int32) *ListEdgeWorkersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeWorkersResponseBody) SetPageSize(v int32) *ListEdgeWorkersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeWorkersResponseBody) SetRequestId(v string) *ListEdgeWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeWorkersResponseBody) SetTotalCount(v int64) *ListEdgeWorkersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeWorkersResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeWorkersResponseBodyInstances struct {
	// The creation time.
	//
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// hive-58mq6jynvgxxmlid3pt39x6gk-0
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// as-d135ca4425c24b99b79cd0b6c552cac9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The package ID.
	//
	// example:
	//
	// pk-db3394401cc8403f85e4d72d99b52449
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The specification.
	//
	// example:
	//
	// crs.xic.s1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status.
	//
	// example:
	//
	// Idle
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEdgeWorkersResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeWorkersResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListEdgeWorkersResponseBodyInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListEdgeWorkersResponseBodyInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListEdgeWorkersResponseBodyInstances) GetHiveId() *string {
	return s.HiveId
}

func (s *ListEdgeWorkersResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEdgeWorkersResponseBodyInstances) GetPlanId() *string {
	return s.PlanId
}

func (s *ListEdgeWorkersResponseBodyInstances) GetSpec() *string {
	return s.Spec
}

func (s *ListEdgeWorkersResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListEdgeWorkersResponseBodyInstances) SetCreationTime(v string) *ListEdgeWorkersResponseBodyInstances {
	s.CreationTime = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) SetExpireTime(v string) *ListEdgeWorkersResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) SetHiveId(v string) *ListEdgeWorkersResponseBodyInstances {
	s.HiveId = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) SetInstanceId(v string) *ListEdgeWorkersResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) SetPlanId(v string) *ListEdgeWorkersResponseBodyInstances {
	s.PlanId = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) SetSpec(v string) *ListEdgeWorkersResponseBodyInstances {
	s.Spec = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) SetStatus(v string) *ListEdgeWorkersResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListEdgeWorkersResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
