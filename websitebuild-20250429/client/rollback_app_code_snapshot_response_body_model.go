// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppCodeSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RollbackAppCodeSnapshotResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RollbackAppCodeSnapshotResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RollbackAppCodeSnapshotResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RollbackAppCodeSnapshotResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RollbackAppCodeSnapshotResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RollbackAppCodeSnapshotResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RollbackAppCodeSnapshotResponseBodyModule) *RollbackAppCodeSnapshotResponseBody
	GetModule() *RollbackAppCodeSnapshotResponseBodyModule
	SetRequestId(v string) *RollbackAppCodeSnapshotResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RollbackAppCodeSnapshotResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RollbackAppCodeSnapshotResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RollbackAppCodeSnapshotResponseBody
	GetSynchro() *bool
}

type RollbackAppCodeSnapshotResponseBody struct {
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// abc
	DynamicMessage *string                                    `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                              `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *RollbackAppCodeSnapshotResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s RollbackAppCodeSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppCodeSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackAppCodeSnapshotResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RollbackAppCodeSnapshotResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RollbackAppCodeSnapshotResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RollbackAppCodeSnapshotResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RollbackAppCodeSnapshotResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RollbackAppCodeSnapshotResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RollbackAppCodeSnapshotResponseBody) GetModule() *RollbackAppCodeSnapshotResponseBodyModule {
	return s.Module
}

func (s *RollbackAppCodeSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackAppCodeSnapshotResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RollbackAppCodeSnapshotResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RollbackAppCodeSnapshotResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RollbackAppCodeSnapshotResponseBody) SetAccessDeniedDetail(v string) *RollbackAppCodeSnapshotResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetAllowRetry(v bool) *RollbackAppCodeSnapshotResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetAppName(v string) *RollbackAppCodeSnapshotResponseBody {
	s.AppName = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetDynamicCode(v string) *RollbackAppCodeSnapshotResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetDynamicMessage(v string) *RollbackAppCodeSnapshotResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetErrorArgs(v []interface{}) *RollbackAppCodeSnapshotResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetModule(v *RollbackAppCodeSnapshotResponseBodyModule) *RollbackAppCodeSnapshotResponseBody {
	s.Module = v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetRequestId(v string) *RollbackAppCodeSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetRootErrorCode(v string) *RollbackAppCodeSnapshotResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetRootErrorMsg(v string) *RollbackAppCodeSnapshotResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) SetSynchro(v bool) *RollbackAppCodeSnapshotResponseBody {
	s.Synchro = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RollbackAppCodeSnapshotResponseBodyModule struct {
	// example:
	//
	// - Test ProduceCommand ops tool.
	ChangeLog *string `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty"`
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 123
	LogicalNumber *int32 `json:"LogicalNumber,omitempty" xml:"LogicalNumber,omitempty"`
}

func (s RollbackAppCodeSnapshotResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppCodeSnapshotResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) GetChangeLog() *string {
	return s.ChangeLog
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) GetLogicalNumber() *int32 {
	return s.LogicalNumber
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) SetChangeLog(v string) *RollbackAppCodeSnapshotResponseBodyModule {
	s.ChangeLog = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) SetGmtCreate(v string) *RollbackAppCodeSnapshotResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) SetLogicalNumber(v int32) *RollbackAppCodeSnapshotResponseBodyModule {
	s.LogicalNumber = &v
	return s
}

func (s *RollbackAppCodeSnapshotResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
