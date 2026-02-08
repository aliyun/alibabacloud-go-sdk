// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMaterialFileDetailResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryMaterialFileDetailResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryMaterialFileDetailResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryMaterialFileDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryMaterialFileDetailResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryMaterialFileDetailResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryMaterialFileDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryMaterialFileDetailResponseBody
	GetErrorMsg() *string
	SetModule(v *AppMaterialFile) *QueryMaterialFileDetailResponseBody
	GetModule() *AppMaterialFile
	SetRequestId(v string) *QueryMaterialFileDetailResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryMaterialFileDetailResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryMaterialFileDetailResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *QueryMaterialFileDetailResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryMaterialFileDetailResponseBody
	GetSynchro() *bool
}

type QueryMaterialFileDetailResponseBody struct {
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
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// {\\"TotalPageNum\\": 1, \\"ResultLimit\\": False, \\"CurrentPageNum\\": 0, \\"PageSize\\": 0, \\"TotalItemNum\\": 0}
	Module *AppMaterialFile `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryMaterialFileDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMaterialFileDetailResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryMaterialFileDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryMaterialFileDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryMaterialFileDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryMaterialFileDetailResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryMaterialFileDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMaterialFileDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryMaterialFileDetailResponseBody) GetModule() *AppMaterialFile {
	return s.Module
}

func (s *QueryMaterialFileDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMaterialFileDetailResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryMaterialFileDetailResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryMaterialFileDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMaterialFileDetailResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryMaterialFileDetailResponseBody) SetAccessDeniedDetail(v string) *QueryMaterialFileDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetAllowRetry(v bool) *QueryMaterialFileDetailResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetAppName(v string) *QueryMaterialFileDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetDynamicCode(v string) *QueryMaterialFileDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetDynamicMessage(v string) *QueryMaterialFileDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetErrorArgs(v []interface{}) *QueryMaterialFileDetailResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetErrorCode(v string) *QueryMaterialFileDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetErrorMsg(v string) *QueryMaterialFileDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetModule(v *AppMaterialFile) *QueryMaterialFileDetailResponseBody {
	s.Module = v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetRequestId(v string) *QueryMaterialFileDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetRootErrorCode(v string) *QueryMaterialFileDetailResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetRootErrorMsg(v string) *QueryMaterialFileDetailResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetSuccess(v bool) *QueryMaterialFileDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) SetSynchro(v bool) *QueryMaterialFileDetailResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryMaterialFileDetailResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}
