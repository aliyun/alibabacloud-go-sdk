// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSeoTrendsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSeoTrendsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSeoTrendsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSeoTrendsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSeoTrendsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSeoTrendsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSeoTrendsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppSeoTrendsResponseBodyModule) *GetAppSeoTrendsResponseBody
	GetModule() *GetAppSeoTrendsResponseBodyModule
	SetRequestId(v string) *GetAppSeoTrendsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSeoTrendsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSeoTrendsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSeoTrendsResponseBody
	GetSynchro() *bool
}

type GetAppSeoTrendsResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application name
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *GetAppSeoTrendsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Backup parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppSeoTrendsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoTrendsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSeoTrendsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSeoTrendsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSeoTrendsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSeoTrendsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSeoTrendsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSeoTrendsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSeoTrendsResponseBody) GetModule() *GetAppSeoTrendsResponseBodyModule {
	return s.Module
}

func (s *GetAppSeoTrendsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSeoTrendsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSeoTrendsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSeoTrendsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSeoTrendsResponseBody) SetAccessDeniedDetail(v string) *GetAppSeoTrendsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetAllowRetry(v bool) *GetAppSeoTrendsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetAppName(v string) *GetAppSeoTrendsResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetDynamicCode(v string) *GetAppSeoTrendsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetDynamicMessage(v string) *GetAppSeoTrendsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetErrorArgs(v []interface{}) *GetAppSeoTrendsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetModule(v *GetAppSeoTrendsResponseBodyModule) *GetAppSeoTrendsResponseBody {
	s.Module = v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetRequestId(v string) *GetAppSeoTrendsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetRootErrorCode(v string) *GetAppSeoTrendsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetRootErrorMsg(v string) *GetAppSeoTrendsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) SetSynchro(v bool) *GetAppSeoTrendsResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSeoTrendsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppSeoTrendsResponseBodyModule struct {
	// Baidu
	//
	// example:
	//
	// baiduIndexTrends
	BaiduIndexTrends *string `json:"BaiduIndexTrends,omitempty" xml:"BaiduIndexTrends,omitempty"`
	// Bing
	//
	// example:
	//
	// bingIndexTrends
	BingIndexTrends *string `json:"BingIndexTrends,omitempty" xml:"BingIndexTrends,omitempty"`
	// Google
	//
	// example:
	//
	// googleIndexTrends
	GoogleIndexTrends *string `json:"GoogleIndexTrends,omitempty" xml:"GoogleIndexTrends,omitempty"`
}

func (s GetAppSeoTrendsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoTrendsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppSeoTrendsResponseBodyModule) GetBaiduIndexTrends() *string {
	return s.BaiduIndexTrends
}

func (s *GetAppSeoTrendsResponseBodyModule) GetBingIndexTrends() *string {
	return s.BingIndexTrends
}

func (s *GetAppSeoTrendsResponseBodyModule) GetGoogleIndexTrends() *string {
	return s.GoogleIndexTrends
}

func (s *GetAppSeoTrendsResponseBodyModule) SetBaiduIndexTrends(v string) *GetAppSeoTrendsResponseBodyModule {
	s.BaiduIndexTrends = &v
	return s
}

func (s *GetAppSeoTrendsResponseBodyModule) SetBingIndexTrends(v string) *GetAppSeoTrendsResponseBodyModule {
	s.BingIndexTrends = &v
	return s
}

func (s *GetAppSeoTrendsResponseBodyModule) SetGoogleIndexTrends(v string) *GetAppSeoTrendsResponseBodyModule {
	s.GoogleIndexTrends = &v
	return s
}

func (s *GetAppSeoTrendsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
