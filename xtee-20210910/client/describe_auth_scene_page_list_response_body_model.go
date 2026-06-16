// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthScenePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAuthScenePageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int64) *DescribeAuthScenePageListResponseBody
	GetCurrentPage() *int64
	SetPageSize(v int64) *DescribeAuthScenePageListResponseBody
	GetPageSize() *int64
	SetResultObject(v []*DescribeAuthScenePageListResponseBodyResultObject) *DescribeAuthScenePageListResponseBody
	GetResultObject() []*DescribeAuthScenePageListResponseBodyResultObject
	SetTotalItem(v int64) *DescribeAuthScenePageListResponseBody
	GetTotalItem() *int64
	SetTotalPage(v int64) *DescribeAuthScenePageListResponseBody
	GetTotalPage() *int64
}

type DescribeAuthScenePageListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The returned object.
	ResultObject []*DescribeAuthScenePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 6
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 9
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeAuthScenePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthScenePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthScenePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthScenePageListResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeAuthScenePageListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAuthScenePageListResponseBody) GetResultObject() []*DescribeAuthScenePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAuthScenePageListResponseBody) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeAuthScenePageListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAuthScenePageListResponseBody) SetRequestId(v string) *DescribeAuthScenePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBody) SetCurrentPage(v int64) *DescribeAuthScenePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBody) SetPageSize(v int64) *DescribeAuthScenePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBody) SetResultObject(v []*DescribeAuthScenePageListResponseBodyResultObject) *DescribeAuthScenePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAuthScenePageListResponseBody) SetTotalItem(v int64) *DescribeAuthScenePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBody) SetTotalPage(v int64) *DescribeAuthScenePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBody) Validate() error {
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

type DescribeAuthScenePageListResponseBodyResultObject struct {
	// The number of General Policy policies.
	//
	// example:
	//
	// 10
	CommonRuleCount *int64 `json:"commonRuleCount,omitempty" xml:"commonRuleCount,omitempty"`
	// The number of custom policies.
	//
	// example:
	//
	// 10
	CustomRuleCount *int64 `json:"customRuleCount,omitempty" xml:"customRuleCount,omitempty"`
	// The event code.
	//
	// example:
	//
	// de_avypfd8253
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// 注册风险识别服务
	SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	// The service name.
	//
	// example:
	//
	// account_abuse
	ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty"`
}

func (s DescribeAuthScenePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthScenePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) GetCommonRuleCount() *int64 {
	return s.CommonRuleCount
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) GetCustomRuleCount() *int64 {
	return s.CustomRuleCount
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) GetServerName() *string {
	return s.ServerName
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) SetCommonRuleCount(v int64) *DescribeAuthScenePageListResponseBodyResultObject {
	s.CommonRuleCount = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) SetCustomRuleCount(v int64) *DescribeAuthScenePageListResponseBodyResultObject {
	s.CustomRuleCount = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) SetEventCode(v string) *DescribeAuthScenePageListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) SetSceneName(v string) *DescribeAuthScenePageListResponseBodyResultObject {
	s.SceneName = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) SetServerName(v string) *DescribeAuthScenePageListResponseBodyResultObject {
	s.ServerName = &v
	return s
}

func (s *DescribeAuthScenePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
