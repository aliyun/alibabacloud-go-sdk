// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContact(v *BookRequestContact) *BookRequest
	GetContact() *BookRequestContact
	SetOutOrderNum(v string) *BookRequest
	GetOutOrderNum() *string
	SetPassengerAncillaryPurchaseMapList(v []*BookRequestPassengerAncillaryPurchaseMapList) *BookRequest
	GetPassengerAncillaryPurchaseMapList() []*BookRequestPassengerAncillaryPurchaseMapList
	SetPassengerList(v []*BookRequestPassengerList) *BookRequest
	GetPassengerList() []*BookRequestPassengerList
	SetSolutionId(v string) *BookRequest
	GetSolutionId() *string
}

type BookRequest struct {
	// The contact information.
	//
	// This parameter is required.
	Contact *BookRequestContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// The external order number.
	//
	// This parameter is required.
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// The mapping between passengers and ancillary purchases.
	PassengerAncillaryPurchaseMapList []*BookRequestPassengerAncillaryPurchaseMapList `json:"passenger_ancillary_purchase_map_list,omitempty" xml:"passenger_ancillary_purchase_map_list,omitempty" type:"Repeated"`
	// The list of passengers.
	//
	// This parameter is required.
	PassengerList []*BookRequestPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// solution_id.
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s BookRequest) String() string {
	return dara.Prettify(s)
}

func (s BookRequest) GoString() string {
	return s.String()
}

func (s *BookRequest) GetContact() *BookRequestContact {
	return s.Contact
}

func (s *BookRequest) GetOutOrderNum() *string {
	return s.OutOrderNum
}

func (s *BookRequest) GetPassengerAncillaryPurchaseMapList() []*BookRequestPassengerAncillaryPurchaseMapList {
	return s.PassengerAncillaryPurchaseMapList
}

func (s *BookRequest) GetPassengerList() []*BookRequestPassengerList {
	return s.PassengerList
}

func (s *BookRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *BookRequest) SetContact(v *BookRequestContact) *BookRequest {
	s.Contact = v
	return s
}

func (s *BookRequest) SetOutOrderNum(v string) *BookRequest {
	s.OutOrderNum = &v
	return s
}

func (s *BookRequest) SetPassengerAncillaryPurchaseMapList(v []*BookRequestPassengerAncillaryPurchaseMapList) *BookRequest {
	s.PassengerAncillaryPurchaseMapList = v
	return s
}

func (s *BookRequest) SetPassengerList(v []*BookRequestPassengerList) *BookRequest {
	s.PassengerList = v
	return s
}

func (s *BookRequest) SetSolutionId(v string) *BookRequest {
	s.SolutionId = &v
	return s
}

