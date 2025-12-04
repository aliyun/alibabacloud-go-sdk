// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotlineTransferRegisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgreement(v string) *SubmitHotlineTransferRegisterRequest
	GetAgreement() *string
	SetHotlineNumber(v string) *SubmitHotlineTransferRegisterRequest
	GetHotlineNumber() *string
	SetOperatorIdentityCard(v string) *SubmitHotlineTransferRegisterRequest
	GetOperatorIdentityCard() *string
	SetOperatorMail(v string) *SubmitHotlineTransferRegisterRequest
	GetOperatorMail() *string
	SetOperatorMailVerifyCode(v string) *SubmitHotlineTransferRegisterRequest
	GetOperatorMailVerifyCode() *string
	SetOperatorMobile(v string) *SubmitHotlineTransferRegisterRequest
	GetOperatorMobile() *string
	SetOperatorMobileVerifyCode(v string) *SubmitHotlineTransferRegisterRequest
	GetOperatorMobileVerifyCode() *string
	SetOperatorName(v string) *SubmitHotlineTransferRegisterRequest
	GetOperatorName() *string
	SetOwnerId(v int64) *SubmitHotlineTransferRegisterRequest
	GetOwnerId() *int64
	SetQualificationId(v string) *SubmitHotlineTransferRegisterRequest
	GetQualificationId() *string
	SetResourceOwnerAccount(v string) *SubmitHotlineTransferRegisterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitHotlineTransferRegisterRequest
	GetResourceOwnerId() *int64
	SetTransferPhoneNumberInfos(v []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) *SubmitHotlineTransferRegisterRequest
	GetTransferPhoneNumberInfos() []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos
}

type SubmitHotlineTransferRegisterRequest struct {
	// The authenticity of the commitment. Valid values:
	//
	// 	- **true**: The commitment is authentic.
	//
	// 	- **false**: The commitment is not authentic.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Agreement *string `json:"Agreement,omitempty" xml:"Agreement,omitempty"`
	// The China 400 number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400****
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The ID card number of the handler.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5****************9
	OperatorIdentityCard *string `json:"OperatorIdentityCard,omitempty" xml:"OperatorIdentityCard,omitempty"`
	// The email address of the handler.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	OperatorMail *string `json:"OperatorMail,omitempty" xml:"OperatorMail,omitempty"`
	// The verification code that is received by the mailbox of the handler.
	//
	// example:
	//
	// 1234
	OperatorMailVerifyCode *string `json:"OperatorMailVerifyCode,omitempty" xml:"OperatorMailVerifyCode,omitempty"`
	// The mobile phone number of the handler.
	//
	// This parameter is required.
	//
	// example:
	//
	// 158****7230
	OperatorMobile *string `json:"OperatorMobile,omitempty" xml:"OperatorMobile,omitempty"`
	// The verification code that is received by the mobile phone of the handler.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	OperatorMobileVerifyCode *string `json:"OperatorMobileVerifyCode,omitempty" xml:"OperatorMobileVerifyCode,omitempty"`
	// The name of the handler.
	//
	// This parameter is required.
	//
	// example:
	//
	// A***
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](https://help.aliyun.com/document_detail/393548.html) operation to obtain the qualification ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000004933****
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The registration information about the China 400 number.
	//
	// This parameter is required.
	TransferPhoneNumberInfos []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos `json:"TransferPhoneNumberInfos,omitempty" xml:"TransferPhoneNumberInfos,omitempty" type:"Repeated"`
}

func (s SubmitHotlineTransferRegisterRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotlineTransferRegisterRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterRequest) GetAgreement() *string {
	return s.Agreement
}

func (s *SubmitHotlineTransferRegisterRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *SubmitHotlineTransferRegisterRequest) GetOperatorIdentityCard() *string {
	return s.OperatorIdentityCard
}

func (s *SubmitHotlineTransferRegisterRequest) GetOperatorMail() *string {
	return s.OperatorMail
}

func (s *SubmitHotlineTransferRegisterRequest) GetOperatorMailVerifyCode() *string {
	return s.OperatorMailVerifyCode
}

func (s *SubmitHotlineTransferRegisterRequest) GetOperatorMobile() *string {
	return s.OperatorMobile
}

func (s *SubmitHotlineTransferRegisterRequest) GetOperatorMobileVerifyCode() *string {
	return s.OperatorMobileVerifyCode
}

func (s *SubmitHotlineTransferRegisterRequest) GetOperatorName() *string {
	return s.OperatorName
}

func (s *SubmitHotlineTransferRegisterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitHotlineTransferRegisterRequest) GetQualificationId() *string {
	return s.QualificationId
}

func (s *SubmitHotlineTransferRegisterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitHotlineTransferRegisterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitHotlineTransferRegisterRequest) GetTransferPhoneNumberInfos() []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	return s.TransferPhoneNumberInfos
}

func (s *SubmitHotlineTransferRegisterRequest) SetAgreement(v string) *SubmitHotlineTransferRegisterRequest {
	s.Agreement = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetHotlineNumber(v string) *SubmitHotlineTransferRegisterRequest {
	s.HotlineNumber = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorIdentityCard(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorIdentityCard = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMail(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMail = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMailVerifyCode(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMailVerifyCode = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMobile(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMobile = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMobileVerifyCode(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMobileVerifyCode = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorName(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorName = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOwnerId(v int64) *SubmitHotlineTransferRegisterRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetQualificationId(v string) *SubmitHotlineTransferRegisterRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetResourceOwnerAccount(v string) *SubmitHotlineTransferRegisterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetResourceOwnerId(v int64) *SubmitHotlineTransferRegisterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetTransferPhoneNumberInfos(v []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) *SubmitHotlineTransferRegisterRequest {
	s.TransferPhoneNumberInfos = v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) Validate() error {
	if s.TransferPhoneNumberInfos != nil {
		for _, item := range s.TransferPhoneNumberInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos struct {
	// The ID card number of the number owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500***
	IdentityCard *string `json:"IdentityCard,omitempty" xml:"IdentityCard,omitempty"`
	// The China 400 number that you want to submit for registration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1580000****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The real name or company name of the number owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// A***
	PhoneNumberOwnerName *string `json:"PhoneNumberOwnerName,omitempty" xml:"PhoneNumberOwnerName,omitempty"`
}

func (s SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) GetIdentityCard() *string {
	return s.IdentityCard
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) GetPhoneNumberOwnerName() *string {
	return s.PhoneNumberOwnerName
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) SetIdentityCard(v string) *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	s.IdentityCard = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) SetPhoneNumber(v string) *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	s.PhoneNumber = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) SetPhoneNumberOwnerName(v string) *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	s.PhoneNumberOwnerName = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) Validate() error {
	return dara.Validate(s)
}
