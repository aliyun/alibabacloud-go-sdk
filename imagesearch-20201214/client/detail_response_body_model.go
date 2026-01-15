// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *DetailResponseBodyInstance) *DetailResponseBody
	GetInstance() *DetailResponseBodyInstance
	SetRequestId(v string) *DetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetailResponseBody
	GetSuccess() *bool
}

type DetailResponseBody struct {
	// The details about the instance.
	Instance *DetailResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetailResponseBody) GoString() string {
	return s.String()
}

func (s *DetailResponseBody) GetInstance() *DetailResponseBodyInstance {
	return s.Instance
}

func (s *DetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetailResponseBody) SetInstance(v *DetailResponseBodyInstance) *DetailResponseBody {
	s.Instance = v
	return s
}

func (s *DetailResponseBody) SetRequestId(v string) *DetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailResponseBody) SetSuccess(v bool) *DetailResponseBody {
	s.Success = &v
	return s
}

func (s *DetailResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetailResponseBodyInstance struct {
	// The capacity of the plan. Unit: × 10,000 images.
	//
	// example:
	//
	// 10
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// imagesearchName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of queries per second supported by the plan.
	//
	// example:
	//
	// 1
	Qps *int32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The information about the region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The Image Search model.
	//
	// 0: commodity search. 1: generic image search.
	//
	// example:
	//
	// 0
	ServiceType *int32 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The number of images.
	//
	// example:
	//
	// 10063
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The time when the instance was created. Unit: milliseconds.
	//
	// example:
	//
	// 1620382716000
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	// The time when the instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1623081600000
	UtcExpireTime *string `json:"UtcExpireTime,omitempty" xml:"UtcExpireTime,omitempty"`
}

func (s DetailResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s DetailResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DetailResponseBodyInstance) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DetailResponseBodyInstance) GetName() *string {
	return s.Name
}

func (s *DetailResponseBodyInstance) GetQps() *int32 {
	return s.Qps
}

func (s *DetailResponseBodyInstance) GetRegion() *string {
	return s.Region
}

func (s *DetailResponseBodyInstance) GetServiceType() *int32 {
	return s.ServiceType
}

func (s *DetailResponseBodyInstance) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DetailResponseBodyInstance) GetUtcCreate() *string {
	return s.UtcCreate
}

func (s *DetailResponseBodyInstance) GetUtcExpireTime() *string {
	return s.UtcExpireTime
}

func (s *DetailResponseBodyInstance) SetCapacity(v int32) *DetailResponseBodyInstance {
	s.Capacity = &v
	return s
}

func (s *DetailResponseBodyInstance) SetName(v string) *DetailResponseBodyInstance {
	s.Name = &v
	return s
}

func (s *DetailResponseBodyInstance) SetQps(v int32) *DetailResponseBodyInstance {
	s.Qps = &v
	return s
}

func (s *DetailResponseBodyInstance) SetRegion(v string) *DetailResponseBodyInstance {
	s.Region = &v
	return s
}

func (s *DetailResponseBodyInstance) SetServiceType(v int32) *DetailResponseBodyInstance {
	s.ServiceType = &v
	return s
}

func (s *DetailResponseBodyInstance) SetTotalCount(v int64) *DetailResponseBodyInstance {
	s.TotalCount = &v
	return s
}

func (s *DetailResponseBodyInstance) SetUtcCreate(v string) *DetailResponseBodyInstance {
	s.UtcCreate = &v
	return s
}

func (s *DetailResponseBodyInstance) SetUtcExpireTime(v string) *DetailResponseBodyInstance {
	s.UtcExpireTime = &v
	return s
}

func (s *DetailResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
