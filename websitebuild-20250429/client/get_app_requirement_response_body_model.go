// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRequirementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppRequirementResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppRequirementResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppRequirementResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppRequirementResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppRequirementResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppRequirementResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppRequirementResponseBodyModule) *GetAppRequirementResponseBody
	GetModule() *GetAppRequirementResponseBodyModule
	SetRequestId(v string) *GetAppRequirementResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppRequirementResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppRequirementResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppRequirementResponseBody
	GetSynchro() *bool
}

type GetAppRequirementResponseBody struct {
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
	DynamicMessage *string                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                        `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppRequirementResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetAppRequirementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppRequirementResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppRequirementResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppRequirementResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppRequirementResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppRequirementResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppRequirementResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppRequirementResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppRequirementResponseBody) GetModule() *GetAppRequirementResponseBodyModule {
	return s.Module
}

func (s *GetAppRequirementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppRequirementResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppRequirementResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppRequirementResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppRequirementResponseBody) SetAccessDeniedDetail(v string) *GetAppRequirementResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetAllowRetry(v bool) *GetAppRequirementResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetAppName(v string) *GetAppRequirementResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetDynamicCode(v string) *GetAppRequirementResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetDynamicMessage(v string) *GetAppRequirementResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetErrorArgs(v []interface{}) *GetAppRequirementResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppRequirementResponseBody) SetModule(v *GetAppRequirementResponseBodyModule) *GetAppRequirementResponseBody {
	s.Module = v
	return s
}

func (s *GetAppRequirementResponseBody) SetRequestId(v string) *GetAppRequirementResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetRootErrorCode(v string) *GetAppRequirementResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetRootErrorMsg(v string) *GetAppRequirementResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppRequirementResponseBody) SetSynchro(v bool) *GetAppRequirementResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppRequirementResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppRequirementResponseBodyModule struct {
	// example:
	//
	// content
	Prd *string `json:"Prd,omitempty" xml:"Prd,omitempty"`
}

func (s GetAppRequirementResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppRequirementResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppRequirementResponseBodyModule) GetPrd() *string {
	return s.Prd
}

func (s *GetAppRequirementResponseBodyModule) SetPrd(v string) *GetAppRequirementResponseBodyModule {
	s.Prd = &v
	return s
}

func (s *GetAppRequirementResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
