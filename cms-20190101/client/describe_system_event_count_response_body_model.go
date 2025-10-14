// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSystemEventCountResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSystemEventCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSystemEventCountResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSystemEventCountResponseBody
	GetSuccess() *string
	SetSystemEventCounts(v *DescribeSystemEventCountResponseBodySystemEventCounts) *DescribeSystemEventCountResponseBody
	GetSystemEventCounts() *DescribeSystemEventCountResponseBodySystemEventCounts
}

type DescribeSystemEventCountResponseBody struct {
	// The HTTP status codes.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7A7B776-0ACE-5A93-9B07-DE8008D9CCDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the system event.
	SystemEventCounts *DescribeSystemEventCountResponseBodySystemEventCounts `json:"SystemEventCounts,omitempty" xml:"SystemEventCounts,omitempty" type:"Struct"`
}

func (s DescribeSystemEventCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSystemEventCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSystemEventCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemEventCountResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSystemEventCountResponseBody) GetSystemEventCounts() *DescribeSystemEventCountResponseBodySystemEventCounts {
	return s.SystemEventCounts
}

func (s *DescribeSystemEventCountResponseBody) SetCode(v string) *DescribeSystemEventCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetMessage(v string) *DescribeSystemEventCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetRequestId(v string) *DescribeSystemEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetSuccess(v string) *DescribeSystemEventCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetSystemEventCounts(v *DescribeSystemEventCountResponseBodySystemEventCounts) *DescribeSystemEventCountResponseBody {
	s.SystemEventCounts = v
	return s
}

func (s *DescribeSystemEventCountResponseBody) Validate() error {
	if s.SystemEventCounts != nil {
		if err := s.SystemEventCounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSystemEventCountResponseBodySystemEventCounts struct {
	SystemEventCount []*DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount `json:"SystemEventCount,omitempty" xml:"SystemEventCount,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventCountResponseBodySystemEventCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventCountResponseBodySystemEventCounts) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponseBodySystemEventCounts) GetSystemEventCount() []*DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	return s.SystemEventCount
}

func (s *DescribeSystemEventCountResponseBodySystemEventCounts) SetSystemEventCount(v []*DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) *DescribeSystemEventCountResponseBodySystemEventCounts {
	s.SystemEventCount = v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCounts) Validate() error {
	if s.SystemEventCount != nil {
		for _, item := range s.SystemEventCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount struct {
	// The description of the system event.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 17285****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ECS-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The level of the system event. Valid values:
	//
	// 	- Critical
	//
	// 	- Warn
	//
	// 	- Info
	//
	// example:
	//
	// Info
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the system event.
	//
	// example:
	//
	// Instance:StateChange
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of times that the system event has occurred.
	//
	// example:
	//
	// 3
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The name of the cloud service in which the system event occurred.
	//
	// example:
	//
	// ECS
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-rj99xc6cptkk64ml****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the system event.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the system event occurred. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1635993751000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetContent() *string {
	return s.Content
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetLevel() *string {
	return s.Level
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetName() *string {
	return s.Name
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetNum() *int64 {
	return s.Num
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetProduct() *string {
	return s.Product
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetStatus() *string {
	return s.Status
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GetTime() *int64 {
	return s.Time
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetContent(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Content = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetGroupId(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetInstanceName(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.InstanceName = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetLevel(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetName(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetNum(v int64) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Num = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetProduct(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetRegionId(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.RegionId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetResourceId(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.ResourceId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetStatus(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetTime(v int64) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Time = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) Validate() error {
	return dara.Validate(s)
}
