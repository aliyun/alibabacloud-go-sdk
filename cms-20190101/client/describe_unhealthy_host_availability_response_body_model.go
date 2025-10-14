// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnhealthyHostAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeUnhealthyHostAvailabilityResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeUnhealthyHostAvailabilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeUnhealthyHostAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeUnhealthyHostAvailabilityResponseBody
	GetSuccess() *bool
	SetUnhealthyList(v *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) *DescribeUnhealthyHostAvailabilityResponseBody
	GetUnhealthyList() *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList
}

type DescribeUnhealthyHostAvailabilityResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ACBDBB40-DFB6-4F4C-8957-51FFB233969C
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The unhealthy instances that are detected by the specified availability monitoring tasks.
	UnhealthyList *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList `json:"UnhealthyList,omitempty" xml:"UnhealthyList,omitempty" type:"Struct"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) GetUnhealthyList() *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList {
	return s.UnhealthyList
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetCode(v string) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetMessage(v string) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetRequestId(v string) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetSuccess(v bool) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetUnhealthyList(v *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.UnhealthyList = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) Validate() error {
	if s.UnhealthyList != nil {
		if err := s.UnhealthyList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList struct {
	NodeTaskInstance []*DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance `json:"NodeTaskInstance,omitempty" xml:"NodeTaskInstance,omitempty" type:"Repeated"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) GetNodeTaskInstance() []*DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance {
	return s.NodeTaskInstance
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) SetNodeTaskInstance(v []*DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList {
	s.NodeTaskInstance = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) Validate() error {
	if s.NodeTaskInstance != nil {
		for _, item := range s.NodeTaskInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance struct {
	// The ID of the availability monitoring task.
	//
	// example:
	//
	// 123456
	Id           *int64                                                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceList *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) GetId() *int64 {
	return s.Id
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) GetInstanceList() *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList {
	return s.InstanceList
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) SetId(v int64) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance {
	s.Id = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) SetInstanceList(v *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance {
	s.InstanceList = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) Validate() error {
	if s.InstanceList != nil {
		if err := s.InstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) GetString_() []*string {
	return s.String_
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) SetString_(v []*string) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList {
	s.String_ = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) Validate() error {
	return dara.Validate(s)
}