func (s *BookRequest) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	if s.PassengerAncillaryPurchaseMapList != nil {
		for _, item := range s.PassengerAncillaryPurchaseMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PassengerList != nil {
		for _, item := range s.PassengerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BookRequestContact struct {
	// The email address.
	//
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The first name.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The last name.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// The country calling code.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 183******96
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s BookRequestContact) String() string {
	return dara.Prettify(s)
}

func (s BookRequestContact) GoString() string {
	return s.String()
}

func (s *BookRequestContact) GetEmail() *string {
	return s.Email
}

func (s *BookRequestContact) GetFirstName() *string {
	return s.FirstName
}

func (s *BookRequestContact) GetLastName() *string {
	return s.LastName
}

func (s *BookRequestContact) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *BookRequestContact) GetMobilePhoneNum() *string {
	return s.MobilePhoneNum
}

func (s *BookRequestContact) SetEmail(v string) *BookRequestContact {
	s.Email = &v
	return s
}

func (s *BookRequestContact) SetFirstName(v string) *BookRequestContact {
	s.FirstName = &v
	return s
}

func (s *BookRequestContact) SetLastName(v string) *BookRequestContact {
	s.LastName = &v
	return s
}

func (s *BookRequestContact) SetMobileCountryCode(v string) *BookRequestContact {
	s.MobileCountryCode = &v
	return s
}

func (s *BookRequestContact) SetMobilePhoneNum(v string) *BookRequestContact {
	s.MobilePhoneNum = &v
	return s
}

func (s *BookRequestContact) Validate() error {
	return dara.Validate(s)
}

type BookRequestPassengerAncillaryPurchaseMapList struct {
	// The ancillary product object for the booking request.
	BookAncillaryReqItem *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem `json:"book_ancillary_req_item,omitempty" xml:"book_ancillary_req_item,omitempty" type:"Struct"`
	// The list of passengers who purchase the same ancillary product.
	PassengerList []*BookRequestPassengerAncillaryPurchaseMapListPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
}

func (s BookRequestPassengerAncillaryPurchaseMapList) String() string {
	return dara.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapList) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) GetBookAncillaryReqItem() *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem {
	return s.BookAncillaryReqItem
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) GetPassengerList() []*BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	return s.PassengerList
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) SetBookAncillaryReqItem(v *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) *BookRequestPassengerAncillaryPurchaseMapList {
	s.BookAncillaryReqItem = v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) SetPassengerList(v []*BookRequestPassengerAncillaryPurchaseMapListPassengerList) *BookRequestPassengerAncillaryPurchaseMapList {
	s.PassengerList = v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) Validate() error {
	if s.BookAncillaryReqItem != nil {
		if err := s.BookAncillaryReqItem.Validate(); err != nil {
			return err
		}
	}
	if s.PassengerList != nil {
		for _, item := range s.PassengerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem struct {
	// The ancillary product ID.
	//
	// example:
	//
	// MDY2NTAxLCJleHAiOjE2NxNzM3MDEsIm5ix
	AncillaryId *string `json:"ancillary_id,omitempty" xml:"ancillary_id,omitempty"`
	// The ancillary product type. Currently supported value: 4 (paid baggage). More types will be supported in the future.
	//
	// example:
	//
	// 4
	AncillaryType *int32 `json:"ancillary_type,omitempty" xml:"ancillary_type,omitempty"`
}

func (s BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) String() string {
	return dara.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) GetAncillaryId() *string {
	return s.AncillaryId
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) GetAncillaryType() *int32 {
	return s.AncillaryType
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) SetAncillaryId(v string) *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem {
	s.AncillaryId = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) SetAncillaryType(v int32) *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem {
	s.AncillaryType = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) Validate() error {
	return dara.Validate(s)
}

type BookRequestPassengerAncillaryPurchaseMapListPassengerList struct {
	// The date of birth in yyyyMMdd format.
	//
	// example:
	//
	// 20020320
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The credential information.
	Credential *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The first name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The gender. Valid values:
	//
	// - 0: MALE
	//
	// - 1: FEMALE.
	//
	// example:
	//
	// 1
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// The last name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// The country calling code for the mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 182******92
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// The nationality.
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// The passenger type. Valid values:
	//
	// - 0: adult
	//
	// - 1: child
	//
	// - 8: infant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerList) String() string {
	return dara.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerList) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetCredential() *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	return s.Credential
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetFirstName() *string {
	return s.FirstName
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetLastName() *string {
	return s.LastName
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) GetType() *int32 {
	return s.Type
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetBirthday(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Birthday = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetCredential(v *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Credential = v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetFirstName(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.FirstName = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetGender(v int32) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Gender = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetLastName(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.LastName = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetMobileCountryCode(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetMobilePhoneNumber(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetNationality(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Nationality = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetType(v int32) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Type = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential struct {
	// The place of issue. Use a two-letter country code.
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// The credential number.
	//
	// example:
	//
	// E1***5673
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// The credential type. Valid values:
	//
	// - 0: ID card
	//
	// - 1: passport
	//
	// - 2: student ID
	//
	// - 3: military ID
	//
	// - 4: Home Return Permit
	//
	// - 5: Taiwan Compatriot Permit
	//
	// - 6: Hong Kong and Macau Travel Permit
	//
	// - 7: international seafarer certificate
	//
	// - 8: foreigner permanent residence permit
	//
	// - 10: police officer certificate
	//
	// - 11: soldier certificate
	//
	// - 12: Taiwan Travel Permit
	//
	// - 13: Taiwan Entry Permit
	//
	// - 14: household register
	//
	// - 15: birth certificate
	//
	// - 16: driver license
	//
	// - 17: Hong Kong and Macau resident residence permit
	//
	// - 18: Taiwan resident residence permit.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// The expiration date of the credential.
	//
	// example:
	//
	// 20290102
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) String() string {
	return dara.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) GetCredentialNum() *string {
	return s.CredentialNum
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) GetCredentialType() *int32 {
	return s.CredentialType
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetCertIssuePlace(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetCredentialNum(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetCredentialType(v int32) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetExpireDate(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.ExpireDate = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) Validate() error {
	return dara.Validate(s)
}

type BookRequestPassengerList struct {
	// The date of birth in yyyyMMdd format.
	//
	// example:
	//
	// 20200320
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The credential information.
	Credential *BookRequestPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The first name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The gender. Valid values:
	//
	// - 0: MALE
	//
	// - 1: FEMALE.
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// The last name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// The country calling code for the mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 183******95
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// The nationality. Use a two-letter country code.
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// The passenger type. Valid values:
	//
	// - 0: adult
	//
	// - 1: child
	//
	// - 8: infant.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BookRequestPassengerList) String() string {
	return dara.Prettify(s)
}

func (s BookRequestPassengerList) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *BookRequestPassengerList) GetCredential() *BookRequestPassengerListCredential {
	return s.Credential
}

func (s *BookRequestPassengerList) GetFirstName() *string {
	return s.FirstName
}

func (s *BookRequestPassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *BookRequestPassengerList) GetLastName() *string {
	return s.LastName
}

func (s *BookRequestPassengerList) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *BookRequestPassengerList) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *BookRequestPassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *BookRequestPassengerList) GetType() *int32 {
	return s.Type
}

func (s *BookRequestPassengerList) SetBirthday(v string) *BookRequestPassengerList {
	s.Birthday = &v
	return s
}

func (s *BookRequestPassengerList) SetCredential(v *BookRequestPassengerListCredential) *BookRequestPassengerList {
	s.Credential = v
	return s
}

func (s *BookRequestPassengerList) SetFirstName(v string) *BookRequestPassengerList {
	s.FirstName = &v
	return s
}

func (s *BookRequestPassengerList) SetGender(v int32) *BookRequestPassengerList {
	s.Gender = &v
	return s
}

func (s *BookRequestPassengerList) SetLastName(v string) *BookRequestPassengerList {
	s.LastName = &v
	return s
}

func (s *BookRequestPassengerList) SetMobileCountryCode(v string) *BookRequestPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *BookRequestPassengerList) SetMobilePhoneNumber(v string) *BookRequestPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *BookRequestPassengerList) SetNationality(v string) *BookRequestPassengerList {
	s.Nationality = &v
	return s
}

func (s *BookRequestPassengerList) SetType(v int32) *BookRequestPassengerList {
	s.Type = &v
	return s
}

func (s *BookRequestPassengerList) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BookRequestPassengerListCredential struct {
	// The place of issue. Use a two-letter country code.
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// The credential number.
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// The credential type. Valid values:
	//
	// - 0: ID card
	//
	// - 1: passport
	//
	// - 4: Home Return Permit
	//
	// - 5: Taiwan Compatriot Permit
	//
	// - 6: Hong Kong and Macau Travel Permit
	//
	// - 12: Taiwan Travel Permit
	//
	// - 19: no credential.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// The expiration date of the credential.
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s BookRequestPassengerListCredential) String() string {
	return dara.Prettify(s)
}

func (s BookRequestPassengerListCredential) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerListCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *BookRequestPassengerListCredential) GetCredentialNum() *string {
	return s.CredentialNum
}

func (s *BookRequestPassengerListCredential) GetCredentialType() *int32 {
	return s.CredentialType
}

func (s *BookRequestPassengerListCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *BookRequestPassengerListCredential) SetCertIssuePlace(v string) *BookRequestPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *BookRequestPassengerListCredential) SetCredentialNum(v string) *BookRequestPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *BookRequestPassengerListCredential) SetCredentialType(v int32) *BookRequestPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *BookRequestPassengerListCredential) SetExpireDate(v string) *BookRequestPassengerListCredential {
	s.ExpireDate = &v
	return s
}

func (s *BookRequestPassengerListCredential) Validate() error {
	return dara.Validate(s)
}
