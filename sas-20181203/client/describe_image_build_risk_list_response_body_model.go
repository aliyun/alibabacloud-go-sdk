// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBuildRiskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageBuildRiskListResponseBody
	GetCode() *string
	SetData(v *DescribeImageBuildRiskListResponseBodyData) *DescribeImageBuildRiskListResponseBody
	GetData() *DescribeImageBuildRiskListResponseBodyData
	SetMessage(v string) *DescribeImageBuildRiskListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeImageBuildRiskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageBuildRiskListResponseBody
	GetSuccess() *bool
}

type DescribeImageBuildRiskListResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeImageBuildRiskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageBuildRiskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageBuildRiskListResponseBody) GetData() *DescribeImageBuildRiskListResponseBodyData {
	return s.Data
}

func (s *DescribeImageBuildRiskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageBuildRiskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBuildRiskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageBuildRiskListResponseBody) SetCode(v string) *DescribeImageBuildRiskListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBody) SetData(v *DescribeImageBuildRiskListResponseBodyData) *DescribeImageBuildRiskListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageBuildRiskListResponseBody) SetMessage(v string) *DescribeImageBuildRiskListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBody) SetRequestId(v string) *DescribeImageBuildRiskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBody) SetSuccess(v bool) *DescribeImageBuildRiskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBuildRiskListResponseBodyData struct {
	// The risks.
	List []*DescribeImageBuildRiskListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageBuildRiskListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s DescribeImageBuildRiskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskListResponseBodyData) GetList() []*DescribeImageBuildRiskListResponseBodyDataList {
	return s.List
}

func (s *DescribeImageBuildRiskListResponseBodyData) GetPageInfo() *DescribeImageBuildRiskListResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeImageBuildRiskListResponseBodyData) SetList(v []*DescribeImageBuildRiskListResponseBodyDataList) *DescribeImageBuildRiskListResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyData) SetPageInfo(v *DescribeImageBuildRiskListResponseBodyDataPageInfo) *DescribeImageBuildRiskListResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type DescribeImageBuildRiskListResponseBodyDataList struct {
	// The number of affected images.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1723710827000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1723710827999
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The type key of the risk.
	//
	// example:
	//
	// other
	RiskClass *string `json:"RiskClass,omitempty" xml:"RiskClass,omitempty"`
	// The type name of the risk.
	//
	// example:
	//
	// other
	RiskClassName *string `json:"RiskClassName,omitempty" xml:"RiskClassName,omitempty"`
	// The key of the risk. You can call the [DescribeImageBuildRiskList](~~~~) operation to obtain the value of **RiskKey**.
	//
	// example:
	//
	// no_user
	RiskKey *string `json:"RiskKey,omitempty" xml:"RiskKey,omitempty"`
	// The rule name of the risk.
	//
	// example:
	//
	// no_user
	RiskKeyName *string `json:"RiskKeyName,omitempty" xml:"RiskKeyName,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The number of unprocessed images.
	//
	// example:
	//
	// 1
	UnprocessedNum *int32 `json:"UnprocessedNum,omitempty" xml:"UnprocessedNum,omitempty"`
}

func (s DescribeImageBuildRiskListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetRiskClass() *string {
	return s.RiskClass
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetRiskClassName() *string {
	return s.RiskClassName
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetRiskKey() *string {
	return s.RiskKey
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetRiskKeyName() *string {
	return s.RiskKeyName
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) GetUnprocessedNum() *int32 {
	return s.UnprocessedNum
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetCount(v int32) *DescribeImageBuildRiskListResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetFirstScanTime(v int64) *DescribeImageBuildRiskListResponseBodyDataList {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetLastScanTime(v int64) *DescribeImageBuildRiskListResponseBodyDataList {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetRiskClass(v string) *DescribeImageBuildRiskListResponseBodyDataList {
	s.RiskClass = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetRiskClassName(v string) *DescribeImageBuildRiskListResponseBodyDataList {
	s.RiskClassName = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetRiskKey(v string) *DescribeImageBuildRiskListResponseBodyDataList {
	s.RiskKey = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetRiskKeyName(v string) *DescribeImageBuildRiskListResponseBodyDataList {
	s.RiskKeyName = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetRiskLevel(v string) *DescribeImageBuildRiskListResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) SetUnprocessedNum(v int32) *DescribeImageBuildRiskListResponseBodyDataList {
	s.UnprocessedNum = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeImageBuildRiskListResponseBodyDataPageInfo struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 109
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageBuildRiskListResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeImageBuildRiskListResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeImageBuildRiskListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) SetTotalCount(v int32) *DescribeImageBuildRiskListResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageBuildRiskListResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
