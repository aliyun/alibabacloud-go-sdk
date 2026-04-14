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
	if s.IpControlInfos != nil {
		if err := s.IpControlInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.IpControlInfo != nil {
		for _, item := range s.IpControlInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpControlsResponseBodyIpControlInfosIpControlInfo struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	IpControlType *string `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
