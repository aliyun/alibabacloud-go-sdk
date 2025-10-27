// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineItemListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineItemInfos(v []*DescribeImageBaselineItemListResponseBodyBaselineItemInfos) *DescribeImageBaselineItemListResponseBody
	GetBaselineItemInfos() []*DescribeImageBaselineItemListResponseBodyBaselineItemInfos
	SetPageInfo(v *DescribeImageBaselineItemListResponseBodyPageInfo) *DescribeImageBaselineItemListResponseBody
	GetPageInfo() *DescribeImageBaselineItemListResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageBaselineItemListResponseBody
	GetRequestId() *string
}

type DescribeImageBaselineItemListResponseBody struct {
	// An array that consists of baseline check items.
	BaselineItemInfos []*DescribeImageBaselineItemListResponseBodyBaselineItemInfos `json:"BaselineItemInfos,omitempty" xml:"BaselineItemInfos,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageBaselineItemListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageBaselineItemListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineItemListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineItemListResponseBody) GetBaselineItemInfos() []*DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	return s.BaselineItemInfos
}

func (s *DescribeImageBaselineItemListResponseBody) GetPageInfo() *DescribeImageBaselineItemListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageBaselineItemListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBaselineItemListResponseBody) SetBaselineItemInfos(v []*DescribeImageBaselineItemListResponseBodyBaselineItemInfos) *DescribeImageBaselineItemListResponseBody {
	s.BaselineItemInfos = v
	return s
}

func (s *DescribeImageBaselineItemListResponseBody) SetPageInfo(v *DescribeImageBaselineItemListResponseBodyPageInfo) *DescribeImageBaselineItemListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageBaselineItemListResponseBody) SetRequestId(v string) *DescribeImageBaselineItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBody) Validate() error {
	if s.BaselineItemInfos != nil {
		for _, item := range s.BaselineItemInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBaselineItemListResponseBodyBaselineItemInfos struct {
	// The alias of the baseline type.
	//
	// example:
	//
	// AccessKey pair leak
	BaselineClassAlias *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	// The key of the baseline type.
	//
	// example:
	//
	// ak_leak
	BaselineClassKey *string `json:"BaselineClassKey,omitempty" xml:"BaselineClassKey,omitempty"`
	// The alias of the baseline check item.
	//
	// example:
	//
	// AccessKey pair leak
	BaselineItemAlias *string `json:"BaselineItemAlias,omitempty" xml:"BaselineItemAlias,omitempty"`
	// The key of the baseline check item.
	//
	// example:
	//
	// ak_leak
	BaselineItemKey *string `json:"BaselineItemKey,omitempty" xml:"BaselineItemKey,omitempty"`
	// The alias of the baseline.
	//
	// example:
	//
	// AccessKey pair leak
	BaselineNameAlias *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	// The key of the baseline name.
	//
	// example:
	//
	// ak_leak
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The status of the baseline risks. Valid values:
	//
	// 	- **0**: unfixed
	//
	// 	- **1**: fixed
	//
	// 	- **2**: pending verification
	//
	// 	- **3**: fixing failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the baseline check item is added to the whitelist. Valid values:
	//
	// 	- **0**: The baseline check item is not added to the whitelist.
	//
	// 	- **1**: The baseline check item is added to the whitelist.
	//
	// example:
	//
	// 0
	WhiteList *int32 `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s DescribeImageBaselineItemListResponseBodyBaselineItemInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetBaselineClassAlias() *string {
	return s.BaselineClassAlias
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetBaselineClassKey() *string {
	return s.BaselineClassKey
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetBaselineItemAlias() *string {
	return s.BaselineItemAlias
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetBaselineItemKey() *string {
	return s.BaselineItemKey
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetBaselineNameAlias() *string {
	return s.BaselineNameAlias
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetBaselineNameKey() *string {
	return s.BaselineNameKey
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) GetWhiteList() *int32 {
	return s.WhiteList
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetBaselineClassAlias(v string) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.BaselineClassAlias = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetBaselineClassKey(v string) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.BaselineClassKey = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetBaselineItemAlias(v string) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.BaselineItemAlias = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetBaselineItemKey(v string) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.BaselineItemKey = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetBaselineNameAlias(v string) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.BaselineNameAlias = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetBaselineNameKey(v string) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.BaselineNameKey = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetStatus(v int32) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.Status = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) SetWhiteList(v int32) *DescribeImageBaselineItemListResponseBodyBaselineItemInfos {
	s.WhiteList = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyBaselineItemInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeImageBaselineItemListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 253
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageBaselineItemListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineItemListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) SetCount(v int32) *DescribeImageBaselineItemListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageBaselineItemListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageBaselineItemListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageBaselineItemListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageBaselineItemListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
