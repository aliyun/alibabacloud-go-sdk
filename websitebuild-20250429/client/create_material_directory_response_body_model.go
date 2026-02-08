// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaterialDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateMaterialDirectoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateMaterialDirectoryResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateMaterialDirectoryResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateMaterialDirectoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateMaterialDirectoryResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateMaterialDirectoryResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *CreateMaterialDirectoryResponseBody
	GetErrorCode() *string
	SetModule(v *CreateMaterialDirectoryResponseBodyModule) *CreateMaterialDirectoryResponseBody
	GetModule() *CreateMaterialDirectoryResponseBodyModule
	SetRequestId(v string) *CreateMaterialDirectoryResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateMaterialDirectoryResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateMaterialDirectoryResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *CreateMaterialDirectoryResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *CreateMaterialDirectoryResponseBody
	GetSynchro() *bool
}

type CreateMaterialDirectoryResponseBody struct {
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
	// ish-intelligence-store-platform-admin-web
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
	ErrorCode *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Module    *CreateMaterialDirectoryResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateMaterialDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterialDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaterialDirectoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateMaterialDirectoryResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateMaterialDirectoryResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateMaterialDirectoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateMaterialDirectoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateMaterialDirectoryResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateMaterialDirectoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMaterialDirectoryResponseBody) GetModule() *CreateMaterialDirectoryResponseBodyModule {
	return s.Module
}

func (s *CreateMaterialDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaterialDirectoryResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateMaterialDirectoryResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateMaterialDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMaterialDirectoryResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateMaterialDirectoryResponseBody) SetAccessDeniedDetail(v string) *CreateMaterialDirectoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetAllowRetry(v bool) *CreateMaterialDirectoryResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetAppName(v string) *CreateMaterialDirectoryResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetDynamicCode(v string) *CreateMaterialDirectoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetDynamicMessage(v string) *CreateMaterialDirectoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetErrorArgs(v []interface{}) *CreateMaterialDirectoryResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetErrorCode(v string) *CreateMaterialDirectoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetModule(v *CreateMaterialDirectoryResponseBodyModule) *CreateMaterialDirectoryResponseBody {
	s.Module = v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetRequestId(v string) *CreateMaterialDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetRootErrorCode(v string) *CreateMaterialDirectoryResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetRootErrorMsg(v string) *CreateMaterialDirectoryResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetSuccess(v bool) *CreateMaterialDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) SetSynchro(v bool) *CreateMaterialDirectoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMaterialDirectoryResponseBodyModule struct {
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s CreateMaterialDirectoryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterialDirectoryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateMaterialDirectoryResponseBodyModule) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateMaterialDirectoryResponseBodyModule) SetDirectoryId(v string) *CreateMaterialDirectoryResponseBodyModule {
	s.DirectoryId = &v
	return s
}

func (s *CreateMaterialDirectoryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
