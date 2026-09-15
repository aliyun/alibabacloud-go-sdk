// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorAccountAKRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkType(v string) *AddCloudVendorAccountAKRequest
	GetAkType() *string
	SetAuthModules(v []*string) *AddCloudVendorAccountAKRequest
	GetAuthModules() []*string
	SetCtdrCloudUserId(v string) *AddCloudVendorAccountAKRequest
	GetCtdrCloudUserId() *string
	SetDomain(v string) *AddCloudVendorAccountAKRequest
	GetDomain() *string
	SetExtendInfo(v string) *AddCloudVendorAccountAKRequest
	GetExtendInfo() *string
	SetLang(v string) *AddCloudVendorAccountAKRequest
	GetLang() *string
	SetRegions(v []*string) *AddCloudVendorAccountAKRequest
	GetRegions() []*string
	SetSecretId(v string) *AddCloudVendorAccountAKRequest
	GetSecretId() *string
	SetSecretKey(v string) *AddCloudVendorAccountAKRequest
	GetSecretKey() *string
	SetSubscriptionIds(v []*string) *AddCloudVendorAccountAKRequest
	GetSubscriptionIds() []*string
	SetTenantId(v string) *AddCloudVendorAccountAKRequest
	GetTenantId() *string
	SetVendor(v string) *AddCloudVendorAccountAKRequest
	GetVendor() *string
	SetVendorAuthAlias(v string) *AddCloudVendorAccountAKRequest
	GetVendorAuthAlias() *string
}

type AddCloudVendorAccountAKRequest struct {
	// The AccessKey (AK) type. Valid values:
	//
	// - **primary**: Primary account.
	//
	// - **sub**: Sub-account.
	//
	// - **ctdr**: Agentic SOC.
	//
	// 	Warning: When the vendor is **CHAITIN**, **FORTINET**, **THREATBOOK**, or **WIZ**, set this parameter to ctdr.</warning>
	//
	// This parameter is required.
	//
	// example:
	//
	// primary
	AkType *string `json:"AkType,omitempty" xml:"AkType,omitempty"`
	// The list of AK-associated modules.
	AuthModules []*string `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	// The account ID.
	//
	// > The account ID of the connected cloud vendor. This parameter is required when the permission description includes Cloud Threat Detection and Response (CTDR).
	//
	// example:
	//
	// azure_demo_1
	CtdrCloudUserId *string `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The account domain for access. Valid values:
	//
	// - **china**: China
	//
	// - **global**: Global
	//
	// - **europe**: Huawei Cloud Europe
	//
	// > This parameter is valid only when **Vendor*	- is set to **HUAWEICLOUD**, **Azure**, **AWS**, **VOLCENGINE**, **KingsoftCloud**, **UCloud**, or **BaiduCloud**, and is required. Set this parameter to **china*	- for KingsoftCloud and BaiduCloud, and to **global*	- for UCloud.
	//
	// example:
	//
	// global
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The extended information.
	//
	// > Used to record extended information for different vendors.
	//
	// >Google Cloud is accessed through a service account. ExtendInfo stores the JSON-formatted service key file, excluding the private_key_id and private_key fields. The file contains the following fields: type, project_id, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url, and universe_domain.
	//
	// example:
	//
	// {\\"product\\":\\"webFirewall\\",\\"remark\\":\\"remark\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of regions used for AK information verification. This parameter is valid only when Vendor is set to AWS.
	//
	// >Call the [ListCloudVendorRegions](~~ListCloudVendorRegions~~) operation to obtain this parameter.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The AK parameter ID. Valid values:
	//
	// 1. When AkType is set to primary:
	//
	// - **Tencent**: AccessKeyId of the primary account
	//
	// - **HUAWEICLOUD**: AccessKeyId of the primary account
	//
	// - **Azure**: ClientId
	//
	// - **AWS**: AccessKeyId of the primary account
	//
	// - **VOLCENGINE**: AccessKeyId of the primary account
	//
	// 2. When AkType is set to sub:
	//
	// - **Tencent**: AccessKeyId of the sub-account
	//
	// - **HUAWEICLOUD**: AccessKeyId of the sub-account
	//
	// - **Azure**: ClientId
	//
	// - **AWS**: AccessKeyId of the sub-account
	//
	// - **VOLCENGINE**: AccessKeyId of the sub-account
	//
	// - **google**: private_key_id
	//
	// >If AkType is set to **primary**, this value is the SecretID of the primary account on the third-party cloud. If AkType is set to **sub**, this value is the Access Key ID of the sub-account on the third-party cloud. For **Azure**, no distinction is made, and this value is the **appId*	- in the authentication information. Google Cloud is accessed through a service account. AkType defaults to sub, and this value is the private_key_id property value from the JSON-formatted service key file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45GLRV4SOT0YFB****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The AK parameter secret. Valid values:
	//
	// 1. When AkType is set to primary:
	//
	// - **Tencent**: SecretAccessKey of the primary account
	//
	// - **HUAWEICLOUD**: SecretAccessKey of the primary account
	//
	// - **Azure**: ClientSecret
	//
	// - **AWS**: SecretAccessKey of the primary account
	//
	// 2. When AkType is set to sub:
	//
	// - **Tencent**: SecretAccessKey of the sub-account
	//
	// - **HUAWEICLOUD**: SecretAccessKey of the sub-account
	//
	// - **Azure**: ClientSecret
	//
	// - **AWS**: SecretAccessKey of the sub-account
	//
	// - **google**: private_key
	//
	// >If AkType is set to **primary**, this value is the Secret Access Key of the primary account on the third-party cloud. If AkType is set to **sub**, this value is the Secret Access Key of the sub-account on the third-party cloud. For **Azure**, no distinction is made, and this value is the **password*	- in the authentication information. Google Cloud is accessed through a service account. AkType defaults to sub, and this value is the private_key property value from the JSON-formatted service key file.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE6SLd****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The list of subscription IDs.
	//
	// > This parameter is no longer valid.
	SubscriptionIds []*string `json:"SubscriptionIds,omitempty" xml:"SubscriptionIds,omitempty" type:"Repeated"`
	// The tenant ID. This parameter is valid only when Vendor is set to Azure.
	//
	// example:
	//
	// 95304a97-339b-4de5-9a7d-cdbffaf****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The cloud asset vendor. Valid values:
	//
	// - **Tencent**: Tencent Cloud
	//
	// - **HUAWEICLOUD**: Huawei Cloud
	//
	// - **Azure**: Azure
	//
	// - **AWS**: AWS
	//
	// - **VOLCENGINE**: Volcengine
	//
	// - **google**: Google Cloud
	//
	// - **CHAITIN**: Chaitin Technology
	//
	// - **FORTINET**: Fortinet
	//
	// - **THREATBOOK**: ThreatBook
	//
	// - **KingsoftCloud**: Kingsoft Cloud
	//
	// - **UCloud**: UCloud
	//
	// - **BaiduCloud**: Baidu AI Cloud
	//
	// - **WIZ**: Wiz Security
	//
	// This parameter is required.
	//
	// example:
	//
	// AWS
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The AK account name.
	//
	// >Used to identify the account to which third-party host assets belong.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
}

func (s AddCloudVendorAccountAKRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorAccountAKRequest) GoString() string {
	return s.String()
}

func (s *AddCloudVendorAccountAKRequest) GetAkType() *string {
	return s.AkType
}

func (s *AddCloudVendorAccountAKRequest) GetAuthModules() []*string {
	return s.AuthModules
}

func (s *AddCloudVendorAccountAKRequest) GetCtdrCloudUserId() *string {
	return s.CtdrCloudUserId
}

func (s *AddCloudVendorAccountAKRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddCloudVendorAccountAKRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *AddCloudVendorAccountAKRequest) GetLang() *string {
	return s.Lang
}

func (s *AddCloudVendorAccountAKRequest) GetRegions() []*string {
	return s.Regions
}

func (s *AddCloudVendorAccountAKRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *AddCloudVendorAccountAKRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *AddCloudVendorAccountAKRequest) GetSubscriptionIds() []*string {
	return s.SubscriptionIds
}

func (s *AddCloudVendorAccountAKRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *AddCloudVendorAccountAKRequest) GetVendor() *string {
	return s.Vendor
}

func (s *AddCloudVendorAccountAKRequest) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *AddCloudVendorAccountAKRequest) SetAkType(v string) *AddCloudVendorAccountAKRequest {
	s.AkType = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetAuthModules(v []*string) *AddCloudVendorAccountAKRequest {
	s.AuthModules = v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetCtdrCloudUserId(v string) *AddCloudVendorAccountAKRequest {
	s.CtdrCloudUserId = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetDomain(v string) *AddCloudVendorAccountAKRequest {
	s.Domain = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetExtendInfo(v string) *AddCloudVendorAccountAKRequest {
	s.ExtendInfo = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetLang(v string) *AddCloudVendorAccountAKRequest {
	s.Lang = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetRegions(v []*string) *AddCloudVendorAccountAKRequest {
	s.Regions = v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetSecretId(v string) *AddCloudVendorAccountAKRequest {
	s.SecretId = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetSecretKey(v string) *AddCloudVendorAccountAKRequest {
	s.SecretKey = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetSubscriptionIds(v []*string) *AddCloudVendorAccountAKRequest {
	s.SubscriptionIds = v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetTenantId(v string) *AddCloudVendorAccountAKRequest {
	s.TenantId = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetVendor(v string) *AddCloudVendorAccountAKRequest {
	s.Vendor = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) SetVendorAuthAlias(v string) *AddCloudVendorAccountAKRequest {
	s.VendorAuthAlias = &v
	return s
}

func (s *AddCloudVendorAccountAKRequest) Validate() error {
	return dara.Validate(s)
}
