// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPreBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *ConfirmPreBillResponseBodyModule) *ConfirmPreBillResponseBody
	GetModule() *ConfirmPreBillResponseBodyModule
	SetMorePage(v bool) *ConfirmPreBillResponseBody
	GetMorePage() *bool
	SetRequestId(v string) *ConfirmPreBillResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *ConfirmPreBillResponseBody
	GetResultCode() *int32
	SetResultMsg(v string) *ConfirmPreBillResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *ConfirmPreBillResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ConfirmPreBillResponseBody
	GetTraceId() *string
}

type ConfirmPreBillResponseBody struct {
	// The details of the returned result.
	Module *ConfirmPreBillResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// The pagination token set by the server. Indicates whether more data is available on the next page during pagination.
	//
	// example:
	//
	// true
	MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// 0
	ResultCode *int32 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// The error message.
	//
	// example:
	//
	// 成功
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ConfirmPreBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPreBillResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmPreBillResponseBody) GetModule() *ConfirmPreBillResponseBodyModule {
	return s.Module
}

func (s *ConfirmPreBillResponseBody) GetMorePage() *bool {
	return s.MorePage
}

func (s *ConfirmPreBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmPreBillResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *ConfirmPreBillResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ConfirmPreBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfirmPreBillResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ConfirmPreBillResponseBody) SetModule(v *ConfirmPreBillResponseBodyModule) *ConfirmPreBillResponseBody {
	s.Module = v
	return s
}

func (s *ConfirmPreBillResponseBody) SetMorePage(v bool) *ConfirmPreBillResponseBody {
	s.MorePage = &v
	return s
}

func (s *ConfirmPreBillResponseBody) SetRequestId(v string) *ConfirmPreBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmPreBillResponseBody) SetResultCode(v int32) *ConfirmPreBillResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ConfirmPreBillResponseBody) SetResultMsg(v string) *ConfirmPreBillResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *ConfirmPreBillResponseBody) SetSuccess(v bool) *ConfirmPreBillResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmPreBillResponseBody) SetTraceId(v string) *ConfirmPreBillResponseBody {
	s.TraceId = &v
	return s
}

func (s *ConfirmPreBillResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfirmPreBillResponseBodyModule struct {
	// The batch ID.
	//
	// example:
	//
	// 9999
	BatchId *int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The number of bills that cannot be updated.
	//
	// example:
	//
	// 10
	ForbidUpdateBillCount *int32 `json:"forbid_update_bill_count,omitempty" xml:"forbid_update_bill_count,omitempty"`
	// The details of items that cannot be updated.
	ForbidUpdateDetail []*ConfirmPreBillResponseBodyModuleForbidUpdateDetail `json:"forbid_update_detail,omitempty" xml:"forbid_update_detail,omitempty" type:"Repeated"`
	// The number of matched items.
	//
	// example:
	//
	// 10
	MatchCount *int32 `json:"match_count,omitempty" xml:"match_count,omitempty"`
	// The number of unmatched items.
	//
	// example:
	//
	// 10
	NotMatchCount *int32 `json:"not_match_count,omitempty" xml:"not_match_count,omitempty"`
	// The details of unmatched items.
	NotMatchDetail []*string `json:"not_match_detail,omitempty" xml:"not_match_detail,omitempty" type:"Repeated"`
}

func (s ConfirmPreBillResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPreBillResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ConfirmPreBillResponseBodyModule) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ConfirmPreBillResponseBodyModule) GetForbidUpdateBillCount() *int32 {
	return s.ForbidUpdateBillCount
}

func (s *ConfirmPreBillResponseBodyModule) GetForbidUpdateDetail() []*ConfirmPreBillResponseBodyModuleForbidUpdateDetail {
	return s.ForbidUpdateDetail
}

func (s *ConfirmPreBillResponseBodyModule) GetMatchCount() *int32 {
	return s.MatchCount
}

func (s *ConfirmPreBillResponseBodyModule) GetNotMatchCount() *int32 {
	return s.NotMatchCount
}

func (s *ConfirmPreBillResponseBodyModule) GetNotMatchDetail() []*string {
	return s.NotMatchDetail
}

func (s *ConfirmPreBillResponseBodyModule) SetBatchId(v int64) *ConfirmPreBillResponseBodyModule {
	s.BatchId = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModule) SetForbidUpdateBillCount(v int32) *ConfirmPreBillResponseBodyModule {
	s.ForbidUpdateBillCount = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModule) SetForbidUpdateDetail(v []*ConfirmPreBillResponseBodyModuleForbidUpdateDetail) *ConfirmPreBillResponseBodyModule {
	s.ForbidUpdateDetail = v
	return s
}

func (s *ConfirmPreBillResponseBodyModule) SetMatchCount(v int32) *ConfirmPreBillResponseBodyModule {
	s.MatchCount = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModule) SetNotMatchCount(v int32) *ConfirmPreBillResponseBodyModule {
	s.NotMatchCount = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModule) SetNotMatchDetail(v []*string) *ConfirmPreBillResponseBodyModule {
	s.NotMatchDetail = v
	return s
}

func (s *ConfirmPreBillResponseBodyModule) Validate() error {
	if s.ForbidUpdateDetail != nil {
		for _, item := range s.ForbidUpdateDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConfirmPreBillResponseBodyModuleForbidUpdateDetail struct {
	// The number of items that cannot be updated.
	//
	// example:
	//
	// 10
	CanNotUpdateCount *int32 `json:"can_not_update_count,omitempty" xml:"can_not_update_count,omitempty"`
	// The number of items that can be updated.
	//
	// example:
	//
	// 10
	CanUpdateCount *int32 `json:"can_update_count,omitempty" xml:"can_update_count,omitempty"`
	// The value.
	//
	// example:
	//
	// []
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ConfirmPreBillResponseBodyModuleForbidUpdateDetail) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPreBillResponseBodyModuleForbidUpdateDetail) GoString() string {
	return s.String()
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) GetCanNotUpdateCount() *int32 {
	return s.CanNotUpdateCount
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) GetCanUpdateCount() *int32 {
	return s.CanUpdateCount
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) GetValue() *string {
	return s.Value
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) SetCanNotUpdateCount(v int32) *ConfirmPreBillResponseBodyModuleForbidUpdateDetail {
	s.CanNotUpdateCount = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) SetCanUpdateCount(v int32) *ConfirmPreBillResponseBodyModuleForbidUpdateDetail {
	s.CanUpdateCount = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) SetValue(v string) *ConfirmPreBillResponseBodyModuleForbidUpdateDetail {
	s.Value = &v
	return s
}

func (s *ConfirmPreBillResponseBodyModuleForbidUpdateDetail) Validate() error {
	return dara.Validate(s)
}
