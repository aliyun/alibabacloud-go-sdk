// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMaterialFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UploadMaterialFileResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UploadMaterialFileResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UploadMaterialFileResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UploadMaterialFileResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UploadMaterialFileResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UploadMaterialFileResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *UploadMaterialFileResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UploadMaterialFileResponseBody
	GetErrorMsg() *string
	SetModule(v *AppMaterialFile) *UploadMaterialFileResponseBody
	GetModule() *AppMaterialFile
	SetRequestId(v string) *UploadMaterialFileResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UploadMaterialFileResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UploadMaterialFileResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *UploadMaterialFileResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *UploadMaterialFileResponseBody
	GetSynchro() *bool
}

type UploadMaterialFileResponseBody struct {
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
	// {\\"Success\\": True}
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

func (s UploadMaterialFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UploadMaterialFileResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UploadMaterialFileResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UploadMaterialFileResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UploadMaterialFileResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UploadMaterialFileResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UploadMaterialFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UploadMaterialFileResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UploadMaterialFileResponseBody) GetModule() *AppMaterialFile {
	return s.Module
}

func (s *UploadMaterialFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadMaterialFileResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UploadMaterialFileResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UploadMaterialFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadMaterialFileResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UploadMaterialFileResponseBody) SetAccessDeniedDetail(v string) *UploadMaterialFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetAllowRetry(v bool) *UploadMaterialFileResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetAppName(v string) *UploadMaterialFileResponseBody {
	s.AppName = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetDynamicCode(v string) *UploadMaterialFileResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetDynamicMessage(v string) *UploadMaterialFileResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetErrorArgs(v []interface{}) *UploadMaterialFileResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UploadMaterialFileResponseBody) SetErrorCode(v string) *UploadMaterialFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetErrorMsg(v string) *UploadMaterialFileResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetModule(v *AppMaterialFile) *UploadMaterialFileResponseBody {
	s.Module = v
	return s
}

func (s *UploadMaterialFileResponseBody) SetRequestId(v string) *UploadMaterialFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetRootErrorCode(v string) *UploadMaterialFileResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetRootErrorMsg(v string) *UploadMaterialFileResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetSuccess(v bool) *UploadMaterialFileResponseBody {
	s.Success = &v
	return s
}

func (s *UploadMaterialFileResponseBody) SetSynchro(v bool) *UploadMaterialFileResponseBody {
	s.Synchro = &v
	return s
}

func (s *UploadMaterialFileResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}
