// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitPreBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *BatchSubmitPreBillResponseBodyModule) *BatchSubmitPreBillResponseBody
	GetModule() *BatchSubmitPreBillResponseBodyModule
	SetMorePage(v bool) *BatchSubmitPreBillResponseBody
	GetMorePage() *bool
	SetRequestId(v string) *BatchSubmitPreBillResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *BatchSubmitPreBillResponseBody
	GetResultCode() *int32
	SetResultMsg(v string) *BatchSubmitPreBillResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *BatchSubmitPreBillResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchSubmitPreBillResponseBody
	GetTraceId() *string
}

type BatchSubmitPreBillResponseBody struct {
	// The data.
	Module *BatchSubmitPreBillResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// The pagination token set by the server. Indicates whether more data exists on the next page during pagination.
	//
	// example:
	//
	// true
	MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
	// The request ID.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 21041ce********056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s BatchSubmitPreBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillResponseBody) GetModule() *BatchSubmitPreBillResponseBodyModule {
	return s.Module
}

func (s *BatchSubmitPreBillResponseBody) GetMorePage() *bool {
	return s.MorePage
}

func (s *BatchSubmitPreBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSubmitPreBillResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *BatchSubmitPreBillResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *BatchSubmitPreBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchSubmitPreBillResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchSubmitPreBillResponseBody) SetModule(v *BatchSubmitPreBillResponseBodyModule) *BatchSubmitPreBillResponseBody {
	s.Module = v
	return s
}

func (s *BatchSubmitPreBillResponseBody) SetMorePage(v bool) *BatchSubmitPreBillResponseBody {
	s.MorePage = &v
	return s
}

func (s *BatchSubmitPreBillResponseBody) SetRequestId(v string) *BatchSubmitPreBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSubmitPreBillResponseBody) SetResultCode(v int32) *BatchSubmitPreBillResponseBody {
	s.ResultCode = &v
	return s
}

func (s *BatchSubmitPreBillResponseBody) SetResultMsg(v string) *BatchSubmitPreBillResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *BatchSubmitPreBillResponseBody) SetSuccess(v bool) *BatchSubmitPreBillResponseBody {
	s.Success = &v
	return s
}

func (s *BatchSubmitPreBillResponseBody) SetTraceId(v string) *BatchSubmitPreBillResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchSubmitPreBillResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchSubmitPreBillResponseBodyModule struct {
	// The batch ID. This value may be null if no actionable bills exist.
	//
	// example:
	//
	// 999
	BatchId *int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// The number of bills that cannot be updated.
	//
	// example:
	//
	// 20
	ForbidUpdateBillCount *int32 `json:"forbid_update_bill_count,omitempty" xml:"forbid_update_bill_count,omitempty"`
	// The details of bills that cannot be updated.
	ForbidUpdateDetail []*BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail `json:"forbid_update_detail,omitempty" xml:"forbid_update_detail,omitempty" type:"Repeated"`
	// The number of matched records.
	//
	// example:
	//
	// 10
	MatchCount *int32 `json:"match_count,omitempty" xml:"match_count,omitempty"`
	// The number of unmatched records.
	//
	// example:
	//
	// 10
	NotMatchCount *int32 `json:"not_match_count,omitempty" xml:"not_match_count,omitempty"`
	// The unmatched details.
	NotMatchDetail []*string `json:"not_match_detail,omitempty" xml:"not_match_detail,omitempty" type:"Repeated"`
}

func (s BatchSubmitPreBillResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillResponseBodyModule) GetBatchId() *int64 {
	return s.BatchId
}

func (s *BatchSubmitPreBillResponseBodyModule) GetForbidUpdateBillCount() *int32 {
	return s.ForbidUpdateBillCount
}

func (s *BatchSubmitPreBillResponseBodyModule) GetForbidUpdateDetail() []*BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail {
	return s.ForbidUpdateDetail
}

func (s *BatchSubmitPreBillResponseBodyModule) GetMatchCount() *int32 {
	return s.MatchCount
}

func (s *BatchSubmitPreBillResponseBodyModule) GetNotMatchCount() *int32 {
	return s.NotMatchCount
}

func (s *BatchSubmitPreBillResponseBodyModule) GetNotMatchDetail() []*string {
	return s.NotMatchDetail
}

func (s *BatchSubmitPreBillResponseBodyModule) SetBatchId(v int64) *BatchSubmitPreBillResponseBodyModule {
	s.BatchId = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModule) SetForbidUpdateBillCount(v int32) *BatchSubmitPreBillResponseBodyModule {
	s.ForbidUpdateBillCount = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModule) SetForbidUpdateDetail(v []*BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) *BatchSubmitPreBillResponseBodyModule {
	s.ForbidUpdateDetail = v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModule) SetMatchCount(v int32) *BatchSubmitPreBillResponseBodyModule {
	s.MatchCount = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModule) SetNotMatchCount(v int32) *BatchSubmitPreBillResponseBodyModule {
	s.NotMatchCount = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModule) SetNotMatchDetail(v []*string) *BatchSubmitPreBillResponseBodyModule {
	s.NotMatchDetail = v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModule) Validate() error {
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

type BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail struct {
	// The number of records that cannot be updated.
	//
	// example:
	//
	// 10
	CanNotUpdateCount *int32 `json:"can_not_update_count,omitempty" xml:"can_not_update_count,omitempty"`
	// The number of records that can be updated.
	//
	// example:
	//
	// 10
	CanUpdateCount *int32 `json:"can_update_count,omitempty" xml:"can_update_count,omitempty"`
	// The value.
	//
	// example:
	//
	// "9999"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) GetCanNotUpdateCount() *int32 {
	return s.CanNotUpdateCount
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) GetCanUpdateCount() *int32 {
	return s.CanUpdateCount
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) GetValue() *string {
	return s.Value
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) SetCanNotUpdateCount(v int32) *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail {
	s.CanNotUpdateCount = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) SetCanUpdateCount(v int32) *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail {
	s.CanUpdateCount = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) SetValue(v string) *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail {
	s.Value = &v
	return s
}

func (s *BatchSubmitPreBillResponseBodyModuleForbidUpdateDetail) Validate() error {
	return dara.Validate(s)
}
