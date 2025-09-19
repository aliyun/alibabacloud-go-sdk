// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppDomainDnsRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeAppDomainDnsRecordResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DescribeAppDomainDnsRecordResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DescribeAppDomainDnsRecordResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DescribeAppDomainDnsRecordResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeAppDomainDnsRecordResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DescribeAppDomainDnsRecordResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *DescribeAppDomainDnsRecordResponseBodyModule) *DescribeAppDomainDnsRecordResponseBody
	GetModule() *DescribeAppDomainDnsRecordResponseBodyModule
	SetRequestId(v string) *DescribeAppDomainDnsRecordResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DescribeAppDomainDnsRecordResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DescribeAppDomainDnsRecordResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DescribeAppDomainDnsRecordResponseBody
	GetSynchro() *bool
}

type DescribeAppDomainDnsRecordResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/vs1w2yd3p8264pz/vs1w2yd3p8264pz.diff.zip?Expires=1739592470&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=3DRuMtCeTinVYxo7XAOEIOEmZGE%3D
	DynamicMessage *string                                       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                 `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *DescribeAppDomainDnsRecordResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s DescribeAppDomainDnsRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppDomainDnsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetModule() *DescribeAppDomainDnsRecordResponseBodyModule {
	return s.Module
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DescribeAppDomainDnsRecordResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetAccessDeniedDetail(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetAllowRetry(v bool) *DescribeAppDomainDnsRecordResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetAppName(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetDynamicCode(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetDynamicMessage(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetErrorArgs(v []interface{}) *DescribeAppDomainDnsRecordResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetModule(v *DescribeAppDomainDnsRecordResponseBodyModule) *DescribeAppDomainDnsRecordResponseBody {
	s.Module = v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetRequestId(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetRootErrorCode(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetRootErrorMsg(v string) *DescribeAppDomainDnsRecordResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) SetSynchro(v bool) *DescribeAppDomainDnsRecordResponseBody {
	s.Synchro = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppDomainDnsRecordResponseBodyModule struct {
	// example:
	//
	// *.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// A
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// Maintenance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppDomainDnsRecordResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppDomainDnsRecordResponseBodyModule) GoString() string {
	return s.String()
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) GetHost() *string {
	return s.Host
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) GetRecordType() *string {
	return s.RecordType
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) GetValue() *string {
	return s.Value
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) SetHost(v string) *DescribeAppDomainDnsRecordResponseBodyModule {
	s.Host = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) SetRecordType(v string) *DescribeAppDomainDnsRecordResponseBodyModule {
	s.RecordType = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) SetValue(v string) *DescribeAppDomainDnsRecordResponseBodyModule {
	s.Value = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
