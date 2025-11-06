// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSpecialBizInfoByDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetHttpStatusCode() *int32
	SetModule(v *QueryDomainSpecialBizInfoByDomainResponseBodyModule) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetModule() *QueryDomainSpecialBizInfoByDomainResponseBodyModule
	SetRequestId(v string) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryDomainSpecialBizInfoByDomainResponseBody
	GetSynchro() *bool
}

type QueryDomainSpecialBizInfoByDomainResponseBody struct {
	// Indicates whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application for which the error code is returned.
	//
	// example:
	//
	// test-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// -
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// -
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The array of error parameters that are returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 110001
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP status code that is directly returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned data.
	Module *QueryDomainSpecialBizInfoByDomainResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 97663DFF-D24D-55AE-A577-6CC5AF20B732
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether to perform synchronous processing.
	//
	// example:
	//
	// true
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryDomainSpecialBizInfoByDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizInfoByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetModule() *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	return s.Module
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetAllowRetry(v bool) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetAppName(v string) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetDynamicCode(v string) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetDynamicMessage(v string) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetErrorArgs(v []interface{}) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetErrorCode(v string) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetErrorMsg(v string) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetHttpStatusCode(v int32) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetModule(v *QueryDomainSpecialBizInfoByDomainResponseBodyModule) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.Module = v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetRequestId(v string) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetSuccess(v bool) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) SetSynchro(v bool) *QueryDomainSpecialBizInfoByDomainResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainSpecialBizInfoByDomainResponseBodyModule struct {
	// The review information.
	AuditMsg *string `json:"AuditMsg,omitempty" xml:"AuditMsg,omitempty"`
	// The business name.
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The business ID.
	//
	// example:
	//
	// A83E1D74-E46B-505C-947A-8C6B7E41C012
	BizNo *string `json:"BizNo,omitempty" xml:"BizNo,omitempty"`
	// The business status.
	//
	// example:
	//
	// REGISTRANT_VSP_AUDIT_SUCCESS
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The business type.
	//
	// example:
	//
	// govRegister
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The time when the business was created.
	//
	// example:
	//
	// 2023-01-17 11:31:10
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test003.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The contact information.
	DomainSpecialBizContact *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact `json:"DomainSpecialBizContact,omitempty" xml:"DomainSpecialBizContact,omitempty" type:"Struct"`
	// The certificate information.
	DomainSpecialBizCredentials []*QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials `json:"DomainSpecialBizCredentials,omitempty" xml:"DomainSpecialBizCredentials,omitempty" type:"Repeated"`
	// The time when the business was created.
	//
	// example:
	//
	// 2023-03-21 15:10:04.0
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the business was modified.
	//
	// example:
	//
	// 2023-03-21 15:10:04.0
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key.
	//
	// example:
	//
	// 34083
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 123
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// S20172315BJ37809
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The business status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of business status.
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The time when the business was updated.
	//
	// example:
	//
	// 2020-11-17 18:11:15
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 121000000****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDomainSpecialBizInfoByDomainResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizInfoByDomainResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetAuditMsg() *string {
	return s.AuditMsg
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetBizName() *string {
	return s.BizName
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetBizNo() *string {
	return s.BizNo
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetBizStatus() *string {
	return s.BizStatus
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetBizType() *string {
	return s.BizType
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetDomainSpecialBizContact() *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	return s.DomainSpecialBizContact
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetDomainSpecialBizCredentials() []*QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	return s.DomainSpecialBizCredentials
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetProductId() *string {
	return s.ProductId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetAuditMsg(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.AuditMsg = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetBizName(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.BizName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetBizNo(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.BizNo = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetBizStatus(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.BizStatus = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetBizType(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.BizType = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetCreateTime(v int64) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetDomainName(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.DomainName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetDomainSpecialBizContact(v *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.DomainSpecialBizContact = v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetDomainSpecialBizCredentials(v []*QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.DomainSpecialBizCredentials = v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetGmtCreate(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetGmtModified(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetId(v int64) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.Id = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetOrderId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetProductId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.ProductId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetSaleId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.SaleId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetStatus(v int32) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.Status = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetStatusDesc(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.StatusDesc = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetUpdateTime(v int64) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.UpdateTime = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) SetUserId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModule) Validate() error {
	if s.DomainSpecialBizContact != nil {
		if err := s.DomainSpecialBizContact.Validate(); err != nil {
			return err
		}
	}
	if s.DomainSpecialBizCredentials != nil {
		for _, item := range s.DomainSpecialBizCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact struct {
	// The business ID.
	//
	// example:
	//
	// 258
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The city.
	CCity *string `json:"CCity,omitempty" xml:"CCity,omitempty"`
	// The organization name.
	CCompany *string `json:"CCompany,omitempty" xml:"CCompany,omitempty"`
	// The country code.
	//
	// example:
	//
	// CN
	CCountry *string `json:"CCountry,omitempty" xml:"CCountry,omitempty"`
	// The contact name.
	CName *string `json:"CName,omitempty" xml:"CName,omitempty"`
	// The province.
	CProvince *string `json:"CProvince,omitempty" xml:"CProvince,omitempty"`
	// The address.
	CVenu *string `json:"CVenu,omitempty" xml:"CVenu,omitempty"`
	// The city in English.
	//
	// example:
	//
	// an shan Shi
	ECity *string `json:"ECity,omitempty" xml:"ECity,omitempty"`
	// The organization name in English.
	//
	// example:
	//
	// hebeihuiheguandaozhizaoyouxiangongsi
	ECompany *string `json:"ECompany,omitempty" xml:"ECompany,omitempty"`
	// The contact name in English.
	//
	// example:
	//
	// tong da wei
	EName *string `json:"EName,omitempty" xml:"EName,omitempty"`
	// The province in English.
	//
	// example:
	//
	// liao ning
	EProvince *string `json:"EProvince,omitempty" xml:"EProvince,omitempty"`
	// The address in English.
	//
	// example:
	//
	// tie xi qu xin kai jie 59-4 hao
	EVenu *string `json:"EVenu,omitempty" xml:"EVenu,omitempty"`
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The extended information.
	//
	// example:
	//
	// -
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The fax country code.
	//
	// example:
	//
	// 86
	FaxArea *string `json:"FaxArea,omitempty" xml:"FaxArea,omitempty"`
	// The fax extension number.
	//
	// example:
	//
	// 61284565
	FaxExt *string `json:"FaxExt,omitempty" xml:"FaxExt,omitempty"`
	// The fax number with an area code or mobile number.
	//
	// example:
	//
	// 16604121234
	FaxMain *string `json:"FaxMain,omitempty" xml:"FaxMain,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 156********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The zip code.
	//
	// example:
	//
	// 100000
	Postalcode *string `json:"Postalcode,omitempty" xml:"Postalcode,omitempty"`
	// The contact type. Valid values: 1: individual. 2: enterprise.
	//
	// example:
	//
	// 1
	RegType *int32 `json:"RegType,omitempty" xml:"RegType,omitempty"`
	// The registrant ID.
	//
	// example:
	//
	// 121000002****
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
	// The calling code of the country or region where the domain name contact is located.
	//
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// The telephone extension number.
	//
	// example:
	//
	// 2756
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// The landline number with an area code or mobile number.
	//
	// example:
	//
	// 16604121234
	TelMain *string `json:"TelMain,omitempty" xml:"TelMain,omitempty"`
	// The VSP contact ID.
	//
	// example:
	//
	// 121000001****
	VspContactId *string `json:"VspContactId,omitempty" xml:"VspContactId,omitempty"`
}

func (s QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetBizId() *int64 {
	return s.BizId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetCCity() *string {
	return s.CCity
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetCCompany() *string {
	return s.CCompany
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetCCountry() *string {
	return s.CCountry
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetCName() *string {
	return s.CName
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetCProvince() *string {
	return s.CProvince
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetCVenu() *string {
	return s.CVenu
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetECity() *string {
	return s.ECity
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetECompany() *string {
	return s.ECompany
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetEName() *string {
	return s.EName
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetEProvince() *string {
	return s.EProvince
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetEVenu() *string {
	return s.EVenu
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetEmail() *string {
	return s.Email
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetExtend() *string {
	return s.Extend
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetFaxArea() *string {
	return s.FaxArea
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetFaxExt() *string {
	return s.FaxExt
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetFaxMain() *string {
	return s.FaxMain
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetMobile() *string {
	return s.Mobile
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetPostalcode() *string {
	return s.Postalcode
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetRegType() *int32 {
	return s.RegType
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetTelMain() *string {
	return s.TelMain
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) GetVspContactId() *string {
	return s.VspContactId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetBizId(v int64) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.BizId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetCCity(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.CCity = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetCCompany(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.CCompany = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetCCountry(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.CCountry = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetCName(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.CName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetCProvince(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.CProvince = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetCVenu(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.CVenu = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetECity(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.ECity = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetECompany(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.ECompany = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetEName(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.EName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetEProvince(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.EProvince = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetEVenu(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.EVenu = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetEmail(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.Email = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetExtend(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.Extend = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetFaxArea(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.FaxArea = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetFaxExt(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.FaxExt = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetFaxMain(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.FaxMain = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetMobile(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.Mobile = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetPostalcode(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.Postalcode = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetRegType(v int32) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.RegType = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetRegistrantId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.RegistrantId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetTelArea(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.TelArea = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetTelExt(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.TelExt = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetTelMain(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.TelMain = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) SetVspContactId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact {
	s.VspContactId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizContact) Validate() error {
	return dara.Validate(s)
}

type QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials struct {
	// The ID of the associated workflow.
	//
	// example:
	//
	// T20220831150014000001
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The certificate number.
	//
	// example:
	//
	// 4111111111111110**
	CredentialNo *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	// The certificate type.
	//
	// example:
	//
	// SFZ
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The certificate URL.
	//
	// example:
	//
	// http://test.oss-cn-hangzhou.aliyuncs.com/20170522/1219541161213057_070445190.jpg
	CredentialUrl *string `json:"CredentialUrl,omitempty" xml:"CredentialUrl,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test003.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Indicates whether the certificate belongs to the registrant.
	//
	// example:
	//
	// 1
	HolderCert *int32 `json:"HolderCert,omitempty" xml:"HolderCert,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// S20172315BJ37809
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
}

func (s QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetBizId() *int64 {
	return s.BizId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetCredentialType() *string {
	return s.CredentialType
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetCredentialUrl() *string {
	return s.CredentialUrl
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetHolderCert() *int32 {
	return s.HolderCert
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetBizId(v int64) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.BizId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetCredentialNo(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.CredentialNo = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetCredentialType(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.CredentialType = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetCredentialUrl(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.CredentialUrl = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetDomainName(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.DomainName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetHolderCert(v int32) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.HolderCert = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) SetSaleId(v string) *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials {
	s.SaleId = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainResponseBodyModuleDomainSpecialBizCredentials) Validate() error {
	return dara.Validate(s)
}
