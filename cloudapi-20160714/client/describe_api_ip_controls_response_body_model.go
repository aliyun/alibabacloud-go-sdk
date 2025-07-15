// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiIpControlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiIpControls(v *DescribeApiIpControlsResponseBodyApiIpControls) *DescribeApiIpControlsResponseBody
	GetApiIpControls() *DescribeApiIpControlsResponseBodyApiIpControls
	SetPageNumber(v int32) *DescribeApiIpControlsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiIpControlsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiIpControlsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiIpControlsResponseBody
	GetTotalCount() *int32
}

type DescribeApiIpControlsResponseBody struct {
	// The information about the ACLs. The information is an array of ApiIpControlItem data.
	ApiIpControls *DescribeApiIpControlsResponseBodyApiIpControls `json:"ApiIpControls,omitempty" xml:"ApiIpControls,omitempty" type:"Struct"`
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
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiIpControlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiIpControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponseBody) GetApiIpControls() *DescribeApiIpControlsResponseBodyApiIpControls {
	return s.ApiIpControls
}

func (s *DescribeApiIpControlsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiIpControlsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiIpControlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiIpControlsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiIpControlsResponseBody) SetApiIpControls(v *DescribeApiIpControlsResponseBodyApiIpControls) *DescribeApiIpControlsResponseBody {
	s.ApiIpControls = v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetPageNumber(v int32) *DescribeApiIpControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetPageSize(v int32) *DescribeApiIpControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetRequestId(v string) *DescribeApiIpControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetTotalCount(v int32) *DescribeApiIpControlsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiIpControlsResponseBodyApiIpControls struct {
	ApiIpControlItem []*DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem `json:"ApiIpControlItem,omitempty" xml:"ApiIpControlItem,omitempty" type:"Repeated"`
}

func (s DescribeApiIpControlsResponseBodyApiIpControls) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiIpControlsResponseBodyApiIpControls) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponseBodyApiIpControls) GetApiIpControlItem() []*DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	return s.ApiIpControlItem
}

func (s *DescribeApiIpControlsResponseBodyApiIpControls) SetApiIpControlItem(v []*DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) *DescribeApiIpControlsResponseBodyApiIpControls {
	s.ApiIpControlItem = v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControls) Validate() error {
	return dara.Validate(s)
}

type DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem struct {
	// The ID of the API.
	//
	// example:
	//
	// 46fbb52840d146f186e38e8e70fc8c90
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// testapi
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The time of binding.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	BoundTime *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	// The ID of the ACL.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// testControlName
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
}

func (s DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GetBoundTime() *string {
	return s.BoundTime
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GetIpControlId() *string {
	return s.IpControlId
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GetIpControlName() *string {
	return s.IpControlName
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetApiId(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetApiName(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetBoundTime(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetIpControlId(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.IpControlId = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetIpControlName(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.IpControlName = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) Validate() error {
	return dara.Validate(s)
}
