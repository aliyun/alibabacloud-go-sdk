// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceTempShortUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppInstanceTempShortUrlResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppInstanceTempShortUrlResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppInstanceTempShortUrlResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppInstanceTempShortUrlResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppInstanceTempShortUrlResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppInstanceTempShortUrlResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppInstanceTempShortUrlResponseBodyModule) *GetAppInstanceTempShortUrlResponseBody
	GetModule() *GetAppInstanceTempShortUrlResponseBodyModule
	SetRequestId(v string) *GetAppInstanceTempShortUrlResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppInstanceTempShortUrlResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppInstanceTempShortUrlResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppInstanceTempShortUrlResponseBody
	GetSynchro() *bool
}

type GetAppInstanceTempShortUrlResponseBody struct {
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
	DynamicMessage *string                                       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                 `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppInstanceTempShortUrlResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetAppInstanceTempShortUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceTempShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetModule() *GetAppInstanceTempShortUrlResponseBodyModule {
	return s.Module
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppInstanceTempShortUrlResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetAccessDeniedDetail(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetAllowRetry(v bool) *GetAppInstanceTempShortUrlResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetAppName(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetDynamicCode(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetDynamicMessage(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetErrorArgs(v []interface{}) *GetAppInstanceTempShortUrlResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetModule(v *GetAppInstanceTempShortUrlResponseBodyModule) *GetAppInstanceTempShortUrlResponseBody {
	s.Module = v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetRequestId(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetRootErrorCode(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetRootErrorMsg(v string) *GetAppInstanceTempShortUrlResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) SetSynchro(v bool) *GetAppInstanceTempShortUrlResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceTempShortUrlResponseBodyModule struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 123456
	ExpireAtTime *string `json:"ExpireAtTime,omitempty" xml:"ExpireAtTime,omitempty"`
	// example:
	//
	// http://www.aliyun.com
	TempShortUrl *string `json:"TempShortUrl,omitempty" xml:"TempShortUrl,omitempty"`
	// example:
	//
	// https://bj.download.cycore.cn/zhkt-student-cystore-https/2025/9/31/10/12/9b48342f-e595-4c60-8032-dccd355e7552.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAppInstanceTempShortUrlResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceTempShortUrlResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) GetExpireAtTime() *string {
	return s.ExpireAtTime
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) GetTempShortUrl() *string {
	return s.TempShortUrl
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) GetUrl() *string {
	return s.Url
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) SetExpireAtTime(v string) *GetAppInstanceTempShortUrlResponseBodyModule {
	s.ExpireAtTime = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) SetTempShortUrl(v string) *GetAppInstanceTempShortUrlResponseBodyModule {
	s.TempShortUrl = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) SetUrl(v string) *GetAppInstanceTempShortUrlResponseBodyModule {
	s.Url = &v
	return s
}

func (s *GetAppInstanceTempShortUrlResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
