// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAppRequirementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SaveAppRequirementResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SaveAppRequirementResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SaveAppRequirementResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SaveAppRequirementResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SaveAppRequirementResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SaveAppRequirementResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *SaveAppRequirementResponseBody
	GetModule() *bool
	SetRequestId(v string) *SaveAppRequirementResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SaveAppRequirementResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SaveAppRequirementResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SaveAppRequirementResponseBody
	GetSynchro() *bool
}

type SaveAppRequirementResponseBody struct {
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
	// true
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
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

func (s SaveAppRequirementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAppRequirementResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAppRequirementResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SaveAppRequirementResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SaveAppRequirementResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SaveAppRequirementResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SaveAppRequirementResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SaveAppRequirementResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SaveAppRequirementResponseBody) GetModule() *bool {
	return s.Module
}

func (s *SaveAppRequirementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAppRequirementResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SaveAppRequirementResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SaveAppRequirementResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SaveAppRequirementResponseBody) SetAccessDeniedDetail(v string) *SaveAppRequirementResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetAllowRetry(v bool) *SaveAppRequirementResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetAppName(v string) *SaveAppRequirementResponseBody {
	s.AppName = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetDynamicCode(v string) *SaveAppRequirementResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetDynamicMessage(v string) *SaveAppRequirementResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetErrorArgs(v []interface{}) *SaveAppRequirementResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SaveAppRequirementResponseBody) SetModule(v bool) *SaveAppRequirementResponseBody {
	s.Module = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetRequestId(v string) *SaveAppRequirementResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetRootErrorCode(v string) *SaveAppRequirementResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetRootErrorMsg(v string) *SaveAppRequirementResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SaveAppRequirementResponseBody) SetSynchro(v bool) *SaveAppRequirementResponseBody {
	s.Synchro = &v
	return s
}

func (s *SaveAppRequirementResponseBody) Validate() error {
	return dara.Validate(s)
}
