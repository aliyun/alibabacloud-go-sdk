// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmployeeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEmployeeDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QueryEmployeeDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryEmployeeDetailResponseBody
	GetMessage() *string
	SetModule(v *QueryEmployeeDetailResponseBodyModule) *QueryEmployeeDetailResponseBody
	GetModule() *QueryEmployeeDetailResponseBodyModule
	SetRequestId(v string) *QueryEmployeeDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEmployeeDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryEmployeeDetailResponseBody
	GetTraceId() *string
}

type QueryEmployeeDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// None
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *QueryEmployeeDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s QueryEmployeeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEmployeeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmployeeDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEmployeeDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryEmployeeDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEmployeeDetailResponseBody) GetModule() *QueryEmployeeDetailResponseBodyModule {
	return s.Module
}

func (s *QueryEmployeeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEmployeeDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEmployeeDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryEmployeeDetailResponseBody) SetCode(v string) *QueryEmployeeDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEmployeeDetailResponseBody) SetHttpStatusCode(v int32) *QueryEmployeeDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryEmployeeDetailResponseBody) SetMessage(v string) *QueryEmployeeDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEmployeeDetailResponseBody) SetModule(v *QueryEmployeeDetailResponseBodyModule) *QueryEmployeeDetailResponseBody {
	s.Module = v
	return s
}

func (s *QueryEmployeeDetailResponseBody) SetRequestId(v string) *QueryEmployeeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmployeeDetailResponseBody) SetSuccess(v bool) *QueryEmployeeDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEmployeeDetailResponseBody) SetTraceId(v string) *QueryEmployeeDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryEmployeeDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryEmployeeDetailResponseBodyModule struct {
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// false
	IsLeave *bool `json:"is_leave,omitempty" xml:"is_leave,omitempty"`
	// example:
	//
	// "12138"
	JobNumber *string `json:"job_number,omitempty" xml:"job_number,omitempty"`
	NickName  *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// example:
	//
	// "123456"
	OutDeptId     *string   `json:"out_dept_id,omitempty" xml:"out_dept_id,omitempty"`
	OutDeptIdList []*string `json:"out_dept_id_list,omitempty" xml:"out_dept_id_list,omitempty" type:"Repeated"`
	// example:
	//
	// "abc12138"
	OutEmployeeId *string `json:"out_employee_id,omitempty" xml:"out_employee_id,omitempty"`
	PhoneNo       *string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
	RealName      *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// example:
	//
	// "zhang/san"
	RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
}

func (s QueryEmployeeDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryEmployeeDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryEmployeeDetailResponseBodyModule) GetEmail() *string {
	return s.Email
}

func (s *QueryEmployeeDetailResponseBodyModule) GetIsLeave() *bool {
	return s.IsLeave
}

func (s *QueryEmployeeDetailResponseBodyModule) GetJobNumber() *string {
	return s.JobNumber
}

func (s *QueryEmployeeDetailResponseBodyModule) GetNickName() *string {
	return s.NickName
}

func (s *QueryEmployeeDetailResponseBodyModule) GetOutDeptId() *string {
	return s.OutDeptId
}

func (s *QueryEmployeeDetailResponseBodyModule) GetOutDeptIdList() []*string {
	return s.OutDeptIdList
}

func (s *QueryEmployeeDetailResponseBodyModule) GetOutEmployeeId() *string {
	return s.OutEmployeeId
}

func (s *QueryEmployeeDetailResponseBodyModule) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *QueryEmployeeDetailResponseBodyModule) GetRealName() *string {
	return s.RealName
}

func (s *QueryEmployeeDetailResponseBodyModule) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *QueryEmployeeDetailResponseBodyModule) SetEmail(v string) *QueryEmployeeDetailResponseBodyModule {
	s.Email = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetIsLeave(v bool) *QueryEmployeeDetailResponseBodyModule {
	s.IsLeave = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetJobNumber(v string) *QueryEmployeeDetailResponseBodyModule {
	s.JobNumber = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetNickName(v string) *QueryEmployeeDetailResponseBodyModule {
	s.NickName = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetOutDeptId(v string) *QueryEmployeeDetailResponseBodyModule {
	s.OutDeptId = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetOutDeptIdList(v []*string) *QueryEmployeeDetailResponseBodyModule {
	s.OutDeptIdList = v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetOutEmployeeId(v string) *QueryEmployeeDetailResponseBodyModule {
	s.OutEmployeeId = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetPhoneNo(v string) *QueryEmployeeDetailResponseBodyModule {
	s.PhoneNo = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetRealName(v string) *QueryEmployeeDetailResponseBodyModule {
	s.RealName = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) SetRealNameEn(v string) *QueryEmployeeDetailResponseBodyModule {
	s.RealNameEn = &v
	return s
}

func (s *QueryEmployeeDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
