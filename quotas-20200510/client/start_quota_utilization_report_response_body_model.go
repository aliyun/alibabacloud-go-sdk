// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartQuotaUtilizationReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StartQuotaUtilizationReportResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *StartQuotaUtilizationReportResponseBody
	GetErrorMsg() *string
	SetMessage(v string) *StartQuotaUtilizationReportResponseBody
	GetMessage() *string
	SetNextRefreshTime(v string) *StartQuotaUtilizationReportResponseBody
	GetNextRefreshTime() *string
	SetRefreshSuccess(v bool) *StartQuotaUtilizationReportResponseBody
	GetRefreshSuccess() *bool
	SetRequestId(v string) *StartQuotaUtilizationReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartQuotaUtilizationReportResponseBody
	GetSuccess() *bool
}

type StartQuotaUtilizationReportResponseBody struct {
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// Quota utilization report generation has started
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2026-10-26T08:00:00Z
	NextRefreshTime *string `json:"NextRefreshTime,omitempty" xml:"NextRefreshTime,omitempty"`
	// example:
	//
	// true
	RefreshSuccess *bool `json:"RefreshSuccess,omitempty" xml:"RefreshSuccess,omitempty"`
	// example:
	//
	// 7ED584FB-ECBF-4A2A-969D-F54D25EFABF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartQuotaUtilizationReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartQuotaUtilizationReportResponseBody) GoString() string {
	return s.String()
}

func (s *StartQuotaUtilizationReportResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartQuotaUtilizationReportResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *StartQuotaUtilizationReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartQuotaUtilizationReportResponseBody) GetNextRefreshTime() *string {
	return s.NextRefreshTime
}

func (s *StartQuotaUtilizationReportResponseBody) GetRefreshSuccess() *bool {
	return s.RefreshSuccess
}

func (s *StartQuotaUtilizationReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartQuotaUtilizationReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartQuotaUtilizationReportResponseBody) SetErrorCode(v string) *StartQuotaUtilizationReportResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) SetErrorMsg(v string) *StartQuotaUtilizationReportResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) SetMessage(v string) *StartQuotaUtilizationReportResponseBody {
	s.Message = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) SetNextRefreshTime(v string) *StartQuotaUtilizationReportResponseBody {
	s.NextRefreshTime = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) SetRefreshSuccess(v bool) *StartQuotaUtilizationReportResponseBody {
	s.RefreshSuccess = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) SetRequestId(v string) *StartQuotaUtilizationReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) SetSuccess(v bool) *StartQuotaUtilizationReportResponseBody {
	s.Success = &v
	return s
}

func (s *StartQuotaUtilizationReportResponseBody) Validate() error {
	return dara.Validate(s)
}
