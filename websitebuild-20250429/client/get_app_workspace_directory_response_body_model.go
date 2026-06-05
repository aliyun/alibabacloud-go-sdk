// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppWorkspaceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppWorkspaceDirectoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppWorkspaceDirectoryResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppWorkspaceDirectoryResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppWorkspaceDirectoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppWorkspaceDirectoryResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppWorkspaceDirectoryResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppWorkspaceDirectoryResponseBodyModule) *GetAppWorkspaceDirectoryResponseBody
	GetModule() *GetAppWorkspaceDirectoryResponseBodyModule
	SetRequestId(v string) *GetAppWorkspaceDirectoryResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppWorkspaceDirectoryResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppWorkspaceDirectoryResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppWorkspaceDirectoryResponseBody
	GetSynchro() *bool
}

type GetAppWorkspaceDirectoryResponseBody struct {
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
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                               `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppWorkspaceDirectoryResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetAppWorkspaceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppWorkspaceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetModule() *GetAppWorkspaceDirectoryResponseBodyModule {
	return s.Module
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppWorkspaceDirectoryResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetAccessDeniedDetail(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetAllowRetry(v bool) *GetAppWorkspaceDirectoryResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetAppName(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetDynamicCode(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetDynamicMessage(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetErrorArgs(v []interface{}) *GetAppWorkspaceDirectoryResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetModule(v *GetAppWorkspaceDirectoryResponseBodyModule) *GetAppWorkspaceDirectoryResponseBody {
	s.Module = v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetRequestId(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetRootErrorCode(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetRootErrorMsg(v string) *GetAppWorkspaceDirectoryResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) SetSynchro(v bool) *GetAppWorkspaceDirectoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppWorkspaceDirectoryResponseBodyModule struct {
	// example:
	//
	// 2026
	CurrentTime   *string                                                    `json:"CurrentTime,omitempty" xml:"CurrentTime,omitempty"`
	DirectoryList []*GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList `json:"DirectoryList,omitempty" xml:"DirectoryList,omitempty" type:"Repeated"`
}

func (s GetAppWorkspaceDirectoryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppWorkspaceDirectoryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppWorkspaceDirectoryResponseBodyModule) GetCurrentTime() *string {
	return s.CurrentTime
}

func (s *GetAppWorkspaceDirectoryResponseBodyModule) GetDirectoryList() []*GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList {
	return s.DirectoryList
}

func (s *GetAppWorkspaceDirectoryResponseBodyModule) SetCurrentTime(v string) *GetAppWorkspaceDirectoryResponseBodyModule {
	s.CurrentTime = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBodyModule) SetDirectoryList(v []*GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) *GetAppWorkspaceDirectoryResponseBodyModule {
	s.DirectoryList = v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBodyModule) Validate() error {
	if s.DirectoryList != nil {
		for _, item := range s.DirectoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList struct {
	Children []interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// example:
	//
	// PolarDBInnoDBRedoLogWrites
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Evaluable=true
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) String() string {
	return dara.Prettify(s)
}

func (s GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) GoString() string {
	return s.String()
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) GetChildren() []interface{} {
	return s.Children
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) GetKey() *string {
	return s.Key
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) GetLabel() *string {
	return s.Label
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) SetChildren(v []interface{}) *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList {
	s.Children = v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) SetKey(v string) *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList {
	s.Key = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) SetLabel(v string) *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList {
	s.Label = &v
	return s
}

func (s *GetAppWorkspaceDirectoryResponseBodyModuleDirectoryList) Validate() error {
	return dara.Validate(s)
}
