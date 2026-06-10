// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppCodeWorkspaceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppCodeWorkspaceDetailResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppCodeWorkspaceDetailResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppCodeWorkspaceDetailResponseBodyModule) *GetAppCodeWorkspaceDetailResponseBody
	GetModule() *GetAppCodeWorkspaceDetailResponseBodyModule
	SetRequestId(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppCodeWorkspaceDetailResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppCodeWorkspaceDetailResponseBody
	GetSynchro() *bool
}

type GetAppCodeWorkspaceDetailResponseBody struct {
	// Detailed reason for access denial.
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
	// App name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Faulty parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Application module
	Module *GetAppCodeWorkspaceDetailResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppCodeWorkspaceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppCodeWorkspaceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetModule() *GetAppCodeWorkspaceDetailResponseBodyModule {
	return s.Module
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppCodeWorkspaceDetailResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetAccessDeniedDetail(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetAllowRetry(v bool) *GetAppCodeWorkspaceDetailResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetAppName(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetDynamicCode(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetDynamicMessage(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetErrorArgs(v []interface{}) *GetAppCodeWorkspaceDetailResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetModule(v *GetAppCodeWorkspaceDetailResponseBodyModule) *GetAppCodeWorkspaceDetailResponseBody {
	s.Module = v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetRequestId(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetRootErrorCode(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetRootErrorMsg(v string) *GetAppCodeWorkspaceDetailResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) SetSynchro(v bool) *GetAppCodeWorkspaceDetailResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppCodeWorkspaceDetailResponseBodyModule struct {
	// 11111
	//
	// example:
	//
	// 111
	ActiveLogicalNumber *int32  `json:"ActiveLogicalNumber,omitempty" xml:"ActiveLogicalNumber,omitempty"`
	CommitHash          *string `json:"CommitHash,omitempty" xml:"CommitHash,omitempty"`
	// true
	//
	// example:
	//
	// 1
	IsDirty *bool `json:"IsDirty,omitempty" xml:"IsDirty,omitempty"`
	// 1111
	//
	// example:
	//
	// 111
	MaxLogicalNumber *int32 `json:"MaxLogicalNumber,omitempty" xml:"MaxLogicalNumber,omitempty"`
	// Site ID, which can be obtained by invoking the [ListSites](~~ListSites~~) API.
	//
	// example:
	//
	// 865181640657408
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// API Guide information.
	Snapshots []*GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s GetAppCodeWorkspaceDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppCodeWorkspaceDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) GetActiveLogicalNumber() *int32 {
	return s.ActiveLogicalNumber
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) GetCommitHash() *string {
	return s.CommitHash
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) GetIsDirty() *bool {
	return s.IsDirty
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) GetMaxLogicalNumber() *int32 {
	return s.MaxLogicalNumber
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) GetSnapshots() []*GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots {
	return s.Snapshots
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) SetActiveLogicalNumber(v int32) *GetAppCodeWorkspaceDetailResponseBodyModule {
	s.ActiveLogicalNumber = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) SetCommitHash(v string) *GetAppCodeWorkspaceDetailResponseBodyModule {
	s.CommitHash = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) SetIsDirty(v bool) *GetAppCodeWorkspaceDetailResponseBodyModule {
	s.IsDirty = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) SetMaxLogicalNumber(v int32) *GetAppCodeWorkspaceDetailResponseBodyModule {
	s.MaxLogicalNumber = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) SetSiteId(v string) *GetAppCodeWorkspaceDetailResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) SetSnapshots(v []*GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) *GetAppCodeWorkspaceDetailResponseBodyModule {
	s.Snapshots = v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModule) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots struct {
	// SDK change log
	//
	// example:
	//
	// - Add Params.
	ChangeLog *string `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty"`
	// Creation UTC time in ISO8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-06-10T09:40:36.562Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Logical value
	//
	// example:
	//
	// 111
	LogicalNumber *int32 `json:"LogicalNumber,omitempty" xml:"LogicalNumber,omitempty"`
}

func (s GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) String() string {
	return dara.Prettify(s)
}

func (s GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) GoString() string {
	return s.String()
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) GetChangeLog() *string {
	return s.ChangeLog
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) GetLogicalNumber() *int32 {
	return s.LogicalNumber
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) SetChangeLog(v string) *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots {
	s.ChangeLog = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) SetGmtCreateTime(v string) *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) SetLogicalNumber(v int32) *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots {
	s.LogicalNumber = &v
	return s
}

func (s *GetAppCodeWorkspaceDetailResponseBodyModuleSnapshots) Validate() error {
	return dara.Validate(s)
}
