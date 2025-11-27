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
	// The queried tasks.
	Items []*DescribeHistoryTasksStatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
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
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHistoryTasksStatResponseBodyItems struct {
	// The status of the task. Valid values:
	//
	// 	- **Scheduled**
	//
	// 	- **Running**
	//
	// 	- **Succeed**
	//
	// 	- **Failed**
	//
	// 	- **Cancelling**
	//
	// 	- **Canceled**
	//
	// 	- **Waiting**
	//
	// example:
	//
	// Succeed,Waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 20
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
