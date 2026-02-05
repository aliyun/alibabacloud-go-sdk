// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*GetScheduledInstancesResponseBodyInstances) *GetScheduledInstancesResponseBody
	GetInstances() []*GetScheduledInstancesResponseBodyInstances
	SetMessage(v string) *GetScheduledInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *GetScheduledInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *GetScheduledInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetScheduledInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScheduledInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetScheduledInstancesResponseBody
	GetTotalCount() *int64
}

type GetScheduledInstancesResponseBody struct {
	Instances []*GetScheduledInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetScheduledInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledInstancesResponseBody) GetInstances() []*GetScheduledInstancesResponseBodyInstances {
	return s.Instances
}

func (s *GetScheduledInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetScheduledInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetScheduledInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScheduledInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetScheduledInstancesResponseBody) SetInstances(v []*GetScheduledInstancesResponseBodyInstances) *GetScheduledInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *GetScheduledInstancesResponseBody) SetMessage(v string) *GetScheduledInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledInstancesResponseBody) SetPageNumber(v int64) *GetScheduledInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetScheduledInstancesResponseBody) SetPageSize(v int64) *GetScheduledInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetScheduledInstancesResponseBody) SetRequestId(v string) *GetScheduledInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledInstancesResponseBody) SetSuccess(v bool) *GetScheduledInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *GetScheduledInstancesResponseBody) SetTotalCount(v int64) *GetScheduledInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetScheduledInstancesResponseBody) Validate() error {
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

type GetScheduledInstancesResponseBodyInstances struct {
	// example:
	//
	// MySQL
	EngineType   *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceDesc *string `json:"InstanceDesc,omitempty" xml:"InstanceDesc,omitempty"`
	// example:
	//
	// rm-2zep6e5u6l2yu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetScheduledInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *GetScheduledInstancesResponseBodyInstances) GetEngineType() *string {
	return s.EngineType
}

func (s *GetScheduledInstancesResponseBodyInstances) GetInstanceDesc() *string {
	return s.InstanceDesc
}

func (s *GetScheduledInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScheduledInstancesResponseBodyInstances) GetRegion() *string {
	return s.Region
}

func (s *GetScheduledInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *GetScheduledInstancesResponseBodyInstances) SetEngineType(v string) *GetScheduledInstancesResponseBodyInstances {
	s.EngineType = &v
	return s
}

func (s *GetScheduledInstancesResponseBodyInstances) SetInstanceDesc(v string) *GetScheduledInstancesResponseBodyInstances {
	s.InstanceDesc = &v
	return s
}

func (s *GetScheduledInstancesResponseBodyInstances) SetInstanceId(v string) *GetScheduledInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *GetScheduledInstancesResponseBodyInstances) SetRegion(v string) *GetScheduledInstancesResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *GetScheduledInstancesResponseBodyInstances) SetStatus(v string) *GetScheduledInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *GetScheduledInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
