// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppFileContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppFileContentResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppFileContentResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppFileContentResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppFileContentResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppFileContentResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppFileContentResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppFileContentResponseBodyModule) *GetAppFileContentResponseBody
	GetModule() *GetAppFileContentResponseBodyModule
	SetRequestId(v string) *GetAppFileContentResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppFileContentResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppFileContentResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppFileContentResponseBody
	GetSynchro() *bool
}

type GetAppFileContentResponseBody struct {
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
	Module         *GetAppFileContentResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppFileContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppFileContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppFileContentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppFileContentResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppFileContentResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppFileContentResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppFileContentResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppFileContentResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppFileContentResponseBody) GetModule() *GetAppFileContentResponseBodyModule {
	return s.Module
}

func (s *GetAppFileContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppFileContentResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppFileContentResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppFileContentResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppFileContentResponseBody) SetAccessDeniedDetail(v string) *GetAppFileContentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetAllowRetry(v bool) *GetAppFileContentResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetAppName(v string) *GetAppFileContentResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetDynamicCode(v string) *GetAppFileContentResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetDynamicMessage(v string) *GetAppFileContentResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetErrorArgs(v []interface{}) *GetAppFileContentResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppFileContentResponseBody) SetModule(v *GetAppFileContentResponseBodyModule) *GetAppFileContentResponseBody {
	s.Module = v
	return s
}

func (s *GetAppFileContentResponseBody) SetRequestId(v string) *GetAppFileContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetRootErrorCode(v string) *GetAppFileContentResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetRootErrorMsg(v string) *GetAppFileContentResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppFileContentResponseBody) SetSynchro(v bool) *GetAppFileContentResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppFileContentResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppFileContentResponseBodyModule struct {
	// example:
	//
	// domain cnamensdxdjq.com sdxdjq.com.a1.initrr.comn*.sdxdjq.com all.sdxdjq.com.a1.initrr.comn
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// currentTime
	CurrentTime *string `json:"CurrentTime,omitempty" xml:"CurrentTime,omitempty"`
}

func (s GetAppFileContentResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppFileContentResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppFileContentResponseBodyModule) GetContent() *string {
	return s.Content
}

func (s *GetAppFileContentResponseBodyModule) GetCurrentTime() *string {
	return s.CurrentTime
}

func (s *GetAppFileContentResponseBodyModule) SetContent(v string) *GetAppFileContentResponseBodyModule {
	s.Content = &v
	return s
}

func (s *GetAppFileContentResponseBodyModule) SetCurrentTime(v string) *GetAppFileContentResponseBodyModule {
	s.CurrentTime = &v
	return s
}

func (s *GetAppFileContentResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
