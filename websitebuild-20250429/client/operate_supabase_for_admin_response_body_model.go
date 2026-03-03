// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSupabaseForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OperateSupabaseForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *OperateSupabaseForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *OperateSupabaseForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *OperateSupabaseForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *OperateSupabaseForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *OperateSupabaseForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *OperateSupabaseForAdminResponseBody
	GetModule() *string
	SetRequestId(v string) *OperateSupabaseForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *OperateSupabaseForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *OperateSupabaseForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *OperateSupabaseForAdminResponseBody
	GetSynchro() *bool
}

type OperateSupabaseForAdminResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s OperateSupabaseForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateSupabaseForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *OperateSupabaseForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OperateSupabaseForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *OperateSupabaseForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *OperateSupabaseForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *OperateSupabaseForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *OperateSupabaseForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *OperateSupabaseForAdminResponseBody) GetModule() *string {
	return s.Module
}

func (s *OperateSupabaseForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateSupabaseForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *OperateSupabaseForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *OperateSupabaseForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *OperateSupabaseForAdminResponseBody) SetAccessDeniedDetail(v string) *OperateSupabaseForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetAllowRetry(v bool) *OperateSupabaseForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetAppName(v string) *OperateSupabaseForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetDynamicCode(v string) *OperateSupabaseForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetDynamicMessage(v string) *OperateSupabaseForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetErrorArgs(v []interface{}) *OperateSupabaseForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetModule(v string) *OperateSupabaseForAdminResponseBody {
	s.Module = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetRequestId(v string) *OperateSupabaseForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetRootErrorCode(v string) *OperateSupabaseForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetRootErrorMsg(v string) *OperateSupabaseForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) SetSynchro(v bool) *OperateSupabaseForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *OperateSupabaseForAdminResponseBody) Validate() error {
	return dara.Validate(s)
}
