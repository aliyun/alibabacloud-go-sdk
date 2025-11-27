// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeHistoryEventsStatResponseBodyItems) *DescribeHistoryEventsStatResponseBody
	GetItems() []*DescribeHistoryEventsStatResponseBodyItems
	SetRequestId(v string) *DescribeHistoryEventsStatResponseBody
	GetRequestId() *string
}

type DescribeHistoryEventsStatResponseBody struct {
	// The event.
	Items []*DescribeHistoryEventsStatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E52666CC-330E-418A-8E5B-A19E3FB42D13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHistoryEventsStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsStatResponseBody) GetItems() []*DescribeHistoryEventsStatResponseBodyItems {
	return s.Items
}

func (s *DescribeHistoryEventsStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryEventsStatResponseBody) SetItems(v []*DescribeHistoryEventsStatResponseBodyItems) *DescribeHistoryEventsStatResponseBody {
	s.Items = v
	return s
}

func (s *DescribeHistoryEventsStatResponseBody) SetRequestId(v string) *DescribeHistoryEventsStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryEventsStatResponseBody) Validate() error {
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

type DescribeHistoryEventsStatResponseBodyItems struct {
	// The system event category. For more information, see [View the event history of an ApsaraDB RDS instance](https://help.aliyun.com/document_detail/129759.html).
	//
	// example:
	//
	// Exception
	EventCategory *string `json:"EventCategory,omitempty" xml:"EventCategory,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryEventsStatResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsStatResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsStatResponseBodyItems) GetEventCategory() *string {
	return s.EventCategory
}

func (s *DescribeHistoryEventsStatResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryEventsStatResponseBodyItems) SetEventCategory(v string) *DescribeHistoryEventsStatResponseBodyItems {
	s.EventCategory = &v
	return s
}

func (s *DescribeHistoryEventsStatResponseBodyItems) SetTotalCount(v int32) *DescribeHistoryEventsStatResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryEventsStatResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
