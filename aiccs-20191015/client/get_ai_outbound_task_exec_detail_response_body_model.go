// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskExecDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiOutboundTaskExecDetailResponseBody
	GetCode() *string
	SetData(v *GetAiOutboundTaskExecDetailResponseBodyData) *GetAiOutboundTaskExecDetailResponseBody
	GetData() *GetAiOutboundTaskExecDetailResponseBodyData
	SetMessage(v string) *GetAiOutboundTaskExecDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiOutboundTaskExecDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiOutboundTaskExecDetailResponseBody
	GetSuccess() *bool
}

type GetAiOutboundTaskExecDetailResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAiOutboundTaskExecDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAiOutboundTaskExecDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskExecDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskExecDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiOutboundTaskExecDetailResponseBody) GetData() *GetAiOutboundTaskExecDetailResponseBodyData {
	return s.Data
}

func (s *GetAiOutboundTaskExecDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiOutboundTaskExecDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiOutboundTaskExecDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiOutboundTaskExecDetailResponseBody) SetCode(v string) *GetAiOutboundTaskExecDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBody) SetData(v *GetAiOutboundTaskExecDetailResponseBodyData) *GetAiOutboundTaskExecDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBody) SetMessage(v string) *GetAiOutboundTaskExecDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBody) SetRequestId(v string) *GetAiOutboundTaskExecDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBody) SetSuccess(v bool) *GetAiOutboundTaskExecDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskExecDetailResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// true
	HasNextPage *bool                                              `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	List        []*GetAiOutboundTaskExecDetailResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 199
	TotalResults *int32 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s GetAiOutboundTaskExecDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskExecDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) GetHasNextPage() *bool {
	return s.HasNextPage
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) GetList() []*GetAiOutboundTaskExecDetailResponseBodyDataList {
	return s.List
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) GetTotalResults() *int32 {
	return s.TotalResults
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) SetCurrentPage(v int32) *GetAiOutboundTaskExecDetailResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) SetHasNextPage(v bool) *GetAiOutboundTaskExecDetailResponseBodyData {
	s.HasNextPage = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) SetList(v []*GetAiOutboundTaskExecDetailResponseBodyDataList) *GetAiOutboundTaskExecDetailResponseBodyData {
	s.List = v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) SetPageSize(v int32) *GetAiOutboundTaskExecDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) SetTotalResults(v int32) *GetAiOutboundTaskExecDetailResponseBodyData {
	s.TotalResults = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiOutboundTaskExecDetailResponseBodyDataList struct {
	// example:
	//
	// 1
	BatchVersion *int32 `json:"BatchVersion,omitempty" xml:"BatchVersion,omitempty"`
	// example:
	//
	// 123
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	// example:
	//
	// 1
	CallCount *int32 `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// example:
	//
	// 123456
	CaseId *int64 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// example:
	//
	// 1632289999000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 未接通
	LastCallResult *string `json:"LastCallResult,omitempty" xml:"LastCallResult,omitempty"`
	// example:
	//
	// 150****0000
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 待呼叫
	StatusDesc *int32 `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s GetAiOutboundTaskExecDetailResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskExecDetailResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetBatchVersion() *int32 {
	return s.BatchVersion
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetBizData() *string {
	return s.BizData
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetCallCount() *int32 {
	return s.CallCount
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetCaseId() *int64 {
	return s.CaseId
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetLastCallResult() *string {
	return s.LastCallResult
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetStatus() *int32 {
	return s.Status
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) GetStatusDesc() *int32 {
	return s.StatusDesc
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetBatchVersion(v int32) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.BatchVersion = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetBizData(v string) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.BizData = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetCallCount(v int32) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.CallCount = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetCaseId(v int64) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.CaseId = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetCreateTime(v int64) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetLastCallResult(v string) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.LastCallResult = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetPhoneNum(v string) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.PhoneNum = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetStatus(v int32) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) SetStatusDesc(v int32) *GetAiOutboundTaskExecDetailResponseBodyDataList {
	s.StatusDesc = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
