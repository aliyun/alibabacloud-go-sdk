// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIStaffPreviewUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAIStaffPreviewUrlResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAIStaffPreviewUrlResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAIStaffPreviewUrlResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAIStaffPreviewUrlResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAIStaffPreviewUrlResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAIStaffPreviewUrlResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAIStaffPreviewUrlResponseBodyModule) *GetAIStaffPreviewUrlResponseBody
	GetModule() *GetAIStaffPreviewUrlResponseBodyModule
	SetRequestId(v string) *GetAIStaffPreviewUrlResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAIStaffPreviewUrlResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAIStaffPreviewUrlResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAIStaffPreviewUrlResponseBody
	GetSynchro() *bool
}

type GetAIStaffPreviewUrlResponseBody struct {
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
	// abc
	DynamicMessage *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                           `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAIStaffPreviewUrlResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetAIStaffPreviewUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIStaffPreviewUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIStaffPreviewUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAIStaffPreviewUrlResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAIStaffPreviewUrlResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAIStaffPreviewUrlResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAIStaffPreviewUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAIStaffPreviewUrlResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAIStaffPreviewUrlResponseBody) GetModule() *GetAIStaffPreviewUrlResponseBodyModule {
	return s.Module
}

func (s *GetAIStaffPreviewUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIStaffPreviewUrlResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAIStaffPreviewUrlResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAIStaffPreviewUrlResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAIStaffPreviewUrlResponseBody) SetAccessDeniedDetail(v string) *GetAIStaffPreviewUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetAllowRetry(v bool) *GetAIStaffPreviewUrlResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetAppName(v string) *GetAIStaffPreviewUrlResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetDynamicCode(v string) *GetAIStaffPreviewUrlResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetDynamicMessage(v string) *GetAIStaffPreviewUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetErrorArgs(v []interface{}) *GetAIStaffPreviewUrlResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetModule(v *GetAIStaffPreviewUrlResponseBodyModule) *GetAIStaffPreviewUrlResponseBody {
	s.Module = v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetRequestId(v string) *GetAIStaffPreviewUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetRootErrorCode(v string) *GetAIStaffPreviewUrlResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetRootErrorMsg(v string) *GetAIStaffPreviewUrlResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) SetSynchro(v bool) *GetAIStaffPreviewUrlResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIStaffPreviewUrlResponseBodyModule struct {
	UrlMap map[string]*string `json:"UrlMap,omitempty" xml:"UrlMap,omitempty"`
}

func (s GetAIStaffPreviewUrlResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAIStaffPreviewUrlResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAIStaffPreviewUrlResponseBodyModule) GetUrlMap() map[string]*string {
	return s.UrlMap
}

func (s *GetAIStaffPreviewUrlResponseBodyModule) SetUrlMap(v map[string]*string) *GetAIStaffPreviewUrlResponseBodyModule {
	s.UrlMap = v
	return s
}

func (s *GetAIStaffPreviewUrlResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
