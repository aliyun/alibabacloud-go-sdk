// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiTrafficControlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiTrafficControls(v *DescribeApiTrafficControlsResponseBodyApiTrafficControls) *DescribeApiTrafficControlsResponseBody
	GetApiTrafficControls() *DescribeApiTrafficControlsResponseBodyApiTrafficControls
	SetPageNumber(v int32) *DescribeApiTrafficControlsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiTrafficControlsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiTrafficControlsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiTrafficControlsResponseBody
	GetTotalCount() *int32
}

type DescribeApiTrafficControlsResponseBody struct {
	// The returned throttling policy information. It is an array consisting of ApiTrafficControlItem data.
	ApiTrafficControls *DescribeApiTrafficControlsResponseBodyApiTrafficControls `json:"ApiTrafficControls,omitempty" xml:"ApiTrafficControls,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiTrafficControlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponseBody) GetApiTrafficControls() *DescribeApiTrafficControlsResponseBodyApiTrafficControls {
	return s.ApiTrafficControls
}

func (s *DescribeApiTrafficControlsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiTrafficControlsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiTrafficControlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiTrafficControlsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiTrafficControlsResponseBody) SetApiTrafficControls(v *DescribeApiTrafficControlsResponseBodyApiTrafficControls) *DescribeApiTrafficControlsResponseBody {
	s.ApiTrafficControls = v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetPageNumber(v int32) *DescribeApiTrafficControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetPageSize(v int32) *DescribeApiTrafficControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetRequestId(v string) *DescribeApiTrafficControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetTotalCount(v int32) *DescribeApiTrafficControlsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) Validate() error {
	if s.ApiTrafficControls != nil {
		if err := s.ApiTrafficControls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiTrafficControlsResponseBodyApiTrafficControls struct {
	ApiTrafficControlItem []*DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem `json:"ApiTrafficControlItem,omitempty" xml:"ApiTrafficControlItem,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControls) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControls) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControls) GetApiTrafficControlItem() []*DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	return s.ApiTrafficControlItem
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControls) SetApiTrafficControlItem(v []*DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) *DescribeApiTrafficControlsResponseBodyApiTrafficControls {
	s.ApiTrafficControlItem = v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControls) Validate() error {
	if s.ApiTrafficControlItem != nil {
		for _, item := range s.ApiTrafficControlItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem struct {
	// The ID of the API.
	//
	// example:
	//
	// 46fbb52840d146f186e38e8e70fc8c90
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// API operation
	//
	// example:
	//
	// testapi
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The binding time of the throttling policy.
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
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	// The name of the throttling policy.
	//
	// example:
	//
	// backendsignature
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GetTrafficControlName() *string {
	return s.TrafficControlName
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetApiId(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetApiName(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetBoundTime(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetTrafficControlId(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetTrafficControlName(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.TrafficControlName = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) Validate() error {
	return dara.Validate(s)
}
