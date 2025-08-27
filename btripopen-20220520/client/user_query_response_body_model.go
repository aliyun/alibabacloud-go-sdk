// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UserQueryResponseBody
	GetCode() *string
	SetMessage(v string) *UserQueryResponseBody
	GetMessage() *string
	SetModule(v *UserQueryResponseBodyModule) *UserQueryResponseBody
	GetModule() *UserQueryResponseBodyModule
	SetRequestId(v string) *UserQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UserQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UserQueryResponseBody
	GetTraceId() *string
}

type UserQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Module  *UserQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 2f624a6316366024344424669e3279
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s UserQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UserQueryResponseBody) GoString() string {
	return s.String()
}

func (s *UserQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UserQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UserQueryResponseBody) GetModule() *UserQueryResponseBodyModule {
	return s.Module
}

func (s *UserQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UserQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UserQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UserQueryResponseBody) SetCode(v string) *UserQueryResponseBody {
	s.Code = &v
	return s
}

func (s *UserQueryResponseBody) SetMessage(v string) *UserQueryResponseBody {
	s.Message = &v
	return s
}

func (s *UserQueryResponseBody) SetModule(v *UserQueryResponseBodyModule) *UserQueryResponseBody {
	s.Module = v
	return s
}

func (s *UserQueryResponseBody) SetRequestId(v string) *UserQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UserQueryResponseBody) SetSuccess(v bool) *UserQueryResponseBody {
	s.Success = &v
	return s
}

func (s *UserQueryResponseBody) SetTraceId(v string) *UserQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *UserQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type UserQueryResponseBodyModule struct {
	// example:
	//
	// true
	HasMore *bool                               `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Items   []*UserQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 9YN+jxa7PcxbNUTISeKjEw==
	PageToken *string `json:"page_token,omitempty" xml:"page_token,omitempty"`
	// example:
	//
	// 0
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s UserQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UserQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UserQueryResponseBodyModule) GetHasMore() *bool {
	return s.HasMore
}

func (s *UserQueryResponseBodyModule) GetItems() []*UserQueryResponseBodyModuleItems {
	return s.Items
}

func (s *UserQueryResponseBodyModule) GetPageToken() *string {
	return s.PageToken
}

func (s *UserQueryResponseBodyModule) GetTotal() *int64 {
	return s.Total
}

func (s *UserQueryResponseBodyModule) SetHasMore(v bool) *UserQueryResponseBodyModule {
	s.HasMore = &v
	return s
}

func (s *UserQueryResponseBodyModule) SetItems(v []*UserQueryResponseBodyModuleItems) *UserQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *UserQueryResponseBodyModule) SetPageToken(v string) *UserQueryResponseBodyModule {
	s.PageToken = &v
	return s
}

func (s *UserQueryResponseBodyModule) SetTotal(v int64) *UserQueryResponseBodyModule {
	s.Total = &v
	return s
}

func (s *UserQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type UserQueryResponseBodyModuleItems struct {
	EmployeeNick *string `json:"employee_nick,omitempty" xml:"employee_nick,omitempty"`
	// example:
	//
	// 0
	LeaveStatus *int32 `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
	// example:
	//
	// 123
	ThirdPartEmployeeId *string `json:"third_part_employee_id,omitempty" xml:"third_part_employee_id,omitempty"`
	// example:
	//
	// 001
	ThirdPartJobNo *string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
}

func (s UserQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s UserQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *UserQueryResponseBodyModuleItems) GetEmployeeNick() *string {
	return s.EmployeeNick
}

func (s *UserQueryResponseBodyModuleItems) GetLeaveStatus() *int32 {
	return s.LeaveStatus
}

func (s *UserQueryResponseBodyModuleItems) GetThirdPartEmployeeId() *string {
	return s.ThirdPartEmployeeId
}

func (s *UserQueryResponseBodyModuleItems) GetThirdPartJobNo() *string {
	return s.ThirdPartJobNo
}

func (s *UserQueryResponseBodyModuleItems) SetEmployeeNick(v string) *UserQueryResponseBodyModuleItems {
	s.EmployeeNick = &v
	return s
}

func (s *UserQueryResponseBodyModuleItems) SetLeaveStatus(v int32) *UserQueryResponseBodyModuleItems {
	s.LeaveStatus = &v
	return s
}

func (s *UserQueryResponseBodyModuleItems) SetThirdPartEmployeeId(v string) *UserQueryResponseBodyModuleItems {
	s.ThirdPartEmployeeId = &v
	return s
}

func (s *UserQueryResponseBodyModuleItems) SetThirdPartJobNo(v string) *UserQueryResponseBodyModuleItems {
	s.ThirdPartJobNo = &v
	return s
}

func (s *UserQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
