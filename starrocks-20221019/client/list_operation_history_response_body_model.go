// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListOperationHistoryResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*ListOperationHistoryResponseBodyData) *ListOperationHistoryResponseBody
	GetData() []*ListOperationHistoryResponseBodyData
	SetErrCode(v string) *ListOperationHistoryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListOperationHistoryResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListOperationHistoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListOperationHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOperationHistoryResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListOperationHistoryResponseBody
	GetTotal() *int32
}

type ListOperationHistoryResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*ListOperationHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 832
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOperationHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListOperationHistoryResponseBody) GetData() []*ListOperationHistoryResponseBodyData {
	return s.Data
}

func (s *ListOperationHistoryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListOperationHistoryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListOperationHistoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOperationHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOperationHistoryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListOperationHistoryResponseBody) SetAccessDeniedDetail(v string) *ListOperationHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetData(v []*ListOperationHistoryResponseBodyData) *ListOperationHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationHistoryResponseBody) SetErrCode(v string) *ListOperationHistoryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetErrMessage(v string) *ListOperationHistoryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetHttpStatusCode(v int32) *ListOperationHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetRequestId(v string) *ListOperationHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetSuccess(v bool) *ListOperationHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetTotal(v int32) *ListOperationHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *ListOperationHistoryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationHistoryResponseBodyData struct {
	// example:
	//
	// FE enable = true
	AfterValue *string `json:"AfterValue,omitempty" xml:"AfterValue,omitempty"`
	// example:
	//
	// FE enable = false
	BeforeValue *string `json:"BeforeValue,omitempty" xml:"BeforeValue,omitempty"`
	// example:
	//
	// 1742179008000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1742179008000
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// example:
	//
	// c-cd7a3a6f2186d5c9
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperationDetail *string `json:"OperationDetail,omitempty" xml:"OperationDetail,omitempty"`
	// example:
	//
	// op-f49743caa809****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// example:
	//
	// COMPLETED
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// example:
	//
	// upgrade_version
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s ListOperationHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryResponseBodyData) GetAfterValue() *string {
	return s.AfterValue
}

func (s *ListOperationHistoryResponseBodyData) GetBeforeValue() *string {
	return s.BeforeValue
}

func (s *ListOperationHistoryResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListOperationHistoryResponseBodyData) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *ListOperationHistoryResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationHistoryResponseBodyData) GetOperationDetail() *string {
	return s.OperationDetail
}

func (s *ListOperationHistoryResponseBodyData) GetOperationId() *string {
	return s.OperationId
}

func (s *ListOperationHistoryResponseBodyData) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ListOperationHistoryResponseBodyData) GetOperationType() *string {
	return s.OperationType
}

func (s *ListOperationHistoryResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *ListOperationHistoryResponseBodyData) SetAfterValue(v string) *ListOperationHistoryResponseBodyData {
	s.AfterValue = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetBeforeValue(v string) *ListOperationHistoryResponseBodyData {
	s.BeforeValue = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetGmtCreate(v int64) *ListOperationHistoryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetGmtEnd(v int64) *ListOperationHistoryResponseBodyData {
	s.GmtEnd = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetInstanceId(v string) *ListOperationHistoryResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationDetail(v string) *ListOperationHistoryResponseBodyData {
	s.OperationDetail = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationId(v string) *ListOperationHistoryResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationStatus(v string) *ListOperationHistoryResponseBodyData {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationType(v string) *ListOperationHistoryResponseBodyData {
	s.OperationType = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetProgress(v int32) *ListOperationHistoryResponseBodyData {
	s.Progress = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
