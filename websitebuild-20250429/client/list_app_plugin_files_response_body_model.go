// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppPluginFilesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppPluginFilesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppPluginFilesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppPluginFilesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppPluginFilesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppPluginFilesResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListAppPluginFilesResponseBodyModule) *ListAppPluginFilesResponseBody
	GetModule() *ListAppPluginFilesResponseBodyModule
	SetRequestId(v string) *ListAppPluginFilesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppPluginFilesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppPluginFilesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppPluginFilesResponseBody
	GetSynchro() *bool
}

type ListAppPluginFilesResponseBody struct {
	// The details of the permission verification failure.
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
	// The dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *ListAppPluginFilesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppPluginFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPluginFilesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppPluginFilesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppPluginFilesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppPluginFilesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppPluginFilesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppPluginFilesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppPluginFilesResponseBody) GetModule() *ListAppPluginFilesResponseBodyModule {
	return s.Module
}

func (s *ListAppPluginFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPluginFilesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppPluginFilesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppPluginFilesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppPluginFilesResponseBody) SetAccessDeniedDetail(v string) *ListAppPluginFilesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetAllowRetry(v bool) *ListAppPluginFilesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetAppName(v string) *ListAppPluginFilesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetDynamicCode(v string) *ListAppPluginFilesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetDynamicMessage(v string) *ListAppPluginFilesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetErrorArgs(v []interface{}) *ListAppPluginFilesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetModule(v *ListAppPluginFilesResponseBodyModule) *ListAppPluginFilesResponseBody {
	s.Module = v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetRequestId(v string) *ListAppPluginFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetRootErrorCode(v string) *ListAppPluginFilesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetRootErrorMsg(v string) *ListAppPluginFilesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) SetSynchro(v bool) *ListAppPluginFilesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppPluginFilesResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppPluginFilesResponseBodyModule struct {
	// The task execution progress.
	Files []*ListAppPluginFilesResponseBodyModuleFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The application instance version.
	//
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppPluginFilesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginFilesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppPluginFilesResponseBodyModule) GetFiles() []*ListAppPluginFilesResponseBodyModuleFiles {
	return s.Files
}

func (s *ListAppPluginFilesResponseBodyModule) GetVersion() *string {
	return s.Version
}

func (s *ListAppPluginFilesResponseBodyModule) SetFiles(v []*ListAppPluginFilesResponseBodyModuleFiles) *ListAppPluginFilesResponseBodyModule {
	s.Files = v
	return s
}

func (s *ListAppPluginFilesResponseBodyModule) SetVersion(v string) *ListAppPluginFilesResponseBodyModule {
	s.Version = &v
	return s
}

func (s *ListAppPluginFilesResponseBodyModule) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppPluginFilesResponseBodyModuleFiles struct {
	// The file content, returned when POP displays SKILL.md.
	//
	// example:
	//
	// domain cnamen
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The file path.
	//
	// example:
	//
	// /
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// blob / tree
	//
	// example:
	//
	// baidu
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The number of returned results. Default value: 10.
	//
	// example:
	//
	// 5
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListAppPluginFilesResponseBodyModuleFiles) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginFilesResponseBodyModuleFiles) GoString() string {
	return s.String()
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) GetContent() *string {
	return s.Content
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) GetFilePath() *string {
	return s.FilePath
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) GetFileType() *string {
	return s.FileType
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) GetSize() *int64 {
	return s.Size
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) SetContent(v string) *ListAppPluginFilesResponseBodyModuleFiles {
	s.Content = &v
	return s
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) SetFilePath(v string) *ListAppPluginFilesResponseBodyModuleFiles {
	s.FilePath = &v
	return s
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) SetFileType(v string) *ListAppPluginFilesResponseBodyModuleFiles {
	s.FileType = &v
	return s
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) SetSize(v int64) *ListAppPluginFilesResponseBodyModuleFiles {
	s.Size = &v
	return s
}

func (s *ListAppPluginFilesResponseBodyModuleFiles) Validate() error {
	return dara.Validate(s)
}
