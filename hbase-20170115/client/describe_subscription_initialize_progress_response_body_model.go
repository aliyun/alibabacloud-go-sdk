// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInitializeProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeSubscriptionInitializeProgressResponseBodyItems) *DescribeSubscriptionInitializeProgressResponseBody
	GetItems() []*DescribeSubscriptionInitializeProgressResponseBodyItems
	SetPageNumber(v int32) *DescribeSubscriptionInitializeProgressResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSubscriptionInitializeProgressResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSubscriptionInitializeProgressResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeSubscriptionInitializeProgressResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSubscriptionInitializeProgressResponseBody struct {
	Items            []*DescribeSubscriptionInitializeProgressResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber       *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                                     `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                                     `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) GetItems() []*DescribeSubscriptionInitializeProgressResponseBodyItems {
	return s.Items
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetItems(v []*DescribeSubscriptionInitializeProgressResponseBodyItems) *DescribeSubscriptionInitializeProgressResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetPageNumber(v int32) *DescribeSubscriptionInitializeProgressResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetPageRecordCount(v int32) *DescribeSubscriptionInitializeProgressResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetRequestId(v string) *DescribeSubscriptionInitializeProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) SetTotalRecordCount(v int32) *DescribeSubscriptionInitializeProgressResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBody) Validate() error {
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

type DescribeSubscriptionInitializeProgressResponseBodyItems struct {
	FinishTime     *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
}

func (s DescribeSubscriptionInitializeProgressResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInitializeProgressResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetFinishTime(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.FinishTime = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetProgress(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetStatus(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) SetSubscriptionId(v string) *DescribeSubscriptionInitializeProgressResponseBodyItems {
	s.SubscriptionId = &v
	return s
}

func (s *DescribeSubscriptionInitializeProgressResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
