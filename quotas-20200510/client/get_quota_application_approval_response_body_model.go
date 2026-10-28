// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaApplicationApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GetQuotaApplicationApprovalResponseBody
	GetAllowRetry() *bool
	SetDynamicCode(v string) *GetQuotaApplicationApprovalResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetQuotaApplicationApprovalResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetQuotaApplicationApprovalResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *GetQuotaApplicationApprovalResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetQuotaApplicationApprovalResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetQuotaApplicationApprovalResponseBody
	GetHttpStatusCode() *int32
	SetModule(v *GetQuotaApplicationApprovalResponseBodyModule) *GetQuotaApplicationApprovalResponseBody
	GetModule() *GetQuotaApplicationApprovalResponseBodyModule
	SetRequestId(v string) *GetQuotaApplicationApprovalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQuotaApplicationApprovalResponseBody
	GetSuccess() *bool
}

type GetQuotaApplicationApprovalResponseBody struct {
	// Specifies whether retries are allowed. Valid values:
	//
	// - false: Retries are not allowed.
	//
	// - true: Retries are allowed.
	//
	// example:
	//
	// false
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// PARAMETER.ILLEGALL
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// Parameter[%s] is required.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response object.
	Module *GetQuotaApplicationApprovalResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7ED584FB-ECBF-4A2A-969D-F54D25EFABF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQuotaApplicationApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetQuotaApplicationApprovalResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetQuotaApplicationApprovalResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetQuotaApplicationApprovalResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetQuotaApplicationApprovalResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQuotaApplicationApprovalResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetQuotaApplicationApprovalResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQuotaApplicationApprovalResponseBody) GetModule() *GetQuotaApplicationApprovalResponseBodyModule {
	return s.Module
}

func (s *GetQuotaApplicationApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaApplicationApprovalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQuotaApplicationApprovalResponseBody) SetAllowRetry(v bool) *GetQuotaApplicationApprovalResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetDynamicCode(v string) *GetQuotaApplicationApprovalResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetDynamicMessage(v string) *GetQuotaApplicationApprovalResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetErrorArgs(v []interface{}) *GetQuotaApplicationApprovalResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetErrorCode(v string) *GetQuotaApplicationApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetErrorMsg(v string) *GetQuotaApplicationApprovalResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetHttpStatusCode(v int32) *GetQuotaApplicationApprovalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetModule(v *GetQuotaApplicationApprovalResponseBodyModule) *GetQuotaApplicationApprovalResponseBody {
	s.Module = v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetRequestId(v string) *GetQuotaApplicationApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) SetSuccess(v bool) *GetQuotaApplicationApprovalResponseBody {
	s.Success = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaApplicationApprovalResponseBodyModule struct {
	// The average approval time for the quota application. Unit: hours.
	//
	// example:
	//
	// 24
	ApprovalHours *int32 `json:"ApprovalHours,omitempty" xml:"ApprovalHours,omitempty"`
	// The approval reminder interval for the quota application. Unit: hours.
	//
	// example:
	//
	// 24
	ReminderIntervalHours *int32 `json:"ReminderIntervalHours,omitempty" xml:"ReminderIntervalHours,omitempty"`
	// Specifies whether approval reminders are supported for the quota application. Valid values:
	//
	// - false: Not supported.
	//
	// - true: Supported.
	//
	// example:
	//
	// false
	SupportReminder *bool `json:"SupportReminder,omitempty" xml:"SupportReminder,omitempty"`
	// The reason why approval reminders are not supported for the quota application.
	//
	// example:
	//
	// can only be remind once within the interval
	UnsupportReminderReason *string `json:"UnsupportReminderReason,omitempty" xml:"UnsupportReminderReason,omitempty"`
}

func (s GetQuotaApplicationApprovalResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaApplicationApprovalResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) GetApprovalHours() *int32 {
	return s.ApprovalHours
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) GetReminderIntervalHours() *int32 {
	return s.ReminderIntervalHours
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) GetSupportReminder() *bool {
	return s.SupportReminder
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) GetUnsupportReminderReason() *string {
	return s.UnsupportReminderReason
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetApprovalHours(v int32) *GetQuotaApplicationApprovalResponseBodyModule {
	s.ApprovalHours = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetReminderIntervalHours(v int32) *GetQuotaApplicationApprovalResponseBodyModule {
	s.ReminderIntervalHours = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetSupportReminder(v bool) *GetQuotaApplicationApprovalResponseBodyModule {
	s.SupportReminder = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) SetUnsupportReminderReason(v string) *GetQuotaApplicationApprovalResponseBodyModule {
	s.UnsupportReminderReason = &v
	return s
}

func (s *GetQuotaApplicationApprovalResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
