// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetPageNumber(v int32) *ListInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A8AED3C8-F57B-5D71-9A34-4A170287533F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
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

type ListInstancesResponseBodyInstances struct {
	ApplicableOperations []*string `json:"ApplicableOperations,omitempty" xml:"ApplicableOperations,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dc437bba-5a25-4bbc-b4c2-f268864bebb5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1582266750353
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// xxx
	ModifyUserName       *string   `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NluServiceParamsJson *string   `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	// example:
	//
	// Published
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnionInstanceId *string `json:"UnionInstanceId,omitempty" xml:"UnionInstanceId,omitempty"`
	UnionSource     *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetApplicableOperations() []*string {
	return s.ApplicableOperations
}

func (s *ListInstancesResponseBodyInstances) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *ListInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListInstancesResponseBodyInstances) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListInstancesResponseBodyInstances) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyInstances) GetNluServiceParamsJson() *string {
	return s.NluServiceParamsJson
}

func (s *ListInstancesResponseBodyInstances) GetNumbers() []*string {
	return s.Numbers
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) GetUnionInstanceId() *string {
	return s.UnionInstanceId
}

func (s *ListInstancesResponseBodyInstances) GetUnionSource() *string {
	return s.UnionSource
}

func (s *ListInstancesResponseBodyInstances) SetApplicableOperations(v []*string) *ListInstancesResponseBodyInstances {
	s.ApplicableOperations = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetConcurrency(v int64) *ListInstancesResponseBodyInstances {
	s.Concurrency = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetModifyTime(v int64) *ListInstancesResponseBodyInstances {
	s.ModifyTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetModifyUserName(v string) *ListInstancesResponseBodyInstances {
	s.ModifyUserName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetName(v string) *ListInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetNluServiceParamsJson(v string) *ListInstancesResponseBodyInstances {
	s.NluServiceParamsJson = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetNumbers(v []*string) *ListInstancesResponseBodyInstances {
	s.Numbers = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUnionInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.UnionInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUnionSource(v string) *ListInstancesResponseBodyInstances {
	s.UnionSource = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
