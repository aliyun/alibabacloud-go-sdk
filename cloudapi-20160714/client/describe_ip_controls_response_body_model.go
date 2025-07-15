// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpControlsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlInfos(v *DescribeIpControlsResponseBodyIpControlInfos) *DescribeIpControlsResponseBody
	GetIpControlInfos() *DescribeIpControlsResponseBodyIpControlInfos
	SetPageNumber(v int32) *DescribeIpControlsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpControlsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIpControlsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpControlsResponseBody
	GetTotalCount() *int32
}

type DescribeIpControlsResponseBody struct {
	// The information about the ACL. The information is an array that consists of IpControlInfo data. The information does not include specific policies.
	IpControlInfos *DescribeIpControlsResponseBodyIpControlInfos `json:"IpControlInfos,omitempty" xml:"IpControlInfos,omitempty" type:"Struct"`
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

func (s DescribeIpControlsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponseBody) GetIpControlInfos() *DescribeIpControlsResponseBodyIpControlInfos {
	return s.IpControlInfos
}

func (s *DescribeIpControlsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpControlsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpControlsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpControlsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpControlsResponseBody) SetIpControlInfos(v *DescribeIpControlsResponseBodyIpControlInfos) *DescribeIpControlsResponseBody {
	s.IpControlInfos = v
	return s
}

func (s *DescribeIpControlsResponseBody) SetPageNumber(v int32) *DescribeIpControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlsResponseBody) SetPageSize(v int32) *DescribeIpControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlsResponseBody) SetRequestId(v string) *DescribeIpControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpControlsResponseBody) SetTotalCount(v int32) *DescribeIpControlsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpControlsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIpControlsResponseBodyIpControlInfos struct {
	IpControlInfo []*DescribeIpControlsResponseBodyIpControlInfosIpControlInfo `json:"IpControlInfo,omitempty" xml:"IpControlInfo,omitempty" type:"Repeated"`
}

func (s DescribeIpControlsResponseBodyIpControlInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlsResponseBodyIpControlInfos) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponseBodyIpControlInfos) GetIpControlInfo() []*DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	return s.IpControlInfo
}

func (s *DescribeIpControlsResponseBodyIpControlInfos) SetIpControlInfo(v []*DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) *DescribeIpControlsResponseBodyIpControlInfos {
	s.IpControlInfo = v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeIpControlsResponseBodyIpControlInfosIpControlInfo struct {
	// The time when the ACL was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-17T05:48:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the ACL.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ACL.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// testControl11
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	// The type of the ACL.
	//
	// example:
	//
	// ALLOW
	IpControlType *string `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
	// The time when the ACL was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-17T06:00:38Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the region in which the ACL is deployed.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetIpControlId() *string {
	return s.IpControlId
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetIpControlName() *string {
	return s.IpControlName
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetIpControlType() *string {
	return s.IpControlType
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetCreateTime(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetDescription(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.Description = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetIpControlId(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.IpControlId = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetIpControlName(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.IpControlName = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetIpControlType(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.IpControlType = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetModifiedTime(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetRegionId(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) Validate() error {
	return dara.Validate(s)
}
