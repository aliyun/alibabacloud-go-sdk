// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpControlPolicyItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlPolicyItems(v *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) *DescribeIpControlPolicyItemsResponseBody
	GetIpControlPolicyItems() *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems
	SetPageNumber(v int32) *DescribeIpControlPolicyItemsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpControlPolicyItemsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIpControlPolicyItemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpControlPolicyItemsResponseBody
	GetTotalCount() *int32
}

type DescribeIpControlPolicyItemsResponseBody struct {
	// The information about policies. The information is an array of IpControlPolicyItem data.
	IpControlPolicyItems *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems `json:"IpControlPolicyItems,omitempty" xml:"IpControlPolicyItems,omitempty" type:"Struct"`
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
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpControlPolicyItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponseBody) GetIpControlPolicyItems() *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems {
	return s.IpControlPolicyItems
}

func (s *DescribeIpControlPolicyItemsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpControlPolicyItemsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpControlPolicyItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpControlPolicyItemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetIpControlPolicyItems(v *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) *DescribeIpControlPolicyItemsResponseBody {
	s.IpControlPolicyItems = v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetPageNumber(v int32) *DescribeIpControlPolicyItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetPageSize(v int32) *DescribeIpControlPolicyItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetRequestId(v string) *DescribeIpControlPolicyItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetTotalCount(v int32) *DescribeIpControlPolicyItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems struct {
	IpControlPolicyItem []*DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem `json:"IpControlPolicyItem,omitempty" xml:"IpControlPolicyItem,omitempty" type:"Repeated"`
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) GetIpControlPolicyItem() []*DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	return s.IpControlPolicyItem
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) SetIpControlPolicyItem(v []*DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems {
	s.IpControlPolicyItem = v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem struct {
	// The ID of the application.
	//
	// example:
	//
	// 11112
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The IP addresses or CIDR blocks.
	//
	// example:
	//
	// 113.125.XX.XX;101.11.XX.XX
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	// The time when the policy was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-17T06:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the policy was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-17T06:25:13Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// P151617000829241
	PolicyItemId *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GetAppId() *string {
	return s.AppId
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GetCidrIp() *string {
	return s.CidrIp
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GetPolicyItemId() *string {
	return s.PolicyItemId
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetAppId(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.AppId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetCidrIp(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.CidrIp = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetCreateTime(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetModifiedTime(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetPolicyItemId(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.PolicyItemId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) Validate() error {
	return dara.Validate(s)
}
