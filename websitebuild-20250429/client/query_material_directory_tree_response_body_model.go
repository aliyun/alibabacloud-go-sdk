// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialDirectoryTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMaterialDirectoryTreeResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryMaterialDirectoryTreeResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryMaterialDirectoryTreeResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryMaterialDirectoryTreeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryMaterialDirectoryTreeResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryMaterialDirectoryTreeResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryMaterialDirectoryTreeResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryMaterialDirectoryTreeResponseBody
	GetErrorMsg() *string
	SetModule(v *QueryMaterialDirectoryTreeResponseBodyModule) *QueryMaterialDirectoryTreeResponseBody
	GetModule() *QueryMaterialDirectoryTreeResponseBodyModule
	SetRequestId(v string) *QueryMaterialDirectoryTreeResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryMaterialDirectoryTreeResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryMaterialDirectoryTreeResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *QueryMaterialDirectoryTreeResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryMaterialDirectoryTreeResponseBody
	GetSynchro() *bool
}

type QueryMaterialDirectoryTreeResponseBody struct {
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
	// dewuApp
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
	ErrorMsg *string                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Module   *QueryMaterialDirectoryTreeResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s QueryMaterialDirectoryTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialDirectoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetModule() *QueryMaterialDirectoryTreeResponseBodyModule {
	return s.Module
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMaterialDirectoryTreeResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetAccessDeniedDetail(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetAllowRetry(v bool) *QueryMaterialDirectoryTreeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetAppName(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetDynamicCode(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetDynamicMessage(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetErrorArgs(v []interface{}) *QueryMaterialDirectoryTreeResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetErrorCode(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetErrorMsg(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetModule(v *QueryMaterialDirectoryTreeResponseBodyModule) *QueryMaterialDirectoryTreeResponseBody {
	s.Module = v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetRequestId(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetRootErrorCode(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetRootErrorMsg(v string) *QueryMaterialDirectoryTreeResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetSuccess(v bool) *QueryMaterialDirectoryTreeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) SetSynchro(v bool) *QueryMaterialDirectoryTreeResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMaterialDirectoryTreeResponseBodyModule struct {
	DirectoryList []*AppMaterialDirectory `json:"DirectoryList,omitempty" xml:"DirectoryList,omitempty" type:"Repeated"`
}

func (s QueryMaterialDirectoryTreeResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialDirectoryTreeResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryMaterialDirectoryTreeResponseBodyModule) GetDirectoryList() []*AppMaterialDirectory {
	return s.DirectoryList
}

func (s *QueryMaterialDirectoryTreeResponseBodyModule) SetDirectoryList(v []*AppMaterialDirectory) *QueryMaterialDirectoryTreeResponseBodyModule {
	s.DirectoryList = v
	return s
}

func (s *QueryMaterialDirectoryTreeResponseBodyModule) Validate() error {
	if s.DirectoryList != nil {
		for _, item := range s.DirectoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
