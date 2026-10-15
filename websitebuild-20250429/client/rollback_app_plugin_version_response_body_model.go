// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppPluginVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RollbackAppPluginVersionResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RollbackAppPluginVersionResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RollbackAppPluginVersionResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RollbackAppPluginVersionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RollbackAppPluginVersionResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RollbackAppPluginVersionResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RollbackAppPluginVersionResponseBodyModule) *RollbackAppPluginVersionResponseBody
	GetModule() *RollbackAppPluginVersionResponseBodyModule
	SetRequestId(v string) *RollbackAppPluginVersionResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RollbackAppPluginVersionResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RollbackAppPluginVersionResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RollbackAppPluginVersionResponseBody
	GetSynchro() *bool
}

type RollbackAppPluginVersionResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *RollbackAppPluginVersionResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// System exception.
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s RollbackAppPluginVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppPluginVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackAppPluginVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RollbackAppPluginVersionResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RollbackAppPluginVersionResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RollbackAppPluginVersionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RollbackAppPluginVersionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RollbackAppPluginVersionResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RollbackAppPluginVersionResponseBody) GetModule() *RollbackAppPluginVersionResponseBodyModule {
	return s.Module
}

func (s *RollbackAppPluginVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackAppPluginVersionResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RollbackAppPluginVersionResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RollbackAppPluginVersionResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RollbackAppPluginVersionResponseBody) SetAccessDeniedDetail(v string) *RollbackAppPluginVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetAllowRetry(v bool) *RollbackAppPluginVersionResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetAppName(v string) *RollbackAppPluginVersionResponseBody {
	s.AppName = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetDynamicCode(v string) *RollbackAppPluginVersionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetDynamicMessage(v string) *RollbackAppPluginVersionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetErrorArgs(v []interface{}) *RollbackAppPluginVersionResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetModule(v *RollbackAppPluginVersionResponseBodyModule) *RollbackAppPluginVersionResponseBody {
	s.Module = v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetRequestId(v string) *RollbackAppPluginVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetRootErrorCode(v string) *RollbackAppPluginVersionResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetRootErrorMsg(v string) *RollbackAppPluginVersionResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) SetSynchro(v bool) *RollbackAppPluginVersionResponseBody {
	s.Synchro = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RollbackAppPluginVersionResponseBodyModule struct {
	// The status.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackAppPluginVersionResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppPluginVersionResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RollbackAppPluginVersionResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *RollbackAppPluginVersionResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackAppPluginVersionResponseBodyModule) SetStatus(v string) *RollbackAppPluginVersionResponseBodyModule {
	s.Status = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBodyModule) SetSuccess(v bool) *RollbackAppPluginVersionResponseBodyModule {
	s.Success = &v
	return s
}

func (s *RollbackAppPluginVersionResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
