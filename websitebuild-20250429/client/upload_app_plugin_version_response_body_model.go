// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAppPluginVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UploadAppPluginVersionResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UploadAppPluginVersionResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UploadAppPluginVersionResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UploadAppPluginVersionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UploadAppPluginVersionResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UploadAppPluginVersionResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UploadAppPluginVersionResponseBodyModule) *UploadAppPluginVersionResponseBody
	GetModule() *UploadAppPluginVersionResponseBodyModule
	SetRequestId(v string) *UploadAppPluginVersionResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UploadAppPluginVersionResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UploadAppPluginVersionResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UploadAppPluginVersionResponseBody
	GetSynchro() *bool
}

type UploadAppPluginVersionResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic code. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/gl3d6l3817id8p1/gl3d6l3817id8p1.diff.zip?Expires=1750392068&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=Bcj3eohy8nmlSQ7AAGdq7JZoLjM%3D
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The application module.
	Module *UploadAppPluginVersionResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The error message.
	//
	// example:
	//
	// System exception.
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UploadAppPluginVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAppPluginVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAppPluginVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UploadAppPluginVersionResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UploadAppPluginVersionResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UploadAppPluginVersionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UploadAppPluginVersionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UploadAppPluginVersionResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UploadAppPluginVersionResponseBody) GetModule() *UploadAppPluginVersionResponseBodyModule {
	return s.Module
}

func (s *UploadAppPluginVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadAppPluginVersionResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UploadAppPluginVersionResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UploadAppPluginVersionResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UploadAppPluginVersionResponseBody) SetAccessDeniedDetail(v string) *UploadAppPluginVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetAllowRetry(v bool) *UploadAppPluginVersionResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetAppName(v string) *UploadAppPluginVersionResponseBody {
	s.AppName = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetDynamicCode(v string) *UploadAppPluginVersionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetDynamicMessage(v string) *UploadAppPluginVersionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetErrorArgs(v []interface{}) *UploadAppPluginVersionResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetModule(v *UploadAppPluginVersionResponseBodyModule) *UploadAppPluginVersionResponseBody {
	s.Module = v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetRequestId(v string) *UploadAppPluginVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetRootErrorCode(v string) *UploadAppPluginVersionResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetRootErrorMsg(v string) *UploadAppPluginVersionResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) SetSynchro(v bool) *UploadAppPluginVersionResponseBody {
	s.Synchro = &v
	return s
}

func (s *UploadAppPluginVersionResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadAppPluginVersionResponseBodyModule struct {
	// The commit ID.
	//
	// example:
	//
	// xxx
	CommitSha *string `json:"CommitSha,omitempty" xml:"CommitSha,omitempty"`
	// The bound API gateway plug-in ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The signature version, such as OSS4-HMAC-SHA256.
	//
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UploadAppPluginVersionResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UploadAppPluginVersionResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UploadAppPluginVersionResponseBodyModule) GetCommitSha() *string {
	return s.CommitSha
}

func (s *UploadAppPluginVersionResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *UploadAppPluginVersionResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *UploadAppPluginVersionResponseBodyModule) GetVersion() *string {
	return s.Version
}

func (s *UploadAppPluginVersionResponseBodyModule) SetCommitSha(v string) *UploadAppPluginVersionResponseBodyModule {
	s.CommitSha = &v
	return s
}

func (s *UploadAppPluginVersionResponseBodyModule) SetPluginId(v string) *UploadAppPluginVersionResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *UploadAppPluginVersionResponseBodyModule) SetStatus(v string) *UploadAppPluginVersionResponseBodyModule {
	s.Status = &v
	return s
}

func (s *UploadAppPluginVersionResponseBodyModule) SetVersion(v string) *UploadAppPluginVersionResponseBodyModule {
	s.Version = &v
	return s
}

func (s *UploadAppPluginVersionResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
