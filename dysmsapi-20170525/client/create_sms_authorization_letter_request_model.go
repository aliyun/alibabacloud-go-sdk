// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAuthorizationLetterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorization(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorization() *string
	SetAuthorizationLetterExpDate(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorizationLetterExpDate() *string
	SetAuthorizationLetterName(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorizationLetterName() *string
	SetAuthorizationLetterPic(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorizationLetterPic() *string
	SetOrganizationCode(v string) *CreateSmsAuthorizationLetterRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *CreateSmsAuthorizationLetterRequest
	GetOwnerId() *int64
	SetProxyAuthorization(v string) *CreateSmsAuthorizationLetterRequest
	GetProxyAuthorization() *string
	SetResourceOwnerAccount(v string) *CreateSmsAuthorizationLetterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsAuthorizationLetterRequest
	GetResourceOwnerId() *int64
	SetSignList(v []*string) *CreateSmsAuthorizationLetterRequest
	GetSignList() []*string
}

type CreateSmsAuthorizationLetterRequest struct {
	// The authorizing party, that is, the owner of the signature. Only the middle dot `·`, Chinese `【】（）`, English `()`, and spaces are supported. Other symbols or purely numeric input are not allowed. The length cannot exceed 150 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 菜鸟网络科技有限公司
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// The validity period of the authorization letter. Format: `YYYY-MM-DD~YYYY-MM-DD`.
	//
	// > The recommended validity period is 1 to 3 years. Set a reasonable time period and avoid making the validity period too long or too short.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2026-01-01
	AuthorizationLetterExpDate *string `json:"AuthorizationLetterExpDate,omitempty" xml:"AuthorizationLetterExpDate,omitempty"`
	// The name of the authorization letter. The name cannot be the same as any of your other authorization letters. Only Chinese characters, English characters, or a combination with numbers are supported. Symbols or purely numeric input are not allowed. The length cannot exceed 100 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 菜鸟公司授权书
	AuthorizationLetterName *string `json:"AuthorizationLetterName,omitempty" xml:"AuthorizationLetterName,omitempty"`
	// The fileKey of the authorization letter.
	//
	// 1. The authorization letter file uploaded to OSS. Download the [Authorization Letter Template](https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20250414/bvpcmo/%E6%8E%88%E6%9D%83%E5%A7%94%E6%89%98%E4%B9%A6%E6%A8%A1%E7%89%88.doc), then fill it out and stamp it according to the [specifications](https://help.aliyun.com/document_detail/56741.html) before uploading. File upload requirements:
	//
	// - The name of the file to be uploaded cannot contain Chinese characters or special characters.
	//
	// - Only images in JPG, PNG, GIF, or JPEG format are supported, and the image must not exceed 5 MB.
	//
	// 2. To obtain the fileKey, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	AuthorizationLetterPic *string `json:"AuthorizationLetterPic,omitempty" xml:"AuthorizationLetterPic,omitempty"`
	// The unified social credit code of the authorizing party. The length cannot exceed 150 characters. The credit code must be consistent with the unified social credit code field in the qualification information bound to the signature. Otherwise, the signature creation will fail.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9****************A
	OrganizationCode *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The authorized party, that is, the signature applicant. Only the middle dot `·`, Chinese `【】（）`, English `()`, and spaces are supported. Other symbols or purely numeric input are not allowed. The length cannot exceed 150 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 支付宝（中国）网络技术有限公司
	ProxyAuthorization   *string `json:"ProxyAuthorization,omitempty" xml:"ProxyAuthorization,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of authorized signatures. The number of signatures cannot exceed 100.
	//
	// > We recommend that you authorize all signatures that may be used at once in the authorization letter. This prevents subsequent signature applications from falling outside the scope of the authorization letter, which would cause review failure and require you to supplement the authorization letter.
	//
	// This parameter is required.
	SignList []*string `json:"SignList,omitempty" xml:"SignList,omitempty" type:"Repeated"`
}

func (s CreateSmsAuthorizationLetterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAuthorizationLetterRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorizationLetterExpDate() *string {
	return s.AuthorizationLetterExpDate
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorizationLetterName() *string {
	return s.AuthorizationLetterName
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorizationLetterPic() *string {
	return s.AuthorizationLetterPic
}

func (s *CreateSmsAuthorizationLetterRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *CreateSmsAuthorizationLetterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsAuthorizationLetterRequest) GetProxyAuthorization() *string {
	return s.ProxyAuthorization
}

func (s *CreateSmsAuthorizationLetterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsAuthorizationLetterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsAuthorizationLetterRequest) GetSignList() []*string {
	return s.SignList
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorization(v string) *CreateSmsAuthorizationLetterRequest {
	s.Authorization = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorizationLetterExpDate(v string) *CreateSmsAuthorizationLetterRequest {
	s.AuthorizationLetterExpDate = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorizationLetterName(v string) *CreateSmsAuthorizationLetterRequest {
	s.AuthorizationLetterName = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorizationLetterPic(v string) *CreateSmsAuthorizationLetterRequest {
	s.AuthorizationLetterPic = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetOrganizationCode(v string) *CreateSmsAuthorizationLetterRequest {
	s.OrganizationCode = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetOwnerId(v int64) *CreateSmsAuthorizationLetterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetProxyAuthorization(v string) *CreateSmsAuthorizationLetterRequest {
	s.ProxyAuthorization = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetResourceOwnerAccount(v string) *CreateSmsAuthorizationLetterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetResourceOwnerId(v int64) *CreateSmsAuthorizationLetterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetSignList(v []*string) *CreateSmsAuthorizationLetterRequest {
	s.SignList = v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) Validate() error {
	return dara.Validate(s)
}
