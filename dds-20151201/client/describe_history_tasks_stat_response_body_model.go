// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeHistoryTasksStatResponseBodyItems) *DescribeHistoryTasksStatResponseBody
	GetItems() []*DescribeHistoryTasksStatResponseBodyItems
	SetRequestId(v string) *DescribeHistoryTasksStatResponseBody
	GetRequestId() *string
}

type DescribeHistoryTasksStatResponseBody struct {
	// The task objects.
	Items []*DescribeHistoryTasksStatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FC724D23-XXXX-XXXX-ABB1-606C935AE7FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHistoryTasksStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatResponseBody) GetItems() []*DescribeHistoryTasksStatResponseBodyItems {
	return s.Items
}

func (s *DescribeHistoryTasksStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryTasksStatResponseBody) SetItems(v []*DescribeHistoryTasksStatResponseBodyItems) *DescribeHistoryTasksStatResponseBody {
	s.Items = v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetRequestId(v string) *DescribeHistoryTasksStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHistoryTasksStatResponseBodyItems struct {
	// The task status. Valid values:
	//
	// 	- Scheduled: The task is waiting to be executed.
	//
	// 	- Running: The task is running.
	//
	// 	- Succeed: The task is successful.
	//
	// 	- Failed: The task failed.
	//
	// 	- Cancelling: The task is being terminated.
	//
	// 	- Canceled: The task has been terminated.
	//
	// 	- Waiting: The task is waiting for scheduled time.
	//
	// example:
	//
	// Succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of tasks in a specified state.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryTasksStatResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksStatResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryTasksStatResponseBodyItems) SetStatus(v string) *DescribeHistoryTasksStatResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyItems) SetTotalCount(v int32) *DescribeHistoryTasksStatResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
