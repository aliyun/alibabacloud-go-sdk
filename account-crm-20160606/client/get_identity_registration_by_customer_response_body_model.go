// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityRegistrationByCustomerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityRegistrationByCustomerResponseBody
	GetCode() *string
	SetData(v *GetIdentityRegistrationByCustomerResponseBodyData) *GetIdentityRegistrationByCustomerResponseBody
	GetData() *GetIdentityRegistrationByCustomerResponseBodyData
	SetMessage(v string) *GetIdentityRegistrationByCustomerResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityRegistrationByCustomerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityRegistrationByCustomerResponseBody
	GetSuccess() *bool
}

type GetIdentityRegistrationByCustomerResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetIdentityRegistrationByCustomerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIdentityRegistrationByCustomerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityRegistrationByCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityRegistrationByCustomerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityRegistrationByCustomerResponseBody) GetData() *GetIdentityRegistrationByCustomerResponseBodyData {
	return s.Data
}

func (s *GetIdentityRegistrationByCustomerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityRegistrationByCustomerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityRegistrationByCustomerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityRegistrationByCustomerResponseBody) SetCode(v string) *GetIdentityRegistrationByCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBody) SetData(v *GetIdentityRegistrationByCustomerResponseBodyData) *GetIdentityRegistrationByCustomerResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBody) SetMessage(v string) *GetIdentityRegistrationByCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBody) SetRequestId(v string) *GetIdentityRegistrationByCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBody) SetSuccess(v bool) *GetIdentityRegistrationByCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityRegistrationByCustomerResponseBodyData struct {
	AccountType       *int32  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ApplicationId     *int64  `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplyStatus       *string `json:"ApplyStatus,omitempty" xml:"ApplyStatus,omitempty"`
	AuditCode         *string `json:"AuditCode,omitempty" xml:"AuditCode,omitempty"`
	CustomerId        *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	DocBackPic        *string `json:"DocBackPic,omitempty" xml:"DocBackPic,omitempty"`
	DocFrontPic       *string `json:"DocFrontPic,omitempty" xml:"DocFrontPic,omitempty"`
	DocNum            *string `json:"DocNum,omitempty" xml:"DocNum,omitempty"`
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	FullName          *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	RegisteredAddress *string `json:"RegisteredAddress,omitempty" xml:"RegisteredAddress,omitempty"`
	RegisteredCountry *string `json:"RegisteredCountry,omitempty" xml:"RegisteredCountry,omitempty"`
	RegisteredNum     *string `json:"RegisteredNum,omitempty" xml:"RegisteredNum,omitempty"`
	Tel               *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s GetIdentityRegistrationByCustomerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityRegistrationByCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetAccountType() *int32 {
	return s.AccountType
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetApplyStatus() *string {
	return s.ApplyStatus
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetAuditCode() *string {
	return s.AuditCode
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetDocBackPic() *string {
	return s.DocBackPic
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetDocFrontPic() *string {
	return s.DocFrontPic
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetDocNum() *string {
	return s.DocNum
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetFullName() *string {
	return s.FullName
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetRegisteredAddress() *string {
	return s.RegisteredAddress
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetRegisteredCountry() *string {
	return s.RegisteredCountry
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetRegisteredNum() *string {
	return s.RegisteredNum
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) GetTel() *string {
	return s.Tel
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetAccountType(v int32) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetApplicationId(v int64) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetApplyStatus(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.ApplyStatus = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetAuditCode(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.AuditCode = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetCustomerId(v int64) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetDocBackPic(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.DocBackPic = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetDocFrontPic(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.DocFrontPic = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetDocNum(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.DocNum = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetDocType(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.DocType = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetEmail(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetFullName(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.FullName = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetRegisteredAddress(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.RegisteredAddress = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetRegisteredCountry(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.RegisteredCountry = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetRegisteredNum(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.RegisteredNum = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) SetTel(v string) *GetIdentityRegistrationByCustomerResponseBodyData {
	s.Tel = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
