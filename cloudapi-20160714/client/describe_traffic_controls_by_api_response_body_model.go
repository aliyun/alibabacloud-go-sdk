// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficControlsByApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTrafficControlsByApiResponseBody
	GetRequestId() *string
	SetTrafficControlItems(v *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) *DescribeTrafficControlsByApiResponseBody
	GetTrafficControlItems() *DescribeTrafficControlsByApiResponseBodyTrafficControlItems
}

type DescribeTrafficControlsByApiResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned throttling policy information. It is an array consisting of TrafficControlItem data.
	TrafficControlItems *DescribeTrafficControlsByApiResponseBodyTrafficControlItems `json:"TrafficControlItems,omitempty" xml:"TrafficControlItems,omitempty" type:"Struct"`
}

func (s DescribeTrafficControlsByApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrafficControlsByApiResponseBody) GetTrafficControlItems() *DescribeTrafficControlsByApiResponseBodyTrafficControlItems {
	return s.TrafficControlItems
}

func (s *DescribeTrafficControlsByApiResponseBody) SetRequestId(v string) *DescribeTrafficControlsByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBody) SetTrafficControlItems(v *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) *DescribeTrafficControlsByApiResponseBody {
	s.TrafficControlItems = v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsByApiResponseBodyTrafficControlItems struct {
	TrafficControlItem []*DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem `json:"TrafficControlItem,omitempty" xml:"TrafficControlItem,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItems) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) GetTrafficControlItem() []*DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	return s.TrafficControlItem
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) SetTrafficControlItem(v []*DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) *DescribeTrafficControlsByApiResponseBodyTrafficControlItems {
	s.TrafficControlItem = v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem struct {
	// The binding time of the policy.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	BoundTime *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	// The ID of the throttling policy.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	TrafficControlItemId *string `json:"TrafficControlItemId,omitempty" xml:"TrafficControlItemId,omitempty"`
	// The name of the throttling policy.
	//
	// example:
	//
	// mysecret
	TrafficControlItemName *string `json:"TrafficControlItemName,omitempty" xml:"TrafficControlItemName,omitempty"`
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) GetTrafficControlItemId() *string {
	return s.TrafficControlItemId
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) GetTrafficControlItemName() *string {
	return s.TrafficControlItemName
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) SetBoundTime(v string) *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) SetTrafficControlItemId(v string) *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	s.TrafficControlItemId = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) SetTrafficControlItemName(v string) *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	s.TrafficControlItemName = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) Validate() error {
	return dara.Validate(s)
}
