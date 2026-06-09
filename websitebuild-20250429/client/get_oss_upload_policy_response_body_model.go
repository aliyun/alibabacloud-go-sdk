// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetOssUploadPolicyResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetOssUploadPolicyResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetOssUploadPolicyResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetOssUploadPolicyResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetOssUploadPolicyResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetOssUploadPolicyResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetOssUploadPolicyResponseBodyModule) *GetOssUploadPolicyResponseBody
	GetModule() *GetOssUploadPolicyResponseBodyModule
	SetRequestId(v string) *GetOssUploadPolicyResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetOssUploadPolicyResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetOssUploadPolicyResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetOssUploadPolicyResponseBody
	GetSynchro() *bool
}

type GetOssUploadPolicyResponseBody struct {
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
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                         `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetOssUploadPolicyResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetOssUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadPolicyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetOssUploadPolicyResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetOssUploadPolicyResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetOssUploadPolicyResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetOssUploadPolicyResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetOssUploadPolicyResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetOssUploadPolicyResponseBody) GetModule() *GetOssUploadPolicyResponseBodyModule {
	return s.Module
}

func (s *GetOssUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssUploadPolicyResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetOssUploadPolicyResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetOssUploadPolicyResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetOssUploadPolicyResponseBody) SetAccessDeniedDetail(v string) *GetOssUploadPolicyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetAllowRetry(v bool) *GetOssUploadPolicyResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetAppName(v string) *GetOssUploadPolicyResponseBody {
	s.AppName = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetDynamicCode(v string) *GetOssUploadPolicyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetDynamicMessage(v string) *GetOssUploadPolicyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetErrorArgs(v []interface{}) *GetOssUploadPolicyResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetModule(v *GetOssUploadPolicyResponseBodyModule) *GetOssUploadPolicyResponseBody {
	s.Module = v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetRequestId(v string) *GetOssUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetRootErrorCode(v string) *GetOssUploadPolicyResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetRootErrorMsg(v string) *GetOssUploadPolicyResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) SetSynchro(v bool) *GetOssUploadPolicyResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetOssUploadPolicyResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssUploadPolicyResponseBodyModule struct {
	// example:
	//
	// pano_src/100070-2679478/images/
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// *.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// ***
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// ****************************
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// x-oss-credential
	//
	// example:
	//
	// 123123
	XossCredential *string `json:"XossCredential,omitempty" xml:"XossCredential,omitempty"`
	// x-oss-date
	//
	// example:
	//
	// 20260101
	XossDate *string `json:"XossDate,omitempty" xml:"XossDate,omitempty"`
}

func (s GetOssUploadPolicyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadPolicyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetOssUploadPolicyResponseBodyModule) GetDir() *string {
	return s.Dir
}

func (s *GetOssUploadPolicyResponseBodyModule) GetHost() *string {
	return s.Host
}

func (s *GetOssUploadPolicyResponseBodyModule) GetPolicy() *string {
	return s.Policy
}

func (s *GetOssUploadPolicyResponseBodyModule) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetOssUploadPolicyResponseBodyModule) GetSignature() *string {
	return s.Signature
}

func (s *GetOssUploadPolicyResponseBodyModule) GetVersion() *string {
	return s.Version
}

func (s *GetOssUploadPolicyResponseBodyModule) GetXossCredential() *string {
	return s.XossCredential
}

func (s *GetOssUploadPolicyResponseBodyModule) GetXossDate() *string {
	return s.XossDate
}

func (s *GetOssUploadPolicyResponseBodyModule) SetDir(v string) *GetOssUploadPolicyResponseBodyModule {
	s.Dir = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetHost(v string) *GetOssUploadPolicyResponseBodyModule {
	s.Host = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetPolicy(v string) *GetOssUploadPolicyResponseBodyModule {
	s.Policy = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetSecurityToken(v string) *GetOssUploadPolicyResponseBodyModule {
	s.SecurityToken = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetSignature(v string) *GetOssUploadPolicyResponseBodyModule {
	s.Signature = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetVersion(v string) *GetOssUploadPolicyResponseBodyModule {
	s.Version = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetXossCredential(v string) *GetOssUploadPolicyResponseBodyModule {
	s.XossCredential = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) SetXossDate(v string) *GetOssUploadPolicyResponseBodyModule {
	s.XossDate = &v
	return s
}

func (s *GetOssUploadPolicyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
