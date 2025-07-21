// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIsvCustomerTermsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessDesc(v string) *SubmitIsvCustomerTermsRequest
	GetBusinessDesc() *string
	SetContactMail(v string) *SubmitIsvCustomerTermsRequest
	GetContactMail() *string
	SetCountryId(v string) *SubmitIsvCustomerTermsRequest
	GetCountryId() *string
	SetCustName(v string) *SubmitIsvCustomerTermsRequest
	GetCustName() *string
	SetCustSpaceId(v string) *SubmitIsvCustomerTermsRequest
	GetCustSpaceId() *string
	SetIsvTerms(v string) *SubmitIsvCustomerTermsRequest
	GetIsvTerms() *string
	SetOfficeAddress(v string) *SubmitIsvCustomerTermsRequest
	GetOfficeAddress() *string
	SetOwnerId(v int64) *SubmitIsvCustomerTermsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SubmitIsvCustomerTermsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitIsvCustomerTermsRequest
	GetResourceOwnerId() *int64
}

type SubmitIsvCustomerTermsRequest struct {
	// The business scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// Marketing products
	BusinessDesc *string `json:"BusinessDesc,omitempty" xml:"BusinessDesc,omitempty"`
	// The enterprise mail.
	//
	// This parameter is required.
	//
	// example:
	//
	// partner@aliyun.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// The country code.
	//
	// >  For more information about country codes, see [Country codes](https://help.aliyun.com/document_detail/608210.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// CN
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The enterprise name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	CustName *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The ISV or Client agreement.
	//
	// This parameter is required.
	//
	// example:
	//
	// isvTerms.pdf
	IsvTerms *string `json:"IsvTerms,omitempty" xml:"IsvTerms,omitempty"`
	// The enterprise address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	OfficeAddress        *string `json:"OfficeAddress,omitempty" xml:"OfficeAddress,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SubmitIsvCustomerTermsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIsvCustomerTermsRequest) GoString() string {
	return s.String()
}

func (s *SubmitIsvCustomerTermsRequest) GetBusinessDesc() *string {
	return s.BusinessDesc
}

func (s *SubmitIsvCustomerTermsRequest) GetContactMail() *string {
	return s.ContactMail
}

func (s *SubmitIsvCustomerTermsRequest) GetCountryId() *string {
	return s.CountryId
}

func (s *SubmitIsvCustomerTermsRequest) GetCustName() *string {
	return s.CustName
}

func (s *SubmitIsvCustomerTermsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SubmitIsvCustomerTermsRequest) GetIsvTerms() *string {
	return s.IsvTerms
}

func (s *SubmitIsvCustomerTermsRequest) GetOfficeAddress() *string {
	return s.OfficeAddress
}

func (s *SubmitIsvCustomerTermsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitIsvCustomerTermsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitIsvCustomerTermsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitIsvCustomerTermsRequest) SetBusinessDesc(v string) *SubmitIsvCustomerTermsRequest {
	s.BusinessDesc = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetContactMail(v string) *SubmitIsvCustomerTermsRequest {
	s.ContactMail = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetCountryId(v string) *SubmitIsvCustomerTermsRequest {
	s.CountryId = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetCustName(v string) *SubmitIsvCustomerTermsRequest {
	s.CustName = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetCustSpaceId(v string) *SubmitIsvCustomerTermsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetIsvTerms(v string) *SubmitIsvCustomerTermsRequest {
	s.IsvTerms = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetOfficeAddress(v string) *SubmitIsvCustomerTermsRequest {
	s.OfficeAddress = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetOwnerId(v int64) *SubmitIsvCustomerTermsRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetResourceOwnerAccount(v string) *SubmitIsvCustomerTermsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) SetResourceOwnerId(v int64) *SubmitIsvCustomerTermsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitIsvCustomerTermsRequest) Validate() error {
	return dara.Validate(s)
}
