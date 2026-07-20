// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchQueryDepartmentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BatchQueryDepartmentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchQueryDepartmentResponseBody
	GetMessage() *string
	SetModule(v *BatchQueryDepartmentResponseBodyModule) *BatchQueryDepartmentResponseBody
	GetModule() *BatchQueryDepartmentResponseBodyModule
	SetRequestId(v string) *BatchQueryDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchQueryDepartmentResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchQueryDepartmentResponseBody
	GetTraceId() *string
}

type BatchQueryDepartmentResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 成功
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *BatchQueryDepartmentResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-****-****-****-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s BatchQueryDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchQueryDepartmentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchQueryDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchQueryDepartmentResponseBody) GetModule() *BatchQueryDepartmentResponseBodyModule {
	return s.Module
}

func (s *BatchQueryDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchQueryDepartmentResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchQueryDepartmentResponseBody) SetCode(v string) *BatchQueryDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryDepartmentResponseBody) SetHttpStatusCode(v int32) *BatchQueryDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchQueryDepartmentResponseBody) SetMessage(v string) *BatchQueryDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *BatchQueryDepartmentResponseBody) SetModule(v *BatchQueryDepartmentResponseBodyModule) *BatchQueryDepartmentResponseBody {
	s.Module = v
	return s
}

func (s *BatchQueryDepartmentResponseBody) SetRequestId(v string) *BatchQueryDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryDepartmentResponseBody) SetSuccess(v bool) *BatchQueryDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *BatchQueryDepartmentResponseBody) SetTraceId(v string) *BatchQueryDepartmentResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchQueryDepartmentResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchQueryDepartmentResponseBodyModule struct {
	// example:
	//
	// true
	HasMore *bool                                          `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Items   []*BatchQueryDepartmentResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// NjE1OTgwOTY
	NextCursorToken *string `json:"next_cursor_token,omitempty" xml:"next_cursor_token,omitempty"`
	// example:
	//
	// 0
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s BatchQueryDepartmentResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryDepartmentResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BatchQueryDepartmentResponseBodyModule) GetHasMore() *bool {
	return s.HasMore
}

func (s *BatchQueryDepartmentResponseBodyModule) GetItems() []*BatchQueryDepartmentResponseBodyModuleItems {
	return s.Items
}

func (s *BatchQueryDepartmentResponseBodyModule) GetNextCursorToken() *string {
	return s.NextCursorToken
}

func (s *BatchQueryDepartmentResponseBodyModule) GetTotal() *int64 {
	return s.Total
}

func (s *BatchQueryDepartmentResponseBodyModule) SetHasMore(v bool) *BatchQueryDepartmentResponseBodyModule {
	s.HasMore = &v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModule) SetItems(v []*BatchQueryDepartmentResponseBodyModuleItems) *BatchQueryDepartmentResponseBodyModule {
	s.Items = v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModule) SetNextCursorToken(v string) *BatchQueryDepartmentResponseBodyModule {
	s.NextCursorToken = &v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModule) SetTotal(v int64) *BatchQueryDepartmentResponseBodyModule {
	s.Total = &v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModule) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQueryDepartmentResponseBodyModuleItems struct {
	// example:
	//
	// 电磁继电器装配SL10线
	DeptName              *string   `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	ManagerEmployeeIdList []*string `json:"manager_employee_id_list,omitempty" xml:"manager_employee_id_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1335
	OutDeptId *string `json:"out_dept_id,omitempty" xml:"out_dept_id,omitempty"`
}

func (s BatchQueryDepartmentResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryDepartmentResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) GetDeptName() *string {
	return s.DeptName
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) GetManagerEmployeeIdList() []*string {
	return s.ManagerEmployeeIdList
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) SetDeptName(v string) *BatchQueryDepartmentResponseBodyModuleItems {
	s.DeptName = &v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) SetManagerEmployeeIdList(v []*string) *BatchQueryDepartmentResponseBodyModuleItems {
	s.ManagerEmployeeIdList = v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) SetOutDeptId(v string) *BatchQueryDepartmentResponseBodyModuleItems {
	s.OutDeptId = &v
	return s
}

func (s *BatchQueryDepartmentResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
