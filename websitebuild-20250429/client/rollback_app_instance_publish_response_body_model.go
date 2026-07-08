// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppInstancePublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RollbackAppInstancePublishResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RollbackAppInstancePublishResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RollbackAppInstancePublishResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RollbackAppInstancePublishResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RollbackAppInstancePublishResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RollbackAppInstancePublishResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RollbackAppInstancePublishResponseBodyModule) *RollbackAppInstancePublishResponseBody
	GetModule() *RollbackAppInstancePublishResponseBodyModule
	SetRequestId(v string) *RollbackAppInstancePublishResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RollbackAppInstancePublishResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RollbackAppInstancePublishResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RollbackAppInstancePublishResponseBody
	GetSynchro() *bool
}

type RollbackAppInstancePublishResponseBody struct {
	// The details about the access denial.
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
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
	Module *RollbackAppInstancePublishResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s RollbackAppInstancePublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppInstancePublishResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackAppInstancePublishResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RollbackAppInstancePublishResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RollbackAppInstancePublishResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RollbackAppInstancePublishResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RollbackAppInstancePublishResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RollbackAppInstancePublishResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RollbackAppInstancePublishResponseBody) GetModule() *RollbackAppInstancePublishResponseBodyModule {
	return s.Module
}

func (s *RollbackAppInstancePublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackAppInstancePublishResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RollbackAppInstancePublishResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RollbackAppInstancePublishResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RollbackAppInstancePublishResponseBody) SetAccessDeniedDetail(v string) *RollbackAppInstancePublishResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetAllowRetry(v bool) *RollbackAppInstancePublishResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetAppName(v string) *RollbackAppInstancePublishResponseBody {
	s.AppName = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetDynamicCode(v string) *RollbackAppInstancePublishResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetDynamicMessage(v string) *RollbackAppInstancePublishResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetErrorArgs(v []interface{}) *RollbackAppInstancePublishResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetModule(v *RollbackAppInstancePublishResponseBodyModule) *RollbackAppInstancePublishResponseBody {
	s.Module = v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetRequestId(v string) *RollbackAppInstancePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetRootErrorCode(v string) *RollbackAppInstancePublishResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetRootErrorMsg(v string) *RollbackAppInstancePublishResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) SetSynchro(v bool) *RollbackAppInstancePublishResponseBody {
	s.Synchro = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RollbackAppInstancePublishResponseBodyModule struct {
	// The publish order ID.
	//
	// example:
	//
	// 123
	PublishOrderId *int64 `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
}

func (s RollbackAppInstancePublishResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppInstancePublishResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RollbackAppInstancePublishResponseBodyModule) GetPublishOrderId() *int64 {
	return s.PublishOrderId
}

func (s *RollbackAppInstancePublishResponseBodyModule) SetPublishOrderId(v int64) *RollbackAppInstancePublishResponseBodyModule {
	s.PublishOrderId = &v
	return s
}

func (s *RollbackAppInstancePublishResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
