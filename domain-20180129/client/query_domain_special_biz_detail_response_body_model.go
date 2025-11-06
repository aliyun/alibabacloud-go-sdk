// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSpecialBizDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *QueryDomainSpecialBizDetailResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryDomainSpecialBizDetailResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryDomainSpecialBizDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryDomainSpecialBizDetailResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryDomainSpecialBizDetailResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryDomainSpecialBizDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryDomainSpecialBizDetailResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *QueryDomainSpecialBizDetailResponseBody
	GetHttpStatusCode() *int32
	SetModule(v *QueryDomainSpecialBizDetailResponseBodyModule) *QueryDomainSpecialBizDetailResponseBody
	GetModule() *QueryDomainSpecialBizDetailResponseBodyModule
	SetRequestId(v string) *QueryDomainSpecialBizDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDomainSpecialBizDetailResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryDomainSpecialBizDetailResponseBody
	GetSynchro() *bool
}

type QueryDomainSpecialBizDetailResponseBody struct {
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
	Module *QueryDomainSpecialBizDetailResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A83E1D74-E46B-505C-947A-8C6B7E41C011
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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

func (s QueryDomainSpecialBizDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetModule() *QueryDomainSpecialBizDetailResponseBodyModule {
	return s.Module
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDomainSpecialBizDetailResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetAllowRetry(v bool) *QueryDomainSpecialBizDetailResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetAppName(v string) *QueryDomainSpecialBizDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetDynamicCode(v string) *QueryDomainSpecialBizDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetDynamicMessage(v string) *QueryDomainSpecialBizDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetErrorArgs(v []interface{}) *QueryDomainSpecialBizDetailResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetErrorCode(v string) *QueryDomainSpecialBizDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetErrorMsg(v string) *QueryDomainSpecialBizDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetHttpStatusCode(v int32) *QueryDomainSpecialBizDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetModule(v *QueryDomainSpecialBizDetailResponseBodyModule) *QueryDomainSpecialBizDetailResponseBody {
	s.Module = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetRequestId(v string) *QueryDomainSpecialBizDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetSuccess(v bool) *QueryDomainSpecialBizDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) SetSynchro(v bool) *QueryDomainSpecialBizDetailResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainSpecialBizDetailResponseBodyModule struct {
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
	DomainSpecialBizContact *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact `json:"DomainSpecialBizContact,omitempty" xml:"DomainSpecialBizContact,omitempty" type:"Struct"`
	// The certificate information.
	DomainSpecialBizCredentials []*QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials `json:"DomainSpecialBizCredentials,omitempty" xml:"DomainSpecialBizCredentials,omitempty" type:"Repeated"`
	// The information about the order.
	DomainSpecialOrderResult *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult `json:"DomainSpecialOrderResult,omitempty" xml:"DomainSpecialOrderResult,omitempty" type:"Struct"`
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

func (s QueryDomainSpecialBizDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetAuditMsg() *string {
	return s.AuditMsg
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetBizName() *string {
	return s.BizName
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetBizNo() *string {
	return s.BizNo
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetBizStatus() *string {
	return s.BizStatus
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetBizType() *string {
	return s.BizType
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetDomainSpecialBizContact() *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	return s.DomainSpecialBizContact
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetDomainSpecialBizCredentials() []*QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	return s.DomainSpecialBizCredentials
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetDomainSpecialOrderResult() *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	return s.DomainSpecialOrderResult
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetProductId() *string {
	return s.ProductId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetAuditMsg(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.AuditMsg = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetBizName(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.BizName = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetBizNo(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.BizNo = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetBizStatus(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.BizStatus = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetBizType(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.BizType = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetCreateTime(v int64) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetDomainName(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.DomainName = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetDomainSpecialBizContact(v *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.DomainSpecialBizContact = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetDomainSpecialBizCredentials(v []*QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.DomainSpecialBizCredentials = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetDomainSpecialOrderResult(v *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.DomainSpecialOrderResult = v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetGmtCreate(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetGmtModified(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetId(v int64) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.Id = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetOrderId(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetProductId(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.ProductId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetSaleId(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.SaleId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetStatus(v int32) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.Status = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetStatusDesc(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.StatusDesc = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetUpdateTime(v int64) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.UpdateTime = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) SetUserId(v string) *QueryDomainSpecialBizDetailResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModule) Validate() error {
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
	if s.DomainSpecialOrderResult != nil {
		if err := s.DomainSpecialOrderResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact struct {
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
	// 61284565
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

func (s QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetBizId() *int64 {
	return s.BizId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetCCity() *string {
	return s.CCity
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetCCompany() *string {
	return s.CCompany
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetCCountry() *string {
	return s.CCountry
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetCName() *string {
	return s.CName
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetCProvince() *string {
	return s.CProvince
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetCVenu() *string {
	return s.CVenu
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetECity() *string {
	return s.ECity
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetECompany() *string {
	return s.ECompany
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetEName() *string {
	return s.EName
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetEProvince() *string {
	return s.EProvince
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetEVenu() *string {
	return s.EVenu
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetEmail() *string {
	return s.Email
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetExtend() *string {
	return s.Extend
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetFaxArea() *string {
	return s.FaxArea
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetFaxExt() *string {
	return s.FaxExt
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetFaxMain() *string {
	return s.FaxMain
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetMobile() *string {
	return s.Mobile
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetPostalcode() *string {
	return s.Postalcode
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetRegType() *int32 {
	return s.RegType
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetTelMain() *string {
	return s.TelMain
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) GetVspContactId() *string {
	return s.VspContactId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetBizId(v int64) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.BizId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetCCity(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.CCity = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetCCompany(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.CCompany = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetCCountry(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.CCountry = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetCName(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.CName = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetCProvince(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.CProvince = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetCVenu(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.CVenu = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetECity(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.ECity = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetECompany(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.ECompany = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetEName(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.EName = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetEProvince(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.EProvince = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetEVenu(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.EVenu = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetEmail(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.Email = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetExtend(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.Extend = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetFaxArea(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.FaxArea = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetFaxExt(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.FaxExt = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetFaxMain(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.FaxMain = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetMobile(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.Mobile = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetPostalcode(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.Postalcode = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetRegType(v int32) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.RegType = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetRegistrantId(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.RegistrantId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetTelArea(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.TelArea = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetTelExt(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.TelExt = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetTelMain(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.TelMain = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) SetVspContactId(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact {
	s.VspContactId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizContact) Validate() error {
	return dara.Validate(s)
}

type QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials struct {
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

func (s QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetBizId() *int64 {
	return s.BizId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetCredentialType() *string {
	return s.CredentialType
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetCredentialUrl() *string {
	return s.CredentialUrl
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetHolderCert() *int32 {
	return s.HolderCert
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetBizId(v int64) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.BizId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetCredentialNo(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.CredentialNo = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetCredentialType(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.CredentialType = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetCredentialUrl(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.CredentialUrl = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetDomainName(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.DomainName = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetHolderCert(v int32) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.HolderCert = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) SetSaleId(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials {
	s.SaleId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialBizCredentials) Validate() error {
	return dara.Validate(s)
}

type QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult struct {
	// The cost type.
	//
	// example:
	//
	// activate
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The amount of the order.
	//
	// example:
	//
	// 0.0
	OrderAmount *float64 `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	// The currency.
	//
	// example:
	//
	// CNY
	OrderCurrency *string `json:"OrderCurrency,omitempty" xml:"OrderCurrency,omitempty"`
	// The order ID.
	//
	// example:
	//
	// S201601060123456
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The order status.
	//
	// example:
	//
	// ORDER_WAIT_PAY
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The time when the order was placed.
	//
	// example:
	//
	// 2023-09-19T14:28Z
	OrderTime *string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
	// The validity period.
	//
	// example:
	//
	// 1
	OrderYear *int32 `json:"OrderYear,omitempty" xml:"OrderYear,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// S201601063418719
	SaleId *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	// The suborder ID.
	//
	// example:
	//
	// S201601061234567
	SubOrderId *string `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
}

func (s QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetActionType() *string {
	return s.ActionType
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetOrderAmount() *float64 {
	return s.OrderAmount
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetOrderCurrency() *string {
	return s.OrderCurrency
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetOrderTime() *string {
	return s.OrderTime
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetOrderYear() *int32 {
	return s.OrderYear
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetSaleId() *string {
	return s.SaleId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetActionType(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.ActionType = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetOrderAmount(v float64) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.OrderAmount = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetOrderCurrency(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.OrderCurrency = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetOrderId(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.OrderId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetOrderStatus(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.OrderStatus = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetOrderTime(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.OrderTime = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetOrderYear(v int32) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.OrderYear = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetSaleId(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.SaleId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) SetSubOrderId(v string) *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult {
	s.SubOrderId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailResponseBodyModuleDomainSpecialOrderResult) Validate() error {
	return dara.Validate(s)
}
