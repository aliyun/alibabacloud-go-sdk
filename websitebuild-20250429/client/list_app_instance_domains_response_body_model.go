// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppInstanceDomainsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppInstanceDomainsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppInstanceDomainsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppInstanceDomainsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppInstanceDomainsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppInstanceDomainsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppInstanceDomainsResponseBody
	GetMaxResults() *int32
	SetModule(v *ListAppInstanceDomainsResponseBodyModule) *ListAppInstanceDomainsResponseBody
	GetModule() *ListAppInstanceDomainsResponseBodyModule
	SetNextToken(v string) *ListAppInstanceDomainsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppInstanceDomainsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppInstanceDomainsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppInstanceDomainsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppInstanceDomainsResponseBody
	GetSynchro() *bool
}

type ListAppInstanceDomainsResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace `%s` in the error message of the returned parameter **ErrMessage**.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it means that the input request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Range: 10~100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Response data
	Module *ListAppInstanceDomainsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// 下一个查询开始的Token。没有下一个查询时为空。
	//
	// example:
	//
	// dw+qdTi1EjVSWX/INJdYNw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ID of the request
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
	// Exception message
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

func (s ListAppInstanceDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppInstanceDomainsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppInstanceDomainsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppInstanceDomainsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppInstanceDomainsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppInstanceDomainsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppInstanceDomainsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppInstanceDomainsResponseBody) GetModule() *ListAppInstanceDomainsResponseBodyModule {
	return s.Module
}

