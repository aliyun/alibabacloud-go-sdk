// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAppDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAppDomainResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *BindAppDomainResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *BindAppDomainResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *BindAppDomainResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *BindAppDomainResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *BindAppDomainResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *BindAppDomainResponseBodyModule) *BindAppDomainResponseBody
	GetModule() *BindAppDomainResponseBodyModule
	SetRequestId(v string) *BindAppDomainResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *BindAppDomainResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *BindAppDomainResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *BindAppDomainResponseBody
	GetSynchro() *bool
}

type BindAppDomainResponseBody struct {
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
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the **%s*	- placeholder in the **ErrMessage*	- response element.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The data table module. Valid values:
	//
	// - ABTest: experiment data table
	//
	// - ExperimentTool: experiment tool table
	//
	// - DataDiagnosis: data modeling diagnostics
	Module *BindAppDomainResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// SYSTEM.EROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s BindAppDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAppDomainResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *BindAppDomainResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *BindAppDomainResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *BindAppDomainResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *BindAppDomainResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *BindAppDomainResponseBody) GetModule() *BindAppDomainResponseBodyModule {
	return s.Module
}

func (s *BindAppDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAppDomainResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *BindAppDomainResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *BindAppDomainResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *BindAppDomainResponseBody) SetAccessDeniedDetail(v string) *BindAppDomainResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAppDomainResponseBody) SetAllowRetry(v bool) *BindAppDomainResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *BindAppDomainResponseBody) SetAppName(v string) *BindAppDomainResponseBody {
	s.AppName = &v
	return s
}

func (s *BindAppDomainResponseBody) SetDynamicCode(v string) *BindAppDomainResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BindAppDomainResponseBody) SetDynamicMessage(v string) *BindAppDomainResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BindAppDomainResponseBody) SetErrorArgs(v []interface{}) *BindAppDomainResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *BindAppDomainResponseBody) SetModule(v *BindAppDomainResponseBodyModule) *BindAppDomainResponseBody {
	s.Module = v
	return s
}

func (s *BindAppDomainResponseBody) SetRequestId(v string) *BindAppDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAppDomainResponseBody) SetRootErrorCode(v string) *BindAppDomainResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *BindAppDomainResponseBody) SetRootErrorMsg(v string) *BindAppDomainResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *BindAppDomainResponseBody) SetSynchro(v bool) *BindAppDomainResponseBody {
	s.Synchro = &v
	return s
}

func (s *BindAppDomainResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAppDomainResponseBodyModule struct {
	DnsConflict *BindAppDomainResponseBodyModuleDnsConflict `json:"DnsConflict,omitempty" xml:"DnsConflict,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAppDomainResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponseBodyModule) GetDnsConflict() *BindAppDomainResponseBodyModuleDnsConflict {
	return s.DnsConflict
}

func (s *BindAppDomainResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *BindAppDomainResponseBodyModule) SetDnsConflict(v *BindAppDomainResponseBodyModuleDnsConflict) *BindAppDomainResponseBodyModule {
	s.DnsConflict = v
	return s
}

func (s *BindAppDomainResponseBodyModule) SetSuccess(v bool) *BindAppDomainResponseBodyModule {
	s.Success = &v
	return s
}

func (s *BindAppDomainResponseBodyModule) Validate() error {
	if s.DnsConflict != nil {
		if err := s.DnsConflict.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAppDomainResponseBodyModuleDnsConflict struct {
	CanAutoResolve *bool                                                `json:"CanAutoResolve,omitempty" xml:"CanAutoResolve,omitempty"`
	HasConflict    *bool                                                `json:"HasConflict,omitempty" xml:"HasConflict,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Records        []*BindAppDomainResponseBodyModuleDnsConflictRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s BindAppDomainResponseBodyModuleDnsConflict) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponseBodyModuleDnsConflict) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) GetCanAutoResolve() *bool {
	return s.CanAutoResolve
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) GetHasConflict() *bool {
	return s.HasConflict
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) GetMessage() *string {
	return s.Message
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) GetRecords() []*BindAppDomainResponseBodyModuleDnsConflictRecords {
	return s.Records
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) SetCanAutoResolve(v bool) *BindAppDomainResponseBodyModuleDnsConflict {
	s.CanAutoResolve = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) SetHasConflict(v bool) *BindAppDomainResponseBodyModuleDnsConflict {
	s.HasConflict = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) SetMessage(v string) *BindAppDomainResponseBodyModuleDnsConflict {
	s.Message = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) SetRecords(v []*BindAppDomainResponseBodyModuleDnsConflictRecords) *BindAppDomainResponseBodyModuleDnsConflict {
	s.Records = v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflict) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BindAppDomainResponseBodyModuleDnsConflictRecords struct {
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BindAppDomainResponseBodyModuleDnsConflictRecords) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponseBodyModuleDnsConflictRecords) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) GetHost() *string {
	return s.Host
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) GetRecordType() *string {
	return s.RecordType
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) GetStatus() *string {
	return s.Status
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) GetValue() *string {
	return s.Value
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) SetHost(v string) *BindAppDomainResponseBodyModuleDnsConflictRecords {
	s.Host = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) SetRecordType(v string) *BindAppDomainResponseBodyModuleDnsConflictRecords {
	s.RecordType = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) SetStatus(v string) *BindAppDomainResponseBodyModuleDnsConflictRecords {
	s.Status = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) SetValue(v string) *BindAppDomainResponseBodyModuleDnsConflictRecords {
	s.Value = &v
	return s
}

func (s *BindAppDomainResponseBodyModuleDnsConflictRecords) Validate() error {
	return dara.Validate(s)
}
