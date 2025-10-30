// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneEventPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSceneEventPageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSceneEventPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSceneEventPageListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeSceneEventPageListResponseBodyResultObject) *DescribeSceneEventPageListResponseBody
	GetResultObject() []*DescribeSceneEventPageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSceneEventPageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSceneEventPageListResponseBody
	GetTotalPage() *int32
}

type DescribeSceneEventPageListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of items per page. Default value: 20, minimum value: 1, maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID, each request has a unique value, which facilitates subsequent troubleshooting
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return array
	ResultObject []*DescribeSceneEventPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSceneEventPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneEventPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneEventPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSceneEventPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSceneEventPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSceneEventPageListResponseBody) GetResultObject() []*DescribeSceneEventPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSceneEventPageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSceneEventPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSceneEventPageListResponseBody) SetCurrentPage(v int32) *DescribeSceneEventPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBody) SetPageSize(v int32) *DescribeSceneEventPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBody) SetRequestId(v string) *DescribeSceneEventPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBody) SetResultObject(v []*DescribeSceneEventPageListResponseBodyResultObject) *DescribeSceneEventPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSceneEventPageListResponseBody) SetTotalItem(v int32) *DescribeSceneEventPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBody) SetTotalPage(v int32) *DescribeSceneEventPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneEventPageListResponseBodyResultObject struct {
	// Number of common rules
	//
	// example:
	//
	// 10
	CommonRuleCount *string `json:"commonRuleCount,omitempty" xml:"commonRuleCount,omitempty"`
	// Number of custom rules
	//
	// example:
	//
	// 10
	CustomRuleCount *string `json:"customRuleCount,omitempty" xml:"customRuleCount,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Modification time
	//
	// example:
	//
	// 1565701886000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Modifier
	//
	// example:
	//
	// 公有云用户uid
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Number of custom rules
	//
	// example:
	//
	// 10
	NormalRuleCount *string `json:"normalRuleCount,omitempty" xml:"normalRuleCount,omitempty"`
	// Service code
	//
	// example:
	//
	// device_risk
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// Usage status
	//
	// example:
	//
	// 使用/未使用
	UseStatus *string `json:"useStatus,omitempty" xml:"useStatus,omitempty"`
	// Number of white-box rules
	//
	// example:
	//
	// 10
	WhiteBoxRuleCount *string `json:"whiteBoxRuleCount,omitempty" xml:"whiteBoxRuleCount,omitempty"`
}

func (s DescribeSceneEventPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneEventPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetCommonRuleCount() *string {
	return s.CommonRuleCount
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetCustomRuleCount() *string {
	return s.CustomRuleCount
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetNormalRuleCount() *string {
	return s.NormalRuleCount
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetService() *string {
	return s.Service
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetUseStatus() *string {
	return s.UseStatus
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) GetWhiteBoxRuleCount() *string {
	return s.WhiteBoxRuleCount
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetCommonRuleCount(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.CommonRuleCount = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetCustomRuleCount(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.CustomRuleCount = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetEventCode(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetEventName(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetGmtModified(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetModifier(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.Modifier = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetNormalRuleCount(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.NormalRuleCount = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetService(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.Service = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetUseStatus(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.UseStatus = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) SetWhiteBoxRuleCount(v string) *DescribeSceneEventPageListResponseBodyResultObject {
	s.WhiteBoxRuleCount = &v
	return s
}

func (s *DescribeSceneEventPageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
