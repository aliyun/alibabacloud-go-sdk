// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSeoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSeoStatusResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSeoStatusResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSeoStatusResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSeoStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSeoStatusResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSeoStatusResponseBody
	GetErrorArgs() []interface{}
	SetModule(v []*GetAppSeoStatusResponseBodyModule) *GetAppSeoStatusResponseBody
	GetModule() []*GetAppSeoStatusResponseBodyModule
	SetRequestId(v string) *GetAppSeoStatusResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSeoStatusResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSeoStatusResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSeoStatusResponseBody
	GetSynchro() *bool
}

type GetAppSeoStatusResponseBody struct {
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
	DynamicMessage *string                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                        `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         []*GetAppSeoStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
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

func (s GetAppSeoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSeoStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSeoStatusResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSeoStatusResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSeoStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSeoStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSeoStatusResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSeoStatusResponseBody) GetModule() []*GetAppSeoStatusResponseBodyModule {
	return s.Module
}

func (s *GetAppSeoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSeoStatusResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSeoStatusResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSeoStatusResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSeoStatusResponseBody) SetAccessDeniedDetail(v string) *GetAppSeoStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetAllowRetry(v bool) *GetAppSeoStatusResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetAppName(v string) *GetAppSeoStatusResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetDynamicCode(v string) *GetAppSeoStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetDynamicMessage(v string) *GetAppSeoStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetErrorArgs(v []interface{}) *GetAppSeoStatusResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetModule(v []*GetAppSeoStatusResponseBodyModule) *GetAppSeoStatusResponseBody {
	s.Module = v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetRequestId(v string) *GetAppSeoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetRootErrorCode(v string) *GetAppSeoStatusResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetRootErrorMsg(v string) *GetAppSeoStatusResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) SetSynchro(v bool) *GetAppSeoStatusResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSeoStatusResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppSeoStatusResponseBodyModule struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// stxycw.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-03-26T13:25:41.119+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-07-04T00:47:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// authinfo
	SeAuthInfo *string `json:"SeAuthInfo,omitempty" xml:"SeAuthInfo,omitempty"`
	// example:
	//
	// 0,1,2
	SeIndexStatus *int32 `json:"SeIndexStatus,omitempty" xml:"SeIndexStatus,omitempty"`
	// example:
	//
	// baidu,bing,google
	SeType *string `json:"SeType,omitempty" xml:"SeType,omitempty"`
}

func (s GetAppSeoStatusResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoStatusResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppSeoStatusResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSeoStatusResponseBodyModule) GetDomain() *string {
	return s.Domain
}

func (s *GetAppSeoStatusResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAppSeoStatusResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetAppSeoStatusResponseBodyModule) GetSeAuthInfo() *string {
	return s.SeAuthInfo
}

func (s *GetAppSeoStatusResponseBodyModule) GetSeIndexStatus() *int32 {
	return s.SeIndexStatus
}

func (s *GetAppSeoStatusResponseBodyModule) GetSeType() *string {
	return s.SeType
}

func (s *GetAppSeoStatusResponseBodyModule) SetBizId(v string) *GetAppSeoStatusResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) SetDomain(v string) *GetAppSeoStatusResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) SetGmtCreateTime(v string) *GetAppSeoStatusResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) SetGmtModifiedTime(v string) *GetAppSeoStatusResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) SetSeAuthInfo(v string) *GetAppSeoStatusResponseBodyModule {
	s.SeAuthInfo = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) SetSeIndexStatus(v int32) *GetAppSeoStatusResponseBodyModule {
	s.SeIndexStatus = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) SetSeType(v string) *GetAppSeoStatusResponseBodyModule {
	s.SeType = &v
	return s
}

func (s *GetAppSeoStatusResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
