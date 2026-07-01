// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSignOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendMessage(v string) *CreateDigitalSignOrderRequest
	GetExtendMessage() *string
	SetOrderContext(v map[string]interface{}) *CreateDigitalSignOrderRequest
	GetOrderContext() map[string]interface{}
	SetOrderType(v string) *CreateDigitalSignOrderRequest
	GetOrderType() *string
	SetOwnerId(v int64) *CreateDigitalSignOrderRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *CreateDigitalSignOrderRequest
	GetQualificationId() *int64
	SetQualificationVersion(v int64) *CreateDigitalSignOrderRequest
	GetQualificationVersion() *int64
	SetResourceOwnerAccount(v string) *CreateDigitalSignOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDigitalSignOrderRequest
	GetResourceOwnerId() *int64
	SetSignId(v int64) *CreateDigitalSignOrderRequest
	GetSignId() *int64
	SetSignIndustry(v int64) *CreateDigitalSignOrderRequest
	GetSignIndustry() *int64
	SetSignName(v string) *CreateDigitalSignOrderRequest
	GetSignName() *string
	SetSignSource(v int64) *CreateDigitalSignOrderRequest
	GetSignSource() *int64
	SetSubmitter(v string) *CreateDigitalSignOrderRequest
	GetSubmitter() *string
}

type CreateDigitalSignOrderRequest struct {
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
	OrderContext map[string]interface{} `json:"OrderContext,omitempty" xml:"OrderContext,omitempty"`
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

func (s CreateDigitalSignOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSignOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalSignOrderRequest) GetExtendMessage() *string {
	return s.ExtendMessage
}

func (s *CreateDigitalSignOrderRequest) GetOrderContext() map[string]interface{} {
	return s.OrderContext
}

func (s *CreateDigitalSignOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateDigitalSignOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDigitalSignOrderRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *CreateDigitalSignOrderRequest) GetQualificationVersion() *int64 {
	return s.QualificationVersion
}

func (s *CreateDigitalSignOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDigitalSignOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDigitalSignOrderRequest) GetSignId() *int64 {
	return s.SignId
}

func (s *CreateDigitalSignOrderRequest) GetSignIndustry() *int64 {
	return s.SignIndustry
}

func (s *CreateDigitalSignOrderRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateDigitalSignOrderRequest) GetSignSource() *int64 {
	return s.SignSource
}

func (s *CreateDigitalSignOrderRequest) GetSubmitter() *string {
	return s.Submitter
}

func (s *CreateDigitalSignOrderRequest) SetExtendMessage(v string) *CreateDigitalSignOrderRequest {
	s.ExtendMessage = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetOrderContext(v map[string]interface{}) *CreateDigitalSignOrderRequest {
	s.OrderContext = v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetOrderType(v string) *CreateDigitalSignOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetOwnerId(v int64) *CreateDigitalSignOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetQualificationId(v int64) *CreateDigitalSignOrderRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetQualificationVersion(v int64) *CreateDigitalSignOrderRequest {
	s.QualificationVersion = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetResourceOwnerAccount(v string) *CreateDigitalSignOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetResourceOwnerId(v int64) *CreateDigitalSignOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignId(v int64) *CreateDigitalSignOrderRequest {
	s.SignId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignIndustry(v int64) *CreateDigitalSignOrderRequest {
	s.SignIndustry = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignName(v string) *CreateDigitalSignOrderRequest {
	s.SignName = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignSource(v int64) *CreateDigitalSignOrderRequest {
	s.SignSource = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSubmitter(v string) *CreateDigitalSignOrderRequest {
	s.Submitter = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) Validate() error {
	return dara.Validate(s)
}
