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
	// The AccessKey type. Valid values:
	//
	// - **primary**: Primary account
	//
	// - **sub**: Sub-account
	//
	// - **ctdr**: Agentic SOC
	//
	// 	Warning: When the vendor is **CHAITIN**, **FORTINET**, **THREATBOOK**, or **WIZ**, set this parameter to **ctdr**. When the vendor is **OracleCloud**, **google**, **BaiduCloud**, or **StateCloud**, set this parameter to **sub**.
	//
	// This parameter is required.
	//
	// example:
	//
	// primary
	AkType *string `json:"AkType,omitempty" xml:"AkType,omitempty"`
	// The list of modules associated with the AccessKey.
	//
	// When AkType is set to ctdr, this parameter is required and at least one module must be provided, such as SIEM, HOST, or CSPM. If this parameter is not provided, the API returns HTTP 400 with the error code -101.
	AuthModules []*string `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	// The account ID.
	//
	// > The account ID of the connected cloud vendor. This parameter is required when the permission description includes Cloud Threat Detection and Response.
	//
	// example:
	//
	// azure_demo_1
	CtdrCloudUserId *string `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The domain to which the connected account belongs. Valid values:
	//
	// - **china**: China
	//
	// - **global**: Global
	//
	// - **europe**: Huawei Cloud Europe
	//
	// > This parameter is valid and required only when **Vendor*	- is set to **HUAWEICLOUD**, **Azure**, **AWS**, **VOLCENGINE**, **KingsoftCloud**, **UCloud**, **BaiduCloud**, or **OracleCloud**. For **KingsoftCloud*	- and **BaiduCloud**, set this parameter to **china**. For **OracleCloud**, set this parameter to **global**.
	//
	// example:
	//
	// global
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The extension information.
	//
	// > This parameter is used to record the extension information of different vendors.
	//
	// Google Cloud is accessed through a service account. ExtendInfo stores the JSON-formatted service key file, excluding the private_key_id and zprivate_key fields. The file contains the following fields: type, project_id, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url, and universe_domain. When the vendor is **OracleCloud**, this parameter contains the **fingerprint*	- key, which indicates the fingerprint.
	//
	// example:
	//
	// {\\"product\\":\\"webFirewall\\",\\"remark\\":\\"remark\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of regions used to verify the AccessKey information. This parameter is valid only when **Vendor*	- is set to **AWS*	- or **OracleCloud**.
	//
	// >You can call the [ListCloudVendorRegions](~~ListCloudVendorRegions~~) operation to query the regions.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The AccessKey ID. Valid values:
	//
	// - If AkType is set to primary:
	//
	//   - **Tencent**: AccessKey ID of the primary account
	//
	//   - **HUAWEICLOUD**: AccessKey ID of the primary account
	//
	//   - **Azure**: Client ID
	//
	//   - **AWS**: AccessKey ID of the primary account
	//
	//   - **VOLCENGINE**: AccessKey ID of the primary account
	//
	//   - **KingsoftCloud**: AccessKey ID of the primary account
	//
	//   - **UCloud**: AccessKey ID of the primary account
	//
	// - If AkType is set to sub:
	//
	//   - **Tencent**: AccessKey ID of the sub-account
	//
	//   - **HUAWEICLOUD**: AccessKey ID of the sub-account
	//
	//   - **Azure**: Client ID
	//
	//   - **AWS**: AccessKey ID of the sub-account
	//
	//   - **VOLCENGINE**: AccessKey ID of the sub-account
	//
	//   - **google**: private_key_id
	//
	//   - **OracleCloud**: OCID of the sub-user
	//
	//   - **BaiduCloud**: AccessKey ID of the sub-account
	//
	//   - **KingsoftCloud**: AccessKey ID of the sub-account
	//
	//   - **UCloud**: AccessKey ID of the sub-account
	//
	//   - **StateCloud**: AccessKey ID of the sub-account
	//
	// >If AkType is set to **primary**, this value is the AccessKey ID of the primary account in the third-party cloud. If AkType is set to **sub**, this value is the AccessKey ID of the sub-account in the third-party cloud. For **Azure**, this parameter is not differentiated and the value is the **appId*	- of the authentication information. Google Cloud is accessed through a service account, and AkType defaults to **sub**. The value is obtained from the private_key_id property in the JSON-formatted service key file. If the vendor is **OracleCloud**, this value is the OCID of the sub-user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45GLRV4SOT0YFB****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The AccessKey secret. Valid values:
	//
	// - If AkType is set to primary:
	//
	//   - **Tencent**: AccessKey secret of the primary account
	//
	//   - **HUAWEICLOUD**: AccessKey secret of the primary account
	//
	//   - **Azure**: Client secret
	//
	//   - **AWS**: AccessKey secret of the primary account
	//
	//   - **VOLCENGINE**: AccessKey secret of the primary account
	//
	//   - **KingsoftCloud**: AccessKey secret of the primary account
	//
	//   - **UCloud**: AccessKey secret of the primary account
	//
	// - If AkType is set to sub:
	//
	//   - **Tencent**: AccessKey secret of the sub-account
	//
	//   - **HUAWEICLOUD**: AccessKey secret of the sub-account
	//
	//   - **Azure**: Client secret
	//
	//   - **AWS**: AccessKey secret of the sub-account
	//
	//   - **google**: private_key
	//
	//   - **OracleCloud**: private_key of the sub-user
	//
	//   - **BaiduCloud**: AccessKey secret of the sub-account
	//
	//   - **KingsoftCloud**: AccessKey secret of the sub-account
	//
	//   - **UCloud**: AccessKey secret of the sub-account
	//
	//   - **StateCloud**: AccessKey secret of the sub-account
	//
	// >If AkType is set to **primary**, this value is the AccessKey secret of the primary account in the third-party cloud. If AkType is set to **sub**, this value is the AccessKey secret of the sub-account in the third-party cloud. For **Azure**, this parameter is not differentiated and the value is the **password*	- of the authentication information. Google Cloud is accessed through a service account, and AkType defaults to **sub**. The value is obtained from the private_key property in the JSON-formatted service key file. For **OracleCloud**, this value is obtained from the content of the corresponding private key file.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE6SLd****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The list of subscription IDs.
	//
	// > This parameter is deprecated.
	SubscriptionIds []*string `json:"SubscriptionIds,omitempty" xml:"SubscriptionIds,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// > This parameter is valid only when Vendor is set to **Azure*	- or **OracleCloud**.
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
	// - **CHAITIN**: Chaitin Tech
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
	// - **WIZ**: Wiz
	//
	// - **OracleCloud**: Oracle Cloud
	//
	// - **StateCloud**: China Telecom Cloud
	//
	// This parameter is required.
	//
	// example:
	//
	// AWS
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the AccessKey account.
	//
	// > This parameter is used to identify the account to which the third-party host assets belong.
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
