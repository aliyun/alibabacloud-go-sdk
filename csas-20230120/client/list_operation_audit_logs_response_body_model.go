// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationAuditLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListOperationAuditLogsResponseBodyLogs) *ListOperationAuditLogsResponseBody
	GetLogs() []*ListOperationAuditLogsResponseBodyLogs
	SetRequestId(v string) *ListOperationAuditLogsResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListOperationAuditLogsResponseBody
	GetTotalNum() *int64
}

type ListOperationAuditLogsResponseBody struct {
	// The list of administrator operation audit logs, sorted by operation time in descending order.
	Logs []*ListOperationAuditLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// E9D4B681-0E79-57B7-AF0D-4A675D40141C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of logs that match the query conditions.
	//
	// example:
	//
	// 4
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListOperationAuditLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationAuditLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationAuditLogsResponseBody) GetLogs() []*ListOperationAuditLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListOperationAuditLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationAuditLogsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListOperationAuditLogsResponseBody) SetLogs(v []*ListOperationAuditLogsResponseBodyLogs) *ListOperationAuditLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListOperationAuditLogsResponseBody) SetRequestId(v string) *ListOperationAuditLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationAuditLogsResponseBody) SetTotalNum(v int64) *ListOperationAuditLogsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListOperationAuditLogsResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationAuditLogsResponseBodyLogs struct {
	// The post-operation snapshot. This value is recorded as-is by the audit framework without localization. This field is empty for historical logs that are not integrated with the audit framework.
	//
	// example:
	//
	// {"RequestId":"01A03244-5BAD-5FAA-93D6-E4F4A1A2****"}
	AfterAction *string `json:"AfterAction,omitempty" xml:"AfterAction,omitempty"`
	// The pre-operation snapshot. This value is recorded as-is by the audit framework without localization. This field is empty for historical logs that are not integrated with the audit framework.
	//
	// example:
	//
	// "pa-application-ea73352b4b75****"
	BeforeAction *string `json:"BeforeAction,omitempty" xml:"BeforeAction,omitempty"`
	// The error code when the operation failed. This field is empty when the operation succeeded.
	//
	// example:
	//
	// ResourceNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message when the operation failed. This field is empty when the operation succeeded.
	//
	// example:
	//
	// the specified resource is not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The event source type. Valid values:
	//
	// - **console**: console call.
	//
	// - **sdk**: SDK call.
	//
	// example:
	//
	// console
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The operation function module. The return value is localized based on the request language.
	//
	// example:
	//
	// Office Data Protection
	OperationFunc *string `json:"OperationFunc,omitempty" xml:"OperationFunc,omitempty"`
	// The operation page. The return value is localized based on the request language.
	//
	// example:
	//
	// Peripheral Management
	OperationPage *string `json:"OperationPage,omitempty" xml:"OperationPage,omitempty"`
	// The operation time.
	//
	// example:
	//
	// 2026-08-24 13:38:06
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// The operation type. The return value is localized based on the request language.
	//
	// example:
	//
	// Modify peripheral control policy
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The Alibaba Cloud account ID (AliUid) of the operator.
	//
	// example:
	//
	// 1234****
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOperationAuditLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListOperationAuditLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetAfterAction() *string {
	return s.AfterAction
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetBeforeAction() *string {
	return s.BeforeAction
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetEventType() *string {
	return s.EventType
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetOperationFunc() *string {
	return s.OperationFunc
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetOperationPage() *string {
	return s.OperationPage
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetOperationTime() *string {
	return s.OperationTime
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetOperationType() *string {
	return s.OperationType
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListOperationAuditLogsResponseBodyLogs) GetSuccess() *bool {
	return s.Success
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetAfterAction(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.AfterAction = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetBeforeAction(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.BeforeAction = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetErrorCode(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.ErrorCode = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetErrorMessage(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.ErrorMessage = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetEventType(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.EventType = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetOperationFunc(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.OperationFunc = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetOperationPage(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.OperationPage = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetOperationTime(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.OperationTime = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetOperationType(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.OperationType = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetOperatorId(v string) *ListOperationAuditLogsResponseBodyLogs {
	s.OperatorId = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) SetSuccess(v bool) *ListOperationAuditLogsResponseBodyLogs {
	s.Success = &v
	return s
}

func (s *ListOperationAuditLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