func (s *ListAppInstanceDomainsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppInstanceDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppInstanceDomainsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppInstanceDomainsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppInstanceDomainsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppInstanceDomainsResponseBody) SetAccessDeniedDetail(v string) *ListAppInstanceDomainsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetAllowRetry(v bool) *ListAppInstanceDomainsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetAppName(v string) *ListAppInstanceDomainsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetDynamicCode(v string) *ListAppInstanceDomainsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetDynamicMessage(v string) *ListAppInstanceDomainsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetErrorArgs(v []interface{}) *ListAppInstanceDomainsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetMaxResults(v int32) *ListAppInstanceDomainsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetModule(v *ListAppInstanceDomainsResponseBodyModule) *ListAppInstanceDomainsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetNextToken(v string) *ListAppInstanceDomainsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetRequestId(v string) *ListAppInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetRootErrorCode(v string) *ListAppInstanceDomainsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetRootErrorMsg(v string) *ListAppInstanceDomainsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) SetSynchro(v bool) *ListAppInstanceDomainsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModule struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// Query result.
	Data []*ListAppInstanceDomainsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Decision weight
	Next *ListAppInstanceDomainsResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// Whether there is a next page.
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Whether there is a previous page
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// Apart from pagination limits, the server can process up to the most recent 1000 records. If the result exceeds 1000 records, **ResultLimit*	- will be **true**, please narrow down the time range and search again; otherwise, **ResultLimit*	- will be **false**.
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 0
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetData() []*ListAppInstanceDomainsResponseBodyModuleData {
	return s.Data
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetNext() *ListAppInstanceDomainsResponseBodyModuleNext {
	return s.Next
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListAppInstanceDomainsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetCurrentPageNum(v int32) *ListAppInstanceDomainsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetData(v []*ListAppInstanceDomainsResponseBodyModuleData) *ListAppInstanceDomainsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetNext(v *ListAppInstanceDomainsResponseBodyModuleNext) *ListAppInstanceDomainsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetNextPage(v bool) *ListAppInstanceDomainsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetPageSize(v int32) *ListAppInstanceDomainsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetPrePage(v bool) *ListAppInstanceDomainsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetResultLimit(v bool) *ListAppInstanceDomainsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetTotalItemNum(v int32) *ListAppInstanceDomainsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) SetTotalPageNum(v int32) *ListAppInstanceDomainsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleData struct {
	// Domain SSL certificate information
	Certificate *ListAppInstanceDomainsResponseBodyModuleDataCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// Instance creation time (required, format: yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 1683256054000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain name
	//
	// example:
	//
	// kaibaidu.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Binding overall status
	//
	// example:
	//
	// ACTIVE
	OverallStatus *string `json:"OverallStatus,omitempty" xml:"OverallStatus,omitempty"`
	// Domain ownership information
	Ownership *ListAppInstanceDomainsResponseBodyModuleDataOwnership `json:"Ownership,omitempty" xml:"Ownership,omitempty" type:"Struct"`
	// Domain resolution information
	Resolution *ListAppInstanceDomainsResponseBodyModuleDataResolution `json:"Resolution,omitempty" xml:"Resolution,omitempty" type:"Struct"`
	// Domain verification information
	Verification *ListAppInstanceDomainsResponseBodyModuleDataVerification `json:"Verification,omitempty" xml:"Verification,omitempty" type:"Struct"`
}

func (s ListAppInstanceDomainsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetCertificate() *ListAppInstanceDomainsResponseBodyModuleDataCertificate {
	return s.Certificate
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetDomainName() *string {
	return s.DomainName
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetOverallStatus() *string {
	return s.OverallStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetOwnership() *ListAppInstanceDomainsResponseBodyModuleDataOwnership {
	return s.Ownership
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetResolution() *ListAppInstanceDomainsResponseBodyModuleDataResolution {
	return s.Resolution
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) GetVerification() *ListAppInstanceDomainsResponseBodyModuleDataVerification {
	return s.Verification
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetCertificate(v *ListAppInstanceDomainsResponseBodyModuleDataCertificate) *ListAppInstanceDomainsResponseBodyModuleData {
	s.Certificate = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetCreateTime(v string) *ListAppInstanceDomainsResponseBodyModuleData {
	s.CreateTime = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetDomainName(v string) *ListAppInstanceDomainsResponseBodyModuleData {
	s.DomainName = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetOverallStatus(v string) *ListAppInstanceDomainsResponseBodyModuleData {
	s.OverallStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetOwnership(v *ListAppInstanceDomainsResponseBodyModuleDataOwnership) *ListAppInstanceDomainsResponseBodyModuleData {
	s.Ownership = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetResolution(v *ListAppInstanceDomainsResponseBodyModuleDataResolution) *ListAppInstanceDomainsResponseBodyModuleData {
	s.Resolution = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) SetVerification(v *ListAppInstanceDomainsResponseBodyModuleDataVerification) *ListAppInstanceDomainsResponseBodyModuleData {
	s.Verification = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleData) Validate() error {
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	if s.Ownership != nil {
		if err := s.Ownership.Validate(); err != nil {
			return err
		}
	}
	if s.Resolution != nil {
		if err := s.Resolution.Validate(); err != nil {
			return err
		}
	}
	if s.Verification != nil {
		if err := s.Verification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleDataCertificate struct {
	// Certificate name.
	//
	// example:
	//
	// 2024
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// Certificate status
	//
	// example:
	//
	// ACTIVE
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// Certificate type
	//
	// example:
	//
	// self-signed
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// Certificate expiration date
	//
	// example:
	//
	// 4885718400000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleDataCertificate) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleDataCertificate) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) GetCertificateName() *string {
	return s.CertificateName
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) SetCertificateName(v string) *ListAppInstanceDomainsResponseBodyModuleDataCertificate {
	s.CertificateName = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) SetCertificateStatus(v string) *ListAppInstanceDomainsResponseBodyModuleDataCertificate {
	s.CertificateStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) SetCertificateType(v string) *ListAppInstanceDomainsResponseBodyModuleDataCertificate {
	s.CertificateType = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) SetEndTime(v string) *ListAppInstanceDomainsResponseBodyModuleDataCertificate {
	s.EndTime = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataCertificate) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleDataOwnership struct {
	// Account owner type: CURRENT OTHER
	//
	// example:
	//
	// 1813244684017878
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// Registrar type: ALIYUN OTHER
	//
	// example:
	//
	// ROS
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The root domain name of the domain.
	//
	// example:
	//
	// tjouya.cn
	RootDomain *string `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleDataOwnership) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleDataOwnership) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) GetAccount() *string {
	return s.Account
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) GetProvider() *string {
	return s.Provider
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) GetRootDomain() *string {
	return s.RootDomain
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) SetAccount(v string) *ListAppInstanceDomainsResponseBodyModuleDataOwnership {
	s.Account = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) SetProvider(v string) *ListAppInstanceDomainsResponseBodyModuleDataOwnership {
	s.Provider = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) SetRootDomain(v string) *ListAppInstanceDomainsResponseBodyModuleDataOwnership {
	s.RootDomain = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataOwnership) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleDataResolution struct {
	// DNS record information
	DnsRecord *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// code: 400, invalid unionId request id: 09CC0F8B-49C2-7EFB-81E8-9AEF49068D02
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Domain resolution status
	//
	// example:
	//
	// SUCCESSFUL
	ResolutionStatus *string `json:"ResolutionStatus,omitempty" xml:"ResolutionStatus,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleDataResolution) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleDataResolution) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) GetDnsRecord() *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord {
	return s.DnsRecord
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) GetResolutionStatus() *string {
	return s.ResolutionStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) SetDnsRecord(v *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) *ListAppInstanceDomainsResponseBodyModuleDataResolution {
	s.DnsRecord = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) SetErrorMsg(v string) *ListAppInstanceDomainsResponseBodyModuleDataResolution {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) SetResolutionStatus(v string) *ListAppInstanceDomainsResponseBodyModuleDataResolution {
	s.ResolutionStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolution) Validate() error {
	if s.DnsRecord != nil {
		if err := s.DnsRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord struct {
	// Host record
	//
	// example:
	//
	// portal-dev.bambulab.net
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Record type
	//
	// example:
	//
	// 1
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// Record value
	//
	// example:
	//
	// Maintenance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) GetHost() *string {
	return s.Host
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) GetRecordType() *string {
	return s.RecordType
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) SetHost(v string) *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord {
	s.Host = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) SetRecordType(v string) *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord {
	s.RecordType = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) SetValue(v string) *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord {
	s.Value = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataResolutionDnsRecord) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleDataVerification struct {
	// DNS record guide that the user needs to configure.
	DnsRecord *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty" type:"Struct"`
	// Error message.
	//
	// example:
	//
	// code: 400, invalid unionId request id: 2270AB0B-6FD0-7F72-9DC5-7A02B67CBB3B
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Ownership verification status.
	//
	// example:
	//
	// SUCCESSFUL
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	// Ownership verification status
	//
	// example:
	//
	// NoAliyunServiceRoleForWebsiteBuildPublishAuth
	VerificationStatusCode *string `json:"VerificationStatusCode,omitempty" xml:"VerificationStatusCode,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleDataVerification) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleDataVerification) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) GetDnsRecord() *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord {
	return s.DnsRecord
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) GetVerificationStatusCode() *string {
	return s.VerificationStatusCode
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) SetDnsRecord(v *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) *ListAppInstanceDomainsResponseBodyModuleDataVerification {
	s.DnsRecord = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) SetErrorMsg(v string) *ListAppInstanceDomainsResponseBodyModuleDataVerification {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) SetVerificationStatus(v string) *ListAppInstanceDomainsResponseBodyModuleDataVerification {
	s.VerificationStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) SetVerificationStatusCode(v string) *ListAppInstanceDomainsResponseBodyModuleDataVerification {
	s.VerificationStatusCode = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerification) Validate() error {
	if s.DnsRecord != nil {
		if err := s.DnsRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord struct {
	// Host record
	//
	// example:
	//
	// ${host}
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Record type
	//
	// example:
	//
	// 1
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// Record value
	//
	// example:
	//
	// 159.138.94.138
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) GetHost() *string {
	return s.Host
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) GetRecordType() *string {
	return s.RecordType
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) SetHost(v string) *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord {
	s.Host = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) SetRecordType(v string) *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord {
	s.RecordType = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) SetValue(v string) *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord {
	s.Value = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleDataVerificationDnsRecord) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleNext struct {
	// Domain SSL certificate information
	Certificate *ListAppInstanceDomainsResponseBodyModuleNextCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// Instance creation time (required, format: yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 1741572465000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain name
	//
	// example:
	//
	// usdcoin.xin
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Binding overall status
	//
	// example:
	//
	// ACTIVE
	OverallStatus *string `json:"OverallStatus,omitempty" xml:"OverallStatus,omitempty"`
	// Domain ownership information
	Ownership *ListAppInstanceDomainsResponseBodyModuleNextOwnership `json:"Ownership,omitempty" xml:"Ownership,omitempty" type:"Struct"`
	// Domain resolution information
	Resolution *ListAppInstanceDomainsResponseBodyModuleNextResolution `json:"Resolution,omitempty" xml:"Resolution,omitempty" type:"Struct"`
	// Domain verification information
	Verification *ListAppInstanceDomainsResponseBodyModuleNextVerification `json:"Verification,omitempty" xml:"Verification,omitempty" type:"Struct"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetCertificate() *ListAppInstanceDomainsResponseBodyModuleNextCertificate {
	return s.Certificate
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetDomainName() *string {
	return s.DomainName
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetOverallStatus() *string {
	return s.OverallStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetOwnership() *ListAppInstanceDomainsResponseBodyModuleNextOwnership {
	return s.Ownership
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetResolution() *ListAppInstanceDomainsResponseBodyModuleNextResolution {
	return s.Resolution
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) GetVerification() *ListAppInstanceDomainsResponseBodyModuleNextVerification {
	return s.Verification
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetCertificate(v *ListAppInstanceDomainsResponseBodyModuleNextCertificate) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.Certificate = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetCreateTime(v string) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.CreateTime = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetDomainName(v string) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.DomainName = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetOverallStatus(v string) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.OverallStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetOwnership(v *ListAppInstanceDomainsResponseBodyModuleNextOwnership) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.Ownership = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetResolution(v *ListAppInstanceDomainsResponseBodyModuleNextResolution) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.Resolution = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) SetVerification(v *ListAppInstanceDomainsResponseBodyModuleNextVerification) *ListAppInstanceDomainsResponseBodyModuleNext {
	s.Verification = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNext) Validate() error {
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	if s.Ownership != nil {
		if err := s.Ownership.Validate(); err != nil {
			return err
		}
	}
	if s.Resolution != nil {
		if err := s.Resolution.Validate(); err != nil {
			return err
		}
	}
	if s.Verification != nil {
		if err := s.Verification.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleNextCertificate struct {
	// Certificate name.
	//
	// example:
	//
	// jfztkg202502
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// Certificate status
	//
	// example:
	//
	// ACTIVE
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// Certificate type
	//
	// example:
	//
	// Server
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// 证书的到期日期
	//
	// example:
	//
	// 2025-01-15T02:04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNextCertificate) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNextCertificate) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) GetCertificateName() *string {
	return s.CertificateName
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) SetCertificateName(v string) *ListAppInstanceDomainsResponseBodyModuleNextCertificate {
	s.CertificateName = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) SetCertificateStatus(v string) *ListAppInstanceDomainsResponseBodyModuleNextCertificate {
	s.CertificateStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) SetCertificateType(v string) *ListAppInstanceDomainsResponseBodyModuleNextCertificate {
	s.CertificateType = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) SetEndTime(v string) *ListAppInstanceDomainsResponseBodyModuleNextCertificate {
	s.EndTime = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextCertificate) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleNextOwnership struct {
	// Account owner type: CURRENT OTHER
	//
	// example:
	//
	// 拾肆gavin
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// Registrar type: ALIYUN OTHER
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNextOwnership) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNextOwnership) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextOwnership) GetAccount() *string {
	return s.Account
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextOwnership) GetProvider() *string {
	return s.Provider
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextOwnership) SetAccount(v string) *ListAppInstanceDomainsResponseBodyModuleNextOwnership {
	s.Account = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextOwnership) SetProvider(v string) *ListAppInstanceDomainsResponseBodyModuleNextOwnership {
	s.Provider = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextOwnership) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleNextResolution struct {
	// DNS record information
	DnsRecord *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// aliuid:1133664521498319 assumeOssRole not exist,serviceName:aliyunesarealtimelogpushossrole
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Domain resolution status
	//
	// example:
	//
	// SUCCESSFUL
	ResolutionStatus *string `json:"ResolutionStatus,omitempty" xml:"ResolutionStatus,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNextResolution) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNextResolution) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) GetDnsRecord() *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord {
	return s.DnsRecord
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) GetResolutionStatus() *string {
	return s.ResolutionStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) SetDnsRecord(v *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) *ListAppInstanceDomainsResponseBodyModuleNextResolution {
	s.DnsRecord = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) SetErrorMsg(v string) *ListAppInstanceDomainsResponseBodyModuleNextResolution {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) SetResolutionStatus(v string) *ListAppInstanceDomainsResponseBodyModuleNextResolution {
	s.ResolutionStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolution) Validate() error {
	if s.DnsRecord != nil {
		if err := s.DnsRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord struct {
	// Host record
	//
	// example:
	//
	// 172.16.6.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Record type
	//
	// example:
	//
	// A
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// Record value
	//
	// example:
	//
	// 2025032000000054nuba0r0b0a79y70c1c983tsz09xw9hg3p04kqxmbt4g2p65h
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) GetHost() *string {
	return s.Host
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) GetRecordType() *string {
	return s.RecordType
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) SetHost(v string) *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord {
	s.Host = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) SetRecordType(v string) *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord {
	s.RecordType = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) SetValue(v string) *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord {
	s.Value = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextResolutionDnsRecord) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceDomainsResponseBodyModuleNextVerification struct {
	// Guide for DNS records that the user needs to configure.
	DnsRecord *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty" type:"Struct"`
	// Error message.
	//
	// example:
	//
	// aliuid:1133664521498319 assumeOssRole not exist,serviceName:aliyunesarealtimelogpushossrole
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Ownership verification status.
	//
	// example:
	//
	// SUCCESSFUL
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNextVerification) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNextVerification) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) GetDnsRecord() *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord {
	return s.DnsRecord
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) SetDnsRecord(v *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) *ListAppInstanceDomainsResponseBodyModuleNextVerification {
	s.DnsRecord = v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) SetErrorMsg(v string) *ListAppInstanceDomainsResponseBodyModuleNextVerification {
	s.ErrorMsg = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) SetVerificationStatus(v string) *ListAppInstanceDomainsResponseBodyModuleNextVerification {
	s.VerificationStatus = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerification) Validate() error {
	if s.DnsRecord != nil {
		if err := s.DnsRecord.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord struct {
	// Host record
	//
	// example:
	//
	// www.abc.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Record type
	//
	// example:
	//
	// 5
	RecordType *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// Record value
	//
	// example:
	//
	// faHuBlyPcodSjBvBJYpm3-9W_cCSowLLQ4zAUyguEGM
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) GetHost() *string {
	return s.Host
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) GetRecordType() *string {
	return s.RecordType
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) SetHost(v string) *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord {
	s.Host = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) SetRecordType(v string) *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord {
	s.RecordType = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) SetValue(v string) *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord {
	s.Value = &v
	return s
}

func (s *ListAppInstanceDomainsResponseBodyModuleNextVerificationDnsRecord) Validate() error {
	return dara.Validate(s)
}
