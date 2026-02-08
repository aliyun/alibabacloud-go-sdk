// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileSummaryInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryMaterialFileSummaryInfoResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryMaterialFileSummaryInfoResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryMaterialFileSummaryInfoResponseBodyModule) *QueryMaterialFileSummaryInfoResponseBody
	GetModule() *QueryMaterialFileSummaryInfoResponseBodyModule
	SetRequestId(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryMaterialFileSummaryInfoResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryMaterialFileSummaryInfoResponseBody
	GetSynchro() *bool
}

type QueryMaterialFileSummaryInfoResponseBody struct {
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
	// SYSTEM_ERROR
	DynamicMessage *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                   `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *QueryMaterialFileSummaryInfoResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.EROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryMaterialFileSummaryInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetModule() *QueryMaterialFileSummaryInfoResponseBodyModule {
	return s.Module
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryMaterialFileSummaryInfoResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetAccessDeniedDetail(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetAllowRetry(v bool) *QueryMaterialFileSummaryInfoResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetAppName(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetDynamicCode(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetDynamicMessage(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetErrorArgs(v []interface{}) *QueryMaterialFileSummaryInfoResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetModule(v *QueryMaterialFileSummaryInfoResponseBodyModule) *QueryMaterialFileSummaryInfoResponseBody {
	s.Module = v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetRequestId(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetRootErrorCode(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetRootErrorMsg(v string) *QueryMaterialFileSummaryInfoResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) SetSynchro(v bool) *QueryMaterialFileSummaryInfoResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMaterialFileSummaryInfoResponseBodyModule struct {
	// example:
	//
	// 10
	FileNum *int64 `json:"FileNum,omitempty" xml:"FileNum,omitempty"`
	// example:
	//
	// 23M
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s QueryMaterialFileSummaryInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileSummaryInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileSummaryInfoResponseBodyModule) GetFileNum() *int64 {
	return s.FileNum
}

func (s *QueryMaterialFileSummaryInfoResponseBodyModule) GetStorageSize() *string {
	return s.StorageSize
}

func (s *QueryMaterialFileSummaryInfoResponseBodyModule) SetFileNum(v int64) *QueryMaterialFileSummaryInfoResponseBodyModule {
	s.FileNum = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBodyModule) SetStorageSize(v string) *QueryMaterialFileSummaryInfoResponseBodyModule {
	s.StorageSize = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
