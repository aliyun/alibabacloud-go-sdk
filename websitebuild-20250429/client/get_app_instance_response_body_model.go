// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *AppInstanceAggregate) *GetAppInstanceResponseBody
	GetModule() *AppInstanceAggregate
	SetRequestId(v string) *GetAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppInstanceResponseBody
	GetSynchro() *bool
}

type GetAppInstanceResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/wlr10v1g25549kq/wlr10v1g25549kq.diff.zip?Expires=1730174953&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=F93JAFEuVN63YzNQyUy2xOaOtKs%3D
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"Success\\": True}
	Module *AppInstanceAggregate `json:"Module,omitempty" xml:"Module,omitempty"`
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

func (s GetAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppInstanceResponseBody) GetModule() *AppInstanceAggregate {
	return s.Module
}

func (s *GetAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppInstanceResponseBody) SetAccessDeniedDetail(v string) *GetAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetAllowRetry(v bool) *GetAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetAppName(v string) *GetAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetDynamicCode(v string) *GetAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetDynamicMessage(v string) *GetAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetErrorArgs(v []interface{}) *GetAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppInstanceResponseBody) SetModule(v *AppInstanceAggregate) *GetAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *GetAppInstanceResponseBody) SetRequestId(v string) *GetAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetRootErrorCode(v string) *GetAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetRootErrorMsg(v string) *GetAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppInstanceResponseBody) SetSynchro(v bool) *GetAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}
