// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSignOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendMessage(v string) *CreateDigitalSignOrderShrinkRequest
	GetExtendMessage() *string
	SetOrderContextShrink(v string) *CreateDigitalSignOrderShrinkRequest
	GetOrderContextShrink() *string
	SetOrderType(v string) *CreateDigitalSignOrderShrinkRequest
	GetOrderType() *string
	SetOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetQualificationId() *int64
	SetQualificationVersion(v int64) *CreateDigitalSignOrderShrinkRequest
	GetQualificationVersion() *int64
	SetResourceOwnerAccount(v string) *CreateDigitalSignOrderShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetResourceOwnerId() *int64
	SetSignId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetSignId() *int64
	SetSignIndustry(v int64) *CreateDigitalSignOrderShrinkRequest
	GetSignIndustry() *int64
	SetSignName(v string) *CreateDigitalSignOrderShrinkRequest
	GetSignName() *string
	SetSignSource(v int64) *CreateDigitalSignOrderShrinkRequest
	GetSignSource() *int64
	SetSubmitter(v string) *CreateDigitalSignOrderShrinkRequest
	GetSubmitter() *string
}

type CreateDigitalSignOrderShrinkRequest struct {
	// Reserved for future use.
	//
	// example:
	//
	// example
	ExtendMessage *string `json:"ExtendMessage,omitempty" xml:"ExtendMessage,omitempty"`
	// The qualification information. This object is required when you create a signature, or when you update a signature\\"s qualification information.
	//
	// - qualificationCompanyName: Company name. The name can be up to 150 characters long. It cannot consist of only digits or contain symbols other than the middle dot (·), Chinese brackets (【】), Chinese parentheses (（）), English parentheses (()), and spaces.
	//
	// - `qualificationOrganizationCode`: The 18-character Unified Social Credit Identifier (USCI). It must be an 18-digit code or a code that consists of 18 uppercase or lowercase letters and digits.
	//
	// - `qualificationAdminName`: The name of the agent or legal representative. The name must be in Chinese.
	//
	// - `qualificationAdminIDCard`: The 18-digit ID card number of the agent. Only PRC ID cards are supported.
	//
	// - `qualificationLegalPersonName`: The name of the legal representative or agent.
	//
	// - `qualificationLegalPersonIDCard`: The 18-digit ID card number of the legal representative. Only PRC ID cards are supported.
	//
	// example:
	//
	// {
	//
	//   "qualificationCompanyName": "阿里阿巴",
	//
	//   "qualificationOrganizationCode": "91330106MA2A0XABCD",
	//
	//   "qualificationLegalPersonName": "张三",
	//
	//   "qualificationLegalPersonIDCard": "110105199001011234",
	//
	//   "qualificationAdminName": "李四",
	//
	//   "qualificationAdminIDCard": "11010519900101****"
	//
	// }
	OrderContextShrink *string `json:"OrderContext,omitempty" xml:"OrderContext,omitempty"`
	// The operation to perform on the signature. Valid values:
	//
	// - `UPDATE_DIGITALSMS_SIGN`: Update a signature.
	//
	// - `DELETE_DIGITALSMS_SIGN`: Delete a signature.
	//
	// - `CREATE_DIGITALSMS_SIGN`: Create a signature.
	//
	// example:
	//
	// CREATE_DIGITALSMS_SIGN
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the qualification.
	//
	// example:
	//
	// 41
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The version of the qualification.
	//
	// example:
	//
	// 49
	QualificationVersion *int64  `json:"QualificationVersion,omitempty" xml:"QualificationVersion,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the signature.
	//
	// example:
	//
	// 80
	SignId *int64 `json:"SignId,omitempty" xml:"SignId,omitempty"`
	// The industry type. This parameter is required when you create or update a signature. It is optional when you delete a signature. Valid values:
	//
	// - `0`: General (GENERAL)
	//
	// - `1`: E-commerce and retail (ECOMMERCE)
	//
	// example:
	//
	// 0
	SignIndustry *int64 `json:"SignIndustry,omitempty" xml:"SignIndustry,omitempty"`
	// The signature name. This parameter is required for creating, updating, and deleting signatures.
	//
	// 1. The name must be 2 to 16 characters in length.
	//
	// 2. The name can contain Chinese characters, letters, and digits.
	//
	// - Special characters are not allowed, including $, &, %, #, @, !, ^, \\*, (, ), _, +, -, =, {, }, [, ], |, ;, :, \\", ", <, >, ,, ., /, ?, \\~, and .
	//
	// - The name cannot be only letters.
	//
	// - The name cannot be only digits. Spaces are not allowed.
	//
	// - Emojis are not allowed.
	//
	// example:
	//
	// 阿里云商城
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The signature source. This parameter is required when you create or update a signature. It is optional when you delete a signature. Valid values:
	//
	// - `0`: Enterprises and public institutions
	//
	// - `2`: App
	//
	// example:
	//
	// 0
	SignSource *int64 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The ID of the user who submits the order.
	//
	// example:
	//
	// 110000001750080
	Submitter *string `json:"Submitter,omitempty" xml:"Submitter,omitempty"`
}

func (s CreateDigitalSignOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSignOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalSignOrderShrinkRequest) GetExtendMessage() *string {
	return s.ExtendMessage
}

func (s *CreateDigitalSignOrderShrinkRequest) GetOrderContextShrink() *string {
	return s.OrderContextShrink
}

func (s *CreateDigitalSignOrderShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateDigitalSignOrderShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetQualificationVersion() *int64 {
	return s.QualificationVersion
}

func (s *CreateDigitalSignOrderShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDigitalSignOrderShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignId() *int64 {
	return s.SignId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignIndustry() *int64 {
	return s.SignIndustry
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignSource() *int64 {
	return s.SignSource
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSubmitter() *string {
	return s.Submitter
}

func (s *CreateDigitalSignOrderShrinkRequest) SetExtendMessage(v string) *CreateDigitalSignOrderShrinkRequest {
	s.ExtendMessage = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetOrderContextShrink(v string) *CreateDigitalSignOrderShrinkRequest {
	s.OrderContextShrink = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetOrderType(v string) *CreateDigitalSignOrderShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetQualificationId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetQualificationVersion(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.QualificationVersion = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetResourceOwnerAccount(v string) *CreateDigitalSignOrderShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetResourceOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.SignId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignIndustry(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.SignIndustry = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignName(v string) *CreateDigitalSignOrderShrinkRequest {
	s.SignName = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignSource(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.SignSource = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSubmitter(v string) *CreateDigitalSignOrderShrinkRequest {
	s.Submitter = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
